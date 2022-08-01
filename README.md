# Go CLM Demo

This demo app demonstrates several methods that generate code level
metrics which can be viewed in the [CodeStream Plugin](https://www.codestream.com/)

## Running

To run the app, simply set the environment variables for your 
ingest key and run the docker-compose command.

```shell
export NEW_RELIC_LICENSE_KEY=<your license key> && export NEW_RELIC_HOST=staging-collector.newrelic.com
docker-compose up --build -d
```

This will build and run the go demo app and it will also send traffic
to the app using the `tester.sh` script which runs simple curl commands.

Logs will be available in the `logs` subdirectory. 

## running locally

To run outside of Docker, first grab the develop branch of the 
newrelic golang agent. 

```shell
export NEW_RELIC_LICENSE_KEY=<your license key> && export NEW_RELIC_HOST=staging-collector.newrelic.com
go get github.com/newrelic/go-agent/v3/newrelic@develop
go run server/*
```
