package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdGetThings = &cobra.Command{
	Use:   "things",
	Short: "Get SensorThing Things: sti get things",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get things")
		// todo:get things from service and return them
	},
}
