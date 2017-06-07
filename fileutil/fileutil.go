package fileutil

import (
	"io/ioutil"
	"os"
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

//AddToFile appends a file named fileName at path with line
func AddToFile(text, path, fileName string) {
	fullPath := CreateFullPath(path, fileName)
	//Make sure that the file exists
	if !fileExists(path, fileName) {
		createFile(path, fileName)
	}
	file, err := os.OpenFile(fullPath, os.O_APPEND, 0666)
	checkError(err)
	_, err = file.WriteString(text)
	checkError(err)
}

//EmptyFile removes the contents of a file
func EmptyFile(path, fileName string) {
	if !fileExists(path, fileName) {
		panic("Can't empty non-existent file!")
	}
	ioutil.WriteFile(CreateFullPath(path, fileName), []byte(""), 0644)
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

func createFile(path, fileName string) {
	err := os.Mkdir(path, 0666)
	checkError(err)
	_, err = os.Create(fileName)
	checkError(err)
}

func fileExists(path, fileName string) bool {
	if _, err := os.Stat(CreateFullPath(path, fileName)); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
