# go-octopusdeploy

https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy

[![PkgGoDev](https://pkg.go.dev/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy) ![Run Integration Tests](https://github.com/OctopusDeploy/go-octopusdeploy/workflows/Run%20Integration%20Tests/badge.svg?branch=beta-candidate-01) [![Go Report](https://goreportcard.com/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://goreportcard.com/report/github.com/OctopusDeploy/go-octopusdeploy)

A Go wrapper for the [Octopus Deploy](https://octopus.com/) REST API.

This client is used by the [Octopus Deploy Terraform Provider](https://github.com/OctopusDeploy/terraform-provider-octopusdeploy).

## Go Dependencies

* Dependencies are managed using [Go Modules](https://github.com/golang/go/wiki/Modules#daily-workflow)

## Using the main.go Example

```bash
export OCTOPUS_URL=http://localhost:8081/
export OCTOPUS_APIKEY=API-FAKEAPIKEYFAKEAPIKEY

go run main.go # creates a project
```

## Contributing

💻 Want to out? Check out [CONTRIBUTING.md](CONTRIBUTING.md).
