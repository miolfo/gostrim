package inputhandler

import "testing"

func TestGetCommandAndParameter(t *testing.T) {
	input := "-a test_channel"
	result := getCommandsAndParameters(input)
}
