package inputhandler

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

func getCommandsAndParameters(input string) []string {

}

func addChannel() {

}

func removeChannel() {

}
