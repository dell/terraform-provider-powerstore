# KNOWLEDGE.md — terraform-provider-powerstore

<!-- yaml-metadata-start -->
scope_paths: ["./"]
capture_git_sha: "3d3c6ee5f4a1eeeedb5f2dc055b4d74d0e7a7c06"
status: "current"
auto_update: false
preview_before_apply: true
scaffold_version: "1.0"
# session_state: { is_complete: true }
<!-- yaml-metadata-end -->

<!-- quick-reference-start -->
## Agent Quick Reference

| Section | Heading | Summary | never_again_count |
|---------|---------|---------|-------------------|
| Component Overview | `## Component Overview` | Dell PowerStore block and file storage arrays provider | — |
| Architectural Rationale | `## Architectural Rationale` | Public SDK strategy; Plugin Framework architecture | — |
| Failure Modes & Gotchas | `## Failure Modes & Gotchas` | Endpoint format, SDK versioning, state secrets | 0 |
| Implicit Contracts | `## Implicit Contracts` | Env var precedence, auth validation, TLS defaults | — |
<!-- quick-reference-end -->

## Five Questions Quick Reference

### What does it do?
Terraform provider for Dell PowerStore block and file storage arrays. Exposes 21 resources and 20 data sources covering volumes, hosts, host groups, protection policies, snapshot rules, file systems, file system snapshots, NAS servers, NFS exports, SMB shares, replication rules, replication sessions, storage containers, volume groups, volume group snapshots, volume snapshots, recycle bin, I/O limit rules, file I/O limit rules, QoS policies, metro volumes, metro volume groups, and remote systems
through HashiCorp's Terraform Plugin Framework. Communicates with
the hardware REST API via `github.com/dell/gopowerstore` v1.18.0.

### How do you modify it?
Create `resource_<name>.go` (or `*_resource.go`) implementing
`resource.Resource`, add model structs, register in `provider.go`,
add unit tests with mockey mocks, add acceptance tests, create
example HCL, and run `make generate` for docs.

### What breaks?
**Endpoint must end with `/api/rest`** — omitting this suffix causes authentication failures that report misleading HTTP errors. Acceptance tests against live hardware create real
resources — failed test runs may leave orphaned resources. State files
contain secrets — use encrypted remote backends.

### What depends on it?
Terraform Core (gRPC go-plugin), `github.com/dell/gopowerstore` v1.18.0,
`hashicorp/terraform-plugin-framework` v1.13.0.

### What's undocumented?
The `client.Client` struct wraps two clients: `PStoreClient` (`*gopowerstore.ClientIMPL`) for primary SDK operations, and `GenClient` (`*clientgen.APIClient`) for OpenAPI-generated endpoints not yet in the SDK. Both are initialized in `Configure()`. The `clientgen/` code is auto-generated — run `make build_client` to regenerate from OpenAPI specs in `clientgen_utils/openapi_specs/`.

---

## Component Overview

Terraform provider for Dell PowerStore block and file storage arrays.
21 resources and 20 data sources covering volumes, hosts, host groups, protection policies, snapshot rules, file systems, file system snapshots, NAS servers, NFS exports, SMB shares, replication rules, replication sessions, storage containers, volume groups, volume group snapshots, volume snapshots, recycle bin, I/O limit rules, file I/O limit rules, QoS policies, metro volumes, metro volume groups, and remote systems. Resources use `resource_<name>.go` naming. Each struct holds a `*client.Client` reference.

---

## Architectural Rationale

The provider follows the standard Terraform Plugin Framework architecture
— a standalone Go binary communicating with Terraform Core over gRPC.

**SDK strategy (Public):** Uses a public, versioned Go module on GitHub. Provider and SDK release independently. Update via `go get github.com/dell/gopowerstore@<version>; go mod tidy`.

All providers in the Dell Terraform family share this architecture:
Terraform Plugin Framework interfaces, `resource.Resource` for CRUD
resources, `datasource.DataSource` for read-only queries, models with
`tfsdk` struct tags, and mockey-based unit testing.

### Evolution

TBD — requires SME input on how the architecture changed over time.

