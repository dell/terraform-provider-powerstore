# Architecture: terraform-provider-powerstore

## Metadata

<!-- yaml-metadata-start -->
scope_paths: ["./"]
capture_git_sha: "3d3c6ee5f4a1eeeedb5f2dc055b4d74d0e7a7c06"
status: "current"
auto_update: false
preview_before_apply: true
scaffold_version: "1.0"
<!-- yaml-metadata-end -->

---

## Purpose and Structure

Terraform provider for Dell PowerStore block and file storage arrays.
Implements 21 managed resources and 20 data sources using HashiCorp's
Terraform Plugin Framework, enabling infrastructure-as-code management
of PowerStore arrays via their REST API.

The provider is a standalone Go binary that communicates with Terraform
Core over gRPC (go-plugin protocol). It uses two SDK clients: the
public `gopowerstore` SDK for core operations and an OpenAPI-generated
client (`clientgen/`) for additional API coverage.

---

## Components

| Component | Path | Responsibility |
|-----------|------|---------------|
| Entry point | `main.go` | `providerserver.Serve` — starts gRPC server |
| Provider | `powerstore/provider.go` | Schema, Configure, resource/datasource registration |
| Resources | `powerstore/resource_*.go` | CRUD lifecycle for 21 managed resources |
| Data sources | `powerstore/datasource_*.go` | Read-only queries for 20 data sources |
| Helpers | `powerstore/helper/` | Type converters, diag helpers, list utilities |
| Custom types | `powerstore/customtypes/` | Custom Terraform types (e.g. nfshostset) |
| Client wrapper | `client/` | Wraps `gopowerstore` SDK + OpenAPI-generated client |
| Models | `models/` | Terraform state model structs (`tfsdk` tags) |
| JSON models | `models/jsonmodel/` | JSON serialization model structs |
| OpenAPI client | `clientgen/` | Auto-generated REST client code |
| OpenAPI specs | `clientgen_utils/openapi_specs/` | OpenAPI JSON specifications |
| Examples | `examples/` | HCL configurations for resources and data sources |
| Docs | `docs/` | Generated provider documentation |

---

## Key Behaviors

### Authentication

**GIVEN** a user configures the provider with endpoint, username,
and password (via HCL block or environment variables)
**WHEN** `Configure()` runs
**THEN** (1) env vars `POWERSTORE_ENDPOINT`, `POWERSTORE_USERNAME`,
`POWERSTORE_PASSWORD`, `POWERSTORE_INSECURE`, `POWERSTORE_TIMEOUT`
override HCL values, (2) SDK clients are initialized with a default
120-second timeout, (3) a dummy `GetVolumes()` call validates
authentication before any resource operations proceed

### Resource CRUD Lifecycle

**GIVEN** a resource definition in HCL
**WHEN** `terraform apply` runs
**THEN** the resource's `Create()` reads the plan into a model struct,
calls the SDK to create the resource, maps the API response back to
Terraform state, and sets `resp.State`

### Drift Detection

**GIVEN** a resource exists in Terraform state
**WHEN** `terraform plan` or `terraform refresh` runs
**THEN** `Read()` calls the SDK to fetch current state from the array,
compares it with stored state, and updates the state if drifted

### Import

**GIVEN** a resource exists on the array but not in Terraform state
**WHEN** `terraform import powerstore_<resource>.<name> <id>` runs
**THEN** `ImportState()` fetches the resource by ID and populates state

---

## Interfaces

### Provider Configuration Schema

| Attribute | Type | Env Var | Default | Description |
|-----------|------|---------|---------|-------------|
| `endpoint` | string | `POWERSTORE_ENDPOINT` | — | IP or FQDN (must end with `/api/rest`) |
| `username` | string | `POWERSTORE_USERNAME` | — | API username |
| `password` | string (sensitive) | `POWERSTORE_PASSWORD` | — | API password |
| `insecure` | bool | `POWERSTORE_INSECURE` | `false` | Skip TLS verification |
| `timeout` | int64 | `POWERSTORE_TIMEOUT` | `120` | Request timeout (seconds) |

### SDK Client Layer

The `client.Client` struct wraps two clients:
- `PStoreClient` (`*gopowerstore.ClientIMPL`) — primary SDK
- `GenClient` (`*clientgen.APIClient`) — OpenAPI-generated fallback

---

## Dependencies

| Depends On | For |
|------------|-----|
| `github.com/dell/gopowerstore` v1.18.0 | PowerStore REST API SDK |
| `clientgen/` (local) | OpenAPI-generated client for extended API |
| `hashicorp/terraform-plugin-framework` v1.13.0 | Core provider interfaces |
| `hashicorp/terraform-plugin-framework-validators` | Attribute validation |
| `hashicorp/terraform-plugin-log` | Structured logging |
| `hashicorp/terraform-plugin-testing` | Acceptance test harness |
| `bytedance/mockey` | Unit test function-level mocking |
| `stretchr/testify` | Test assertions |

---

## Known Constraints

1. **Terraform Plugin Framework only** — no SDK v2 code.
2. **CGO_ENABLED=0** — static binaries for all platforms.
3. **Sensitive attributes marked** — credentials never in plan output.
4. **ImportState required** — all resources support `terraform import`.
5. **Environment variable fallback** — all credentials support env vars.
6. **Acceptance tests gated** — never run without `TF_ACC=1`.
7. **Endpoint format** — must end with `/api/rest`.
8. **Dual SDK strategy** — `gopowerstore` for primary ops,
   `clientgen` for endpoints not yet in the SDK.

---

## Change History

| Date | Feature | What Changed | Author |
|------|---------|-------------|--------|
| 2026-06-10 | Initial architecture | Provider-specific architecture extracted from generic multi-provider doc | architecture-agent |
