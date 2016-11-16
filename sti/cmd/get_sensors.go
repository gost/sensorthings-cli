package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdGetSensors = &cobra.Command{
	Use:   "sensors",
	Short: "Get SensorThing Sensors: sti get sensors",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get sensors")
		// todo:get sensors from service and return them
	},
}
