// Code generated by go-bindata.
// sources:
// _assets/bash_comp
// _assets/fish_comp
// _assets/fish_wrapper
// _assets/sh_wrapper
// _assets/zsh_comp
// DO NOT EDIT!

package scripts

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

var __assetsBash_comp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x54\xdf\x6b\xdb\x30\x10\x7e\xf7\x5f\x71\x78\x86\xc6\x03\x51\x92\x3d\x06\xef\xa5\x5b\x61\xb0\xd1\xd2\x3e\x94\x91\x05\xa1\xda\xe7\x58\xa0\x48\xae\xe4\xb6\x6c\x4d\xff\xf7\x9d\x6c\xc7\xb1\x49\xe2\x41\x17\x68\x20\x96\x74\x3f\xbe\xef\xbb\x13\xba\x0f\x90\x9a\x75\xa9\xb0\x42\xa8\x84\x5d\x61\xe5\x02\xfe\x90\x4b\xde\x1e\x26\x31\xbc\x04\x00\xca\xa4\x42\xc1\xb7\xcb\xdb\x24\x3a\xfb\xa5\xcf\xc8\x72\x71\xf5\xe3\xfa\xe6\xeb\xf5\xf7\x9f\xc9\x04\xa2\x89\xc7\x58\xa1\x06\x76\x07\x61\x34\xa1\x7c\x60\x0a\x66\x9f\xcf\x33\x7c\x3a\xd7\x8f\x4a\xc5\x21\x30\x46\xae\x69\x18\x43\x1c\xbc\x06\x41\x8f\xd6\x94\x95\x34\xba\xa5\xa5\xc3\x90\xd3\x1b\xe8\xe0\x97\x24\x64\x02\x58\xe6\xa1\xd9\x1a\x98\x05\xe6\x08\xd5\x15\xa8\xc8\xc0\x68\x29\x69\x79\x42\xeb\x08\x2e\x1c\x97\xe8\xe1\x46\x24\x69\x53\x15\x52\xaf\x1a\x49\xda\x68\x6c\x25\x1d\x47\xdc\x81\xac\x85\xd4\x90\x49\x57\x8a\x2a\x2d\xd0\xd6\x18\x83\x8a\xd2\x47\x3b\x84\x8a\xe9\xc8\xa9\xd9\xdc\x03\xf2\x67\x63\x33\xc7\xef\x7f\x73\x8b\x79\x1b\x2b\x73\x58\x2c\x20\xf2\x29\xfc\xe2\xee\xea\xe6\x0b\x30\x7c\x80\x29\x2c\x97\x73\xa8\x0a\xd4\x14\x02\xb0\xdf\x51\x30\xb6\xbb\x53\x1f\xd1\xc0\x84\x11\x81\x86\x90\x24\xc0\x3e\x12\x02\xf4\x20\x48\xc6\xf6\x0e\xda\xb0\xda\x8c\xca\x61\xdf\xdf\x62\xf6\x43\x72\x19\xf4\xe2\x7a\x52\xee\x85\xc3\x0c\x8c\xee\x6e\xd9\xfb\x53\x32\x52\xf6\x4b\x5d\x90\xaf\xe7\x76\x31\x5d\xbe\x86\x20\xb7\x32\x58\x16\xb7\xbb\x83\xc5\xaf\x10\x66\xc3\xe2\xc7\xd5\x75\x0a\x9b\xdf\x7c\xbe\xa5\x11\xa3\x34\xd4\xe3\x19\x6c\x36\x7b\xc6\x4f\x07\xb9\x73\xa9\x30\x93\x76\x9c\x50\xfd\x9b\xf0\x14\x75\xad\x4f\x4e\x83\xea\x30\xd0\x7f\xf4\xc2\xbe\xb3\xc8\xed\xdb\x1e\x57\xe9\x36\xcd\x68\xd9\x74\xa3\x65\x27\x7b\x1f\x62\x97\xd7\xcc\xa5\x37\x95\x78\x7c\xca\xd0\x6b\x2a\x48\x26\x7d\xfe\xb8\x62\x30\xbf\xde\xd4\x82\xe3\x44\x7e\x0f\xcf\x56\x94\x25\xda\x83\x3c\xfb\xbd\x42\x27\xd2\xa0\x76\xd0\x18\xec\xde\x3f\xbb\xac\xbb\x04\xf4\x0f\xfe\x06\x00\x00\xff\xff\xcf\xc1\x9a\x90\x6b\x06\x00\x00")

