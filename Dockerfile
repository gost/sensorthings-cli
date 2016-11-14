FROM golang:latest
RUN apt-get update
RUN go get github.com/geodan/sensorthings-cli/cmd/sti
ENTRYPOINT ["sti"]