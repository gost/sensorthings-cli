

# sensorthings-cli
CLI for performing SensorThings tasks

## Installation

### Go environment
``` 
$ go get github.com/geodan/sensorthings-cli/sti

$ sti
```

### Docker environment
``` 
$ docker run geodan/sensorthings-cli

with parameters:

$ docker run geodan/sensorthings-cli version
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

```
$ sti get things 
```
Get the SensorThings Sensors

status: todo

```
$ sti get sensors 
```

Get the SensorThings Datastreams

status: todo

```
$ sti get datastreams 
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








