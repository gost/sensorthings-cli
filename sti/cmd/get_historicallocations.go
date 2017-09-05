package cmd

import (
	"github.com/gost/core"
	"github.com/spf13/cobra"
)

var cmdGetHistoricalLocations = &cobra.Command{
	Use:   "historicallocations",
	Short: "Get SensorThing HistoricalLocations: sti get historicallocations",

	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"Iot_id", "Name"}
		getSTEntities(core.EntityTypeHistoricalLocation, fields)
	},
}
