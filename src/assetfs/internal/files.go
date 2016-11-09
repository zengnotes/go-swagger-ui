// Code generated by go-bindata.
// sources:
// fs.go
// DO NOT EDIT!

package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _fsGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\xcd\x6e\xdb\x38\x10\x3e\x4b\x4f\x31\xf5\xa1\x90\x02\x55\xc9\x02\x8b\x3d\x04\xf0\xa1\x68\x1b\x20\x40\x9b\x2d\xd6\x7b\x33\x82\x85\x2c\x91\x36\x61\x89\x14\x48\xe6\x6f\x5d\xbf\xfb\xce\x0c\x29\x59\xb1\x1d\xa0\xdd\xa0\x4d\x44\x72\xe6\x9b\xf9\xe6\x87\xc3\xcb\x4b\xf8\x5e\xd5\xdb\x6a\x2d\x40\x3a\xe8\xad\x79\x54\x8d\x70\x50\x69\xd8\x78\xdf\x97\x37\xaa\x15\x8b\x17\xe7\x45\x07\xd2\x58\x58\x9b\x0f\x2b\xa5\x9b\xca\x57\x50\x39\x27\xbc\x4b\xfb\x51\x39\x4d\x55\xd7\x1b\xeb\x21\x4b\x93\xd9\xea\xc5\x0b\x37\xc3\x0f\x2d\xfc\x25\x21\xd1\xb7\xe1\x9d\xbe\xf2\x1b\xfa\xeb\x50\x96\xff\x7a\xab\xf4\x9a\x8f\xbc\xea\xc4\x2c\xcd\xd3\x14\xbd\xfa\x48\xf8\x0b\x6f\xac\x98\x78\x55\xd7\xc2\x39\xf0\x26\x5a\x87\xb5\xd0\xc2\x56\x5e\x34\xb0\x7a\x99\x38\x57\xa6\xfe\xa5\x17\x53\x08\x34\xf2\x50\x7b\xd8\xa5\xc9\x5d\xd5\x21\x92\x7c\xd0\x75\x96\xc3\xf2\x3e\x58\x4f\x93\xcf\xc4\x29\x6c\x4b\xe4\xac\x51\x0a\xc2\x59\x0e\xd9\xf2\x9e\xf8\x14\x20\xac\x35\x36\x4f\x93\x5b\x2d\xcd\x9b\xc2\xc6\x71\xd4\x48\x66\xd4\xd8\x1f\x28\xdd\x2c\x40\x21\x13\x94\xaf\xbc\xaa\x81\xf4\xc1\x09\xfb\x28\x2c\xf8\x4d\xe5\x01\x83\xd8\x8a\x4e\x68\x64\xe7\x37\xe2\x24\x0b\x4a\x7b\x61\x65\x55\x8b\x29\x47\xc4\x3c\x10\xec\xa0\xab\xfa\x65\x70\xe7\xfe\x82\x03\x45\xe4\xc8\x07\xd6\x18\x77\x26\x3a\x2b\xc0\x9f\xc0\x32\x4d\x14\xb1\x9b\xb0\x88\xde\xdf\x89\xa7\xc1\x58\x6d\x05\x06\x9d\xab\x64\xd8\x92\xd6\x74\xc8\xaa\x55\xce\x83\x91\x4c\x8b\xce\x1b\x22\xa1\x2c\x70\x3d\x94\x29\x57\x9b\x47\x06\x3a\x04\x61\xdd\x9a\x55\x50\xa5\xaa\x28\xbf\x55\xbe\xde\x14\xac\x46\xfe\xf7\x80\x0e\x0a\xe8\x44\xa5\x39\xe9\xbc\x47\x18\x46\x4a\x0e\x4e\x1f\xb1\x32\x51\xae\x4b\x68\xd0\x4e\x6f\x85\x54\xcf\x79\xc0\x24\x11\x8e\x2f\x25\xa8\x4c\x29\x5f\x13\x16\x99\x84\x8b\x43\x85\x14\x23\x58\x88\x5c\x11\x3d\x58\x19\xd3\x62\x52\x2f\xa2\xd2\x90\x51\x0a\x9a\x13\x3d\x5c\xcf\xa3\x3c\xa5\x1d\xa9\x6d\x16\xa2\xaf\xb0\x22\xb9\x4c\x7a\x77\x38\x77\xe5\xa2\x6f\x95\xcf\xa2\x15\x84\x17\x7d\x4e\xc9\x42\x89\xae\xda\x8a\xec\x6c\xd2\x50\x82\xda\xee\x9f\x02\xa4\x26\x49\x5b\x69\xea\xb5\x92\x8b\x38\x63\x2f\x92\xc7\xca\xc2\x16\x86\x3a\x4e\x94\x1c\x99\xbc\x9b\xc3\x6c\xc6\x32\x89\xd9\xb2\xe7\x84\x71\x88\xf4\xc1\x19\xa9\x73\x92\x42\xdd\x77\x66\x0b\x3f\x7e\xb0\x2c\xaa\x6b\xd5\x06\xfd\xa4\x36\xda\x2b\xfd\x20\x68\xb1\x8f\xb2\x21\x42\xe1\x5c\xf6\x27\x5c\xef\x32\x19\x78\x16\x80\x3d\x92\xf5\x2e\x67\x23\xc9\x16\xe6\x20\xfb\x25\xed\xc9\x3e\xff\xf0\xdb\xfd\x80\x49\xff\xd5\xd0\x36\x04\x27\x4b\xaa\xbf\x2c\x78\x87\x06\x8f\xbd\xb2\xc2\x3f\x20\x4f\x5c\xb3\x4a\x84\xa0\x1b\x60\x02\x41\x71\xfc\x55\x08\x14\xdc\x4e\xc3\x87\xd1\x9f\xc3\x76\xc0\x57\xb6\x08\x75\x35\x44\x33\xa4\x76\xb4\xc1\x67\x73\xd4\xc6\xdb\x48\x3c\x97\x1b\xdf\xb5\x11\x87\x4a\x74\x0e\x1f\x9b\xe6\x3b\xd7\x69\x36\xbb\x9c\x15\x01\xe2\x53\x8b\x55\x9e\xe1\x79\x08\x52\xb7\xc4\xcf\x7b\x94\x7d\x3f\xd6\xc2\x6e\x75\x1d\x5b\x74\xb7\x2f\x80\xe2\x74\x4d\x25\xcf\x11\x22\xbd\x21\x82\xdb\x53\x0b\xc1\xb3\x6e\xb9\x3d\x45\x0c\xc1\x0a\x68\xf4\x1b\x11\xf0\x5f\x8c\xca\xfb\x58\xf5\xbb\x0e\x2d\x62\x88\xe2\x4d\xf0\x67\x2f\xf4\x2f\xdc\x53\xdc\x77\x19\x4e\x97\xa1\x89\x72\x46\xc8\x5e\x5f\x9b\x23\xc0\xb4\xc5\x24\x2e\x9e\xf1\x4a\xe1\x46\x92\xae\xec\x96\xa4\x84\x05\x43\x95\x1a\x4f\x28\xb0\xd3\x2c\x62\x23\x7e\xb1\xf6\xce\xf8\x2f\x74\x3e\xa5\xa3\xc5\x13\x19\xc8\x64\xb9\xc2\x98\x94\x44\x37\x9f\xf2\xba\xe1\x4b\x2b\x08\xbb\x93\xcb\x4c\x69\x26\x1a\x39\x9c\xa5\xc5\x00\x93\xb1\xc2\x14\xc6\x0e\x1f\x76\x43\x3b\x10\x1b\xca\xb5\xa2\xf3\xab\xd0\xe7\xaf\x9a\x1c\xcf\x99\x9b\x5c\x2a\x4a\x9b\xd4\x4c\x85\xa6\x66\xb9\x08\x6d\x96\xc9\x7c\xe4\x26\x23\x87\xaf\x98\x9a\x81\x01\xb9\x8b\xa6\xd6\x7e\x43\x2c\x68\xb5\x56\x8f\x78\xce\xf5\xf9\x13\x74\xbe\x1e\x27\x09\x93\xfa\x3f\xb2\x72\x35\xcd\x01\x53\x2f\x57\x3c\x12\xd9\xe4\x90\x14\x7e\x54\x0c\x83\xf6\x78\x04\xe5\x10\xee\x44\x5a\x13\x34\xb7\x76\x98\x29\x78\xa1\xff\x25\xaa\x46\x58\x46\x38\x04\xe4\xfd\xa8\x40\xae\x04\x91\x6b\x9c\x72\xb6\xc0\xe5\x80\x1b\x6a\xbe\x20\x07\x5f\x8d\x47\xb6\x73\x18\x8f\x17\xc1\x54\x00\x49\x93\xa3\xd9\x18\x23\x37\x71\x31\x87\x4f\xad\x71\x02\x2b\x81\x4b\x99\x3d\x1e\x2b\xf4\x2d\x15\x42\xc7\x3e\xce\x6a\xf3\x80\x61\xc6\x50\xf3\xc3\xe3\xcc\x6b\xe2\x35\xdc\x50\xf0\xb7\xfa\xb1\x6a\x55\xf3\x16\xfa\x02\x9f\x1b\xd9\xf9\xd7\xc9\x04\x4f\x4e\x0e\xa3\xab\x1c\x94\x78\xcf\x0c\x53\x26\x5a\x68\x86\xfd\x1c\x68\x20\x21\x7c\x2c\xfb\xf0\xb3\x8b\x85\x38\x4c\xc8\x26\x87\xfd\xa9\xe6\x42\xfd\x4b\x9a\xc8\xf7\x8f\xdf\xe1\x58\xf3\xea\x9c\xc6\x37\xd3\x90\x46\x64\x42\xab\x9f\xd0\xf8\x5b\xb1\x83\xf4\xc2\x2c\xe9\xfb\xa0\xc1\x5b\x77\xe6\x29\x3b\xeb\xde\xad\xfb\x8c\x39\xc9\xf9\x1d\x70\xe2\x1e\x3f\x4f\xce\x71\x7a\x71\x81\x52\xb8\x02\x77\xfb\xa9\x12\x8d\x9e\x7d\xfa\x5f\x00\x00\x00\xff\xff\xc1\x0f\x16\x3f\x78\x0b\x00\x00")

func fsGoBytes() ([]byte, error) {
	return bindataRead(
		_fsGo,
		"fs.go",
	)
}

func fsGo() (*asset, error) {
	bytes, err := fsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fs.go", size: 2936, mode: os.FileMode(420), modTime: time.Unix(1461271662, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"fs.go": fsGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"fs.go": &bintree{fsGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

