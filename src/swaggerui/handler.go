//go:generate sh -c "go-bindata -pkg=internal -o internal/files.go `find third_party -type d`"

package swaggerui

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"assetfs"
	"swaggerui/internal"
)


var index = `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link rel="icon" type="image/png" href="third_party/swagger-ui/images/favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="third_party/swagger-ui/images/favicon-16x16.png" sizes="16x16" />
  <link href='third_party/swagger-ui/css/typography.css' media='screen' rel='stylesheet' type='text/css'/>
  <link href='third_party/swagger-ui/css/reset.css' media='screen' rel='stylesheet' type='text/css'/>
  <link href='third_party/swagger-ui/css/screen.css' media='screen' rel='stylesheet' type='text/css'/>
  <link href='third_party/swagger-ui/css/reset.css' media='print' rel='stylesheet' type='text/css'/>
  <link href='third_party/swagger-ui/css/print.css' media='print' rel='stylesheet' type='text/css'/>
  <script src='third_party/swagger-ui/lib/jquery-1.8.0.min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/jquery.slideto.min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/jquery.wiggle.min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/jquery.ba-bbq.min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/handlebars-2.0.0.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/js-yaml.min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/lodash.min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/backbone-min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/swagger-ui.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/highlight.9.1.0.pack.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/highlight.9.1.0.pack_extended.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/jsoneditor.min.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/marked.js' type='text/javascript'></script>
  <script src='third_party/swagger-ui/lib/swagger-oauth.js' type='text/javascript'></script>

  <!-- Some basic translations -->
  <!-- <script src='third_party/swagger-ui/lang/translator.js' type='text/javascript'></script> -->
  <!-- <script src='third_party/swagger-ui/lang/ru.js' type='text/javascript'></script> -->
  <!-- <script src='third_party/swagger-ui/lang/en.js' type='text/javascript'></script> -->

  <script type="text/javascript">
    $(function () {
      var url = window.location.search.match(/url=([^&]+)/);
      if (url && url.length > 1) {
        url = decodeURIComponent(url[1]);
      } else {
        url = "/api.json";
      }

      hljs.configure({
        highlightSizeThreshold: 5000
      });

      // Pre load translate...
      if(window.SwaggerTranslator) {
        window.SwaggerTranslator.translate();
      }
      window.swaggerUi = new SwaggerUi({
        url: url,
        validatorUrl: null,
        dom_id: "swagger-ui-container",
        supportedSubmitMethods: ['get', 'post', 'put', 'delete', 'patch'],
        onComplete: function(swaggerApi, swaggerUi){
          if(typeof initOAuth == "function") {
            initOAuth({
              clientId: "your-client-id",
              clientSecret: "your-client-secret-if-required",
              realm: "your-realms",
              appName: "your-app-name",
              scopeSeparator: ",",
              additionalQueryStringParams: {}
            });
          }

          if(window.SwaggerTranslator) {
            window.SwaggerTranslator.translate();
          }
        },
        onFailure: function(data) {
          log("Unable to Load SwaggerUI");
        },
        docExpansion: "list",
        jsonEditor: false,
        defaultModelRendering: 'schema',
        showRequestHeaders: false
      });

      $('#input_apiKey').change(function() {
        var key = $('#input_apiKey')[0].value;
        if(key && key.trim() != "") {
          swaggerUi.api.clientAuthorizations.add("key", new SwaggerClient.ApiKeyAuthorization("Authorization", "Bearer " + key, "header"));
        }
      })

      window.swaggerUi.load();

      function log() {
        if ('console' in window) {
          console.log.apply(console, arguments);
        }
      }
  });
  </script>
</head>

<body class="swagger-section">
<div id="message-bar" class="swagger-ui-wrap" data-sw-translate>&nbsp;</div>
<div id="swagger-ui-container" class="swagger-ui-wrap"></div>
</body>
</html>`

// BasePath is the base path of the embedded swagger-ui.
const BasePath = "/third_party/swagger-ui/"

// SpecFile is the name of the swagger JSON file to serve.
const SpecFile = "/api.json"

// Handler returns an HTTP handler that serves the
// swagger-ui under /third_party/swagger-ui.
func Handler(basepath string, data io.ReadSeeker) http.Handler {
	if basepath == "" {
		basepath = "/"
	}
	as := &assetfs.AssetStore{
		Names: internal.AssetNames,
		Data:  internal.Asset,
		Info:  internal.AssetInfo,
	}
	//解压到对应的内存
	fs, err := assetfs.New(as)
	if err != nil {
		panic(fmt.Sprintf("failed to create static fs: %v", err))
	}
	mux := http.NewServeMux()
	fsh := http.FileServer(http.FileSystem(fs))
	if basepath != "/" {
		fsh = http.StripPrefix(basepath, fsh)
	}
	//p := assetfs.AddPrefix(basepath, BasePath)
	f := assetfs.AddPrefix(basepath, SpecFile)
	mux.HandleFunc(basepath, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == basepath {
			w.Write([]byte(index))
			return
		}
		fsh.ServeHTTP(w, r)
	})
	mux.Handle(f, &handler{modTime: time.Now(), body: data})
	return mux
}

type handler struct {
	modTime time.Time
	body    io.ReadSeeker
}

func (f *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, SpecFile, f.modTime, f.body)
}
