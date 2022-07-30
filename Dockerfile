FROM golang:1.18.4-alpine as build
WORKDIR /app
COPY server ./server
COPY go.mod go.sum ./
RUN go get github.com/newrelic/go-agent/v3/newrelic@develop
RUN go build server/main.go

FROM alpine

WORKDIR /app

COPY --from=build /app/main /app

RUN \
  apk add --no-cache \
    curl \
    bash

COPY --chmod=0755 entrypoint.sh /
COPY --chmod=0755 tester.sh /app

ENTRYPOINT ["/entrypoint.sh"]
