package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the root command
type RootCmd struct {
	configFile   string
	cobraCommand *cobra.Command
}

// NewCmdRoot root of new commands
func NewCmdRoot(out io.Writer) *cobra.Command {
	cmd := rootCommand.cobraCommand
	cmd.AddCommand(NewCmdVersion(out))
	return cmd
}

func init() {
	cobra.OnInitialize()
	NewCmdRoot(os.Stdout)
}

var rootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "sti",
		Short: "sti is the sensorthings-cli to perform SensorThings tasks.",
		Long:  "sti is the sensorthings-cli to perform SensorThings tasks.",
	},
}

func main() {
	rootCommand.cobraCommand.Execute()
}
