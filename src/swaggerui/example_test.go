package swaggerui_test

import (
	"net/http"
	"strings"

	"swaggerui"
)

func Example() {
	http.Handle("/", swaggerui.Handler("/", strings.NewReader("{}")))
}