---

## Failure Modes & Gotchas

### 1. Endpoint URL format

**Endpoint must end with `/api/rest`** — omitting this suffix causes authentication failures that report misleading HTTP errors.

### 2. Sensitive attributes must be marked

All credential fields must have `Sensitive: true` in the schema.
Without this, passwords appear in `terraform plan` output and state
files. This is enforced by code convention, not by the framework.

### 3. State file contains secrets

Terraform state files contain full resource representations including
credentials. Always use encrypted remote backends (S3+KMS, Terraform
Cloud) in production.

### 4. OpenAPI client regeneration

`make build_client` regenerates `clientgen/` from the OpenAPI spec. Requires the OpenAPI Generator CLI JAR and a filtered spec in `clientgen_utils/openapi_specs/`. Do not hand-edit files in `clientgen/`.

### 5. Default sector size

The volume resource defaults `sector_size` to 512 bytes (constant `defaultSectorSize`). This is hardcoded and not configurable via provider config.

### Never Again

No incident-derived constraints recorded. If you know of past
incidents affecting this component, please record them during the
next Knowledge Extraction session.

### Evolution

TBD — requires SME input.

---

## Performance Characteristics

TBD — requires SME input for bottlenecks, scaling limits, tuning
parameters, benchmarks, and known performance cliffs.

### Evolution

TBD — requires SME input.

---

## Implicit Contracts

**Environment variable precedence:** env vars (`POWERSTORE_*`)
override HCL provider block values when both are set. This is
implemented in `Configure()` and is not documented as an explicit
contract.

**Authentication validation:** `Configure()` makes a dummy API call
to validate credentials before any resource operations proceed. If
this call fails, all resource operations are blocked.

**TLS verification default:** `insecure` defaults to `false` —
TLS verification is on by default. Setting `insecure = true` is
a lab-only setting and must never be used in production.

**Acceptance test gating:** tests guarded by `TF_ACC=1` — never
run without live hardware credentials. Tests create real resources
that must be cleaned up manually if the test run fails.

### Evolution

TBD — requires SME input.

---

## Threading & Synchronization

Terraform Plugin Framework handles concurrency at the provider level.
Individual resource operations are not concurrent by default.

### Evolution

TBD — requires SME input.

---

## Build System & Configuration

Standard Makefile targets shared across all Dell Terraform providers:

| Target | Purpose | Hardware Required |
|--------|---------|-------------------|
| `make build` | Compile provider binary | No |
| `make install` | Install to `~/.terraform.d/plugins/` | No |
| `make test` | Run unit tests | No |
| `make testacc` | Run acceptance tests | **Yes** |
| `make check` | Format, lint, vet | No |
| `make gosec` | Security scan | No |
| `make cover` | Generate coverage report | No |
| `make generate` | Generate documentation | No |

GoReleaser configuration: CGO_ENABLED=0, platforms (freebsd, windows,
linux, darwin), architectures (amd64, 386, arm, arm64).

### Evolution

TBD — requires SME input.

---

## Operational Knowledge

**Unit tests:** `bytedance/mockey` for runtime function patching.
No hardware required. Run with `make test`.

**Acceptance tests:** `terraform-plugin-testing` against live hardware.
Creates real resources. Run with `TF_ACC=1 make testacc`. Clean up
manually if tests fail mid-run.

### Evolution

TBD — requires SME input.

---

## General Context

### Open Issues

TBD — requires code scanning for TODO/FIXME/HACK markers.

### Glossary

| Term | Definition |
|------|------------|
| Plugin Framework | HashiCorp's Terraform Plugin Framework (`terraform-plugin-framework`) |
| mockey | `bytedance/mockey` — runtime function patching for unit tests |
| POWERSTORE | Environment variable prefix for this provider |

---

## References

- [Terraform Plugin Framework Docs](https://developer.hashicorp.com/terraform/plugin/framework)
- [Dell Terraform Registry](https://registry.terraform.io/namespaces/dell)

---

## Governance Spec Discrepancies

No discrepancies detected between code/SME knowledge and loaded
governance specs.
