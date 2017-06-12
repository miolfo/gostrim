package fileutil

import "strings"

var (
	channelSeparator    byte = ','
	channelListFilePath      = "conf"
	channelListFileName      = "channels.txt"
)

//AddFollowedChannel adds channel with channelName to the list of followed channels
func AddFollowedChannel(channelName string) {
	//TODO: handle special characters
	addToFile(channelName+string(channelSeparator), channelListFilePath, channelListFileName)
}

//RemoveFollowedChannel removes a channelName named channel from the list of followed channels
func RemoveFollowedChannel(channelName string) {
	channels := getChannelList()
	emptyFile(channelListFilePath, channelListFileName)
	for _, channel := range channels {
		if channelName != channel && channel != "" {
			addToFile(channel+string(channelSeparator), channelListFilePath, channelListFileName)
		}
	}

}

func getChannelList() []string {
	content := GetFileContents(channelListFilePath, channelListFileName)
	list := strings.Split(content, string(channelSeparator))
	return list
}

func setFilePath(filePath string) {
	channelListFilePath = filePath
}

func setFileName(fileName string) {
	channelListFileName = fileName
}

func getChannelSeparator() byte {
	return channelSeparator
}

func setChannelSeparator(separator byte) {
	channelSeparator = separator
}
