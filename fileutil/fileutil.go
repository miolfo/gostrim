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

//AddLineToFile appends a file named fileName at path with line
func AddLineToFile(line, path, fileName string) {
	fullPath := CreateFullPath(path, fileName)
	//Make sure that the file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		createFile(path, fileName)
	}
	file, err := os.OpenFile(fullPath, os.O_APPEND, 0666)
	checkError(err)
	_, err = file.WriteString(line + "\n")
	checkError(err)
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
