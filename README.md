# go-octopusdeploy

[![PkgGoDev](https://pkg.go.dev/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy) ![Run Integration Tests](https://github.com/OctopusDeploy/go-octopusdeploy/workflows/Run%20Integration%20Tests/badge.svg?branch=beta-candidate-01) [![Go Report](https://goreportcard.com/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://goreportcard.com/report/github.com/OctopusDeploy/go-octopusdeploy)

A Go client for the [Octopus Deploy](https://octopus.com/) API.

This client is used by the [Octopus Deploy Terraform Provider](https://github.com/OctopusDeploy/terraform-provider-octopusdeploy).

## Supported Octopus Types and API Operations

| Service       | Operations     |
| :- | :- |
| Accounts | `Add`, `DeleteByID` `GetAll`, `GetByID`, `GetByIDs`, `GetByAccountType`, `GetByName`, `GetByPartialName`, `GetUsages`, `Update` |
| Action Templates | `Add`, `DeleteByID`, `GetAll`, `GetCategories`, `GetByID`, `GetByName`, `Search`, `Update` |
| Channels | `Add`, `DeleteByID`, `GetAll`, `GetByID`, `GetByPartialName`, `GetProject`, `GetReleases`, `Update` |
| Feeds | `Add`, `DeleteByID`, `GetAll`, `GetByID`, `GetByPartialName`, `Update` |

## Contributing

💻 Want to out? Check out [CONTRIBUTING.md](CONTRIBUTING.md).
