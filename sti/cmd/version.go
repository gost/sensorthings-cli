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
		// configPath := os.Getenv("HOME")
		viper.SetConfigName(".sti.yaml")

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
