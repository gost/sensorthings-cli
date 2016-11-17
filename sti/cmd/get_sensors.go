package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Sensor structure
type Sensor struct {
	Name string
}

// SensorsResponse structure
type SensorsResponse struct {
	Count int
	Value []Thing
}

var cmdGetSensors = &cobra.Command{
	Use:   "sensors",
	Short: "Get SensorThing Sensors: sti get sensors",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.IsSet("st_server") {
			stServer := viper.GetString("st_server")
			url := stServer + "/Sensors"
			sensorsResponse := new(SensorsResponse)
			fmt.Println("Url: " + url)
			getJson(url, &sensorsResponse)
			fmt.Println("Number of sensors: " + strconv.Itoa(sensorsResponse.Count))
			for i := 0; i < len(sensorsResponse.Value); i++ {
				fmt.Println(sensorsResponse.Value[i].Name)
			}
		}
	},
}
