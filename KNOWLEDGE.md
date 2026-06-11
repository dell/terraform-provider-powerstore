# KNOWLEDGE.md — terraform-provider-powerstore

<!-- yaml-metadata-start -->
scope_paths: ["./"]
capture_git_sha: "4030193a32ec6334d07026f292726100f5d7b933"
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
| Failure Modes & Gotchas | `## Failure Modes & Gotchas` | Endpoint format, SDK versioning, state secrets, state corruption, auth edge cases | 3 |
| Implicit Contracts | `## Implicit Contracts` | Env var precedence, auth validation, TLS defaults | — |
<!-- quick-reference-end -->

## Five Questions Quick Reference

### What does it do?
Terraform provider for Dell PowerStore block and file storage arrays. Exposes 22 resources and 20 data sources covering volumes, hosts, host groups, protection policies, snapshot rules, file systems, file system snapshots, NAS servers, NFS exports, SMB shares, replication rules, replication sessions, storage containers, volume groups, volume group snapshots, volume snapshots, recycle bin, I/O limit rules, file I/O limit rules, QoS policies, metro volumes, metro volume groups, and remote systems
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

Originally built on Terraform Plugin SDK v2, then migrated to
Terraform Plugin Framework. The `clientgen/` OpenAPI-generated client
was introduced when the PowerStore API surface grew large and complex
enough that manual REST code became unmaintainable. Major refactor
patterns over time include:

- Client abstraction cleanup
- Model-driven design
- Error handling standardization
- Async / polling improvements
- Testing maturity (mockey adoption)

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

### 6. State corruption

State corruption has occurred in production. Large state files with
many managed resources increase the risk. Always use remote backends
with locking (S3+DynamoDB, Terraform Cloud) to prevent concurrent
state writes.

### 7. Authentication edge cases

Authentication edge cases exist — credential rotation during active
Terraform runs, expired tokens, and network timeouts during the
`Configure()` validation call can leave the provider in an
unrecoverable state requiring `terraform init` re-run.

### 8. Resource cleanup failures

Failed acceptance test runs or interrupted `terraform destroy` can
leave orphaned resources on the PowerStore array. These must be
cleaned up manually via the array management UI or REST API.

### Never Again

#### NA-001: State corruption from concurrent applies
- **Impact:** State file corruption when multiple engineers ran
  `terraform apply` simultaneously without state locking.
- **Constraint:** Must use remote backend with locking enabled.
- **Applies to:** All Dell Terraform providers.

#### NA-002: Auth failure masking
- **Impact:** Misleading error messages when endpoint URL missing
  `/api/rest` suffix.
- **Constraint:** Endpoint format validated in `Configure()`.
- **Applies to:** terraform-provider-powerstore.

#### NA-003: Orphaned resources from test failures
- **Impact:** Acceptance test resources left on array after test
  failure, consuming storage capacity.
- **Constraint:** Manual cleanup required; `TF_ACC=1` gating.
- **Applies to:** All Dell Terraform providers.

### Evolution

Failure modes evolved with the SDK v2 → Plugin Framework migration.
Error handling was standardized during the model-driven design
refactor. The dual SDK strategy (`gopowerstore` + `clientgen`)
introduced a new failure surface around client version mismatches.

---

## Performance Characteristics

**Large state files:** Performance degrades with many managed
resources in a single state file. Recommend splitting into multiple
Terraform workspaces or state files when managing >100 resources.

**API rate limiting:** PowerStore arrays enforce API rate limits.
Bulk operations (e.g., creating many volumes) may hit these limits,
causing transient 429 errors. The `gopowerstore` SDK handles retries
internally, but long-running applies may timeout.

**Timeout tuning:** The default 120-second timeout
(`POWERSTORE_TIMEOUT`) may be insufficient for bulk operations or
slow network conditions. Increase for large deployments.

### Evolution

Timeout was made configurable via environment variable after
production deployments hit the original hardcoded limit.

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

Environment variable precedence was established during the SDK v2
era and carried forward into Plugin Framework. The authentication
validation call was added after production incidents with invalid
credentials causing cascading resource failures.

---

## Threading & Synchronization

Terraform Plugin Framework handles concurrency at the provider level.
Individual resource operations are not concurrent by default, but
Terraform Core may invoke multiple resource operations in parallel
during `terraform apply` (controlled by `-parallelism` flag,
default 10).

**Concurrent API access:** Multiple resources hitting the same
PowerStore API endpoint simultaneously can cause contention. The
`gopowerstore` SDK is shared across all resource operations within
a single provider instance.

**Dual SDK race conditions:** Both `PStoreClient` and `GenClient`
are initialized in `Configure()` and shared. No mutex protects
concurrent access — the SDK clients are expected to be thread-safe,
but edge cases exist under high parallelism.

### Evolution

Migration from SDK v2 to Plugin Framework changed the concurrency
model. SDK v2 serialized all operations; Plugin Framework allows
parallel resource operations. The dual-client architecture
(`gopowerstore` + `clientgen`) introduced additional concurrency
surface area.

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

Build system evolved from basic `go build` to Makefile with
linting, security scanning (gosec), and GoReleaser for
cross-platform releases. Testing maturity improved from minimal
acceptance tests to comprehensive mockey-based unit tests.

---

## Operational Knowledge

**Unit tests:** `bytedance/mockey` for runtime function patching.
No hardware required. Run with `make test`.

**Acceptance tests:** `terraform-plugin-testing` against live hardware.
Creates real resources. Run with `TF_ACC=1 make testacc`. Clean up
manually if tests fail mid-run.

### Evolution

Operational patterns matured with the mockey adoption for unit
tests, reducing dependence on live hardware for development
feedback loops.

---

## General Context

### Open Issues

No TODO/FIXME/HACK markers found in non-test source files.

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
