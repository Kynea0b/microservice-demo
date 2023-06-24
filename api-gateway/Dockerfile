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

RUN go build -o api-gateway

##
## Deploy
##
FROM debian:12-slim

WORKDIR /

COPY --from=build app/api-gateway .

EXPOSE 8080

ENTRYPOINT ["/api-gateway"]