func _assetsBash_compBytes() ([]byte, error) {
	return bindataRead(
		__assetsBash_comp,
		"_assets/bash_comp",
	)
}

func _assetsBash_comp() (*asset, error) {
	bytes, err := _assetsBash_compBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_assets/bash_comp", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __assetsFish_comp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x93\x4d\xae\xdb\x20\x1c\xc4\xf7\x39\xc5\x2c\x2a\xf1\xb2\x40\xaf\x49\x97\x51\x4f\x12\xa5\x16\x35\xb8\x46\xc2\x40\x00\x27\x4a\x4f\x5f\xc0\x1f\xb1\x55\x29\x95\xed\x26\x9b\xc8\xcc\xf0\xd3\x30\xfc\xa9\x5a\x5d\x06\x69\x34\x8a\xa2\x92\xbe\x2e\x4a\xd3\x58\x25\x82\x28\xae\x95\x2c\x02\x73\xbf\x44\x38\xed\x00\x2f\x02\xca\x86\xe3\x23\xea\x0d\xd3\x5c\x49\x2d\x40\x8d\x2d\xf7\x49\x95\x15\xce\x49\x6a\x75\xc0\x97\x68\xdb\x83\x8a\x2b\x0e\xb8\x24\x11\x70\x22\xb4\x4e\xe3\x6b\xfa\x12\x9a\x8f\x3b\x92\xf5\x7c\xbc\xe0\x3b\x08\xe5\xe4\x5f\xee\xbf\xf8\xc7\x61\x47\x94\x4b\xa3\x03\x93\xda\x83\xd2\x11\x4b\x15\x68\x03\xea\x3a\xd3\x1c\x3c\xa2\xfb\xbf\x5e\x3b\x9c\x76\x79\x61\xf7\xb2\x96\x4a\x3a\x1f\x8c\x7d\x4b\x31\x8b\x82\xf8\xd2\x49\x1b\x8a\xf0\xb0\x62\x6d\x96\x6f\x93\x12\xe7\x37\x42\x7d\x2d\x94\x1a\xaf\xe5\x3f\xf6\x97\xc1\x85\x66\xcd\xea\xd4\xc7\xf7\xa4\x1e\x52\x82\x96\x88\x41\x41\x2b\xd0\x57\x2f\x03\x94\x81\x7c\x64\xa7\xc2\x8f\x4f\x2e\x6e\x9f\xba\x55\x6a\x4f\x40\x39\x48\xe7\x21\x8b\xb8\xcf\x6a\x90\xe1\x3f\x99\xaf\x91\x6c\xf8\xed\xeb\x0e\xdb\x1d\xf0\xb4\x04\xfa\x1c\x93\x0c\xbd\x3b\x66\xad\x70\x48\xa6\x9e\x99\x1d\x48\x8e\x65\x79\x87\xa7\x90\x0a\x20\x31\x58\xe4\xc5\x5f\x42\x72\xe9\xad\x62\x0f\xe4\xc5\x25\x69\x67\xc8\x9b\x70\x3e\x8e\x11\x99\x21\x87\xc5\xd5\xd4\x7e\x46\xba\xa0\xd6\xc9\x38\x5b\x5d\x03\x7e\x25\xd3\x80\x30\xd2\xcd\x5b\x66\x32\xce\x31\x5e\xff\x4a\x20\x9f\x02\xb9\xc8\x84\xad\x4c\x35\x65\x2a\xe9\x43\x4f\xdc\x70\xee\x66\x8a\x6c\xcc\x6d\x7b\x48\x37\x25\x3a\x91\x1f\xc3\x56\xa6\x9f\x32\x7d\x6d\xee\xf0\x81\x85\x36\x9d\xfb\x4f\x00\x00\x00\xff\xff\x7b\xf0\xde\xf3\x02\x07\x00\x00")

func _assetsFish_compBytes() ([]byte, error) {
	return bindataRead(
		__assetsFish_comp,
		"_assets/fish_comp",
	)
}

