package fileutil

import (
	"io/ioutil"
	"path/filepath"
)

//GetFileContents gets the contents of the file (name) in path (path)
func GetFileContents(path, fileName string) string {
	path, err := filepath.Abs(CreateFullPath(path, fileName))
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(path)
	return string(data)
}

//CreateFullPath combines the path and filename to an actual path
func CreateFullPath(path, fileName string) string {
	p := path
	if len(path) == 0 {
		p = "/"
	} else if !(path[len(path)-1:] == "/") {
		p = path + "/"
	}
	n := fileName
	if string(n[0]) == "/" {
		n = fileName[1:len(fileName)]
	}
	return p + n
}
