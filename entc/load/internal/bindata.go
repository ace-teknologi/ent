// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5f\x6b\xdb\x3e\x14\x7d\xb6\x3e\xc5\xf9\x99\xfe\xa8\xdd\xa5\x4a\xdb\xb7\x0d\xf2\x50\xda\x0c\x32\xb6\x76\x90\xc2\x1e\xba\x52\x14\xfb\x3a\x11\x75\x24\xef\x4a\x29\x0b\x42\xdf\x7d\x48\x4e\xc2\xf6\x64\x4b\xe7\xdc\xf3\x47\x37\x84\xe9\x85\xb8\xb3\xc3\x9e\xf5\x7a\xe3\x71\x73\x75\xfd\xf1\x72\x60\x72\x64\x3c\x3e\xab\x86\x56\xd6\xbe\x61\x61\x1a\x89\xdb\xbe\x47\x26\x39\x24\x9c\xdf\xa9\x95\xe2\x69\xa3\x1d\x9c\xdd\x71\x43\x68\x6c\x4b\xd0\x0e\xbd\x6e\xc8\x38\x6a\xb1\x33\x2d\x31\xfc\x86\x70\x3b\xa8\x66\x43\xb8\x91\x57\x47\x14\x9d\xdd\x99\x56\x68\x93\xf1\xaf\x8b\xbb\xf9\xc3\x72\x8e\x4e\xf7\x84\xc3\x1d\x5b\xeb\xd1\x6a\xa6\xc6\x5b\xde\xc3\x76\xf0\x7f\x99\x79\x26\x92\xe2\x62\x1a\xa3\x10\x21\xa0\xa5\x4e\x1b\x42\xb9\x55\xda\x94\x88\x51\x4c\xa7\xb8\x4b\x79\xd6\x64\x88\x95\xa7\x16\xab\x3d\xce\xc9\xf8\xe6\x74\x75\x2e\x71\xff\x88\x87\xc7\x27\xcc\xef\x17\x4f\x52\x0c\xaa\x79\x53\x6b\x42\xd2\x10\x42\x6f\x07\xcb\x1e\x95\x28\x4a\xeb\x4a\x51\x94\xab\xbd\xa7\xf4\x13\x02\x3c\x6d\x87\x5e\x79\x42\x39\xb2\x5c\xb6\xcc\xd0\xc0\xda\xf8\x0e\xe5\xff\xbf\x4a\xc8\xef\x07\xc5\x18\x45\x9d\x63\x9e\xad\x94\x23\x7c\x9a\x21\x7f\x8f\x78\x9a\x7d\x57\x0c\xd7\x6c\x68\xab\x1c\x66\x78\x7e\x21\xe3\xe5\xc2\x78\xe2\x4e\x35\x14\xb2\x34\x2b\xb3\x26\x9c\xbd\x4e\x70\x66\xd4\x36\xcb\xc8\x07\xb5\x25\x97\xf4\x8b\x22\x84\xcb\x83\x7e\x8c\x32\x1d\x4e\x51\x5c\x88\xe5\x61\x26\xc6\x49\xd6\x22\xd3\xe2\x32\x46\x11\x85\xe8\x76\xa6\xc9\x9d\xab\x1a\x41\x14\x29\x48\xaf\x0d\x39\x3c\xbf\x3c\xbf\xa4\xd2\xa2\xe8\x2c\xe3\x75\x72\xc8\x97\x7c\xc7\x28\xc7\xbc\x41\x14\xc5\x6a\x02\x62\x4e\xd8\x37\xc5\x6e\xa3\xfa\x65\x06\xab\x91\x53\x8b\xa2\xd0\x5d\x66\xfc\x37\x83\xd1\x7d\x9e\x29\x3a\xa5\xfb\x8a\x98\x13\x9c\x2a\x8c\xbe\x33\xa8\x61\x20\xd3\x56\xf9\x38\xc1\xaa\x16\x09\xb5\x4e\x2e\x7d\x6b\x77\x5e\xfe\x60\xed\xa9\xca\xfb\x90\x5f\xac\x36\x47\xe2\x18\xb7\x2a\x7f\x9a\xb2\xae\xeb\x53\xb7\xa3\x4b\xb2\xb7\x9c\x4b\x8e\x5a\xc4\x3c\x6a\x2d\x3d\x6b\xb3\x4e\x1c\x39\x4f\x9c\xaa\xfe\x90\x45\x32\x71\xfe\x5b\xfb\xea\x3a\xcb\xfd\xb3\xfa\xb1\xd9\xb8\xf9\xc3\x8b\xc6\x28\xfe\x04\x00\x00\xff\xff\x95\x06\x0f\xa4\x50\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 848, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\x5f\x6f\xe3\x46\x0e\x7f\xb6\x3e\x05\x37\x40\x03\x7b\xe1\xca\xbd\xa2\x28\xee\xbc\xe7\x03\x8a\x76\x8b\xe6\x7a\x4d\x17\xdd\xdd\xbe\x2c\x16\x5b\x45\xa2\xec\xd9\x48\x23\x77\x66\x9c\x4d\x9a\xe6\xbb\x1f\x48\xce\x48\x23\x59\x76\xf6\x5f\xfc\x12\x89\x43\x0e\xc9\xdf\x70\x48\xce\x28\x8b\x05\x7c\xdf\x6c\x6f\x8c\x5a\x6f\x1c\x7c\xfd\xd5\x3f\xfe\xf5\xe5\xd6\xa0\x45\xed\xe0\xc7\x2c\xc7\x8b\xa6\xb9\x84\x33\x9d\xa7\xf0\x5d\x55\x01\x33\x59\xa0\x71\x73\x85\x45\x9a\x2c\x16\xf0\x62\xa3\x2c\xd8\x66\x67\x72\x84\xbc\x29\x10\x94\x85\x4a\xe5\xa8\x2d\x16\xb0\xd3\x05\x1a\x70\x1b\x84\xef\xb6\x59\xbe\x41\xf8\x3a\xfd\x2a\x8c\x42\xd9\xec\x74\x41\x53\x28\xcd\x2c\xff\x3b\xfb\xfe\xe9\xf9\xf3\xa7\x50\xaa\x0a\x03\xcd\x34\x8d\x83\x42\x19\xcc\x5d\x63\x6e\xa0\x29\xc1\x45\xfa\x9c\x41\x4c\x93\x64\x9b\xe5\x97\xd9\x1a\xa1\x6a\xb2\x22\x49\x54\xbd\x6d\x8c\x83\x69\x32\x39\x41\x9d\x37\x85\xd2\xeb\xc5\x5b\xdb\xe8\x93\x64\x72\x52\xd6\x8e\xfe\x18\x2c\x2b\xcc\xdd\x49\x92\x4c\x4e\xd6\xca\x6d\x76\x17\x69\xde\xd4\x8b\xd2\x3b\xac\x74\xbe\xbb\xc8\x5c\x63\x16\xa8\x99\xff\x3e\x9e\x85\xcd\x37\x58\x67\x0b\x2c\xd6\xf8\x21\xfc\xa5\xc2\xaa\xf8\x10\x01\xa5\x0b\xbc\x3e\x49\x66\x09\xc1\xf6\x9c\x69\x60\xd0\x2f\x98\x85\x4c\x03\x6a\x97\xfa\x01\xb7\xc9\x1c\xbc\xcb\x2c\xe3\x82\x05\x94\xa6\xa9\x21\x83\xbc\xa9\xb7\x95\xa2\xc5\xb1\x68\xc0\x63\x97\x26\xee\x66\x8b\x61\x4a\xeb\xcc\x2e\x77\x70\x9b\x4c\xce\xb3\x1a\x01\x80\x28\x4a\xaf\x81\x7f\x7f\x10\x9a\xcb\x13\x9d\xd5\x38\x6f\x6a\xe5\xb0\xde\xba\x9b\x93\x3f\x92\xc9\xf7\x8d\x2e\xd5\x1a\xd8\x86\xf0\xec\x99\x73\x7e\xed\xb3\x3f\x2d\xd6\x68\x01\xe0\xd5\xeb\xc7\xf4\x18\xcf\x4d\x40\xda\x3e\xf7\x8f\x84\x95\x65\x6e\x7e\x8c\xb8\x19\xc6\x01\xfb\x19\x21\x85\x96\xd8\xf9\x31\x62\x57\x32\xd4\xe7\xff\xa9\x69\x2e\xbd\x31\xcf\x1a\xab\x9c\x6a\x74\xe0\xdf\xd0\x50\x9f\xfb\x59\x53\xa9\xfc\x06\xe0\xa2\x69\x2a\x80\x1e\x2c\x5b\x1e\xea\xb1\xdf\xf1\x72\xb5\xd3\x16\x68\x73\xa3\x2e\xd0\x42\x06\x6c\x3a\x6c\xc3\x90\x8f\x7a\x59\x6d\xbf\x26\xad\x5c\xb7\x2a\xad\x47\x00\x4a\x3b\x80\xc5\x02\x04\x13\x76\x2d\xcc\x22\x73\x57\xca\xba\x34\x99\xfc\xa2\xae\xb1\x38\xd3\x24\xc2\x46\x2f\x16\x70\xa6\x0b\x95\x67\x0e\x2d\xa8\x32\x12\xa0\x88\xa9\x89\xfb\x4b\xa5\x45\x50\xe9\x33\x3f\xaf\xe8\x62\x52\x5f\x57\xcd\x24\xd1\x25\xee\x8a\x41\xfb\xc1\x29\xf4\x8f\x88\x4d\x11\xdc\x0f\x4d\xf9\xc5\x01\x7a\x4f\x98\x9e\xe9\xb2\xe9\xd8\x1e\xb3\xd7\xe9\x8b\x9b\x2d\xfa\x01\x2f\x48\x4a\xfb\x82\x2f\xb2\x58\xc1\x41\x8d\x2e\x1b\x04\xfa\x73\xf5\x57\x64\xe9\x63\xa5\xdd\xb7\xdf\x8c\xc8\x59\xf5\xd7\x40\xe1\x53\xbd\xab\x6d\xcb\xf6\xea\xf5\x50\x65\xd8\x2d\xc4\xd6\x97\x7c\xa9\xd5\x9f\xbb\x56\x69\x1c\xa6\x3d\xc9\x1d\xb3\xf5\x45\xcf\x55\x55\x65\x17\x15\xde\x23\xaa\x3d\x5b\x5f\xf8\xd7\x2d\x85\x6a\x56\xdd\x23\xdc\x78\xb6\xbe\xf0\x0f\x58\x66\xbb\xca\xdd\x67\x74\x21\x6c\xa3\xb2\xbf\x67\x15\xb9\xad\xb4\x43\x43\x99\xf4\xf6\x6e\x54\xf6\xcd\x15\xf1\x0d\x20\xdb\x16\x99\xc3\x60\xc3\x61\xc8\x98\xed\xcd\xa8\x11\x67\x75\xbd\x73\x2d\x76\x07\xa7\x50\x81\xad\x2f\xfd\x7b\x56\xa9\x82\x32\x3e\x2f\x39\x6f\xb6\x31\xe9\xab\x96\x6d\x10\x65\xae\x31\xd9\x1a\x7f\xc6\x1b\x38\x16\x9d\x56\xd8\xde\x5c\xe2\xcd\x30\xa7\xf9\x3c\xc3\xbf\xc7\xfd\xd7\x38\xbf\x09\x7d\xa0\x1c\x35\x91\xaf\xee\xf1\xdc\x06\xb6\x81\x34\xe7\x3b\xda\x82\xc4\x5b\x67\xdb\x57\x62\x7e\x08\xf8\x20\xcd\x6c\x6f\xf6\x36\xa6\x24\x1c\xae\x21\xfb\xf9\x86\xc9\x1f\x91\x6e\x58\x6e\x24\xdb\xf4\x4d\xda\xcf\x2e\xc1\x8b\x01\xe3\x91\x6c\x32\x60\x1c\x66\x8f\xdf\xb0\x14\xe5\x7d\x3e\x83\xe5\x9b\x7d\xed\xbf\x61\xe9\xd7\x4f\x4a\x6a\xc7\x7c\x20\x3f\xf8\xa5\x3a\x92\x0f\xce\xf4\x15\x1a\x8b\x43\x56\x25\xe4\xa1\xfa\x3f\x77\xca\x60\x31\xe0\x35\x9e\x3c\xb2\x6a\x52\x59\xf6\x97\x4d\xe8\x1f\xb1\x6e\x22\xd8\x2d\x5c\x94\x09\xdb\xb0\x3c\xe2\x6d\x68\x4a\xe2\x7c\x7b\x7f\x53\x32\xc2\x3d\xd6\x94\x44\x5b\xb4\xdd\x9f\xf7\x6c\x4b\x41\xe9\x1c\xdf\xf1\x7a\xe6\x06\xb9\x60\x67\x3a\x20\x42\x46\x09\x2c\xfc\x24\xbd\xc5\xd6\x35\x26\x4d\xca\x9d\xce\x83\xe4\x14\x0b\x78\x4c\x1c\xe9\x0f\x2d\xc7\xcc\x07\xc9\x6d\x32\xd1\x08\xcb\x15\x9c\xd2\xeb\x6d\x32\xa1\xd0\x5c\x0a\x06\x58\xa4\x2f\xb2\xf5\x9c\x68\x37\x5b\x5c\xb6\x34\x8a\xe6\x64\xc2\xbb\xa2\x25\xd2\x0b\x11\x05\xf1\xa5\x10\xe5\x85\xc8\x3e\x8e\x96\x4c\xf6\x2f\x44\x0f\x31\xb3\x24\x7a\x78\x91\x81\xd2\xcf\xcf\x03\xa5\x9f\xff\x2e\x99\xa8\x12\x0c\x96\x64\xb2\x8c\x3c\xe1\xd7\x47\x2b\xd0\xaa\x22\x77\x26\x1a\x89\x0c\xab\xd6\x7d\x83\xe5\x8c\x45\x0d\xba\x9d\xd1\xa0\xb1\x43\x56\x1a\x8b\x7d\x68\xa5\x1d\x62\x6c\xe5\x71\x0c\x5c\x16\x9e\x96\x45\xe8\x23\x62\x78\xa7\xd2\xa9\xce\x01\x8d\xa1\xf7\x5b\xb6\x1c\x8d\x21\xcb\xcb\x22\x7d\x6a\xcc\x74\xf6\x84\x09\x91\xed\xc1\x42\x55\xcd\xa1\xac\x1d\x71\x35\xa6\x9c\x4a\x40\xc1\x17\x7f\x2e\xe1\x8b\xab\x93\x39\xc9\x33\x20\x24\x2e\xae\x59\x46\xe4\x94\x75\xde\xf6\x16\x87\x7f\x65\xb7\x42\xd4\xe8\xf4\x47\x88\x32\xef\xad\x7c\x18\xf1\xcb\xcf\xad\xc8\x32\x1e\x60\x4a\x7f\xbd\xc3\x50\xb7\xe8\xa1\x99\x58\x76\x36\x84\xbe\x21\x99\xb4\xdd\x42\x37\x1a\x28\x34\xea\x4b\xf1\xb2\x9b\x37\x14\x67\x41\x8b\x75\xc7\x45\x7b\xc9\xba\x7b\x65\xbc\xe3\x6c\x6b\xf3\xb2\xf5\xb9\x2d\xc3\xc9\x24\xda\x9a\x4b\x3f\xdc\x51\x68\xbc\x2b\xce\x3c\x5e\xa1\x9e\x96\x45\xda\x51\x67\x3c\x49\x28\x6f\xad\x8e\x96\xc2\xc3\x6d\x99\x6b\x75\xb4\x94\x36\xb0\x6d\xc9\x8b\x01\xab\xfb\x23\xa2\x56\xd6\x52\x12\xe1\xbc\xa7\x48\xa8\x6c\x0c\x84\x38\x39\x99\xd3\x5c\xb4\xe4\xb3\x76\x6e\x6a\x44\x97\x2b\xe0\x0e\x94\xec\xa7\xce\x74\xf6\x44\xe8\x8f\x56\xf0\x15\xab\xb3\x25\xd3\x61\x05\xa7\x34\xc0\xc2\x94\xa9\xe5\x90\xe0\x1b\x1f\xe0\x0e\x0a\xf2\x4c\xc3\x05\x02\x1f\xb4\xb1\x00\xd7\x30\xcf\x1a\x35\x9a\x8c\xf7\x09\x49\xfe\xd8\x18\xc0\xeb\xac\xde\x56\x38\x07\xdd\x38\x3a\xf7\xec\x74\xce\xdd\x45\xa5\x2e\x11\x9c\xaa\x31\x3d\x6f\xde\xa5\x6c\xe5\x9b\x79\xd8\x23\x94\x1a\xd3\x5f\x32\x63\x37\x59\x35\xed\xd6\xdf\xef\x99\x08\x21\x5b\xa6\xbd\x06\x70\x15\x45\x4b\xbc\xed\x6d\x39\x27\x99\x6e\xef\x4b\xb5\xd8\xdf\xfb\x72\xb8\xe1\xbd\x2f\x8f\x63\x7b\x9f\x85\xa7\xaa\xb8\xa6\x9e\xbe\xc0\xeb\x7e\x6e\x95\xa9\x6f\x5b\xdd\xa7\x4c\x20\x6b\xb9\xc6\xf8\xb0\x56\xc5\x35\xb7\x28\xbc\x93\xa4\x9c\x2c\xdb\x01\x79\x1f\xee\x31\x1a\xe9\x76\x58\x1c\xb8\x34\xd2\x0b\xdb\x3b\xef\xa9\xc7\xd0\x1f\xef\x65\xb5\x78\xa5\xa2\xeb\x82\xb6\x67\xa6\xa7\x06\x32\xf8\xef\xf3\x5f\xcf\x49\x98\x8b\xb0\x5f\xe8\x02\x65\xa1\x99\x85\x26\xf0\xc2\xcd\xc5\x5b\xcc\x9d\xff\xe3\x11\xea\x29\x9d\xda\xa0\x9b\x6a\xbb\xd7\x34\x83\xe9\x05\xbc\x7a\x7d\x71\xe3\x24\x8f\x45\x89\xd2\x72\x2e\x13\x59\xc2\x4c\xee\x13\x96\xe1\x68\x2c\xaf\xd3\x59\x5c\x84\x94\x96\x8b\xa2\xa9\xbf\xde\xe1\x2a\xf5\x6b\xe9\x35\xcf\x66\xbc\x17\x58\xe4\x2e\x4e\xc4\x36\xa5\x35\xe7\x33\x6d\x60\x7d\xef\x9c\xec\x9d\x6a\x93\xb2\x1d\xe6\xe4\xa1\x1a\x59\xd1\xcf\xaf\x47\x7a\x93\x56\x57\x56\x22\x07\x55\x50\xd4\x1a\xf2\x39\x74\x51\x92\xa1\x3d\x4a\x9a\x4c\xa6\xd7\xc8\xad\x87\x95\x8d\x28\xc1\x0c\x2b\xc8\xb6\x5b\xd4\xc5\xd4\x13\xe6\x5d\x23\x12\xed\x92\xe9\x6c\xe6\x61\xf2\x57\x32\xb1\x03\xfe\x06\xe7\x21\x5d\xa0\xad\xdb\x3a\xe1\x6d\xf0\x6e\x84\xfb\xa3\xc8\x91\xb3\x60\x64\xbc\xf5\x47\xbd\x19\x2c\x3a\xdf\x2d\x3d\x7c\x6c\xc9\xa5\xd4\xe7\xd7\xe3\x05\x7b\xc9\xd8\xce\x7c\x66\x79\xa9\xeb\x5e\x6e\x91\x04\x61\xa5\x0c\xa8\x2b\xd4\x70\xb1\x2b\x4b\x34\xc0\x29\xc5\x67\xd7\x70\xbf\xc5\x69\x62\x30\xc3\xf4\x62\x57\xfa\x9c\x40\x1d\x94\x10\xe7\x87\x32\x43\x0f\x06\xb6\xb0\x9d\x8e\x26\x9a\x83\x3d\x0e\x04\x1a\x13\x07\x44\xd9\x85\x83\xf5\xd9\x97\x45\xa2\xb6\x2d\xf5\x05\xd0\x8e\xb4\x6e\xfb\x53\xd3\xdc\x51\xf9\x89\xab\x4f\x9b\x75\xf8\xc9\xfa\x2b\x34\xd7\x78\x74\x7c\x53\x1f\xa7\x4b\x0f\xd8\xd4\x82\x87\x65\x06\xc3\xd4\x35\xcc\xaf\x0c\x1b\xd9\xc6\xb3\xf7\xf6\x57\x2f\xe3\x1d\xd9\x5d\x31\x44\x6a\x0e\x75\xb4\x65\xc4\x64\x6e\xb8\xe9\x84\xca\x9d\xc5\x78\x0e\xae\xaf\xdb\xfc\x9b\x4c\x26\xfe\x6c\x14\x5b\xe3\x13\x63\x7d\x3d\xeb\xe0\x1e\x41\xb6\xdf\xfe\x90\xf6\x36\x6e\x75\x14\xb5\x64\x2f\x1b\xfc\xb6\xb7\xa6\x65\xb7\xa2\x13\x6a\x05\xbc\xfe\xae\x8d\xef\xef\x66\x62\x1b\x31\xe5\x43\x6d\x61\x63\xa8\x45\x69\xaf\x54\x56\x70\x1a\x9e\x65\x46\x4e\x27\xbe\x23\x78\x3b\x67\x92\xbf\xb0\x65\xa2\x33\x52\xeb\x27\xd1\x6d\xec\x12\xd4\xbc\x9b\x3c\x04\x6b\x94\xae\x7c\xf3\x00\xb6\x0c\x80\x1c\x2a\x12\x9f\x1b\xf4\x43\xc5\xe1\xa3\xaa\x03\xcf\x7a\xac\x3e\x3c\x80\xf5\x07\xeb\xc2\xa7\x14\x06\x56\x20\xdf\x12\x62\x37\xa4\x38\x7c\xf6\xb8\xef\xec\x67\x95\xc1\x7a\xf9\xcc\x11\xd9\xfe\x93\x18\xf4\x19\xe3\x71\x36\xcc\x7a\xfd\x94\xe7\x03\x55\x72\x9e\x9c\x55\x3e\x22\xe7\xf5\xfa\xa8\x83\x49\xef\x70\x9e\xf9\xe0\xb4\x37\x9e\x45\xde\x2f\x89\x1c\x5e\xd6\xb6\x46\x1c\x4c\x0f\x01\x5b\xe6\xb9\x6f\x97\xef\x61\x3e\x8a\x5d\xdc\x8e\x1c\x84\xee\x50\xa0\x7e\x20\x70\x63\x61\xf8\xbe\x51\xd8\x06\xa1\x04\x56\x1b\x80\x65\x56\xc9\xcd\xd1\xdd\x7b\xbb\xdc\x6b\x8d\x0e\xfa\xec\x3f\xdd\xc5\x4e\xf7\x7b\xaa\xf7\xf0\xda\xa6\xfe\xdb\xe0\x0a\x64\x3a\xcf\x3b\x6e\x66\x09\x72\x45\x34\x83\xae\xab\xe8\xec\x51\x25\x3c\x6a\x0f\xb6\xf0\xf7\xdf\xf4\x76\xa6\xcb\x26\x3d\xdf\xd5\x68\x54\x3e\x9d\x0d\xfa\x19\xb6\x40\xcf\xa1\xb9\x94\x56\x25\x3e\x13\xa7\xd3\xb2\x6a\x32\xf7\xed\x37\xe2\xc5\xa3\xe6\x32\x16\x8e\xf3\xcb\x4e\xe3\xf5\x16\x73\x87\xc5\xe0\xb0\xcf\xf7\x0c\xed\x15\xc3\x52\xee\x18\xe2\x2b\x06\xfb\x4e\xb9\x7c\x03\x4e\xb4\xb3\xa9\x54\xff\x9f\x90\xa6\x3c\xb3\x08\x0e\xfe\xb3\x82\xf8\x83\x9b\xfb\x27\x9c\x9e\x82\x83\x7f\x0f\xc8\xdf\x7e\xb3\xa4\x4c\x36\x3c\xd5\xcb\xc5\x85\x9e\x8d\x4f\xf7\x52\x8d\xcf\xf7\x52\x1d\x9c\x70\xd7\xcd\x38\x96\xb0\xba\x8c\x01\xef\x4c\xb6\xb5\xf1\xd7\x59\x4f\xcf\x74\x21\x7d\x50\x20\xd4\xe8\x36\x4d\x01\xef\x94\xdb\x80\xc1\xbc\xb9\x92\xe6\x17\xb5\xdd\x19\x04\xdd\xc0\x36\xd3\x2a\xb7\xa0\x34\xf8\x4e\x55\xe9\xb5\x4f\x73\x51\x86\x2a\x8b\xe8\x2b\x16\x78\xe2\x0c\x5e\xbd\xee\x3e\xa2\xde\xcd\x60\xea\x93\x51\x44\x1e\x9e\xa4\x0b\xa4\xf6\x9b\xa6\xf7\xf1\xa2\x4a\xb8\xe2\x7d\x29\xc6\x51\x1f\x7b\xd5\x4b\x4e\x7c\xb9\xd2\x0b\x89\x2f\x5e\x04\xef\xc4\xf8\xf6\x0e\x72\x0e\x57\xdc\xe2\x94\x21\x31\x71\x14\x72\xfe\xa7\x4e\x2f\x44\x57\x91\x06\x07\xe6\x03\x74\xa5\x21\xd8\x03\x57\xc8\x9f\x0a\x65\x7c\x06\x8e\xd1\x14\x7a\x00\x93\x2f\xc1\x09\x4b\xe9\x54\x3a\xe2\x43\x20\xd9\xf3\xaf\x07\xa6\x00\x89\xbe\x41\x1a\xc5\x31\x16\xde\x87\x32\x74\x26\x7b\x60\x86\x81\x4f\x85\xb3\x7f\x22\x8f\x01\x0d\x23\x01\x52\xb9\xfb\x22\x4c\x55\xfb\x7f\x18\x2d\xfd\x01\x61\x0d\x9e\x8e\x00\xab\xda\xbe\xed\x18\xb4\xad\x23\x43\x70\xe5\xa4\xb6\x07\xad\x90\x3f\x15\xd8\x63\x27\xb8\xa9\xb4\x7b\x82\xdf\x2f\xdd\x29\xee\x41\xf0\x13\x77\x46\xd0\x13\x23\x8e\x63\x27\x5e\xec\x21\x27\xc5\x7e\x0f\x39\x21\x7f\x2a\x72\xbd\x5e\x26\x0a\x48\xa1\x87\x70\xa4\x37\x8e\x46\x69\x42\x3a\xe2\x03\x42\x29\xfe\x8d\x40\xb9\xf1\xcd\xcf\x31\x28\xbd\xf9\x43\x28\x7d\x6b\xb1\x87\xa5\xa7\x7f\x2a\x98\x47\xbb\xa4\xa9\x6f\x67\x88\xfc\x2c\x6a\x94\x1e\x04\x3c\xef\xd0\x08\x7a\xdb\xd0\x5d\x1d\x83\xcf\x3b\xd2\xe1\xc7\x2e\xb6\x77\x13\x0e\xe2\xdb\x89\x59\xef\x8d\x8f\x0d\x8d\x01\x97\xfe\xac\x74\x31\x9d\xc1\x6a\xd5\x8e\x3f\x73\xdc\x96\x4d\x1c\xac\xc0\xa5\x4f\x2b\xac\xa7\xbd\xbe\xc1\x25\x77\xc9\xff\x03\x00\x00\xff\xff\xdd\xdc\x23\x9c\x4f\x29\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 10575, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