func _assetsFish_comp() (*asset, error) {
	bytes, err := _assetsFish_compBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_assets/fish_comp", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __assetsFish_wrapper = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x2b\xcd\x4b\x2e\xc9\xcc\xcf\x53\x28\x4c\xcb\xb4\xe6\x52\x50\x48\xcd\x2b\x03\x31\x15\x54\x12\x8b\xd2\xcb\x40\x02\xc5\xe5\x99\x25\xc9\x19\x0a\x2a\xc5\x25\x89\x25\xa5\xc5\x20\x11\x05\x85\xe4\xc4\xe2\x54\x05\x75\x23\x75\x6b\x85\xe4\x14\x05\x0d\x98\x16\xdd\x1c\x88\xae\x68\xc3\x58\x4d\x64\x75\x5a\x40\x75\x45\xa9\x25\xa5\x45\x79\xc8\xa6\xa4\xe6\xa5\x58\x73\x01\x09\x2e\x40\x00\x00\x00\xff\xff\x09\x92\xc1\xef\x82\x00\x00\x00")

func _assetsFish_wrapperBytes() ([]byte, error) {
	return bindataRead(
		__assetsFish_wrapper,
		"_assets/fish_wrapper",
	)
}

func _assetsFish_wrapper() (*asset, error) {
	bytes, err := _assetsFish_wrapperBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_assets/fish_wrapper", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __assetsSh_wrapper = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4c\xcb\xd4\xd0\x54\xa8\xe6\x52\x50\x48\xcd\x2b\x53\x00\xf2\x14\x94\x54\x1c\x94\x80\xdc\x9c\xfc\xe4\xc4\x1c\x85\xe4\xfc\x94\x54\x5b\x15\x7b\x20\x3f\x39\xb1\x38\x55\x41\x05\xc4\x57\xc8\xcc\x03\xf2\x15\x14\x8c\x34\x15\x92\x53\x80\xca\x35\x60\x3a\x75\x73\x80\x3c\x43\x25\x4d\x25\x6b\x6b\xb0\x02\x2d\x4d\x85\xa2\xd4\x92\xd2\xa2\x3c\x88\x3e\xb0\x68\x6a\x71\x62\x32\x57\x2d\x17\x20\x00\x00\xff\xff\x0c\x75\x71\x6b\x77\x00\x00\x00")

func _assetsSh_wrapperBytes() ([]byte, error) {
	return bindataRead(
		__assetsSh_wrapper,
		"_assets/sh_wrapper",
	)
}

