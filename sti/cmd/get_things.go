package cmd

import "github.com/spf13/cobra"

var cmdGetThings = &cobra.Command{
	Use:   "things",
	Short: "Get SensorThing Things: sti get things",

	Run: func(cmd *cobra.Command, args []string) {
		getSTEntitys("Things")
	},
}
