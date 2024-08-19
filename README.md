# tucows-grill

## Overview
Welcome to the tucows grilling service where we grill cows of all types!

## Endpoints
The following endpoints are supported:  
`GET  /ingredients/{id}`  get an ingredient by the ingredient id  
`POST /ingredients`       add a new ingredient to the db  
`GET  /total-cost"`       get the total cost of all purchased items for an item id  
`GET  /total-cost-async`  same as total-cost, but uses batching and concurrency for large requests. Both total-cost endpoints will result in the same calculated value.

## Setup instructions
1. open docker  
2. run `docker compose up`  
3. for logging into the mysql cli, run `mysql -h 127.0.0.1 -P 3306 -u root -p`  
4. I used vscode, so the launch.json is already configured with a build. Run the `main.go` configuration.  

All requests are authenticated via JWT. See the `tucows-grill-client` repo for more details.  

## TODO:
- Unit tests  
- Validation middleware  
- Service layer (between handler and data layer). The client is just returning whatever the service is grabbing from the db. In a real world scenario, there would be business logic handled in another service that acts as a middleman between the handler and the service call itself.
- Improve credential security  
- Grill some burgers!  

## Debugging
When trying to debug locally, if you are getting the error message 
```
Failed to continue: "Cannot find Delve debugger. Install from https://github.com/go-delve/delve & ensure it is in your Go tools path, "GOPATH/bin" or "PATH"."
```
it may be a dependency conflict with dlv versions in vs code and your current go version (if you didn't have the latest go version). It may be a brew specific problem, but I ran into this on both my macs that were using brew and had an older version of Go installed (< 1.22). If this is the case, try running the following commands:
```
go uninstall github.com/go-delve/delve/cmd/dlv@latest
brew uninstall go
brew uninstall golangci-lint // only if required to uninstall go
brew install go
go install github.com/go-delve/delve/cmd/dlv@latest
```
