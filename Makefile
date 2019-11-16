default: build test

fmt:
	go fmt github.com/OctopusDeploy/go-octopusdeploy/...

build: enums fmt
	go build

test: fmt
	go test ./octopusdeploy/... -timeout 30s

testacc:
	go test ./integration/... -timeout 90m

enums:
	# add additional files to be enum-ified with addtional -f <file-path> args
	go-enum --noprefix --lower --marshal --names \
		-f octopusdeploy/tenanted_deployment_mode.go \
		-f octopusdeploy/account_type.go
