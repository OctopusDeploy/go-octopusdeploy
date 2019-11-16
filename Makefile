default: build test

fmt:
	go fmt github.com/OctopusDeploy/go-octopusdeploy/...

build: enums fmt
	go build

test: fmt
	go test ./octopusdeploy/... -timeout 30s

testacc:
	go test ./integration/... -timeout 90m
