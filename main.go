package main

import (
	"fmt"
	"os"

	"github.com/miolfo/gostrim/taparse"
)

func main() {
	//Cut off the path to the program from arguments
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("Usage: ./gostrim [channel name]\n")
		os.Exit(1)
	}
	name := args[0]

	online := taparse.IsStreamerOnline(name)
	fmt.Printf("Streamer online: %t\n", online)
}
