package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// NewCmdVersion new command for getting cli version
func NewCmdVersion(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version ",
		Short: "Get SensorThings cli version",
		Run: func(cmd *cobra.Command, args []string) {
			RunVersion(cmd, args, out)
		},
	}

	return cmd
}

// RunVersion prints the version
func RunVersion(cmd *cobra.Command, args []string, out io.Writer) {
	fmt.Printf("0.0.0.1")
}
