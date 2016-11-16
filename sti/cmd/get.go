package cmd

import "github.com/spf13/cobra"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get SensorThings entities",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			exitWithError(cmd.Help())
		}
	},
}

func init() {
	getCmd.AddCommand(cmdGetThings)
	getCmd.AddCommand(cmdGetSensors)
	getCmd.AddCommand(cmdGetDatastreams)
	RootCmd.AddCommand(getCmd)
}
