FROM golang:1.23-alpine3.19

WORKDIR /go/src/Glimmer

COPY . /go/src/Glimmer/

RUN go build -o admin ./main.go

EXPOSE 8080
ENTRYPOINT ./admin
