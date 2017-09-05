package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gost/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

// getCmd represents the get command
var exportCmd = &cobra.Command{
	Use:   "export {$filename}",
	Short: "Export Things, Locations, Datastreams, Sensors and ObservedProperties. HistoricalLocations, Observations and FOIS are ignored for now",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunExport(cmd, args)
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
}

func RunExport(cmd *cobra.Command, args []string) error {
	if viper.IsSet("st_server") {
		if len(args) != 1 {
			return cmd.Help()
		}

		stServer := viper.GetString("st_server")
		var err error

		export := &CliExport{}

		fmt.Print("Fetching Things\n")
		export.Things, err = getAllThings(fmt.Sprintf("%v/%v%v", stServer, core.EntityTypeThing.GetArrayEndpoint(), ""), nil)
		if err != nil {
			return err
		}

		for _, e := range export.Things {
			relLocations, _ := getAllLinks(fmt.Sprintf("%s/%s(%v)/%s%s", stServer, core.EntityTypeThing.GetArrayEndpoint(), e.ID, core.EntityTypeLocation.GetArrayEndpoint(), "?$select=id"), e.ID, nil)
			export.ThingLocations = append(export.ThingLocations, relLocations)

			relDatastreams, _ := getAllLinks(fmt.Sprintf("%s/%s(%v)/%s%s", stServer, core.EntityTypeThing.GetArrayEndpoint(), e.ID, core.EntityTypeDatastream.GetArrayEndpoint(), "?$select=id"), e.ID, nil)
			export.ThingDatastreams = append(export.ThingDatastreams, relDatastreams)
		}

		fmt.Print("Fetching Locations\n")
		export.Locations, err = getAllLocations(stServer+"/"+core.EntityTypeLocation.GetArrayEndpoint()+"?$orderby=id%20asc", nil)
		if err != nil {
			return err
		}

		fmt.Print("Fetching Sensors\n")
		export.Sensors, err = getAllSensors(stServer+"/"+core.EntityTypeSensor.GetArrayEndpoint(), nil)
		if err != nil {
			return err
		}

		fmt.Print("Fetching ObservedProperties\n")
		export.ObservedProperties, err = getAllObservedProperties(stServer+"/"+core.EntityTypeObservedProperty.GetArrayEndpoint(), nil)
		if err != nil {
			return err
		}

		fmt.Print("Fetching Datastreams\n")
		export.Datastreams, err = getAllDatastreams(stServer+"/"+core.EntityTypeDatastream.GetArrayEndpoint(), nil)
		if err != nil {
			return err
		}

		for _, e := range export.Datastreams {
			relSensor, _ := getLink(fmt.Sprintf("%s/%s(%v)/%s%s", stServer, core.EntityTypeDatastream.GetArrayEndpoint(), e.ID, core.EntityTypeSensor.GetEndpoint(), "?$select=id"), e.ID, nil)
			export.DatastreamSensor = append(export.DatastreamSensor, relSensor)

			relObs, _ := getLink(fmt.Sprintf("%s/%s(%v)/%s%s", stServer, core.EntityTypeDatastream.GetArrayEndpoint(), e.ID, core.EntityTypeObservedProperty.GetEndpoint(), "?$select=id"), e.ID, nil)
			export.DatastreamObservedProperty = append(export.DatastreamObservedProperty, relObs)
		}

		export.ClearNav()

		fmt.Print("Writing file\n")

		exportJson, err := json.Marshal(export)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(args[0], exportJson, 0644)
		if err != nil {
			return err
		}

		fmt.Print("-------------\n")
		fmt.Printf("Export finished: %v\n", args[0])
		fmt.Print("-------------\n")
	} else {
		fmt.Println("Please use 'sti login' first...")
	}

	return nil
}

func getAllThings(url string, list []*core.Thing) ([]*core.Thing, error) {
	if list == nil {
		list = make([]*core.Thing, 0)
	}

	response := &ThingsResponse{}

	err := getJson(url, response)
	if err != nil {
		return nil, err
	}

	list = append(list, response.Value...)

	if len(response.NextLink) > 0 {
		list, err = getAllThings(response.NextLink, list)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func getAllLinks(url string, id interface{}, relation *Relation) (*Relation, error) {
	if relation == nil {
		relation = &Relation{}
		relation.EntityID = id
	}

	response := IDArrayResponse{}
	err := getJson(url, &response)
	if err != nil {
		return nil, err
	}

	for _, r := range response.Value {
		relation.LinkedIDs = append(relation.LinkedIDs, r.ID)
	}

	if len(response.NextLink) > 0 {
		relation, err = getAllLinks(response.NextLink, id, relation)
		if err != nil {
			return nil, err
		}
	}

	return relation, nil
}

func getLink(url string, id interface{}, relation *Relation) (*Relation, error) {
	if relation == nil {
		relation = &Relation{}
		relation.EntityID = id
	}

	response := core.BaseEntity{}
	err := getJson(url, &response)
	if err != nil {
		return nil, err
	}

	relation.LinkedIDs = append(relation.LinkedIDs, response.ID)
	return relation, nil
}

func getAllLocations(url string, list []*core.Location) ([]*core.Location, error) {
	if list == nil {
		list = make([]*core.Location, 0)
	}

	response := &LocationsResponse{}
	err := getJson(url, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	list = append(list, response.Value...)

	if len(response.NextLink) > 0 {
		list, err = getAllLocations(response.NextLink, list)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func getAllSensors(url string, list []*core.Sensor) ([]*core.Sensor, error) {
	if list == nil {
		list = make([]*core.Sensor, 0)
	}

	response := &SensorsResponse{}
	err := getJson(url, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	list = append(list, response.Value...)

	if len(response.NextLink) > 0 {
		list, err = getAllSensors(response.NextLink, list)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func getAllObservedProperties(url string, list []*core.ObservedProperty) ([]*core.ObservedProperty, error) {
	if list == nil {
		list = make([]*core.ObservedProperty, 0)
	}

	response := &ObservedPropertiesResponse{}
	err := getJson(url, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	list = append(list, response.Value...)

	if len(response.NextLink) > 0 {
		list, err = getAllObservedProperties(response.NextLink, list)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func getAllDatastreams(url string, list []*core.Datastream) ([]*core.Datastream, error) {
	if list == nil {
		list = make([]*core.Datastream, 0)
	}

	response := &DatastreamsResponse{}
	err := getJson(url, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	list = append(list, response.Value...)

	if len(response.NextLink) > 0 {
		list, err = getAllDatastreams(response.NextLink, list)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}
