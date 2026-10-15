<div align="center">
  <img alt="go-octopusdeploy Logo" src="https://user-images.githubusercontent.com/71493/133961475-fd4d769f-dc32-4723-a9bd-5529c5b12faf.png" height="140" />
  <h3 align="center">go-octopusdeploy</h3>
  <p align="center">Go API Client for <a href="https://octopus.com/">Octopus Deploy</a> 🐙</p>
  <p align="center">
    <a href="https://github.com/OctopusDeploy/go-octopusdeploy/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/OctopusDeploy/go-octopusdeploy.svg?logo=github&style=flat-square"></a>
    <a href="https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy"><img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/OctopusDeploy/go-octopusdeploy"></a>
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

### Rate limiting

Octopus Server can [limit the rate of requests](https://octopus.com/docs/administration/managing-infrastructure/rate-limiting) from each client. When a client sends too many requests, the server returns `429 Too Many Requests` and does not process the request.

By default, the client does not handle rate limits. To make the client obey the rate limits, use `NewClientWithOptions` and set `RateLimit`:

```go
import (
    "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
    "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/ratelimit"
)

credentials, err := client.NewApiKey(apiKey)
if err != nil {
    _ = fmt.Errorf("error creating API key credentials: %v", err)
    return
}

rateLimitOptions := ratelimit.DefaultOptions()

octopusClient, err := client.NewClientWithOptions(nil, apiURL, credentials, spaceID, "", &client.ClientOptions{
    RateLimit: &rateLimitOptions,
})
if err != nil {
    _ = fmt.Errorf("error creating API client: %v", err)
    return
}
```

With rate limiting on, the client does these things:

- When the server returns `429 Too Many Requests`, the client waits and sends the request again. It waits for the time in the `Retry-After` header. If the server does not send this header, it uses exponential backoff.
- When the `Octopus-RateLimit` advisory header shows that no requests remain, the client waits before it sends the next request. This prevents most `429` responses.

You can change these settings in `ratelimit.Options`:

| Setting              | Default    | Description                                                                                                                                         |
| -------------------- | ---------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `MaxRetries`         | `3`        | The maximum number of times to send a request again after a `429` response. Set to `0` to turn off retries.                                         |
| `MaxWait`            | `60s`      | The maximum time to wait before a request. If `Retry-After` is longer than this time, the client does not retry and returns the `429` response.     |
| `UseAdvisoryHeaders` | `true`     | Wait before a request when the advisory headers show that no requests remain.                                                                       |

The client retries a request only if it can send the request body again. For example, the client does not retry a package upload.

If you supply an `http.Client`, the client uses a copy of it and does not change it. To use rate limiting with an `http.Client` of your own in other code, wrap its transport with `ratelimit.NewTransport`.

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

### Linting and formatting

We use [golangci-lint](https://golangci-lint.run) to lint and format the code. The configuration is in `.golangci.yml`.
The pull request workflow runs golangci-lint, and your pull request must pass it.

See the [golangci-lint website](https://golangci-lint.run) for the install instructions.

To format the code, and then find lint issues, run these commands in the root folder:

```shell
golangci-lint fmt ./...
golangci-lint run ./...
```

If you have make, you can run make fmt and make lint instead.

## Release

Once your pull request has been approved:
1. **Squash** (not merge) your PR with a [conventional commit]([url](https://www.conventionalcommits.org/en/v1.0.0/#summary)) message
   - Merged PR's will still make it into the release notes, but the release notes will also be polluted with the  merge commit (see [example]([url](https://github.com/OctopusDeploy/go-octopusdeploy/releases/tag/v2.107.0)))
3. Tag the commit with the next version number - the [release workflow]([url](https://github.com/OctopusDeploy/go-octopusdeploy/actions/workflows/release.yml)) then trigger to create the release in GitHub
   - Since this is a GO library, it doesn't need to be pushed anywhere else to be shared

## 🤝 Contributions

Contributions are welcome! :heart: Please read our [Contributing Guide](CONTRIBUTING.md) for information about how to get involved in this project.
