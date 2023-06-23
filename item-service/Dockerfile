# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.20-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o item-service

##
## Deploy
##
FROM debian:12-slim

WORKDIR /

COPY --from=build app/item-service .

EXPOSE 8080

ENTRYPOINT ["/item-service"]