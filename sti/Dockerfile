FROM golang:latest
RUN apt-get update
RUN mkdir -p /go/src/github.com/geodan/sensorthings-cli/sti
ADD . /go/src/github.com/geodan/sensorthings-cli/sti
WORKDIR /go/src/github.com/geodan/sensorthings-cli/sti
RUN go get .
RUN go build -o /go/bin/sti github.com/geodan/sensorthings-cli/sti
RUN sti login http://gost.geodan.nl/v1.0
ENTRYPOINT ["sti"]