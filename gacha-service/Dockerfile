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

RUN go build -o gacha-service

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build app/gacha-service .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/gacha-service"]