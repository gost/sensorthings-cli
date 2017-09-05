package cmd

import (
	"github.com/gost/core"
	"github.com/spf13/cobra"
)

var cmdGetDatastreams = &cobra.Command{
	Use:   "datastreams",
	Short: "Get SensorThing DataStreams: sti get datastreams",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"ID", "Name"}
		getSTEntities(core.EntityTypeDatastream, fields)
	},
}
