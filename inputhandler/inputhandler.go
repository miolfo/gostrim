package inputhandler

import (
	"errors"
	"fmt"
	"strings"
)

type inputHandler func()

type InputAction struct {
	shortCmd string
	longCmd  string
	action   inputHandler
}

type cmdAndParam struct {
	cmd       string
	parameter string
}

var inputActions = []InputAction{
	InputAction{
		shortCmd: "a",
		longCmd:  "add",
		action:   addChannel,
	},
	InputAction{
		shortCmd: "r",
		longCmd:  "remove",
		action:   removeChannel,
	},
}

//HandleInput takes the command line input from the user and handles it appropriately
func HandleInput(input string) {

}

func getCommandsAndParameters(input string) {
	//Get the first command
	remaining := input
	var err error
	var cmd, param string
	for err == nil {
		remaining, cmd, param, err = getNextCommandAndParameter(remaining)
		fmt.Printf("Command found was %v, param is %v, remaining %v\n", cmd, param, remaining)
	}
}

func getNextCommandAndParameter(input string) (remaining string, cmd string, param string, err error) {
	ind := strings.Index(input, "-")
	if ind == -1 {
		return "", "", "", errors.New("No more arguments")
	}
	c := string(input[ind+1])
	rem := input[ind+2 : len(input)]
	//See if any more arguments remaining
	nextInd := strings.Index(rem, "-")
	var e error
	if nextInd == -1 {
		e = errors.New("No more arguments")
	}
	return rem, c, rem[strings.Index(rem, " "):len(rem)], e
}

func addChannel() {

}

func removeChannel() {

}
