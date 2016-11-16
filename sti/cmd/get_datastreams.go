package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdGetDatastreams = &cobra.Command{
	Use:   "datastreams",
	Short: "Get SensorThing DataStreams: sti get datastreams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get datastreams")
		// todo:get datastreams from service and return them
	},
}
