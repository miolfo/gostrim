package fileutil

import "testing"

var createFullPathTests = []struct {
	name     string
	path     string
	expected string
}{
	{"name.txt", "test", "test/name.txt"},
	{"name.txt", "test/", "test/name.txt"},
	{"name.txt", "", "/name.txt"},
	{"/name.txt", "asd", "asd/name.txt"},
}

func TestCreateFullPath(t *testing.T) {
	for _, cases := range createFullPathTests {
		actual := CreateFullPath(cases.path, cases.name)
		if actual != cases.expected {
			t.Error("Expected " + cases.expected + ", got " + actual)
		}
	}
}

func TestGetFileContents(t *testing.T) {
	content := "test_read_string"
	result := GetFileContents("fileutil_test", "test_read.txt")
	if result != content {
		t.Error("Expected " + content + ", got " + result)
	}
}

func TestAddToFile(t *testing.T) {
	//File should be currently empty
	testString := "test_write_string"
	path := "fileutil_test"
	fileName := "test_add.txt"
	EmptyFile(path, fileName)
	AddToFile(testString, path, fileName)
	content := GetFileContents(path, fileName)
	println(content)
	if content != testString {
		t.Error("File contents should be " + testString + ", was " + content)
	}
	AddToFile(testString, path, fileName)
	content = GetFileContents(path, fileName)
	if content != (testString + testString) {
		t.Error("File contents should be " + testString + testString + ", was " + content)
	}
}
