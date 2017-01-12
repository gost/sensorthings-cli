package cmd

import (
	"fmt"
	"strconv"

	"github.com/oleiade/reflections"
	"github.com/spf13/viper"
)

func getSTEntities(entityType EntityType, fields []string) {
	if viper.IsSet("st_server") {
		stServer := viper.GetString("st_server")
		url := stServer + "/" + entityType.GetArrayEndpoint()
		thingsResponse := new(ThingsResponse)
		fmt.Println("Url: " + url)
		err := getJson(url, &thingsResponse)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Number of " + entityType.GetArrayEndpoint() + " :" + strconv.Itoa(thingsResponse.Count))
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
				fmt.Print(", ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Please use 'sti login' first...")
	}
}
