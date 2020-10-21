# go-octopusdeploy

[![PkgGoDev](https://pkg.go.dev/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy) ![Run Integration Tests](https://github.com/OctopusDeploy/go-octopusdeploy/workflows/Run%20Integration%20Tests/badge.svg?branch=beta-candidate-01) [![Go Report](https://goreportcard.com/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://goreportcard.com/report/github.com/OctopusDeploy/go-octopusdeploy)

A Go client for the [Octopus Deploy](https://octopus.com/) API. This client is used by the [Octopus Deploy Terraform Provider](https://github.com/OctopusDeploy/terraform-provider-octopusdeploy).

## Usage

```go
import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
```

The [Octopus REST API](https://octopus.com/docs/octopus-rest-api) is exposed through service fields of the client. An API key is required to communicate with the API (see [How to Create an API Key](https://octopus.com/docs/octopus-rest-api/how-to-create-an-api-key) for more information).

```go
apiKey := "API-YOUR_API_KEY"
octopusURL := "https://your_octopus_url"
spaceID := "space-id" // can also be blank to assume the default space

apiURL, err := url.Parse(octopusURL)
if err != nil {
    _ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
    return
}

// the first parameter for NewClient can accept a http.Client if you wish to
// override the default; also, the spaceID may be an empty string (i.e. "") if
// you wish to load the default space
client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
if err != nil {
    _ = fmt.Errorf("error creating API client: %v", err)
    return
}
```

Once the client has been initialized, APIs can be targeted through the model and services available:

```go
usernamePasswordAccount := octopusdeploy.NewUsernamePasswordAccount(name)
usernamePasswordAccount.Username = username

createdAccount, err := client.Accounts.Add(usernamePasswordAccount)
if err != nil {
    _ = fmt.Errorf("error adding account: %v", err)
}
```

Operations like `Add`, `DeleteByID`, `GetByID`, and `Update` are supported by most services that are exposed through the client. These operations are configured at runtime since the Octopus REST API is hypermedia-driven.

Numerous code samples that showcase the API and this client are available in the [examples](/examples) directory. There are also many [integration](/integration) and unit tests available to examine that demonstrate the capabilities of this API client.