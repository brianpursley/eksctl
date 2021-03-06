// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/schema.json (42.17kB)

package v1alpha5

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x7d\xfb\x73\xdb\x38\x92\xff\xef\xf9\x2b\x50\xca\xd6\xee\xa4\x4a\xb4\x76\xa6\xea\xfb\xbd\xab\xec\x9c\xeb\x14\x3b\x0f\x5f\x12\xdb\x67\x65\x67\xeb\xd6\x9e\xaa\x40\x64\x8b\xc2\x9a\x04\x38\x00\x68\x47\x99\xc9\xff\x7e\x85\x07\x29\x3e\xc0\xa7\xe8\xc4\xb9\xdf\x6c\x8a\x68\x34\xba\x1b\x1f\x34\x1a\x8d\xe6\xef\x4f\x10\x9a\xfd\x89\xc3\x66\xf6\x1c\xcd\x9e\x2e\x02\xd8\x10\x4a\x24\x61\x54\x2c\x4e\xa2\x54\x48\xe0\x27\x8c\x6e\x48\x38\x9b\xab\x17\xe5\x2e\x01\xf5\x22\x5b\xff\x0b\x7c\x69\x9e\xfd\x49\xf8\x5b\x88\xb1\x7a\xbc\x95\x32\x79\xbe\x58\xfc\x4b\x30\xea\x99\xa7\x1e\xe3\xe1\x22\xe0\x78\x23\xbd\xbf\xfe\xdb\xc2\x3c\x7b\x6a\xda\x15\xba\x9a\x3d\x47\x8a\x0f\x84\x66\x59\x9f\x11\x4b\x83\x7f\x60\xe9\x6f\xf3\x9f\x10\x9a\x25\x9c\x25\xc0\x25\x01\x51\x78\x8a\xd0\xcc\x37\x8d\xde\xb1\x30\x24\x34\x2c\xfd\xd6\x39\xb8\xbc\xa3\xac\x75\xde\xf4\x8b\xfd\xeb\xcb\x7c\xdf\x3f\x6c\x80\x73\x08\x2e\x78\x00\x7c\xf6\x1c\x5d\x37\xf2\x60\x7f\xf8\x35\x6f\x8b\x83\x40\xf7\x8c\xa3\xcb\xe2\x28\x36\x38\x12\x90\xbf\x14\x80\xf0\x39\x49\xd4\x7b\x8a\x63\x9f\x51\x89\x09\x15\xc8\xd7\x2a\x40\x09\xe6\x38\x06\x09\x5c\x20\x0e\x11\x96\x10\x20\xc9\x50\x41\x56\x39\xa1\x4f\x1e\xa1\x12\xa2\x88\xfc\xcb\xdb\xca\x38\xf2\x0e\x25\xfc\xa4\x20\x88\xba\x8e\xea\x82\x6f\x54\x15\x50\xbc\x8e\xe0\xc3\x2e\xa9\xfc\x80\xd0\x8c\x48\x88\xab\x0f\x0b\x26\x27\x24\x57\x7d\xcc\xcb\xbf\x06\xb0\xc1\x69\x24\xd5\x0b\xb3\xc2\x2f\x5f\x8a\xaf\xe5\x24\x30\xe7\x78\x57\xa2\x50\x95\xb8\x66\x0c\xb1\x0d\x8a\xcc\x90\x94\x18\x0c\xcb\xe8\x07\x01\x80\xae\xf7\x83\x46\x01\xf3\xc5\xaf\x3f\x2c\x52\x81\x43\x58\xf8\xea\xf9\xbd\x7a\xee\x59\x4b\xf0\x2c\x89\xc5\x53\xfb\xc0\xc8\xda\x83\x4f\x38\x4e\x22\x10\xcf\x9e\x1d\xa1\x0f\x0c\xf9\x5b\xc6\x04\xa0\x0d\x67\xf1\x73\xf4\x11\x27\xe4\xe3\x1c\x7d\xc4\x69\x40\xa4\xf9\x43\x6e\x81\x4a\xe2\x63\xc9\xb8\x7a\xa0\x34\xc7\x59\x14\x01\x7f\x8f\x29\x0e\x41\x3f\x54\xd3\x2a\x48\x23\xe0\x1f\xcb\x83\xeb\xb0\x82\xae\xc1\xfe\x8c\xd1\x96\xc3\xe6\x3f\x6e\x66\xa3\x07\x79\x33\x3b\xae\x48\xec\xe7\x05\x3e\x76\x8c\xfc\x67\x9f\x05\x70\x8c\x13\xf2\xf3\x42\xff\x35\xcf\x9e\x28\x49\xd4\x9e\x15\x84\x52\xf9\xad\x26\x9f\xca\xef\xb9\xa8\xec\xf3\xb1\xd3\xbd\x68\xc7\x53\xce\x75\xe0\xed\x73\xd2\x8a\x39\x53\xd9\xd0\x19\x3f\x94\xbc\x73\xde\x9b\xf5\xa0\x30\xd9\x39\xfc\x96\x12\x0e\x41\x59\x44\x31\x48\x1c\x60\x89\xeb\xf2\x69\x02\x07\x9c\x90\x5f\x80\x0b\xc3\xf2\xef\xae\x19\xec\x00\x81\x12\x04\x94\x7e\xc8\x8c\xb0\xc4\x96\xf9\xe5\x56\xf8\x32\x3a\x22\x6c\x71\xf7\x23\x8e\x92\x2d\xfe\x7f\x45\xf0\xf8\xf5\x89\x03\x46\x66\xf8\x0e\x93\x08\xaf\x49\x44\xe4\xee\x9f\x8c\x7e\x3b\xfc\x72\x72\xe7\xbb\x16\x4c\x34\x60\xf9\x6b\x85\xc5\x55\x05\xfa\x44\x9a\x24\x8c\xcb\x3e\xe8\xf7\x6c\x10\x24\xad\x06\xc2\x4e\x19\x5f\x2c\x5b\x0a\x62\xdc\x52\xda\x60\x1e\x62\x09\x97\x9c\x6d\x48\xd4\x5b\x83\x6e\x09\xbe\x2a\xd1\x3a\x48\x79\x21\x91\xfd\xb4\xf6\x9a\x48\x37\x05\x82\xe3\x41\x7a\x3f\x5b\xbe\x77\x13\xba\x25\x34\x78\xe0\xc9\x57\x86\x91\xce\x79\x17\x6b\x1c\x0f\xce\x59\x00\xaf\x39\x4b\x93\xc3\xb4\xf6\xbe\x42\x6d\x0a\xa7\x41\xcf\x0e\x45\x31\xd4\xfc\x21\x6d\xb5\xf9\xdc\xd0\xfc\x13\x1a\x7a\x34\x7f\xe3\x19\xc2\x34\x40\xd7\x76\x64\x68\xff\x43\xde\x08\x6e\x85\x67\x7f\xd6\xed\xc4\x14\xf3\xc8\xc1\xc9\xcd\xec\xb8\xca\xb8\x9a\x3d\x9a\xbf\x5a\xfb\x3a\x53\x37\xb3\xe3\xfa\x20\x9a\xa7\x5f\xbe\x22\x0c\x31\xd5\xf7\xb0\x5f\x42\xca\xe4\xe8\x34\x26\x31\xa9\x2d\xbc\x62\x1c\x11\xba\x61\x3c\xc6\xea\x91\x16\x64\x36\x15\x90\xf6\x20\x1d\xda\x76\x99\xc8\x20\x75\x77\xf6\xda\xd3\x16\xfa\x28\x51\x80\xcf\x41\x8a\x97\xd4\xe7\xbb\x8c\x81\x1e\xda\x5c\xd5\x9a\xb9\xa9\x4b\x2c\xd3\x9a\x3e\x5b\x0d\x64\x65\x9a\x38\xc9\xdd\x25\xfe\x20\x5a\xbf\x5c\x9e\x8c\xf5\x08\x35\x72\xce\x9d\xce\x8c\x6b\x0a\x54\xc0\xbb\xc2\xb3\xdb\xca\x5b\x51\xb1\x65\x99\x6b\xf5\x62\xdc\x4e\x44\xab\xc2\xeb\xfa\xaa\x2c\x66\x93\x78\xc4\x18\x09\xa2\x2c\xd8\x7a\xac\x73\xe5\x9f\xae\x01\x71\x48\x22\xec\x43\x80\xee\x89\xdc\x22\xab\x37\xb4\xbc\x3c\xeb\xed\x0b\x0f\x26\xec\xf2\x82\x5f\xd2\x20\x61\x84\x4a\xd1\x67\xd7\x9b\x70\x72\x87\x25\x2c\x7d\x1f\x44\xcd\xb8\x33\x70\x59\x33\x16\x01\x6e\x98\x17\x49\xba\x8e\x88\x3f\x94\xc0\x20\x03\x2e\x33\xd9\xd4\xf7\x24\xaa\xdd\xb2\x28\x10\xf9\x66\x03\x27\x04\x09\xe0\x77\xc0\x11\x58\xa9\x22\xac\x7b\x2b\x82\x5a\x6f\xf5\x8e\x22\xee\x52\xb1\x72\x92\x7a\x28\x37\x9b\x6d\x2c\x78\xf9\x09\xfc\x54\x91\xbb\x62\x11\x2c\xaf\xce\x3b\x1c\xa9\x56\x17\xb5\x42\xed\x12\x78\x4c\x84\x42\x13\xf1\x82\xa5\x34\xc0\x7c\x37\x86\xba\x92\x04\xf1\x95\x8e\x59\x5a\xb6\x5d\x34\x78\xdd\xdc\x4b\x69\x55\xa2\x7a\x90\x2b\x6c\x19\x3c\x40\x80\x05\x0a\x13\x09\x4d\xe1\xc1\xc5\xd9\xe9\xc9\x03\xcd\xbb\xca\x90\xfb\x0f\xa5\xdb\x6a\x2a\xf4\x06\xd8\x96\x6b\xf8\x2d\x76\x34\x21\x2a\xe0\x28\x42\x67\xcb\xf7\x08\x4b\xc9\xc9\x3a\x95\x26\x44\x85\xb3\x09\x3d\x10\x06\xba\xa8\x35\xcc\xfb\x8a\x45\xf7\x40\x01\x4c\x29\x93\xb8\x1c\xcd\x6e\x97\xc5\xc3\xc5\x09\x0a\xd1\x79\x17\x81\xdf\xbf\xb8\xed\x1c\x4b\x89\xfd\xed\x25\x8b\x88\x5f\x9b\x27\x6e\x08\x38\xa3\x11\xa1\x70\xca\xfc\x34\x06\x5a\xeb\xd0\xa5\x0e\x94\x68\xf2\x28\xb0\x6d\xd4\xda\x6b\xfa\x55\x7f\xc9\x2d\x11\xc8\xda\x96\x42\x69\x2d\xfc\x21\x8e\xf0\xf8\x5e\x3a\x25\xb2\xbc\x3a\x7f\x5c\x21\x9f\x08\xaf\x21\xfa\x6e\x8d\x8d\xe2\x18\xc6\x06\x1a\x1a\x09\x8a\x04\xfb\xd3\x52\x4d\x06\x2f\x21\xc3\xe8\x8f\xd8\xef\xd4\xe0\xa9\x6d\x03\x24\x71\xf8\x7d\x99\xc8\xa0\xa5\x53\x1b\x91\xd3\x06\xea\xf3\x64\xee\xc6\xea\xb6\xd9\xde\x84\x8d\x1d\xf6\xd1\xba\x3d\xd2\x0a\x99\x72\xb9\xa4\x88\xe0\xd8\xa2\x99\x05\x33\x94\xed\x32\x75\x20\xc0\xec\x72\x52\x3e\xc6\x8b\x1e\x4a\xbd\xd7\x72\xba\xaa\xda\x7c\xe3\xa2\xca\x07\xba\x81\x83\x8c\x27\x23\x3e\xa1\x32\x8c\xba\x95\x87\x51\x93\xda\x40\xc1\xb7\x51\x72\x09\x59\x47\xc9\x3a\x8f\x64\xaa\xf3\x85\x43\x58\x88\xc4\x74\x1f\xd0\x4c\x0e\xda\x96\x81\x29\x49\x7e\x77\x88\x57\x0c\xfd\x1c\x76\xfa\x35\x11\x8a\x5a\xa5\x38\xf8\x7a\x10\x1c\xcb\x93\x11\x42\xa0\xc0\x71\x94\x6f\xdd\xc7\x6c\xfe\x7b\x11\x73\x4d\xa0\xf3\xe5\x87\x3e\x88\xa4\xf6\x4f\xf7\xb8\xff\x36\x72\x90\x22\x32\xe2\x13\x22\xd2\xf9\xf2\x03\xb2\x64\xcb\x50\x8d\x58\x52\x5e\x00\xfb\xe1\x52\x37\x3d\x97\x70\xfb\x23\x3e\xe6\xd3\xa2\x81\xaf\x3a\xd9\x10\x1f\x4b\x58\xa6\x72\xcb\x38\x91\xbb\x53\xc7\x11\x44\x3f\x4f\xfe\x10\x77\x3d\x8b\x3e\x4d\xed\x3e\xfa\xb7\xe7\x93\xa0\xf2\x20\x4b\xcd\x07\xd3\x47\xd4\xf3\xb2\x7a\x9d\xec\x4f\x66\xf2\x88\x03\x0e\x3c\x46\xa3\xdd\x24\x11\x84\x1e\xe4\x9c\x06\x9f\xae\x29\x0c\x8a\x0d\x8f\x5b\xb2\x1a\x4e\xb4\x40\xde\x33\x7e\xfb\x60\xcb\x94\x09\x08\x3f\x7a\x8e\x07\x59\x74\xa6\x86\xfa\x30\x27\x44\x63\xdb\x89\x76\x9c\x0d\x75\x24\xac\xa5\x0c\x83\xe1\x16\x42\x2e\x73\xfc\xe5\xf2\xa4\x17\xf8\xa6\x92\x2d\xa3\x88\xa9\x29\x7c\x76\x79\xf7\xff\x47\x9d\x54\xf8\x24\xe0\xfd\xf6\xb3\x21\x91\xdb\x74\x7d\xe4\xb3\xf8\x8f\x7b\xc0\x77\xa0\x2c\x40\xfc\x61\x72\x73\xfe\x48\x6e\xc3\x3f\x52\x49\x22\xf1\x07\x49\x28\xc8\xa3\xb3\xcb\x73\x68\x88\xd2\xf8\xcd\x27\x32\x2d\xbd\xd7\xce\x71\xdc\xa8\xfd\x49\x72\x7c\x72\x76\x7a\x75\x58\xac\xfc\x90\xa1\x8e\x3f\x8c\xde\x30\x8e\xf6\xc6\x8a\xd4\x30\x10\x16\x82\xf9\xc4\x6c\x7e\xe7\x08\x8e\xc2\x23\x24\x19\x4a\x05\x98\x63\x2f\x01\x09\xe6\xca\xb2\xf4\xcb\x8a\x40\x66\x6a\xd6\xbe\x90\xa2\x49\x77\x08\x07\xde\x96\xd5\xcd\xb7\x8f\x09\x7f\x45\xb6\x9c\x3a\x25\xa3\x33\x5c\x9c\xe4\x28\xee\x99\xbd\x53\xf0\x34\x5b\x60\xd5\x9c\xb3\x0d\x31\xb9\xaf\x13\x6e\x14\xe0\xa7\x6a\x49\x37\xa9\x12\x93\x3a\x31\x5b\xcc\xcd\x51\xf6\x6a\x7c\x1f\x35\x0b\x4b\x38\x78\x5a\xfa\x10\x20\xd3\x83\xce\x51\x41\xab\xd7\x83\x8d\xb5\x2f\xa9\xee\x91\xd6\xdc\x82\x6e\x73\x59\xb9\x66\x58\x85\xc9\x6c\x12\x60\x0e\x08\x88\xdc\x02\xcf\x56\x85\xc2\x4c\x51\x23\xa9\x4f\xa8\x7d\xca\xc7\x1c\xc9\x2d\x08\xd0\x44\x6e\x61\x07\x01\x5a\xef\xd0\xf2\x9f\xba\x9d\xcf\xe8\x1d\x50\x02\xb4\x14\x5a\xeb\x96\xde\x57\x65\x6c\xe4\xca\x4f\x4a\x69\x1b\x7a\xf9\x6a\x34\x7b\x87\x2e\xdd\x8b\x45\x0f\xf3\x9e\xb7\x2c\xbc\x15\x78\x69\x5b\xec\x5a\x01\x64\x42\xdf\x25\x8c\xd8\x1a\x47\x16\x59\xb5\xe3\x81\xa3\x08\xf9\x5b\x12\x65\x2e\xc8\xa2\x8c\xc9\x03\x5d\x9a\xe1\xf4\x4b\x9e\x4e\x25\x1d\xb3\x5f\x28\xac\x26\x9e\xe9\x02\x5f\xa5\x11\xb2\x8d\x32\x61\x64\x79\x44\x89\x61\xf2\x68\xd0\x54\xea\x45\xa3\xfb\x3c\x63\x70\x96\x42\xdb\xb8\xce\x96\xef\x11\x67\x11\xfc\x45\xa0\xe5\xd5\x79\xb6\x62\x4b\x86\x78\x4a\x51\xc2\x02\x81\x18\x95\x2c\xe3\x79\xd8\x78\x0f\xa2\xdd\x8d\xc4\x10\x81\x2f\x19\x9f\x32\x07\x78\x65\x69\x4e\xe1\xba\x99\xe5\x46\x6b\x9c\xa7\x11\x08\x35\x70\xc3\x33\x52\xbe\x63\xc4\xb0\xce\xd7\xcf\xae\x33\x1c\x20\xe7\xc3\x7a\x1a\xb2\xcc\x3d\x86\x0b\x3e\xf7\x5b\xe2\x6f\xf3\x49\x24\xb6\x2c\x8d\x82\xcc\xb0\x02\x86\xa8\xd9\x87\x22\x9d\x09\xa6\x4f\x8e\xed\xb4\x33\x12\x81\x20\x97\xc9\x11\x3a\xdb\x20\xca\xa8\x9e\x89\x77\x24\x80\x60\xae\x01\x2b\x5b\xf1\xd4\xe2\xa4\x1a\x66\xf1\xc7\x7b\x12\x45\x68\x0d\xaa\xaf\x60\x98\x82\x1e\x09\xcb\x4e\x4d\x3f\xe2\x60\x7b\x49\x86\x7f\x17\xe6\x7e\x8b\xc4\xa1\x1e\xe2\xf2\x1f\x2b\xc4\x41\xb0\x94\xfb\x30\x6c\xf3\x32\x80\xd2\x83\x1c\x71\xba\x00\xdc\x89\x6b\xed\xae\xca\x74\xe1\x7b\x83\x1f\xc2\x9a\x9c\x94\x84\x86\x42\x9b\x4c\x09\x35\x72\x28\x71\x03\x55\x3f\x90\x1a\xd9\x49\x8b\x9f\x90\x43\x76\x2f\x7f\xc1\x1c\x2d\xf7\x76\x1a\x1e\x75\x82\x46\x49\xbc\x6f\xd3\x35\x70\x0a\x12\x04\xd2\x4c\xa3\xdc\x8c\x0a\xeb\x6e\x65\x51\x18\x06\x62\x13\xf4\xd0\x33\xa9\x64\x44\x0e\x48\x13\xa7\x39\x39\x7d\x49\x11\x19\x20\x9e\x50\x12\xe3\xe8\x4f\x74\xe2\xd7\x94\x27\x31\x29\x28\x1c\xe0\xbb\xf4\x85\x84\xb1\x4e\x4b\x06\x08\xaf\x49\xaf\xfc\xbe\x35\x63\x52\x48\x8e\x93\xfa\x0e\x03\x35\x3b\x88\xd9\xcb\x6d\x06\x77\x7d\x46\x85\xc4\x51\x44\x68\x88\x30\xfa\xef\x94\xf8\xb7\x42\x62\x2e\x33\x17\x3f\xbf\x26\x12\x12\xc9\x12\xb1\x78\x4a\xf2\xf7\x3d\xec\xfd\x96\xbf\xef\xd9\xf7\x3d\x42\xbd\x1d\x4b\x79\x76\x59\x6e\xd8\x55\x92\xda\x4d\x91\x91\xbd\xde\xcc\x8e\x3b\xc6\xd5\x7c\xc5\x44\x69\x00\x97\x51\xb9\x45\xc6\x17\xd9\xdb\xad\x42\x7e\x69\x2e\x36\x5f\x41\xc2\xda\x04\xba\x89\xd2\x4f\xd3\x0b\x4c\x51\xbd\x99\x1d\x17\x78\x68\x1e\x3c\x87\x84\xf5\x1b\xb8\xa2\xf3\x3d\x0f\x7a\x10\x64\xf1\xf2\x60\xf7\x36\x32\x6f\x99\xa3\x93\x60\x99\xbd\x1c\xa7\xa3\x11\xae\x13\xef\xe2\xfd\x69\x7d\x2d\x5c\x19\xfc\x6b\x22\x2f\x12\xb5\x45\xdd\x1f\x14\xea\x98\x46\x44\xe8\xad\xfa\x9d\x98\x9c\x54\xf5\x1e\x52\x43\x13\x44\x32\xbe\x3b\x42\xd7\xaf\xb5\x20\xd1\xeb\x94\x04\x6a\xe6\x1b\xb9\x16\xe6\x5b\xe1\x22\x60\x97\x92\xbe\x2a\xe3\x05\x8b\xa8\xf3\x7c\x33\x3b\x2e\x8e\x6b\x6f\x07\x19\x08\x57\x12\x89\x0b\x78\xdc\xe4\x2e\xed\xad\xa6\xc1\xcf\x69\x4a\x97\xdb\x21\xcc\xd7\x44\x72\xcc\x77\xe8\xbf\x56\x17\xe7\x8b\xff\x59\xbe\x7f\x97\x67\x0a\x8b\x39\x12\xa9\xbf\x45\x58\x20\x1d\xcd\x73\x5c\x97\x67\x5c\x67\x94\xeb\x14\x63\x02\x43\x4f\xee\x1e\x92\x01\x87\x87\x94\x09\xb8\x76\xef\x75\xe2\x00\x19\x8e\xc9\x2b\x1c\x93\x68\xda\xec\xd8\xc7\x7d\xe7\x3e\x00\xa1\xc4\x76\x82\x13\xec\x13\xd9\x38\x72\x65\x14\x21\xf0\x03\x2f\x6e\xe7\x9a\x6b\xbc\xba\xad\xa1\x97\xfa\xba\x28\xc5\xa4\x5a\x78\xd4\xfb\x98\xce\x3d\x41\x8c\x3f\xad\xc8\xe7\x46\x89\xb4\x6a\x27\x26\x74\x74\xdb\xc9\x73\x25\x6d\xf4\xdb\xa6\x4a\x38\x0a\x1d\x55\x8f\xe7\x9b\xc8\xeb\xd5\xae\x21\x72\x27\x7a\xd6\x8f\xc8\xcd\x71\xb5\x7a\xf3\xdd\x85\x86\xba\xf3\x30\x59\x94\xc6\xd0\x47\xf5\x6d\xee\x57\x48\x42\xbc\xde\xc9\x81\x01\xa6\x86\x56\x05\xae\xff\xfd\xaf\x53\x05\x93\xf6\xa8\xdd\x84\x23\x2d\x70\xe7\x98\x27\x8e\x69\xe7\x96\x6a\x2b\xc6\x57\xec\xb1\x8e\x44\xad\x93\xa2\x6a\x82\x15\xb0\x9d\x74\x7b\x8b\x29\x7a\xf9\x76\xe5\xd5\x6a\x20\xa0\x0f\x17\xa7\x17\xe8\x17\x1c\x91\x20\x3f\xe0\xa4\x31\x4e\x12\x08\xd0\x86\x80\xf1\x03\x02\x24\xb7\x9c\xdd\x2b\x22\xc0\x39\xeb\x9f\x97\xf6\x30\xbd\x97\xdd\x05\x90\x9c\xf8\xe2\x84\x45\x6a\x4f\x5d\x4e\x49\x6e\xf0\x17\x42\x8e\x69\x1a\x61\xae\x4c\xa3\xb7\xdb\x50\x6c\x34\x25\x56\xc6\x86\xff\x6f\xef\x2e\x0c\x9a\x9e\x45\x69\x38\x06\x33\x89\xe9\xea\xa0\xe9\x7a\x67\x22\xa9\x3e\xd6\x3e\x7f\x76\x0f\x5e\xd7\xa7\xd0\x45\x00\xf6\xa5\x24\x4c\x95\xb3\xad\x94\x89\x78\xbe\x58\xa8\xff\x8e\xf0\xbd\x38\xc2\x31\xfe\xcc\xe8\x91\xcf\xe2\xc5\xf2\x1f\x2b\x5d\x83\xe7\x55\xd6\x66\xa1\x76\x15\x42\x2e\xfe\x2e\x80\x6b\x7f\x7f\x81\xef\x85\xb7\x37\x01\x0f\x0b\xcf\x8e\xc9\xcf\x0d\xec\x48\x59\x7a\xff\xbd\x4d\xd7\x30\xf6\xdb\x91\xaf\xc4\xfa\xcd\xec\xd8\x21\xb9\xfa\x4e\x27\xcb\x74\xec\x11\x72\xfa\xfa\x99\x74\x53\x64\x46\x0d\xb2\x78\x47\xee\xc5\x24\x56\x6e\xf6\x5a\x67\xa7\x1a\xe8\x4e\xce\x4e\xaf\x06\xee\xd2\x8a\x2d\xcb\xea\x7b\xc0\x0d\xd4\x01\x41\xeb\x55\x02\x3e\xd9\xec\xd0\xb5\x9f\x0a\xc9\x62\xb4\x7c\x7f\x56\xa8\x4d\xa8\x9f\x79\x38\x26\x9e\xad\x90\xb5\x78\x36\x47\x37\x3a\xeb\xc4\x13\x22\xbe\x99\x65\xff\xa9\xbf\x18\x47\x37\xfa\xde\x1a\xf1\x6f\x66\x83\x3c\x97\x8c\x89\x7a\x01\xaf\x3a\x03\x6a\xba\xec\x59\x55\xd3\x64\x8e\xfe\xfc\x5b\xca\xe4\xdf\x32\xae\xcc\x7f\xc5\xa7\xd9\x13\xc6\xed\x43\xc3\xa5\xf9\x7b\xe8\xce\xf2\x61\xf6\xab\x22\x6c\x5b\x39\x51\xcb\x1a\xd4\x50\xae\xaa\x46\x6d\xf0\x0a\xf4\x18\xb7\xd3\x6d\xa6\xfc\x8e\xc4\x44\x9a\x32\x49\x26\xac\xaf\xad\x8a\xf8\x68\xf9\xcf\xbd\x49\x2b\x73\xb0\xb0\xbf\x78\xfa\x99\x51\xf0\xf0\x3d\xe6\xe0\x19\xe3\x31\x3f\x0c\x8b\x68\x9a\x6e\x6b\xa6\xdb\xa7\x23\x5b\x38\xa9\xc6\x6d\x73\x8c\x77\xcd\xa4\x8c\x80\x33\xff\x16\x7a\xe6\x8e\xe6\xb8\xf3\xa2\xd8\xd4\x49\xdc\x8f\xb0\x10\xc4\x7f\xc7\x70\xf0\x02\x47\xca\x93\xe7\xe7\x38\x7e\x9c\xca\x5e\xda\xb4\x5f\x40\xfa\xc8\x66\x6d\xf9\x15\x26\x17\x50\x09\x39\x5f\xdd\xc3\x4a\xd2\x5c\xb7\x4a\x07\x13\x6f\x10\xa7\x0e\x82\x9e\x9e\xaf\x0e\xc0\xe7\xeb\x13\x03\x76\x38\x08\x38\x88\xbd\x1d\xdf\x25\xbe\x47\xf3\xbd\xcb\xe2\xa9\x45\xca\xac\xcc\x61\x40\x85\x67\x9b\x3c\x33\xc7\xdd\xca\x99\x3f\x3d\x5f\xa1\x88\xb1\xdb\x72\xd5\xa7\x11\x41\xfb\xfe\xbd\xdf\xcc\x8e\xcb\x23\xd0\x55\xe2\xba\x39\xea\x44\xcc\x29\x62\x68\xb0\x16\x17\x89\x24\x31\xf9\x0c\x8d\xfe\x4b\x43\x4c\xa4\x24\x1f\x53\xd3\x55\xa0\xeb\x97\x2f\x56\x3a\x46\x1e\x93\xcf\xda\x95\xeb\xf4\x7f\x5f\x9e\xfc\x54\xf7\x1c\x61\x2d\x3c\x96\xf1\x55\x71\x6f\xfb\xa8\x2b\x63\xa7\xb7\x2b\xdb\x93\x8b\x9b\xd9\x71\x75\x80\xcd\x48\xf5\x00\xf1\xc9\x69\xae\xab\x39\x08\x5f\x72\xd8\x90\x4f\x0f\x42\x7a\xf2\x98\x6a\x46\x58\x9c\x12\x61\xae\x95\xf5\xae\xa4\xb7\x97\xb4\x93\x86\xb3\xbb\xdb\x74\x0d\x11\xc8\x97\x3a\x41\xb9\x5a\xbf\xb7\xa5\xaf\x01\x05\x5e\x2c\xc4\x91\xcf\x80\x3e\xda\xee\x3e\xda\x2d\x59\xc5\x13\x25\x9f\x09\x0d\x3d\xb9\x05\xcf\xbe\x37\xb0\xb8\x65\x83\x7f\x59\x27\x9b\xa3\x96\x62\xca\xd4\x7f\xb6\x3f\xd9\xea\xcf\x96\xbf\x66\xf3\xff\xee\x43\xdf\x97\x2c\x10\x97\xc0\x95\xcd\x8c\x8b\x80\xff\x5f\x89\x9e\xb3\x3b\xe0\x9c\x04\xf0\x22\x3b\x23\x3e\x61\x71\x8c\x87\x16\xbb\x2d\xd9\xe1\x85\x25\x89\x3e\x9a\x9d\xf6\xc7\xbf\x08\x94\x1f\x41\x27\xca\xad\x30\xaf\x0f\x32\xee\x9c\xa8\xb1\x57\x43\xd9\x9a\x6b\x13\x7d\xe7\x80\x13\x5e\x1b\xeb\xa3\x74\x01\x41\x67\x33\x42\x80\xd6\xb0\x61\x1c\x2a\x23\xcc\x71\xd2\x14\x6e\x82\xda\x25\xde\x3e\x32\x1d\xd9\x45\x83\x58\x0f\x3b\x85\x29\x31\x66\x73\x23\xae\xb3\x8b\x0d\x7b\x67\xac\xd1\x43\x4c\x05\x78\xf6\x75\xcf\xe6\x75\x7a\x1b\xc6\x3d\x0d\xd9\x38\xda\xd7\x73\x7d\xa6\x3d\xb3\xfc\xdf\x41\x02\xb3\x7c\x75\x3a\x8c\xbd\x99\xb9\x99\x1d\xd7\xc7\xa8\x7d\xc8\x16\x26\x7b\x1e\x59\x15\x2f\xd5\xf4\xbc\x5a\xb5\x3f\xbd\x7a\xdd\x70\x47\x70\xdc\x41\x58\x9b\xae\xb3\x14\x0c\x10\x48\x88\x6d\x56\x83\xd2\xe4\x5c\x13\x31\x52\x51\xbd\x89\x3a\x07\xf9\x3d\x1f\xd1\x49\xec\xba\xea\xfb\xfd\x70\xcf\x43\x90\xda\x6e\xbe\x65\x5d\xb9\x7e\x5b\x73\xc3\xac\xd9\x23\x4f\xbc\x31\xef\x45\xda\x29\x41\x73\x98\x68\x8b\x12\x77\xef\xfb\x5a\x68\x9c\x5d\x5c\x36\x6e\xed\x5b\x7d\x14\xd3\xfc\x6d\x2c\xde\xc2\xee\xec\xb4\x77\x49\x98\x1a\x85\x1e\x1b\xa2\x96\xd6\xdf\xc5\x29\x75\x8d\xeb\xe1\x1b\xaa\x52\xf7\xfa\x70\x13\xdd\x61\x4e\x30\x35\xd7\x4b\x9f\xa3\x8f\x37\xb3\x30\xf9\xe9\x66\xf6\x11\x11\x81\x5e\xdb\xf2\x3f\x97\x29\x4f\x98\x00\xb4\x5a\x9d\xa2\x1f\x2c\x77\xcf\xe6\xea\x5d\xc2\x7e\xb4\xef\x5e\x72\x76\x47\x04\x61\x14\x02\xa4\x8c\x41\xbd\xac\x5f\x11\x7e\xf6\xca\x87\x2d\x67\x69\xb8\x4d\x52\x89\xf2\x50\x03\x7a\x73\x6a\x5f\x93\xd9\x6b\x27\x2c\xd2\x8f\x87\x65\x84\xbb\x06\x63\xbc\x3f\x13\xda\x0e\x93\x9f\xcc\x1f\xd9\xae\xa5\x7b\x7c\xc5\xe6\x84\xfd\x58\x6b\xee\x1e\x72\xb1\x95\xf0\xeb\xad\x9a\xa5\x50\x6a\x29\xeb\x2d\x1b\x04\x53\x30\x97\x30\xf9\xa9\xfc\x1b\xd0\x34\xae\x7f\x3c\xa2\xfa\x9a\x02\x4b\xf6\x63\xf5\x91\xf0\xeb\x8f\xe4\x8f\x45\x54\x2c\x7c\x6b\xe2\x49\xc5\x46\x87\x66\x50\x8c\x4e\xa8\x70\xc7\x00\x9a\x03\x1b\x4d\xb1\x94\xbe\xc9\x14\xd5\x74\x88\xd6\xdc\x89\x8a\x5b\xd5\x12\xb2\x73\xec\xf8\x1c\x1b\xc8\xae\x53\x92\xa6\x78\x9e\x1b\xef\xdc\x78\xe2\x46\xd6\x96\x45\xa3\x19\xcc\xdd\xab\x44\xf3\xde\xba\x1e\x33\xa8\xfb\x2a\x7d\x62\xf4\x2d\x3e\x42\xc5\x3b\xad\x04\xe9\x9a\x4e\x17\xba\x76\x84\x7d\xb6\xc8\xee\x68\x78\x7b\x78\xc9\xfe\x38\x49\x9d\xf8\x52\xf2\x72\xa1\x0c\x93\xdc\x62\xa9\xab\x02\xe4\xa7\x2e\x3a\x35\xb9\xee\x4a\xf7\x2c\x19\x3f\xba\x1f\xdd\x4d\xed\x88\xf8\x85\xfb\x94\xa7\xe3\xdb\x78\xcb\x20\x26\xf4\x24\xfb\x5c\xd7\x28\xbf\x26\xbb\x21\x37\x79\x8c\x2f\xaf\xae\x87\xe9\x0e\x5d\x17\xed\x2c\xbf\x95\xb7\x8f\x95\xef\x93\x12\x16\xc5\x37\x3d\x26\x4a\xff\x2f\x9e\x16\x3a\xf1\xd8\xc6\xcb\x28\x0d\x0b\x0a\x96\x58\xab\x87\xcc\x0f\x65\xe6\x66\x76\xec\x1c\xee\x21\x57\x1a\x9c\xfa\x76\xa9\x71\xc2\xb9\xa4\xe3\x1b\x25\x3b\x57\x9b\xc5\xa2\xa5\xa2\x35\x16\x10\xa0\xfd\x97\x45\xfa\x5f\xc9\x3a\xa0\x0b\xf7\x0c\xea\xf9\x05\x86\x47\x5d\xa6\x7b\xbf\x88\xeb\x3b\x29\x83\x0b\x2e\xf4\x3c\x48\x18\x55\xcc\x61\x00\xed\x07\x3b\xb5\x19\xf7\xa1\x86\x61\x7d\xa9\x9d\xe5\x32\x08\x18\xbd\xcc\xee\x4c\x0c\x3e\xd3\x2a\x37\x1f\x39\xe3\xdb\x2a\x4c\x3b\xec\xa4\x45\xcd\x6d\x5a\x1a\x20\xe4\x56\x19\x4d\x08\x3b\x4d\x9f\x61\xd8\xa7\x53\x0d\xc3\x98\x6e\x7a\x8d\x80\xd2\x64\x07\xcd\xe8\x12\xad\xcf\x68\xc8\xc7\x7e\xb9\x07\x27\xc9\x7b\xa8\x47\x13\x87\xb4\xbd\xe4\x70\x47\xe0\x7e\x1c\x89\x54\xb2\x95\x8f\xa3\x91\xae\x84\x0f\x5c\xda\x2f\xa9\x8e\x6b\xdf\xf8\x5d\xca\x5e\xcd\x61\x3d\x4e\xe8\xb0\x19\xd9\xee\x93\x04\x4e\x71\xd4\x92\xed\xd1\xda\x7e\x23\x1a\xcf\x9f\x5b\xdb\x91\x18\x87\xf0\x22\x25\x51\x30\x52\xce\x9f\xae\x9a\x8b\x12\x1f\xf8\x7d\x9a\x12\x6f\x6e\xcb\x6a\x90\x60\x83\x1d\x39\x26\x47\xb3\xcd\x57\x8c\xa1\x22\xeb\x8a\xca\xe7\xce\x59\x5b\x15\x93\xdb\x3c\x1f\x02\xed\x14\xd4\x8c\xbe\x28\xe8\x26\xd2\x80\x6b\x1d\xe9\x04\x0d\x29\xaa\xc5\xc8\x84\x03\xef\x9b\x10\xb1\xdc\xec\x31\x39\x5b\x6a\x57\xce\x49\x73\xfd\x07\x9a\xc6\xeb\xa6\x70\x2e\xa3\xa7\xa0\xb6\xbb\x2f\xb0\x80\x83\xf2\x91\x32\x42\x97\xc0\x7d\xa0\x12\x87\xb0\x5c\xb3\x3b\x38\x98\xae\x48\x98\xb4\x55\xe2\x08\xa3\x2b\xc9\xb1\x84\x70\xdc\x67\xc0\x12\x26\x33\x93\xb9\x64\xac\x9e\xe1\xd0\xcc\xcf\x30\xe8\x28\x19\x8a\x4b\x4f\x5d\xf2\x1f\x28\xd6\xd6\x31\x76\x8b\x72\x42\x0c\x70\xef\x82\xae\x55\xc7\xfb\x13\xe7\xfc\x94\x57\x3d\xf6\xf2\xc7\x03\xee\x78\xb7\x75\x56\x3b\xbe\xad\xf4\x72\x33\x3b\x2e\xb3\xe3\xb8\xad\x50\x3c\x28\xed\xbd\x13\x3b\x3b\x7d\x94\x27\x5a\x86\x39\x10\xc5\xe2\xb3\x59\x98\x13\xd9\xcb\xf2\x36\x05\x60\xdc\x69\xec\xa8\x0e\x1a\x37\x2c\xef\x98\x8f\xa3\x43\xb2\x0b\xec\x17\xaf\x70\x85\x07\xa4\xcc\x3e\xca\x3f\x84\x75\xc8\x50\x47\xd2\x2e\xe8\x55\xf2\xb4\xe1\x5c\x5f\x89\x60\xa5\x4b\x67\x4e\x20\x03\x53\x37\xaa\xc4\xa9\x2d\xe4\x8a\x63\x46\x43\xbd\xd8\xee\x0b\x8e\x22\x42\x47\xe7\x9a\x4c\xdf\x61\xb3\xb4\x06\x61\xf1\x7e\x6a\xba\x85\xec\xb4\xbe\x49\x00\xd1\x67\x54\x72\x16\x89\xda\x5c\x68\x49\x7e\xe8\x13\xee\xeb\x4b\xb3\x01\xd1\x56\x6f\xfa\xed\xfe\x22\x36\x6e\xe7\x65\x8a\x96\xbe\x85\x51\x2b\x74\xde\x78\xec\xe1\x70\x4e\xe0\x12\xcb\xc6\xbd\x57\xab\x8f\xa0\x4b\xda\x95\x4a\xd6\x9e\x7d\xbb\x0c\xb2\xb1\x46\xaf\xb5\xd7\x28\x16\xa7\xb6\x1a\xb5\xd0\x2d\x9c\x89\xf7\x10\x1a\x44\xf6\x69\x3d\xe5\x15\x5e\x1f\x41\x1c\x12\x3f\x19\x42\xbd\x34\x85\x2e\xea\xe5\x99\x9a\xef\x30\xb2\x38\x26\x32\x6b\xf1\x1e\x53\xb2\x01\x51\xcf\xdb\x19\x02\xe9\x27\x9a\xa4\xfd\xf4\x81\xd8\xa2\x57\x51\xfa\x09\xc5\x19\xe5\x6c\x81\x7d\x4d\xa4\xae\x39\x84\x18\x45\xb6\x28\xd1\x20\x1c\x1f\xdf\x8b\x73\x36\xe9\x43\xc1\x03\x12\x1e\x54\x47\xa6\x70\x9e\x64\xe8\x16\x20\x41\x92\x63\xff\x16\xb1\x8d\xe6\xec\x2f\x02\x89\x1d\xf5\x51\xc2\x99\xde\xf3\xfe\xcd\x60\x20\x11\x48\xed\xfb\xee\x70\x64\x3f\x82\x69\x8f\xf8\x08\x0d\x91\xe7\x85\x44\x7a\xaa\x95\x27\x71\xa8\x07\x6a\x1e\x51\x26\x41\x78\x1c\x36\x6a\x55\x52\xc4\x07\xc9\xed\xf1\x30\xfa\x50\x5f\x8c\x2c\x9b\x89\x2d\x91\xb4\x2f\xe5\x77\xbf\x05\xae\xab\x19\x5a\x7b\x30\x96\x63\x2e\xdc\x03\x7a\x03\x51\x8c\xb2\xe9\x60\xbe\x73\xb0\x19\x2a\xe2\x07\xe9\xb3\xc7\xa7\xda\x70\x70\x41\x9b\x2f\x4c\xf6\x99\xba\x6a\x3f\xc6\x53\x5f\x1a\xfe\x24\x2b\x7c\x5d\x27\x66\x81\xf9\x9c\x89\xcf\x41\x67\x86\x6d\x01\x05\x90\x44\x6c\x87\x6e\x61\x87\xb0\xd8\xbf\x3b\x48\x58\x0f\xd1\x65\xbf\xc4\x54\xe5\x45\x29\xd1\x1f\x2a\xb0\x0c\xab\x4b\x6a\x1c\x2c\x03\x37\x95\x91\xcb\x6a\x13\xaa\xd7\x00\xcf\x39\xdb\x5c\x32\x72\x19\xda\x24\xab\xe9\x90\x22\x68\x4a\x3e\x59\x5d\xb9\xbc\xba\xac\xc1\xb0\x42\x59\xe4\x6c\x5a\x95\x2b\xa0\x29\x0c\x52\x18\xd5\xff\xb4\xf4\xeb\x73\x56\x5a\xc5\x1d\x95\xf2\x9b\x16\x71\x96\xca\x24\x95\x3d\x1c\xc9\x36\x53\xbe\xd0\x44\x50\x40\xb8\xae\xfb\xba\xcb\xcb\x4d\x27\x9c\x29\x07\x04\x82\xac\x30\x24\x92\x10\x27\xfa\x22\x1b\xfa\xc1\x7c\xc6\x6f\x5f\xef\x1e\xf9\x26\x27\x65\x58\x72\xc0\x83\xf6\x5d\x40\x83\xa3\xc5\xcf\x85\xca\x98\x4a\x05\x9e\x32\xfb\xc6\x4a\x8f\x26\x07\xef\x00\xa1\xda\x52\xdf\xba\xb8\x26\x5a\x15\xab\x6b\x1e\xa1\x13\x4c\xd1\x1a\x10\x46\x6b\x8e\xa9\xbf\x9d\xeb\x92\xd5\xfa\x93\x19\xda\xb3\xd9\xe2\x52\xdc\xbd\xf7\x07\x08\xa6\xe9\xab\xbb\x40\xbf\xf6\xb0\x0f\x10\x8d\x72\xd9\x15\x0b\x7f\xbf\x7a\x87\xda\x58\x7f\xa5\xd6\xc4\x4f\x38\x4e\x22\x98\x23\x9c\x24\x5e\x00\x77\x83\xe4\x32\x5d\x47\x13\x14\xd4\xb0\x62\x73\x59\xd9\xdc\x39\xa1\xa7\x06\xd9\x00\x24\x26\x91\x2d\x20\xf9\x5b\xad\xe8\x6b\x5e\x6b\x12\xd4\x1b\x55\xf0\xc2\x41\x50\xf4\xac\x0b\xf5\x25\xc7\xa0\xea\x43\xb1\x52\x82\xd1\xab\x72\xb9\xd6\xe6\xfa\xc1\x7a\x62\x1c\x60\xcf\x1f\xb6\x80\x42\x22\xed\x0c\x43\x29\x0d\x80\xdb\x92\xd1\x19\xdf\x95\x10\x30\x89\x40\xe4\x45\xfc\xcd\x4c\x54\xcb\xc9\x9f\xf5\x86\x06\x02\xfb\x99\xad\x18\x0f\x8e\x71\x4d\xc7\x0a\x8e\x93\xbf\x75\xb3\xd3\x09\x16\x10\x63\x72\xe8\xee\x4a\xd3\xb0\xa3\x28\x7e\xf9\x40\x59\x81\x85\x31\x7f\x8b\x69\x38\x30\x51\xfe\x10\xd2\x9d\xe3\x56\x6e\xfc\x81\xcb\xb2\xd2\xe5\x7e\x5d\x2c\xaa\x52\x3b\xcc\x6d\x7a\xbc\xe7\x4a\x8b\x74\xbe\xdf\x52\x2c\x06\x9b\xd1\x43\x75\xdd\x5d\x4d\x11\xcb\xed\xa3\x3c\xa5\xb8\x52\x5e\x1f\xb9\x03\xa4\x39\xd4\x97\x62\x6c\x74\xb8\xe2\xd6\xd9\x72\xf1\xe6\x07\x5d\xb3\x37\x73\x10\xb5\x94\x62\x46\xd5\x7b\xca\xc6\x36\x84\x06\xa8\x50\x77\xbe\x14\xc6\xc0\x49\x12\xed\xac\x20\xaf\x6f\x74\x2a\xaf\x27\x76\x42\x82\xad\x0b\xb4\xc6\x02\x6e\x66\xbf\x0e\xd2\xec\x37\x1d\x83\xb9\x7c\x50\x18\x47\xb9\x92\x90\x1a\x8f\xf9\xeb\xd7\xd6\x1b\x9e\xab\xd5\x9b\x7e\xd1\xd3\x36\x65\xaa\xe6\xd9\x32\x92\x5d\x80\x5c\xad\xde\xe8\xfd\xe5\xfe\xbb\x07\x38\x95\x5b\xa0\x52\x7f\x3a\x77\x90\x9c\x0f\x25\xdf\x39\x4b\x52\x7e\x08\xac\x7e\xb0\x0a\x57\x2c\x29\xef\xc8\x72\x5a\xd3\xbf\xd6\xb5\x4d\xe0\x2d\x2d\xc4\x25\x08\xb0\xfa\x0d\x89\xfc\xcf\x7d\x36\xef\x73\xc6\xc3\x05\xaf\x15\x42\xef\x81\x3c\xdf\x86\xb1\x6e\x89\x8b\xe6\x0c\xa0\x9e\x2b\x99\x22\xf1\x30\x0b\xd9\x48\xca\x13\xf8\xb5\xca\x10\xe7\x35\x6f\xaa\x86\xe7\xae\xb5\xb1\x2a\xdc\x9a\xdf\xd0\x3a\xf7\xbf\x7a\x10\xa2\x5a\x5a\x7d\x5f\x65\xc7\x40\xe4\xc3\x04\x18\xba\x7b\x2d\x79\xbd\x2b\xf0\x39\x48\x61\x6f\xd2\xf4\xca\x35\xba\x85\xdd\xf2\xea\xbc\x7f\x92\x91\x7d\xbf\xef\xc9\xd5\x20\x6b\x6a\xe2\xe5\xd0\xef\x1f\xd5\x73\x30\xde\xbe\x5f\x21\xc8\xa5\x94\x7d\x78\x6b\x78\xad\xd7\x61\xd4\x4b\xba\x1a\x51\x1c\xb2\xa0\xcc\x06\xec\xa9\x9d\xb3\x53\x74\x76\x99\x15\x87\x42\x84\x9a\x0f\xf5\x52\x26\x71\xe9\x76\x57\xe7\xe9\x79\x3b\x99\x27\x99\xaa\xbf\x3c\xf9\xf2\xe4\x7f\x03\x00\x00\xff\xff\x89\x3e\xbf\xfa\xba\xa4\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5e, 0xeb, 0x7c, 0xd9, 0x67, 0x18, 0xb7, 0x30, 0x32, 0xc1, 0x7d, 0x62, 0x17, 0xdd, 0x94, 0x1c, 0x5f, 0xd2, 0x1, 0x89, 0xc5, 0xa9, 0x9, 0xd9, 0x90, 0x29, 0x8c, 0xbb, 0x94, 0xf0, 0x2b, 0xf2}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"schema.json": schemaJson,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
