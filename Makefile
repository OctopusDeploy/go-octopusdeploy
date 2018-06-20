TEST?=$$(go list ./... |grep -v 'vendor')

default: build test

fmt:
	go fmt github.com/MattHodge/go-octopusdeploy/...

build: fmt
	go build

test: fmt
	go test -v -timeout 30s ./...

testacc:
	go test integration/* -v -timeout 90m
