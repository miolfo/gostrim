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

func TestAddLineToFile(t *testing.T) {
	//TODO: test
	t.Error("Test not implemented")
}
