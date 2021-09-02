FROM golang:1.17.0-alpine

ENV CGO_ENABLED=1
ENV GO117MODULE=on
RUN apk add --no-cache git gcc g++

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./cmd/api/main.go