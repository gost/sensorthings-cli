package main

import "github.com/gost/sensorthings-cli/sti/cmd"

var (
	// VERSION of cli tools
	VERSION = "0.0.3"
)

func main() {
	cmd.Execute(VERSION)
}
