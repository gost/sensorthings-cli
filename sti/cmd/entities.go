package cmd

import (
	"fmt"
	"strings"
)

type Entity interface {
	ClearNav()
	GetID() interface{}
}

// BaseEntity is the entry point for an entity
type BaseEntity struct {
	ID      interface{} `json:"@iot.id,omitempty"`
	NavSelf string      `json:"@iot.selfLink,omitempty"`
}

func (b BaseEntity) ClearNav() {
	b.NavSelf = ""
}

func (b BaseEntity) GetID() interface{} {
	return b.ID
}

// Thing structure
type Thing struct {
	BaseEntity
	Name                   string                 `json:"name,omitempty"`
	Description            string                 `json:"description,omitempty"`
	Properties             map[string]interface{} `json:"properties,omitempty"`
	NavLocations           string                 `json:"Locations@iot.navigationLink,omitempty"`
	NavDatastreams         string                 `json:"Datastreams@iot.navigationLink,omitempty"`
	NavHistoricalLocations string                 `json:"HistoricalLocations@iot.navigationLink,omitempty"`
	Locations              []*Location            `json:"Locations,omitempty"`
	Datastreams            []*Datastream          `json:"Datastreams,omitempty"`
	HistoricalLocations    []*HistoricalLocation  `json:"HistoricalLocations,omitempty"`
}

func (t *Thing) ClearNav() {
	t.NavSelf = ""
	t.NavLocations = ""
	t.NavDatastreams = ""
	t.NavHistoricalLocations = ""
}

type Datastream struct {
	BaseEntity
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	UnitOfMeasurement   map[string]interface{} `json:"unitOfMeasurement,omitempty"`
	ObservationType     string                 `json:"observationType,omitempty"`
	ObservedArea        map[string]interface{} `json:"observedArea,omitempty"`
	NavThing            string                 `json:"Thing@iot.navigationLink,omitempty"`
	NavSensor           string                 `json:"Sensor@iot.navigationLink,omitempty"`
	NavObservations     string                 `json:"Observations@iot.navigationLink,omitempty"`
	NavObservedProperty string                 `json:"ObservedProperty@iot.navigationLink,omitempty"`
	Thing               *Thing                 `json:"Thing,omitempty"`
	Sensor              *Sensor                `json:"Sensor,omitempty"`
	Observations        []*Observation         `json:"Observations,omitempty"`
	ObservedProperty    *ObservedProperty      `json:"ObservedProperty,omitempty"`
	PhenomenonTime      string                 `json:"phenomenonTime,omitempty"`
	ResultTime          string                 `json:"resultTime,omitempty"`
}

func (d *Datastream) ClearNav() {
	d.NavSelf = ""
	d.NavThing = ""
	d.NavSensor = ""
	d.NavObservations = ""
	d.NavObservedProperty = ""
}

type Observation struct {
	BaseEntity
	PhenomenonTime       string                 `json:"phenomenonTime,omitempty"`
	Result               interface{}            `json:"result,omitempty"`
	ResultTime           string                 `json:"resultTime,omitempty"`
	ResultQuality        string                 `json:"resultQuality,omitempty"`
	ValidTime            string                 `json:"validTime,omitempty"`
	Parameters           map[string]interface{} `json:"parameters,omitempty"`
	NavDatastream        string                 `json:"Datastream@iot.navigationLink,omitempty"`
	NavFeatureOfInterest string                 `json:"FeatureOfInterest@iot.navigationLink,omitempty"`
	Datastream           *Datastream            `json:"Datastream,omitempty"`
	FeatureOfInterest    *FeatureOfInterest     `json:"FeatureOfInterest,omitempty"`
}

func (o *Observation) ClearNav() {
	o.NavSelf = ""
	o.NavDatastream = ""
	o.NavFeatureOfInterest = ""
}

type FeatureOfInterest struct {
	BaseEntity
	Name               string                 `json:"name,omitempty"`
	Description        string                 `json:"description,omitempty"`
	EncodingType       string                 `json:"encodingType,omitempty"`
	Feature            map[string]interface{} `json:"feature,omitempty"`
	NavObservations    string                 `json:"Observations@iot.navigationLink,omitempty"`
	Observations       []*Observation         `json:"Observations,omitempty"`
	OriginalLocationID interface{}            `json:"-"`
}

func (f *FeatureOfInterest) ClearNav() {
	f.NavSelf = ""
	f.NavObservations = ""
}

type HistoricalLocation struct {
	BaseEntity
	Time         string      `json:"time,omitempty"`
	NavThing     string      `json:"Thing@iot.navigationLink,omitempty"`
	NavLocations string      `json:"Locations@iot.navigationLink,omitempty"`
	Thing        *Thing      `json:"Thing,omitempty"`
	Locations    []*Location `json:"Locations,omitempty"`
}

func (h *HistoricalLocation) ClearNav() {
	h.NavSelf = ""
	h.NavThing = ""
	h.NavLocations = ""
}

