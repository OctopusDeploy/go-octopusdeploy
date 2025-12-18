# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-octopusdeploy is the official Go API client library for [Octopus Deploy](https://octopus.com/), providing programmatic access to the Octopus REST API. The library is hypermedia-driven, meaning API operations are configured at runtime based on Octopus API responses.

**Module path:** `github.com/OctopusDeploy/go-octopusdeploy/v2`

## Common Commands

```bash
# Build
go build -a -race -v ./...

# Run all tests
go test -v ./...

# Run a single test
go test -v ./pkg/accounts -run TestAccountServiceAdd

# Run tests in a specific package
go test -v ./pkg/accounts/...

# Vet
go vet -v ./...
```

## Integration Tests

Integration tests require a live Octopus Deploy instance. Create a `.env` file at the repo root:
```
OCTOPUS_HOST=http://your-octopus-instance-url
OCTOPUS_API_KEY=API-YOURAPIKEY
```

For VS Code, add to `.vscode/settings.json`:
```json
{
    "go.testEnvFile": "${workspaceFolder}/.env"
}
```

## Architecture

### Client Structure (`pkg/client/octopusdeploy.go`)
The central `Client` struct aggregates 70+ domain-specific services (e.g., `client.Accounts`, `client.Projects`, `client.Deployments`). Initialize with:
```go
client, err := client.NewClient(nil, apiURL, apiKey, spaceID)
```

### Service Pattern (`pkg/services/service.go`)
All API services implement the `IService` interface with standardized CRUD operations:
- `Add`, `GetByID`, `Update`, `DeleteByID`
- Services use URI templates for hypermedia-driven API navigation

### Package Organization (`pkg/`)
- `pkg/client/` - Main client entry point
- `pkg/services/` - Base service infrastructure
- `pkg/resources/` - Generic response wrappers with pagination (`Resources[T]`)
- `pkg/constants/` - Service names, operations, URI templates
- `pkg/[domain]/` - Domain-specific packages (accounts, projects, deployments, etc.)

### Error Handling (`internal/errors.go`)
Use factory functions for consistent errors:
- `internal.CreateInvalidParameterError(operation, parameter)`
- `internal.CreateRequiredParameterIsEmptyOrNilError(parameter)`

## Development Patterns

### Enums
When adding enum values, regenerate string representations with `enumer`:
```bash
go install github.com/dmarkham/enumer@latest
# From the package directory containing the enum:
enumer -type=FilterType -json -output filter_type_string.go
```

### Testing Patterns
Unit tests use `stretchr/testify/require` for assertions. Service tests typically follow this structure:
```go
func createAccountService(t *testing.T) *AccountService {
    service := NewAccountService(nil, constants.TestURIAccounts)
    require.NotNil(t, service)
    return service
}

func TestAccountServiceGetByID(t *testing.T) {
    service := createAccountService(t)
    resource, err := service.GetByID("")
    require.Equal(t, internal.CreateInvalidParameterError(...), err)
    require.Nil(t, resource)
}
```

## Commit Guidelines

Use [Conventional Commits](https://www.conventionalcommits.org/): `feat:`, `fix:`, `refactor:`, `test:`, `docs:`

## Releasing

Create a git tag in format `v[major].[minor].[patch]` (e.g., `v1.0.0`). Release automation via goreleaser triggers on tag creation.