package cmd

import (
	"github.com/gost/core"
	"github.com/spf13/cobra"
)

var cmdGetObservations = &cobra.Command{
	Use:   "observations",
	Short: "Get SensorThing Observations: sti get observations",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"ID"}
		getSTEntities(core.EntityTypeObservation, fields)
	},
}
