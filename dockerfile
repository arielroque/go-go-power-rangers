# Build image
FROM golang:1.21.1 AS builder

ENV GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on  go build -v -o main

ENTRYPOINT ["./main"]