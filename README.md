# go-octopusdeploy

[![PkgGoDev](https://pkg.go.dev/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy) ![Run Integration Tests](https://github.com/OctopusDeploy/go-octopusdeploy/workflows/Run%20Integration%20Tests/badge.svg?branch=beta-candidate-01) [![Go Report](https://goreportcard.com/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://goreportcard.com/report/github.com/OctopusDeploy/go-octopusdeploy)

A Go client for the [Octopus Deploy](https://octopus.com/) API. This client is used by the [Octopus Deploy Terraform Provider](https://github.com/OctopusDeploy/terraform-provider-octopusdeploy).

## Usage

```go
import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
```

APIs are exposed as services and made available through the client:

```go
apiKey := "API-YOUR_API_KEY"
octopusURL := "https://your_octopus_url"
spaceID := "space-id" // can also be blank to assume the default space

apiURL, err := url.Parse(octopusURL)
if err != nil {
    _ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
    return
}

// first parameter is http.Client if you wish to override the default
client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
if err != nil {
    _ = fmt.Errorf("error creating API client: %v", err)
    return
}
```

Once the client has been initialized, APIs can be targeted through the model
and services available:

```go
usernamePasswordAccount := octopusdeploy.NewUsernamePasswordAccount(name)
usernamePasswordAccount.Username = username

createdAccount, err := client.Accounts.Add(usernamePasswordAccount)
if err != nil {
    _ = fmt.Errorf("error adding account: %v", err)
}
```

Operations like `Add`, `DeleteByID`, `GetByID`, and `Update` are supported by
most services exposed through the client. Numerous code samples that showcase
the API and this client are available in the [examples](/examples) directory.