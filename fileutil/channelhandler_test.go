package fileutil

import "testing"

const testPath = "fileutil_test"
const testFileName = "test_channelhandler.txt"
const testChannelName = "test channel"

func TestAddFollowedChannel(t *testing.T) {
	setFileName(testFileName)
	setFilePath(testPath)

	emptyFile(testPath, testFileName)
	AddFollowedChannel(testChannelName)
	result := GetFileContents(testPath, testFileName)
	if result != (testChannelName + string(getChannelSeparator())) {
		t.Error("Expected " + testChannelName + string(getChannelSeparator()) + ", was " + result)
	}
}

func TestRemoveFollowedChannel(t *testing.T) {
	setFileName(testFileName)
	setFilePath(testPath)
	emptyFile(testPath, testFileName)

	anotherChannel := "another channel"
	AddFollowedChannel(testChannelName)
	AddFollowedChannel(anotherChannel)
	RemoveFollowedChannel(testChannelName)
	result := GetFileContents(testPath, testFileName)
	if result != anotherChannel+string(getChannelSeparator()) {
		t.Error("Expected " + anotherChannel + string(getChannelSeparator()) + " was " + result)
	}
}
