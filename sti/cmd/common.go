package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func getSTEntitys(name string) {
	if viper.IsSet("st_server") {
		stServer := viper.GetString("st_server")
		url := stServer + "/" + name
		thingsResponse := new(ThingsResponse)
		fmt.Println("Url: " + url)
		getJson(url, &thingsResponse)
		fmt.Println("Number of " + name + " :" + strconv.Itoa(thingsResponse.Count))
		for i := 0; i < len(thingsResponse.Value); i++ {
			e := thingsResponse.Value[i]
			// todo: why is e.iot_id nil???
			fmt.Println(fmt.Sprint(e.iot_id) + ", " + e.Name)
		}
	}
}

// Thing structure
type Thing struct {
	iot_id      interface{} `json:"@iot.id,omitempty"`
	Name        string
	Description string
}

// ThingsResponse structure
type ThingsResponse struct {
	Count int
	Value []Thing
}
