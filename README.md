# go-octopusdeploy [![Build status](https://ci.appveyor.com/api/projects/status/5t5gbqjyl8hpou52?svg=true)](https://ci.appveyor.com/project/MattHodge/go-octopusdeploy)

A Go wrapper for the [Octopus Deploy](https://octopus.com/) REST API.

This exists to be used in the [Octopus Deploy Terraform Provider](https://github.com/OctopusDeploy/terraform-provider-octopusdeploy).

> :warning: The Octopus Deploy REST Client is in heavy development.

## Go Dependencies

* Dependencies are managed using [dep](https://golang.github.io/dep/docs/new-project.html)

```bash
# Vendor new modules
dep ensure
```

## Using the main.go Example

```bash
export OCTOPUS_URL=http://localhost:8081/
export OCTOPUS_APIKEY=API-FAKEAPIKEYFAKEAPIKEY

go run main.go # creates a project
```

## Contributing

💻 Want to help me out? Check out [CONTRIBUTING.md](CONTRIBUTING.md) and hit me up [@MattHodge](https://twitter.com/matthodge).
