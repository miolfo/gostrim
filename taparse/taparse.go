package taparse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

const apikeyPathName = "taparse/ta.key"
const rootURL = "https://api.twitch.tv/kraken/"
const accept = "application/vnd.twitchtv.v5+json"

var (
	apikey, authstring string
	initialized        = false
)

//Initialize the required functionality for accessing Twitch api
func initialize() {
	apikey = getAPIKeyFromFile()
	initialized = true
	authstring = "&client_id=" + apikey
}

//Get the users api key from ta.key
func getAPIKeyFromFile() string {
	path, err := filepath.Abs(apikeyPathName)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Is your API key pasted to gostrim/ta.key?")
	}
	return string(data)
}

//IsStreamerOnline checks if a particular streamer is online
func IsStreamerOnline(name string) bool {
	if !initialized {
		initialize()
	}
	id := getUserID(name)
	streamInfo := getStreamInfo(id)
	return streamInfo.Online()
}

//Get a streamJSON (info about the stream) object of the userID -user
func getStreamInfo(userID int) streamJSON {
	stream := "streams/" + strconv.Itoa(userID)
	query := rootURL + stream

	response := getResponse(query)
	defer response.Body.Close()

	var streamObject streamJSON
	errJSON := json.Unmarshal(getBytes(response), &streamObject)
	if errJSON != nil {
		panic(errJSON)
	}
	return streamObject
}

//Find the int type ID of the user by name
//Returns -1 if user was not found
func getUserID(name string) int {
	users := "users?login=" + name
	query := rootURL + users

	response := getResponse(query)
	defer response.Body.Close()

	var usersObject usersJSON
	errJSON := json.Unmarshal(getBytes(response), &usersObject)
	if errJSON != nil {
		panic(errJSON)
	}
	if len(usersObject.Users) < 1 {
		return -1
	}
	id, err := strconv.Atoi(usersObject.Users[0].ID)
	if err != nil {
		panic(err)
	}
	return id
}

//Get a http.Response object that is returned when query is run
func getResponse(query string) *http.Response {
	client := http.Client{}
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		panic(err)
	}

	addHeader(req)

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return response
}

//Add the headers requested by Twitch api v5
func addHeader(request *http.Request) {
	request.Header.Add("Accept", accept)
	request.Header.Add("Client-ID", apikey)
}

//Turn the http.Response object to string for JSON parsing
func getResponseString(response *http.Response) string {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

//Get the byte array from http.Response
func getBytes(response *http.Response) []byte {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return bytes
}
