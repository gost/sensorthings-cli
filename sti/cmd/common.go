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
		err := getJson(url, &thingsResponse)
		if err != nil {
			fmt.Println(err)
			return
		}
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
				}
				fmt.Printf(", ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Please use 'sti login' first...")
	}
}

// BaseEntity is the entry point for an entity
type BaseEntity struct {
	Iot_id  interface{} `json:"@iot.id,omitempty"`
	NavSelf string      `json:"@iot.selfLink,omitempty"`
}

// Thing structure
type Thing struct {
	BaseEntity
	Name        string
	Description string
}

type Observation struct {
	BaseEntity
	PhenomenonTime string      `json:"phenomenonTime,omitempty"`
	Result         interface{} `json:"result,omitempty"`
}

// Thing structure
type BaseResponse struct {
	Count int
}

// ThingsResponse structure
type ThingsResponse struct {
	BaseResponse
	Value []Thing
}

// ThingsResponse structure
type ObservationsResponse struct {
	BaseResponse
	Value []Observation
}
