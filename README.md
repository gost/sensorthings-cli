# sensorthings-cli
CLI for performing SensorThings tasks

## Installation

### Go environment
``` 
$ go get github.com/geodan/sensorthings-cli/cmd/sti

$ sti
```

### Docker environment
``` 
$ docker run geodan/sensorthings-cli

```

## Features

### Version

Get the cli tool version

```
$ sti version 
```
### Login

Login to the given SensorThings server

```
$ sti login http://gost.geodan.nl/v1.0
```

### Get

Get the SensorThings Things

```
$ sti get things 
```
Get the SensorThings Sensors

```
$ sti get sensors 
```

Get the SensorThings Datastreams

```
$ sti get datastreams 
```

### Create

Create a SensorThings Things

```
$ sti create thing 
```

Create a SensorThings Sensor

```
$ sti create sensor
```

Create a SensorThings Datastream

```
$ sti create datastream
```


### Delete

Delete a SensorThings Things

```
$ sti delete thing {$name}
```

Delete a SensorThings Sensor

```
$ sti delete sensor {$name}
```

Delete a SensorThings Datastream

```
$ sti delete datastream {$name}
```








