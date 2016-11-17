package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Thing structure
type Thing struct {
	Name string
}

// ThingsResponse structure
type ThingsResponse struct {
	Count int
	Value []Thing
}

var cmdGetThings = &cobra.Command{
	Use: "things",

	Run: func(cmd *cobra.Command, args []string) {
		if viper.IsSet("st_server") {
			stServer := viper.GetString("st_server")
			url := stServer + "/Things"
			thingsResponse := new(ThingsResponse)
			fmt.Println("Url: " + url)
			getJson(url, &thingsResponse)
			fmt.Println("Number of things: " + strconv.Itoa(thingsResponse.Count))
			for i := 0; i < len(thingsResponse.Value); i++ {
				fmt.Println(thingsResponse.Value[i].Name)
			}
		}
	},
}
