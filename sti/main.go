package main

import "github.com/geodan/sensorthings-cli/sti/cmd"

var (
	// VERSION of cli tools
	VERSION = "0.0.1"
)

func main() {
	cmd.Execute(VERSION)
}
