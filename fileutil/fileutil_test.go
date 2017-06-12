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
		actual := createFullPath(cases.path, cases.name)
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
	testString := "test_write_string"
	path := "fileutil_test"
	fileName := "test_add.txt"
	emptyFile(path, fileName)
	addToFile(testString, path, fileName)
	content := GetFileContents(path, fileName)
	if content != testString {
		t.Error("File contents should be " + testString + ", was " + content)
	}
	addToFile(testString, path, fileName)
	content = GetFileContents(path, fileName)
	if content != (testString + testString) {
		t.Error("File contents should be " + testString + testString + ", was " + content)
	}
}
