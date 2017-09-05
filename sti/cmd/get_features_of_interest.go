package cmd

import (
	"github.com/gost/core"
	"github.com/spf13/cobra"
)
var cmdGetFeaturesOfInterest = &cobra.Command{
	Use:   "featuresofinterest",
	Short: "Get SensorThing FeaturesOfInterest: sti get featuresofinterest",
	Run: func(cmd *cobra.Command, args []string) {
		fields := []string{"ID", "Name"}
		getSTEntities(core.EntityTypeFeatureOfInterest, fields)
	},
}
