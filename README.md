<div align="center">
  <img alt="go-octopusdeploy Logo" src="https://user-images.githubusercontent.com/71493/133961475-fd4d769f-dc32-4723-a9bd-5529c5b12faf.png" height="140" />
  <h3 align="center">go-octopusdeploy</h3>
  <p align="center">Go API Client for <a href="https://octopus.com/">Octopus Deploy</a> 🐙</p>
  <p align="center">
    <a href="https://github.com/OctopusDeploy/go-octopusdeploy/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/OctopusDeploy/go-octopusdeploy.svg?logo=github&style=flat-square"></a>
    <a href="https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy"><img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/OctopusDeploy/go-octopusdeploy"></a>
    <a href="https://goreportcard.com/badge/github.com/OctopusDeploy/go-octopusdeploy"><img src="https://goreportcard.com/badge/github.com/OctopusDeploy/go-octopusdeploy" alt="Go Report"></a>
  </p>
</div>

---

## Install

```bash
go get "github.com/OctopusDeploy/go-octopusdeploy/v2"
```

## Usage

```go
import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
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
octopusClient, err := client.NewClient(nil, apiURL, apiKey, spaceID)
if err != nil {
    _ = fmt.Errorf("error creating API client: %v", err)
    return
}
```

Once the client has been initialized, APIs can be targeted through the model and services available:

```go
usernamePasswordAccount := accounts.NewUsernamePasswordAccount(name)
usernamePasswordAccount.Username = username

createdAccount, err := accounts.Add(octopusClient, usernamePasswordAccount)
if err != nil {
    _ = fmt.Errorf("error adding account: %v", err)
}
```

Operations like `Add`, `DeleteByID`, `GetByID`, and `Update` are supported by most services that are exposed through the client if not exposed in the package. These operations are configured at runtime since the Octopus REST API is hypermedia-driven.

Numerous code samples that showcase the API and this client are available in the [examples](/examples) directory. There are also many [integration](/test) and unit tests available to examine that demonstrate the capabilities of this API client.

## Testing

### Running Tests Locally using Visual Studio Code

> [!CAUTION]
> The integration tests will create and delete resources on your instance. Use a dedicated instance to run tests

To run the tests against an instance of Octopus Deploy, create a `.env` file at the root directory of this repository

```
OCTOPUS_HOST=http://your-octopus-instance-url
OCTOPUS_API_KEY=API-YOURAPIKEY
```

and add a Visual Studio Code workspace setting (`.vscode/settings.json`) for the test environment settings

```
{
    "go.testEnvFile": "${workspaceFolder}/.env"
}
```

## 🤝 Contributions

Contributions are welcome! :heart: Please read our [Contributing Guide](CONTRIBUTING.md) for information about how to get involved in this project.
