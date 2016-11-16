package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// VERSION is set during build
	VERSION string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get SensorThings cli version",
	Run: func(cmd *cobra.Command, args []string) {

		if err := viper.ReadInConfig(); err == nil {
			if viper.IsSet("st_server") {
				stServer := viper.GetString("st_server")
				fmt.Println("server:" + stServer)
			}
		}
		fmt.Println(VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
