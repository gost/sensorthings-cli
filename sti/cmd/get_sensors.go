package cmd

import (
	"github.com/gost/core"
	"github.com/spf13/cobra"
)

var cmdGetSensors = &cobra.Command{
	Use:   "sensors",
	Short: "Get SensorThing Sensors: sti get sensors",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"ID", "Name"}
		getSTEntities(core.EntityTypeSensor, fields)
	},
}
