package cmd

import (
	"github.com/spf13/cobra"
	"github.com/asaskevich/govalidator"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"net/http"
)

// getCmd represents the get command
var importCmd = &cobra.Command{
	Use:   "import {$filename} {$server}",
	Short: "Import entities to a given SensorThings server",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunImport(cmd, args)
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(importCmd)
}

func RunImport(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return cmd.Help()
	}

	validURL := govalidator.IsURL(args[1])

	if !validURL {
		exitWithError(fmt.Errorf("Not a valid url: %s", args[1]))
	}

	f, err := ioutil.ReadFile(args[0])
	if err != nil {
		exitWithError(fmt.Errorf("Unable to read file %s: %v", args[0], err))
	}

	export := &CliExport{}
	err = json.Unmarshal(f, export)
	if err != nil {
		exitWithError(fmt.Errorf("Unable to parse file %s: %v", args[0], err))
	}

	s := make([]Entity, len(export.Locations))
	for i, v := range export.Locations { s[i] = v }
	locationMap, err := postEntities(fmt.Sprintf("%v/%v", args[1], EntityTypeLocation.GetArrayEndpoint()), s)
	if err != nil {
		return err
	}

	s = make([]Entity, len(export.Sensors))
	for i, v := range export.Sensors { s[i] = v }
	sensorMap, err := postEntities(fmt.Sprintf("%v/%v", args[1], EntityTypeSensor.GetArrayEndpoint()), s)
	if err != nil {
		return err
	}

	s = make([]Entity, len(export.ObservedProperties))
	for i, v := range export.ObservedProperties { s[i] = v }
	observedPropertiesMap, err := postEntities(fmt.Sprintf("%v/%v", args[1], EntityTypeObservedProperty.GetArrayEndpoint()), s)
	if err != nil {
		return err
	}

	thingMap, err := postThings(fmt.Sprintf("%v/%v", args[1], EntityTypeThing.GetArrayEndpoint()), export.Things, export.ThingLocations, locationMap)
	if err != nil {
		return err
	}

	err = postDatastreams(fmt.Sprintf("%v/%v", args[1], EntityTypeDatastream.GetArrayEndpoint()), export.Datastreams, export.ThingDatastreams, export.DatastreamSensor, export.DatastreamObservedProperty, thingMap, sensorMap, observedPropertiesMap)
	if err != nil {
		return err
	}

	fmt.Println("---------------")
	fmt.Println("--Import done--")
	fmt.Println("---------------")

	return nil
}

func postEntities(url string, entities []Entity) (map[interface{}]interface{}, error) {
	idMap := make(map[interface{}]interface{})
	fmt.Printf("Posting %v\n", url)

	for i := len(entities)-1; i >= 0; i-- {
		e := entities[i]
		jsonStr, err := json.Marshal(e)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		be := &BaseEntity{}
		err = json.NewDecoder(resp.Body).Decode(be)
		if err != nil{
			return nil, err
		}

		idMap[e.GetID()] = be.ID
	}

	return idMap, nil
}

func postThings(url string, entities []*Thing, thingLocations []*Relation, locationMap map[interface{}]interface{}) (map[interface{}]interface{}, error) {
	fmt.Printf("Posting %v\n", url)
	idMap := make(map[interface{}]interface{})
	for i := len(entities)-1; i >= 0; i-- {
		e := entities[i]
		e.Locations = make([]*Location, 0)
		for _, r := range thingLocations {
			if r.EntityID == e.ID { // find the connected locations for thing
				for _, id := range r.LinkedIDs {
					locationID, ok := locationMap[id] // get the new id from posted version
					if !ok {
						return nil, fmt.Errorf("Linked thing location %v was not posted", id)
					}

					location := &Location{}
					location.ID = locationID
					e.Locations = append(e.Locations, location)
				}
			}
		}

		jsonStr, err := json.Marshal(e)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		be := &Thing{}
		err = json.NewDecoder(resp.Body).Decode(be)
		if err != nil{
			return nil, err
		}
		idMap[e.GetID()] = be.ID
	}

	return idMap, nil
}

func postDatastreams(url string, entities []*Datastream, thingDatastreams []*Relation, datastreamSensor []*Relation, datastreamObservedProperty []*Relation, thingMap map[interface{}]interface{}, sensorMap map[interface{}]interface{}, observedPropertiesMap map[interface{}]interface{}) error {
	fmt.Printf("Posting %v\n", url)

	for i := len(entities)-1; i >= 0; i-- {
		e := entities[i]

		for _, r := range thingDatastreams {
			for _, id := range r.LinkedIDs{
				if id == e.ID { // if datastream is linked to this thing
					newID, ok := thingMap[r.EntityID]
					if !ok {
						return fmt.Errorf("Thing not posted %v", r.EntityID)
					}
					thing := &Thing{}
					thing.ID = newID
					e.Thing = thing
				}
			}
		}

		for _, ds := range datastreamSensor {
			if ds.EntityID == e.ID {
				if len(ds.LinkedIDs) == 0 {
					return fmt.Errorf("Datastream %v has no Sensor", e.ID)
				}

				newID, ok := sensorMap[ds.LinkedIDs[0]]
				if !ok {
					return fmt.Errorf("Sensor was not posted %v", ds.LinkedIDs[0])
				}

				sensor := &Sensor{}
				sensor.ID = newID
				e.Sensor = sensor
			}
		}

		for _, do := range datastreamObservedProperty {
			if do.EntityID == e.ID {
				if len(do.LinkedIDs) == 0 {
					return fmt.Errorf("Datastream %v has no ObservedProperty", e.ID)
				}

				newID, ok := observedPropertiesMap[do.LinkedIDs[0]]
				if !ok {
					return fmt.Errorf("ObservedProperty was not posted %v", do.LinkedIDs[0])
				}

				op := &ObservedProperty{}
				op.ID = newID
				e.ObservedProperty = op
			}
		}

		jsonStr, err := json.Marshal(e)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		be := &Thing{}
		err = json.NewDecoder(resp.Body).Decode(be)
		if err != nil{
			return err
		}
	}

	return nil
}