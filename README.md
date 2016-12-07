

# sensorthings-cli
CLI for performing SensorThings tasks

![Demo](sti_demo.gif)

## Installation

### Go environment
``` 
$ go get github.com/geodan/sensorthings-cli/sti

$ sti
```

### Docker environment
``` 
$ docker run geodan/sti

with parameters:

$ docker run geodan/sti version
```


## Features

### Version

Get the cli tool version

```
$ sti version 
```

status: implemented

### Login

Login to the given SensorThings server

```
$ sti login http://gost.geodan.nl/v1.0
```
status: implemented


### Get

Get the SensorThings Things

status: partly implemented

```
$ sti get things 
```

Get the SensorThings Sensors

status: partly implemented

```
$ sti get sensors 
```

Get the SensorThings Datastreams

status: partly implemented

```
$ sti get datastreams 
```

Get the SensorThings ObservedProperties

status: partly implemented

```
$ sti get observedproperties 
```
Get the SensorThings Locations

status: partly implemented

```
$ sti get locations 
```
Get the SensorThings Observations

status: partly implemented

```
$ sti get observations 
```
Get the SensorThings HistoricalLocations

status: partly implemented

```
$ sti get historicallocations 
```
Get the SensorThings FeaturesOfInterest

status: partly implemented

```
$ sti get featuresofinterest 
```








### Create

Create a SensorThings Things

```
$ sti create thing 
```

status: todo

Create a SensorThings Sensor

```
$ sti create sensor
```

status: todo

Create a SensorThings Datastream

```
$ sti create datastream
```

status: todo

### Delete

Delete a SensorThings Things

```
$ sti delete thing {$name}
```
status: todo

Delete a SensorThings Sensor

```
$ sti delete sensor {$name}
```
status: todo

Delete a SensorThings Datastream

```
$ sti delete datastream {$name}
```
status: todo








