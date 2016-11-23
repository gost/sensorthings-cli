package cmd

import "github.com/spf13/cobra"

// Sensor structure
type Sensor struct {
	Name string
}

// SensorsResponse structure
type SensorsResponse struct {
	Count int
	Value []Thing
}

var cmdGetSensors = &cobra.Command{
	Use:   "sensors",
	Short: "Get SensorThing Sensors: sti get sensors",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"Iot_id", "Name"}
		getSTEntitys("Sensors", fields)
	},
}
