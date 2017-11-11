// Code generated by go-bindata.
// sources:
// neb.js
// DO NOT EDIT!

package jslib

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

var _nebJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x6d\x6f\xdb\x46\x12\xfe\x2c\xfe\x8a\xf1\x1e\x60\x93\x30\x41\xd9\x69\x5a\x14\xe4\x11\x45\xda\xe4\x2e\x3e\x24\xb6\xe1\x38\xb8\x02\x86\x11\xac\xc4\x91\xc4\x94\xda\xe5\xed\x2e\xad\x18\xaa\xfe\xfb\x61\x5f\x48\x91\x12\x65\xc9\xc6\x05\x07\xd4\x5f\x28\xef\xcb\x33\xcf\x3c\x33\x3b\x3b\xa4\xc0\xff\x54\xb9\xc0\xd4\x9f\x54\x6c\xac\x72\xce\x00\x7d\x15\xb2\x50\x04\xcb\x66\x44\xfa\x3c\xac\x82\x65\x3e\xf1\x8f\xd8\x1d\xbf\xb7\xbf\x94\xf9\xf5\x40\x05\xd0\x54\x3d\x96\xc8\x27\x50\x63\xa5\xa4\xde\x4a\x8e\x8f\xdd\x60\xa2\xf7\x54\xc7\xc7\x34\x10\xa8\x2a\xc1\x80\xfa\x3c\x3c\x3a\x0b\xf4\x78\x5e\x8f\xe5\x6e\x4c\xa3\x4e\x52\x86\x0b\x78\x27\x04\x17\x3e\xf9\x8d\x32\xc6\x15\x4c\x72\x96\xc1\x9c\x67\x55\x81\x70\x42\x4e\xf9\x29\x39\x21\x41\xa2\x66\x82\x2f\x60\x12\x8d\x79\x86\x29\xf9\x78\xf5\xf6\xf3\x87\x77\x5f\x2e\xaf\x6e\xbf\xfc\xe3\xea\xf3\xe5\x5b\x12\x4e\x56\x1a\xaf\x48\x35\xf7\x74\x89\xdf\x4a\x2e\x94\x8c\x97\xab\x55\xa2\x7d\xb8\x3b\xbb\x8f\xc6\xb4\x28\xfc\x22\x72\x53\x61\xcd\xde\x47\xeb\x20\x4b\xcd\xc2\xf3\xfb\x3b\xbc\x4f\x1c\x55\xe9\xb3\x5f\x58\x8c\xc1\x2a\x2c\xc2\xf5\x4e\x0c\xad\x76\x2b\xb7\x4a\x9b\xac\x27\x0d\x8b\xfc\x20\xad\x26\x5c\xf8\x7a\x35\x4f\xcf\x12\xfe\x77\x11\x15\xc8\xa6\x6a\x96\xf0\xd3\xd3\x40\xfa\x42\x0b\xdf\xd0\x58\x05\xfe\xf2\x3c\xbe\x6b\x28\x3b\x88\xd0\xaa\x14\x3a\xdb\xc1\xd2\xf3\x48\x25\x11\xa4\x12\xf9\x58\x91\xc4\xf3\x34\xfe\x9b\x6c\x9e\x33\x48\xa1\x89\xb4\xcf\x70\x14\xc0\xd2\x1b\xa8\x59\x2e\xa3\x2f\x1a\x0c\xa5\x7a\x4f\x59\x56\xa0\x80\x14\x18\x8e\x36\x47\x13\x6f\x95\x78\x9e\x41\x8a\x4a\xc1\x15\xd7\x0e\x46\x0c\x17\x6f\xc6\x63\x5e\x31\xd5\x81\x2f\xa9\x94\xe5\x4c\x50\x89\xc6\x8a\xe6\x50\x52\x41\xe7\x12\x52\x58\x92\xf5\x2c\x89\x61\xfd\xcf\x2a\xf1\x06\xce\x5d\x43\xcb\xd9\xf7\xc9\x14\x15\x09\x81\x0c\x1f\xce\x87\x6b\x7b\x24\x74\x90\x41\x3f\xb3\x8a\x15\x7c\xfc\x47\x1f\x39\x9a\x65\x02\xa5\x0c\xe1\x69\x96\x6e\x19\x89\xa1\xde\xe0\x0d\xe0\xd9\xdc\x4b\x2e\x1b\xf2\x96\xd2\x3e\xe2\x7b\x68\x1f\x44\xf5\x20\x3a\x87\x90\x91\xf9\x94\xdd\x0a\xca\x24\xb5\x2c\xda\x84\x26\x82\xcf\x43\x50\x3c\x84\x07\x5a\x54\x18\x02\xe3\x6c\x8c\x21\x48\x5e\x09\xfd\xa4\x62\xda\xcb\x56\xef\x23\x31\x98\xed\xde\x80\x28\x4e\x62\x8d\xe2\x0d\x88\xc1\x21\xb1\xc3\xf3\x06\xc4\x20\x92\xd8\x21\x7b\x03\x62\xb1\x49\x5c\x1b\xf1\x06\x44\x9b\xd1\x9e\x8b\xa9\xf4\x06\x07\x39\xae\x9d\xda\xeb\x38\xb2\xac\xe5\xf8\xbf\x73\x35\xbb\x6e\xa2\xfd\x6c\x19\xf6\xe5\xda\xf7\x93\x44\xff\xbb\x23\x67\x0f\x54\x4b\xed\x52\x61\x9f\x84\x0e\xad\x23\xd6\x1c\xd5\x8c\x67\x21\xd0\x32\x6f\x76\x6b\x3d\xda\x2c\x36\x0a\x4f\xc3\xaa\x6f\xaf\xb5\x6c\x8b\x60\x5d\x80\x21\xb5\x05\x2f\xf1\x56\xe1\x72\x75\x1f\xbe\x7a\x71\xd9\xbc\xbe\xf8\xdf\x14\xcd\xeb\x8b\x96\x2c\x53\x54\x97\x38\xfa\xa4\xa8\xea\xe6\xd1\x96\x0e\xfd\xf5\x6f\x34\x94\x7a\x2b\x09\xfa\xb0\x19\xcf\xf0\x82\x4d\xf8\x0b\x80\x79\x86\xc3\x9c\x4d\x78\x3f\x30\xb5\x05\x49\x3e\x1f\xb8\xde\xd9\x8f\x3b\xd2\x45\xe8\x6d\x35\x2f\x3b\xc0\x66\x47\xdf\x39\xb1\x95\x3f\x36\x8f\x83\x92\xd7\xe0\x0f\xb3\x6a\x5e\x6e\x65\xeb\x66\x50\x5c\xcd\xdd\x0e\xcc\x21\x85\xf7\x39\x75\xd7\x29\xe2\xe2\xf8\x24\xad\x8d\x2a\xf4\xd7\x28\xbf\xad\x82\xf2\xb4\xf7\xba\x5d\xdb\xe1\xb2\xf3\x55\x4f\xbd\xc4\xd3\x2d\xc7\x9a\xe6\x2c\xb6\x98\x2f\x72\x4c\xf3\xdd\x1f\xcf\x1b\xba\xd8\x15\xd2\x8c\x2a\xda\xe7\x88\x1e\x27\x31\xe8\xc7\x41\x44\x04\x5d\x1c\x2c\xf2\x14\xd5\xaf\xfa\x90\xfc\xfa\xf8\x9e\xca\x59\x87\xce\x8c\xca\x59\x1f\x1d\x3d\x4e\x62\xd0\x8f\x83\xe8\x74\x4d\xec\xa5\xd3\x52\xe7\x06\xc7\x98\x97\xea\x7b\xb1\xda\xb6\xf4\x34\xb9\xff\xe7\x8d\x76\x7d\x51\xdf\x67\x3f\xbc\xe4\x3e\x1b\x0e\x61\x24\xf8\x42\xa2\xf0\xf2\x09\xf8\xee\xed\x64\x91\xb3\x8c\x2f\xe0\x28\x4d\x81\x54\x2c\xc3\x49\xce\x30\x23\x70\x7c\xec\x66\xa2\xdf\x3f\x7e\x78\xaf\x54\x79\x63\x09\x6b\xc7\x00\xba\x63\x90\xf6\xaf\x4d\x60\x38\x84\xaf\x72\x96\x33\x05\xf9\x94\x71\x81\x31\x14\x39\x43\xcd\x44\x5f\x36\xde\x0a\xb0\x90\xd8\x8f\xe8\xdc\xf2\xc9\xb7\x79\x31\x53\xaa\x74\x82\x91\xe0\x50\x23\x2b\x7b\x83\xff\xfe\xfe\xe6\x55\x07\x6e\x26\x5e\x91\x20\x01\x3b\xdb\xb5\xd9\xca\x30\x2e\x55\x08\x2a\x9f\x23\xaf\xd4\xfa\xc6\xd7\xc3\x90\x82\x79\xfc\xf9\x27\x10\x4d\x2c\x1e\xea\x36\x9a\x16\x7a\x30\xfe\xf9\xec\xe7\x33\x92\xb8\xd5\x6e\x3b\xa4\x35\x90\xde\x73\x66\xa3\xdb\x32\xdc\xca\xae\x2f\x0c\x17\x7d\x74\x3a\x59\x42\xe5\x23\x1b\x37\xe9\xef\x74\x49\xbc\x81\x0e\xea\x7a\x6e\xb0\xce\x54\xfd\x76\xad\x65\xf0\x83\x64\x3d\xde\x66\xd7\x22\x9b\x78\x83\x26\x2a\x5b\x18\x1d\xe1\x0d\xda\xca\x92\x98\x43\x0a\xe4\x9f\xef\x6e\x89\xa3\x61\xf9\x46\x8a\x7f\x2e\x4b\x14\xbf\x51\x89\x7e\x00\xa9\x4e\xb1\xeb\xab\x4f\xb7\xc4\x12\x34\x9b\xcc\xff\x6b\xa0\x4a\x14\x35\x21\x23\xf2\xa9\x76\xd9\x1c\x65\x4b\x9a\x97\xc8\xfc\x79\xa8\xd7\xd5\x42\xac\x0f\x7a\x23\xc5\x6e\x81\x0f\x38\xbe\x8f\x05\xa7\xd9\xa6\xbc\x35\xa9\x56\x80\xba\xfb\x26\xb4\x90\xa8\xb9\x28\xf1\x68\xbc\xd3\x32\x38\x30\xe3\x79\x73\xb6\x74\x12\xb4\x27\x88\x93\xa3\xf1\x51\xdf\x10\x36\x54\xad\x48\x74\x67\xff\xf5\xe9\xea\x32\xd2\x27\x9b\x4d\xf3\xc9\x63\x6d\x27\xb0\x9b\x74\x04\xc7\x54\x8d\x67\xe0\xa3\x10\x5c\x58\x78\xfb\xe9\xc4\x0c\x18\xb9\x6b\xf7\x64\x55\xd4\xc7\x4d\xc3\x0b\x94\x25\x67\x12\x6f\xf1\x9b\x6a\x79\xd3\xac\x33\x96\x4b\x2a\x24\xfa\x76\xcc\x64\x41\x63\xaf\x63\xcb\xd9\x69\xa2\xa3\x97\x3f\x15\x1c\x13\xcf\xbd\xe9\xef\x9c\x0d\x41\xdf\xb5\x23\x3a\xfe\xe3\xd9\xb1\x52\xa2\xc2\xa0\x9d\x54\x4c\x20\xcd\x1e\x4d\x33\x36\x9e\x51\x36\xed\xe9\xca\x01\x00\x74\x48\xd7\x3a\xd1\xec\xd1\x35\x8a\x69\x0a\xaf\x75\xc1\xdc\x3c\x59\xba\xa6\x9e\x37\xbb\x01\x0e\x11\x7c\xbd\xd2\x84\x4a\x9f\xbc\xaa\x28\x12\xaf\x9e\x71\x01\x01\xf7\xf7\x74\x5c\xec\x9a\x8d\xe8\xd4\x5b\xcd\xb9\x45\x29\xa9\x71\xf7\xe8\xc8\x41\x1d\x1f\x37\xbf\x23\x4b\x61\x6b\x24\xaa\xb7\xfd\x02\xbd\xc3\x31\x90\x0b\xf6\x40\x8b\x3c\x83\xda\xbb\x18\x08\x9c\xc2\x46\xda\x36\x44\x6b\x4a\x83\xc6\xe7\xe6\x7b\xa0\xc3\x6c\x79\xd3\x48\x51\xc7\xdf\x66\x79\x08\x5d\xbf\x57\x1e\x80\x4e\x35\x80\x75\x98\xd7\x25\xaf\x1b\xdd\x0e\x5a\xeb\x5b\xe4\x98\x33\x86\x76\x9d\xdb\x4a\x02\x43\xd7\x02\xdb\x58\x00\x3c\xff\xac\x6f\x1d\xf5\xed\x9a\xbb\xf7\xa0\x1b\x07\xb7\x0e\xfa\xb6\x30\x96\x70\x7f\x4b\xd1\xbe\x47\x75\x6b\x61\x6f\xc7\xf8\xa7\x70\xf3\xd6\x8d\x7f\x5c\xdd\x87\xaf\x5f\xfa\x16\xbd\xe3\x66\x8f\x86\x2d\x13\xd1\x57\xf3\x7e\xd8\x7a\xeb\x6e\xad\xa3\x65\xee\xe6\xdb\xdf\x32\xdb\x0b\xcc\x87\x87\x16\xc4\x25\x8e\x3a\x61\x16\xeb\xf6\x65\xd0\x3a\xc7\x75\xbd\xea\x7d\x9b\x5f\x5f\xad\xad\xf0\xec\x7c\xf3\x5f\xc0\xf6\xe5\xe8\x1a\x01\x5a\xe6\x6e\xc9\x9b\xeb\x0b\x5f\x0f\x05\x75\x8f\x40\x9d\x2f\x66\x52\xff\xae\xa7\x75\xbc\x2e\x71\xd4\x79\x73\x50\x37\x9b\x66\xfb\x1d\xdc\xe7\x4f\x6f\x2e\x5c\xe2\xc8\xe6\x40\x4b\xcd\xf8\x3c\x5c\xab\x1f\xbf\x0a\xb7\x43\x16\xff\xb0\xba\x0f\x7f\x3c\x24\x31\xb6\xbb\xd1\x29\x07\x64\x0f\x90\x71\x94\x8c\x28\x98\xd1\x07\x04\xca\x32\x60\x88\xd9\x46\xb3\xd1\xee\x58\x37\x5b\xc5\x6e\xe7\x5a\x1f\x02\x67\x38\xda\x6a\x2c\x97\xab\xa4\xdd\x78\x3e\xb1\x72\x6f\xa3\x59\xf7\x99\xae\x2b\xff\xe9\x10\x19\xb6\x74\xdf\xb0\x52\x83\x11\x86\x23\x72\x08\xe0\x3a\xdb\xeb\xf3\x70\x12\x0d\x8b\x7c\x34\x64\x38\x3a\x09\xac\xd2\x19\x67\x0a\xf8\x03\x0a\x91\x67\x08\xd3\x82\x8f\x68\xa1\x2f\x80\x9c\x8e\x0a\xdc\xf1\x36\x70\xd2\x68\x7a\xa2\xaf\x80\xce\x82\xc8\x18\xec\x2e\xaa\x85\x6f\xaf\x70\x49\xb5\x2b\xdb\x5c\xba\x39\xb2\x24\x7e\xbd\xba\xd7\xce\x87\x77\xc6\xf9\xfb\xc0\x1b\x0e\xff\xe6\xbe\x33\x7c\xa4\x65\x99\xb3\xe9\xe7\x9b\x0f\x29\xc3\x51\xf4\x55\x46\x73\x5a\x7a\xff\x0d\x00\x00\xff\xff\x83\xe1\x7b\x64\xd3\x1a\x00\x00")

func nebJsBytes() ([]byte, error) {
	return bindataRead(
		_nebJs,
		"neb.js",
	)
}

func nebJs() (*asset, error) {
	bytes, err := nebJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "neb.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"neb.js": nebJs,
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
	"neb.js": &bintree{nebJs, map[string]*bintree{}},
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