package cmd

// ArrayResponse is the default response format for sending content back
type ArrayResponse struct {
	Count    int    `json:"count,omitempty"`
	NextLink string `json:"@iot.nextLink,omitempty"`
}

// IDArrayResponse structure
type IDArrayResponse struct {
	ArrayResponse
	Value []*BaseEntity `json:"value"`
}

// ThingsResponse structure
type ThingsResponse struct {
	ArrayResponse
	Value []*Thing `json:"value"`
}

// ObservationsResponse structure
type ObservationsResponse struct {
	ArrayResponse
	Value []*Observation `json:"value"`
}

// DatastreamsResponse structure
type DatastreamsResponse struct {
	ArrayResponse
	Value []*Datastream `json:"value"`
}

// LocationsResponse structure
type LocationsResponse struct {
	ArrayResponse
	Value []*Location `json:"value"`
}

// HistoricalLocationsResponse structure
type HistoricalLocationsResponse struct {
	ArrayResponse
	Value []*HistoricalLocation `json:"value"`
}

// SensorsResponse structure
type SensorsResponse struct {
	ArrayResponse
	Value []*Sensor `json:"value"`
}

// ObservedPropertiesResponse structure
type ObservedPropertiesResponse struct {
	ArrayResponse
	Value []*ObservedProperty `json:"value"`
}

// FeaturesOfInterestResponse structure
type FeaturesOfInterestResponse struct {
	ArrayResponse
	Value []*FeatureOfInterest `json:"value"`
}

// CliExport struct for importing and exporting SensorThings data
type CliExport struct {
	Things                     []*Thing            `json:"things"`
	ThingLocations             []*Relation         `json:"thingLocations"`
	ThingDatastreams           []*Relation         `json:"thingDatastreams"`
	Locations                  []*Location         `json:"locations"`
	Sensors                    []*Sensor           `json:"sensors"`
	ObservedProperties         []*ObservedProperty `json:"observedProperties"`
	Datastreams                []*Datastream       `json:"datastreams"`
	DatastreamSensor           []*Relation         `json:"datastreamSensor"`
	DatastreamObservedProperty []*Relation         `json:"datastreamObservedProperty"`
}

// ClearNav clears all navigational links from the entities in CliExport
func (c *CliExport) ClearNav() {
	for _, e := range c.Things {
		e.ClearNav()
	}

	for _, e := range c.Locations {
		e.ClearNav()
	}

	for _, e := range c.Sensors {
		e.ClearNav()
	}

	for _, e := range c.ObservedProperties {
		e.ClearNav()
	}

	for _, e := range c.Datastreams {
		e.ClearNav()
	}
}

type Relation struct {
	EntityID  interface{}   `json:"entityId"`
	LinkedIDs []interface{} `json:"linkedIds"`
}

// ErrorResponse is the default response format for sending errors back
type ErrorResponse struct {
	Error ErrorContent `json:"error"`
}

// ErrorContent holds information on the error that occurred
type ErrorContent struct {
	StatusText string   `json:"status"`
	StatusCode int      `json:"code"`
	Messages   []string `json:"message"`
}
