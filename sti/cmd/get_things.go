package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Thing struct {
	Name string
}

type ThingsResponse struct {
	Count int
	Value []Thing
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

var cmdGetThings = &cobra.Command{
	Use:   "things",
	Short: "Get SensorThing Things: sti get things",
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
