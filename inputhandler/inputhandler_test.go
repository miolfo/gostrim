package inputhandler

import "testing"

func TestGetCommandAndParameter(t *testing.T) {
	input := "-a test_channel -b sing_sing"
	getCommandsAndParameters(input)
}
