TEST?=$$(go list ./... |grep -v 'vendor')

default: build test

fmt:
	go fmt github.com/OctopusDeploy/go-octopusdeploy/...

build: fmt
	go build

test: fmt
	go test -v -timeout 30s ./...

testacc:
	go test integration/*.go -v -timeout 90m
