package cmd

import (
	"github.com/gost/core"
	"github.com/spf13/cobra"
)

var cmdGetLocations = &cobra.Command{
	Use:   "locations",
	Short: "Get SensorThing Locations: sti get locations",

	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"ID", "Name"}
		getSTEntities(core.EntityTypeLocation, fields)
	},
}
