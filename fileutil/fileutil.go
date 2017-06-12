package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//GetFileContents gets the contents of the file (name) in path (path)
func GetFileContents(path, fileName string) string {
	path, err := filepath.Abs(createFullPath(path, fileName))
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(path)
	return string(data)
}

//AddToFile appends a file named fileName at path with line
func addToFile(text, path, fileName string) {
	fullPath := createFullPath(path, fileName)
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
func emptyFile(path, fileName string) {
	if !fileExists(path, fileName) {
		panic("Can't empty non-existent file!")
	}
	ioutil.WriteFile(createFullPath(path, fileName), []byte(""), 0644)
}

//CreateFullPath combines the path and filename to an actual path
func createFullPath(path, fileName string) string {
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
	if _, err := os.Stat(createFullPath(path, fileName)); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
