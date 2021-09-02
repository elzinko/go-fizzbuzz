FROM golang:1.17.0-alpine AS builder
ENV CGO_ENABLED=1
ENV GO117MODULE=on
RUN apk add --no-cache git gcc g++
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/app ./cmd/api/main.go

FROM alpine:3.14.2
RUN apk add ca-certificates
WORKDIR /app
COPY --from=builder /src/out/app /app/fuzzbuzz
COPY --from=builder /src/data /app/data
RUN chmod +x fuzzbuzz
RUN mkdir log
EXPOSE 3333
ENTRYPOINT ./fuzzbuzz