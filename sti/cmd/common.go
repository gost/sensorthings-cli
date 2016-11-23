package cmd

import (
	"fmt"
	"strconv"

	"github.com/oleiade/reflections"
	"github.com/spf13/viper"
)

func getSTEntitys(name string, fields []string) {
	if viper.IsSet("st_server") {
		stServer := viper.GetString("st_server")
		url := stServer + "/" + name
		thingsResponse := new(ThingsResponse)
		fmt.Println("Url: " + url)
		getJson(url, &thingsResponse)
		fmt.Println("Number of " + name + " :" + strconv.Itoa(thingsResponse.Count))
		for i := 0; i < len(thingsResponse.Value); i++ {
			e := thingsResponse.Value[i]
			for j := 0; j < len(fields); j++ {
				fldName := fields[j]
				value, _ := reflections.GetField(e, fldName)
				switch v := value.(type) {
				case int:
					fmt.Printf("%v", v)
				case float64:
					fmt.Printf("%v", v)
				case string:
					fmt.Printf("%v", v)
				default:
					// fmt.Printf("I don't know, ask stackoverflow.")
				}
				fmt.Printf(", ")
			}
			fmt.Println()
		}
	}
}

// Thing structure
type Thing struct {
	Iot_id      interface{} `json:"@iot.id,omitempty"`
	Name        string
	Description string
}

// ThingsResponse structure
type ThingsResponse struct {
	Count int
	Value []Thing
}
