# go-octopusdeploy

[![PkgGoDev](https://pkg.go.dev/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://pkg.go.dev/github.com/OctopusDeploy/go-octopusdeploy) ![Run Integration Tests](https://github.com/OctopusDeploy/go-octopusdeploy/workflows/Run%20Integration%20Tests/badge.svg?branch=beta-candidate-01) [![Go Report](https://goreportcard.com/badge/github.com/OctopusDeploy/go-octopusdeploy)](https://goreportcard.com/report/github.com/OctopusDeploy/go-octopusdeploy)

A Go client for the [Octopus Deploy](https://octopus.com/) API. This client is used by the [Octopus Deploy Terraform Provider](https://github.com/OctopusDeploy/terraform-provider-octopusdeploy).

## Usage

```
import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
```

APIs are exposed as services and made available through the client:

```
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

Once the client has been initialized (as above), APIs can be made:

```
usernamePasswordAccount := octopusdeploy.NewUsernamePasswordAccount(name)
usernamePasswordAccount.Username = username

createdAccount, err := client.Accounts.Add(usernamePasswordAccount)
if err != nil {
    _ = fmt.Errorf("error adding account: %v", err)
}
```

A variety of code samples that showcase the API and this client are available in the [examples](/examples) directory.

## Supported Types and Operations

| Service | Types | Operations |
| :- | :- | :- |
| Accounts | `AmazonWebServicesAccount`<br>`AzureServicePrincipalAccount`<br>`AzureSubscriptionAccount`<br>`SSHKeyAccount`<br>`TokenAccount`<br>`UsernamePasswordAccount` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`GetByIDs`<br>`GetByAccountType`<br>`GetByName`<br>`GetByPartialName`<br>`GetUsages`<br>`Update` |
| Action Templates | `ActionTemplate`<br>`ActionTemplateCategory`<br>`ActionTemplateParameter`<br>`ActionTemplateSearch` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetCategories`<br>`GetByID`<br>`GetByName`<br>`Search`<br>`Update` |
| Channels | `Channel` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`GetByPartialName`<br>`GetProject`<br>`GetReleases`<br>`Update` |
| Feeds | `AwsElasticContainerRegistry`<br>`BuiltInFeed`<br>`DockerContainerRegistry`<br>`GitHubRepositoryFeed`<br>`HelmFeed`<br>`MavenFeed`<br>`NuGetFeed`<br>`OctopusProjectFeed` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`GetByName`<br>`GetByPartialName`<br>`Update` |
| Machines | `DeploymentTarget`<br>`MachinePolicy` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`Update` |
| Runbooks | `Runbook` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`Update` |
| Spaces | `Space` | `Add` |
| Teams | `Team` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`GetByPartialName`<br>`Update` |
| Users | `User`<br>`UserAuthentication` | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetAuthentication`<br>`GetAuthenticationForUser`<br>`GetByID`<br>`GetMe`<br>`GetSpaces`<br>`Update` |
| Worker Pools | * | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`GetByIDs`<br>`GetByName`<br>`GetByPartialName`<br>`Update` |
| Workers | * | `Add`<br>`DeleteByID`<br>`GetAll`<br>`GetByID`<br>`GetByIDs`<br>`GetByName`<br>`GetByPartialName`<br>`Update` |

## Contributing

💻 Want to out? Check out [CONTRIBUTING.md](CONTRIBUTING.md).
