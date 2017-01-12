package cmd

import "github.com/spf13/cobra"

var cmdGetSensors = &cobra.Command{
	Use:   "sensors",
	Short: "Get SensorThing Sensors: sti get sensors",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"ID", "Name"}
		getSTEntities(EntityTypeSensor, fields)
	},
}
