# AGENTS.md - Dell Terraform Provider for PowerStore

## Project Overview

This is the Terraform provider for Dell PowerStore block and file storage arrays. It implements resources and data sources using HashiCorp's Terraform Plugin Framework, enabling infrastructure-as-code management of PowerStore arrays.

- **Language:** Go 1.25
- **Module path:** `terraform-provider-powerstore`
- **Terraform Plugin Framework:** v1.13.0
- **SDK:** `github.com/dell/gopowerstore` v1.18.0
- **Registry address:** `registry.terraform.io/dell/powerstore`
- **License:** Mozilla Public License 2.0

## Architecture

The provider follows the standard Terraform Plugin Framework architecture. It runs as a gRPC server that Terraform Core communicates with to manage PowerStore resources.

### Provider Configuration

The provider authenticates to a PowerStore array using endpoint, username, and password. Configuration can be supplied via HCL provider block or environment variables (`POWERSTORE_ENDPOINT`, `POWERSTORE_USERNAME`, `POWERSTORE_PASSWORD`, `POWERSTORE_INSECURE`, `POWERSTORE_TIMEOUT`).

### SDK Strategy

Uses `gopowerstore` — a public, versioned Go module on GitHub (`github.com/dell/gopowerstore`). The provider and SDK release independently. The provider also uses an OpenAPI-generated client (`clientgen/`) for additional API coverage.

### Resources and Data Sources

The provider exposes approximately 22 resources and 21 data sources covering PowerStore entities such as volumes, hosts, host groups, protection policies, snapshot rules, file systems, NAS servers, NFS exports, and more.

## Directory Structure

```
main.go                           Entry point (providerserver.Serve)
powerstore/
  provider.go                     Provider configuration, resource/datasource registration
  resource_*.go                   Resource implementations (Create, Read, Update, Delete)
  datasource_*.go                 Data source implementations
  helper/                         Shared helper functions
  customtypes/                    Custom Terraform type implementations
models/                           Terraform state model structs
  jsonmodel/                      JSON model structs
client/                           PowerStore API client wrappers
clientgen/                        OpenAPI-generated client code
clientgen_utils/                  OpenAPI spec and generation utilities
  openapi_specs/                  OpenAPI JSON specifications
examples/                         Example HCL configurations
  resources/                      Resource examples
  data-sources/                   Data source examples
  provider/                       Provider configuration example
docs/                             Generated documentation
templates/                        Documentation templates
tools/                            Build and generation tools
about/                            Provider metadata
```

## Build Commands

| Command | Description |
|---------|-------------|
| `make build` | Compile the provider binary |
| `make install` | Build and install to `~/.terraform.d/plugins/` |
| `make test` | Run formatting, linting, vetting, and unit tests |
| `make testacc` | Run acceptance tests (`TF_ACC=1`, requires live hardware) |
| `make check` | Run `terraform fmt`, `gofmt`, `golangci-lint`, `go vet` |
| `make gosec` | Run security scan with `gosec` |
| `make cover` | Generate HTML coverage report |
| `make generate` | Run `go generate` (docs generation) |
| `make build_client` | Regenerate OpenAPI client from spec |

## Testing

### Unit Tests (mockey)

- Test files follow `*_test.go` convention alongside source files in `powerstore/`.
- Frameworks: `github.com/stretchr/testify` (assertions), `github.com/bytedance/mockey` (function-level mocking).
- Run with `make test`.
- No hardware required.

### Acceptance Tests (terraform-plugin-testing)

- **Requires live PowerStore hardware** with credentials set via environment variables.
- Creates real resources — clean up after failures.
- Run with `make testacc` (sets `TF_ACC=1`).
- Tests use `resource.TestCase` with `ProtoV6ProviderFactories`.

### Running Tests

```bash
# Unit tests (no hardware)
make test

# Acceptance tests (requires live hardware)
export POWERSTORE_ENDPOINT="https://10.0.0.1/api/rest"
export POWERSTORE_USERNAME="admin"
export POWERSTORE_PASSWORD="secret"
export POWERSTORE_INSECURE="true"
make testacc
```

## Code Style and Conventions

### Formatting and Linting

- All Go code must pass `gofmt`, `go vet`, and `golangci-lint`.
- Terraform example files must pass `terraform fmt`.

### Code Organization Patterns

- **Resource pattern:** Each resource is a Go file (`resource_<name>.go`) with a struct implementing `resource.Resource` interface methods: `Metadata`, `Schema`, `Configure`, `Create`, `Read`, `Update`, `Delete`, and optionally `ImportState`.
- **Data source pattern:** Each data source is a Go file (`datasource_<name>.go`) implementing `datasource.DataSource`.
- **Models:** Terraform state structs are in `models/` using `tfsdk` struct tags.
- **Helpers:** Shared mapping functions between API types and Terraform models in `powerstore/helper/`.
- **Client layer:** `client/` wraps the `gopowerstore` SDK; `clientgen/` provides OpenAPI-generated methods.

### File Header

All source files must include the Dell copyright and MPL 2.0 license header:

```go
/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
...
*/
```

## Common Development Tasks

### Adding a New Resource

1. Create `powerstore/resource_<name>.go` implementing the `resource.Resource` interface.
2. Create the corresponding model struct in `models/`.
3. Add helper functions in `powerstore/helper/` for mapping between API and Terraform types.
4. Register the resource in `powerstore/provider.go` `Resources()` method.
5. Add unit tests in `powerstore/resource_<name>_test.go` using mockey mocks.
6. Add acceptance tests that exercise the full CRUD lifecycle.
7. Create example HCL in `examples/resources/powerstore_<name>/`.
8. Run `make generate` to produce documentation.

### Adding a New Data Source

1. Create `powerstore/datasource_<name>.go` implementing `datasource.DataSource`.
2. Create the model struct in `models/`.
3. Register in `powerstore/provider.go` `DataSources()` method.
4. Add tests and examples following the same patterns as resources.

### Updating the SDK

```bash
go get github.com/dell/gopowerstore@<version>
go mod tidy
```

### Regenerating the OpenAPI Client

```bash
make build_client
```

Requires the OpenAPI Generator CLI JAR and a filtered spec in `clientgen_utils/openapi_specs/`.

## CI/CD

GitHub Actions workflows in `.github/workflows/`. GoReleaser configuration in `.goreleaser.yaml` builds cross-platform binaries (linux, darwin, windows, freebsd; amd64, arm64, 386, arm).

## Code Ownership

All files are owned by the maintainers defined in `.github/CODEOWNERS`.