func _assetsSh_wrapper() (*asset, error) {
	bytes, err := _assetsSh_wrapperBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_assets/sh_wrapper", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __assetsZsh_comp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\xdb\x6e\xdb\x30\x0c\x7d\xf7\x57\x10\x6e\x01\x25\x1d\x8c\x62\xdd\x5b\x86\x0e\xfb\x87\x3d\x7a\x83\xa1\x5a\x4c\x6d\x40\x96\x3d\x49\x49\x90\x5d\xfe\x7d\xd4\xc5\xb6\x62\x38\x6b\x80\xbe\x24\x16\x79\x78\xc8\x43\x91\xba\xab\xfb\x6e\x10\xb8\x87\x9f\xfb\x36\xcb\xee\x40\xb6\xc6\x1a\xe0\x52\x82\xe5\xfa\x15\x2d\x28\xde\xa1\xc9\x2a\x72\x57\xc1\x62\x36\x5b\xf8\x9d\x01\xc8\xbe\xe6\x23\x6a\x3a\x16\x3c\x5a\x0c\x99\x28\x06\x0a\x09\x4f\x5f\x1e\x05\x1e\x1f\xd5\x81\x48\xff\xc0\xa9\x69\x25\x82\x46\x2e\x22\xf2\x33\x88\x9e\xc0\x30\x06\x7e\x78\xce\xef\xc3\x67\x4e\x66\xd1\x2b\xa4\xbf\x4a\xa0\xa9\x75\xfb\x82\x50\x58\xc7\x5b\x44\x30\xb0\xf0\xc1\xa6\xb4\x7f\x9d\x0a\x85\xa7\xb1\xfe\x5e\xc3\x9e\x32\x06\x05\x64\x8f\x2a\xa2\x88\xea\xf5\xc0\xb5\x80\xfc\x21\x07\x96\x04\x39\xd1\xcc\xb9\x5d\x68\xe4\xa4\x46\x75\x5c\x09\xd8\x13\x23\xf2\xba\x01\x5e\xdb\xb6\x57\x81\xb8\xee\x04\x95\x28\xd1\xe2\x48\x4c\x44\x87\x0e\x15\x95\x58\x18\xd8\xc1\x77\x2f\x11\x80\x3d\xec\x60\x97\x76\x93\x11\xfb\x44\xa1\xd1\x65\x7e\x8b\xe2\x69\x49\x31\xbb\x3e\x8d\xae\x59\xa9\xf3\x26\x29\xdc\x05\xbf\x9d\x60\x25\x43\xc2\xc1\x85\xb8\xb9\xc6\xcb\x42\xd2\x32\x7d\x6f\x2f\xe4\x77\xfd\xf1\xdd\xe2\x27\xd6\x99\xd6\x34\x28\xe5\xcd\xbc\x1e\x5d\xf9\x09\x58\xe9\xab\x9b\xc2\xc1\x56\xf6\x3c\xe0\x9c\x64\x0e\xb9\xd8\x0d\x5a\x06\xef\x31\xcf\x9b\x17\x6e\x1a\x9a\x43\xfa\xf9\x65\x9a\xed\xc5\x40\x33\x8f\x61\x11\x3a\x73\xce\x89\x96\xa4\xce\x46\x9c\x6e\x73\xe1\xa4\xf9\x30\xa0\x5e\x52\xfa\x60\x0f\x64\x01\x3e\xf1\x2a\x5a\xa8\x2b\xbd\x60\x1f\x49\x25\x0b\xd3\xde\xf1\x56\x81\x68\xcd\xc0\x6d\xdd\xa0\xf6\xb1\x2b\x61\xdf\x16\x5d\x2c\x44\x19\xd6\x60\xdc\xc7\x1f\xbb\x69\xe0\xe7\x1d\x49\x1a\x5b\xf0\x92\xa6\x29\xa2\x17\x60\x72\xa4\xc8\xae\x74\xf3\xb1\x0e\x75\x9e\x14\xab\xcb\xb0\x4a\xeb\xe8\xe0\x4b\xf1\xb2\x74\x7b\x71\xa5\x68\xe7\x4a\xc1\xa6\x34\x4d\x7f\x02\x73\xa0\x07\x41\x9f\x13\xb0\xeb\x6d\x0a\x2c\xfc\x9d\x96\x83\x6e\x95\x85\x70\x27\x0b\xea\x78\xf7\x49\x08\x19\x86\xd2\x35\x5e\xf2\x33\xb8\xc3\xff\xf8\x8f\xa8\x0d\xbd\x41\x13\x3e\x9e\xaf\x87\x6c\x8a\xed\xca\xfb\xe3\x6f\x17\xf2\xfb\xaf\x79\xf6\x2f\x00\x00\xff\xff\xf5\xcf\x1d\x27\x10\x06\x00\x00")

func _assetsZsh_compBytes() ([]byte, error) {
	return bindataRead(
		__assetsZsh_comp,
		"_assets/zsh_comp",
	)
}

func _assetsZsh_comp() (*asset, error) {
	bytes, err := _assetsZsh_compBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_assets/zsh_comp", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"_assets/bash_comp": _assetsBash_comp,
	"_assets/fish_comp": _assetsFish_comp,
	"_assets/fish_wrapper": _assetsFish_wrapper,
	"_assets/sh_wrapper": _assetsSh_wrapper,
	"_assets/zsh_comp": _assetsZsh_comp,
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
	"_assets": &bintree{nil, map[string]*bintree{
		"bash_comp": &bintree{_assetsBash_comp, map[string]*bintree{}},
		"fish_comp": &bintree{_assetsFish_comp, map[string]*bintree{}},
		"fish_wrapper": &bintree{_assetsFish_wrapper, map[string]*bintree{}},
		"sh_wrapper": &bintree{_assetsSh_wrapper, map[string]*bintree{}},
		"zsh_comp": &bintree{_assetsZsh_comp, map[string]*bintree{}},
	}},
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

