# Epic 0: Project Scaffolding — Green CI

## Why This Exists

Switchboard has CI configured but zero Go source code. Every future story depends on a passing CI pipeline. This epic exists to produce the minimum viable binary and test suite that turns CI green, so development can begin from a stable baseline.

## Scope

This is NOT product work. This is infrastructure. One story, one PR, done.

## Acceptance Criteria

- [ ] `go mod init` with correct module path
- [ ] `cmd/switchboard/main.go` builds and runs (prints version, exits 0)
- [ ] `--version` flag outputs the version string (supports `-ldflags` injection from CI)
- [ ] At least one test file with table-driven tests achieving >= 75% coverage
- [ ] `just fmt` passes (gofumpt)
- [ ] `just lint` passes (golangci-lint, zero warnings)
- [ ] `just test` passes (go test ./... -v)
- [ ] `just test-race` passes (race detector)
- [ ] CI pipeline goes green on PR and on merge to main
- [ ] Docker E2E test job passes (or is skipped gracefully if Dockerfile.test doesn't exist yet)

## Stories

### Story 0.1: Stub Switchboard Binary and Tests

**Priority:** P0 — Blocks all future work
**Size:** XS (< 1 hour of dev time)

**What:** Create the minimal Go project structure that satisfies CI.

**Deliverables:**
1. `go.mod` — module `github.com/arcaven/switchboard`, Go 1.25+
2. `cmd/switchboard/main.go` — entry point with `--version` flag, prints version and exits
3. `cmd/switchboard/main_test.go` — table-driven tests for version output, exit behavior
4. Justfile updates if needed (verify existing `just` targets work)
5. Dockerfile.test if needed for Docker E2E job (or ensure CI handles its absence)

**Technical Notes:**
- Version string injected via `ldflags`: `-X main.version=$(VERSION)`
- Default version when not injected: `"dev"`
- Keep `main.go` minimal — no internal packages yet
- Tests must use stdlib `testing` only (no testify)
- Use `t.Helper()`, `t.Parallel()`, table-driven pattern per CLAUDE.md

**Definition of Done:**
- PR passes all CI checks
- Merged to main triggers build-binaries, sign-and-notarize, and release jobs
- Binary runs: `./switchboard --version` → prints version string

## Out of Scope

- Any internal packages
- Any networking, routing, or protocol code
- Configuration files or config parsing
- Anything beyond "binary exists, CI passes"

## Dependencies

- None. This is the root of the dependency tree.