type Location struct {
	BaseEntity
	Name                   string                 `json:"name,omitempty"`
	Description            string                 `json:"description,omitempty"`
	EncodingType           string                 `json:"encodingType,omitempty"`
	Location               map[string]interface{} `json:"location,omitempty"`
	NavThings              string                 `json:"Things@iot.navigationLink,omitempty"`
	NavHistoricalLocations string                 `json:"HistoricalLocations@iot.navigationLink,omitempty"`
	Things                 []*Thing               `json:"Things,omitempty"`
	HistoricalLocations    []*HistoricalLocation  `json:"HistoricalLocations,omitempty"`
}

func (l *Location) ClearNav() {
	l.NavSelf = ""
	l.NavHistoricalLocations = ""
	l.NavThings = ""
}

type ObservedProperty struct {
	BaseEntity
	Name           string        `json:"name,omitempty"`
	Description    string        `json:"description,omitempty"`
	Definition     string        `json:"definition,omitempty"`
	NavDatastreams string        `json:"Datastreams@iot.navigationLink,omitempty"`
	Datastreams    []*Datastream `json:"Datastreams,omitempty"`
}

func (o *ObservedProperty) ClearNav() {
	o.NavSelf = ""
	o.NavDatastreams = ""
}

type Sensor struct {
	BaseEntity
	Name           string        `json:"name,omitempty"`
	Description    string        `json:"description,omitempty"`
	EncodingType   string        `json:"encodingType,omitempty"`
	Metadata       string        `json:"metadata,omitempty"`
	NavDatastreams string        `json:"Datastreams@iot.navigationLink,omitempty"`
	Datastreams    []*Datastream `json:"Datastreams,omitempty"`
}

func (s *Sensor) ClearNav() {
	s.NavSelf = ""
	s.NavDatastreams = ""
}

// EntityType holds the name and type of a SensorThings entity.
type EntityType string

// List of all EntityTypes.
const (
	EntityTypeThing              EntityType = "Thing"
	EntityTypeLocation           EntityType = "Location"
	EntityTypeHistoricalLocation EntityType = "HistoricalLocation"
	EntityTypeDatastream         EntityType = "Datastream"
	EntityTypeSensor             EntityType = "Sensor"
	EntityTypeObservedProperty   EntityType = "ObservedProperty"
	EntityTypeObservation        EntityType = "Observation"
	EntityTypeFeatureOfInterest  EntityType = "FeatureOfInterest"
	EntityTypeUnknown            EntityType = "Unknown"
)

// String returns the string representation of the EntityType.
func (e EntityType) String() string {
	return fmt.Sprintf("%s", e)
}

// GetEndpoint returns the single entity endpoint name
func (e EntityType) GetEndpoint() string {
	switch e {
	case EntityTypeThing:
		return "Thing"
	case EntityTypeLocation:
		return "Location"
	case EntityTypeHistoricalLocation:
		return "HistoricalLocation"
	case EntityTypeDatastream:
		return "Datastream"
	case EntityTypeSensor:
		return "Sensor"
	case EntityTypeObservedProperty:
		return "ObservedProperty"
	case EntityTypeFeatureOfInterest:
		return "FeatureOfInterest"
	}

	return ""
}

// GetArrayEndpoint returns the (array) endpoint name for the current EntityType
func (e EntityType) GetArrayEndpoint() string {
	switch e {
	case EntityTypeThing:
		return "Things"
	case EntityTypeLocation:
		return "Locations"
	case EntityTypeHistoricalLocation:
		return "HistoricalLocations"
	case EntityTypeDatastream:
		return "Datastreams"
	case EntityTypeSensor:
		return "Sensors"
	case EntityTypeObservedProperty:
		return "ObservedProperties"
	case EntityTypeFeatureOfInterest:
		return "FeaturesOfInterest"
	}

	return ""
}

// StringEntityMap is a map of strings that map a string to an EntityType
var StringEntityMap = map[string]EntityType{
	"thing": EntityTypeThing, "things": EntityTypeThing,
	"location": EntityTypeLocation, "locations": EntityTypeLocation,
	"historicallocation": EntityTypeHistoricalLocation, "historicallocations": EntityTypeHistoricalLocation,
	"datastream": EntityTypeDatastream, "datastreams": EntityTypeDatastream,
	"sensor": EntityTypeSensor, "sensors": EntityTypeSensor,
	"observedproperty": EntityTypeObservedProperty, "observedproperties": EntityTypeObservedProperty,
	"observation": EntityTypeObservation, "observations": EntityTypeObservation,
	"featureofinterest": EntityTypeFeatureOfInterest, "featuresofinterest": EntityTypeFeatureOfInterest,
}

// EntityTypeFromString returns the EntityType for a given string
// function is case-insensitive
func EntityTypeFromString(e string) (EntityType, error) {
	val, ok := StringEntityMap[strings.ToLower(e)]
	if !ok {
		return EntityTypeUnknown, fmt.Errorf("Unknown entity %s", e)
	}

	return val, nil
}
