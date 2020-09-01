# go-octopusdeploy [![Build status](https://ci.appveyor.com/api/projects/status/5t5gbqjyl8hpou52?svg=true)](https://ci.appveyor.com/project/MattHodge/go-octopusdeploy) ![Build Status](https://github.com/OctopusDeploy/go-octopusdeploy/workflows/Go/badge.svg)

A Go wrapper for the [Octopus Deploy](https://octopus.com/) REST API.

This exists to be used in the [Octopus Deploy Terraform Provider](https://github.com/OctopusDeploy/terraform-provider-octopusdeploy).

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
