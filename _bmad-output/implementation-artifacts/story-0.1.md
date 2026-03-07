# Story 0.1: Stub Switchboard Binary and Tests

Status: ready-for-dev

## Story

As a developer,
I want a minimal switchboard binary with passing CI,
so that all future stories have a stable green baseline to build on.

## Acceptance Criteria

1. `go.mod` exists with module path `github.com/arcaven/switchboard` and Go 1.25+
2. `cmd/switchboard/main.go` compiles and produces a binary
3. `switchboard --version` prints the version string and exits 0
4. Version is injectable via `-ldflags "-X main.version=X.Y.Z"`; defaults to `"dev"` when not injected
5. `just fmt` passes with zero changes needed
6. `just lint` passes with zero warnings
7. `just test` passes — all tests green
8. `just test-race` passes — no race conditions
9. CI quality-gate job passes on PR
10. CI build-binaries job produces `switchboard-darwin-arm64`, `switchboard-darwin-amd64`, `switchboard-linux-amd64`
11. Test coverage meets or exceeds 75% threshold

## Tasks / Subtasks

- [ ] Task 1: Initialize Go module (AC: #1)
  - [ ] `go mod init github.com/arcaven/switchboard`
  - [ ] Verify `go.mod` specifies Go 1.25+
- [ ] Task 2: Create entry point (AC: #2, #3, #4)
  - [ ] `cmd/switchboard/main.go` with `--version` flag
  - [ ] `var version = "dev"` for ldflags injection
  - [ ] Print version to stdout, exit 0
- [ ] Task 3: Create tests (AC: #7, #8, #11)
  - [ ] `cmd/switchboard/main_test.go` with table-driven tests
  - [ ] Test version output format
  - [ ] Test exit behavior
  - [ ] Use `t.Parallel()` where safe
  - [ ] Achieve >= 75% coverage
- [ ] Task 4: Verify toolchain (AC: #5, #6)
  - [ ] Run `just fmt` — zero diff
  - [ ] Run `just lint` — zero warnings
- [ ] Task 5: Verify CI (AC: #9, #10)
  - [ ] Push PR, confirm quality-gate green
  - [ ] Merge to main, confirm build-binaries produces artifacts

## Dev Notes

- This is an XS story — absolute minimum code. No internal packages, no config, no networking.
- **Zero external dependencies.** Stdlib only. No `go.sum` expected.
- The binary does ONE thing: print its version and exit.
- **No args behavior:** prints version and exits (same as `--version`). Keep it dead simple.
- `main()` should be kept small enough that tests can achieve 75% coverage without gymnastics. Extract a `run()` function that `main()` calls, so tests can invoke `run()` directly.
- **`run()` signature:** `run(stdout io.Writer, args []string) error` — writer first (dependency), args second.
- **Flag handling:** Use stdlib `flag` package. It handles both `-version` and `--version`. Do NOT bring in cobra/pflag for a one-flag binary.
- No `init()` functions per CLAUDE.md rules.

### Project Structure Notes

- `cmd/switchboard/` — entry point, already referenced by CI's `go build ./cmd/switchboard`
- No `internal/` packages needed for this story
- Justfile already exists with `fmt`, `lint`, `test`, `test-race` targets — verify they work with the new code
- CI workflow at `.github/workflows/ci.yml` already configured correctly

### Testing Strategy (Red/Green TDD)

**Approach:** Write failing tests first, then implement to pass.

- Extract a testable `run(stdout io.Writer, args []string) error` function from `main()`
- `main()` is just: call `run()`, check error, `os.Exit(1)` on error — ~3 untestable lines, rest is covered via `run()`
- Use `bytes.Buffer` as stdout for output capture
- No external dependencies — stdlib `testing` only

**Red phase — write these tests first:**

1. `TestRun_Version` — `run(stdout, []string{"switchboard", "--version"})` returns nil, stdout contains `"dev"`
2. `TestRun_NoArgs` — `run(stdout, []string{"switchboard"})` returns nil, stdout contains version
3. `TestRun_UnknownFlag` — `run(stdout, []string{"switchboard", "--bogus"})` returns error
4. `TestRun_VersionNonEmpty` — version variable is never empty string (regression guard)

**Green phase — implement `run()` to pass all tests.**

**Coverage target:** With `run()` covering all logic and only `main()` untested, expect ~85%+ coverage. Well above 75% floor.

### References

- [Source: _bmad-output/planning-artifacts/epic-0-project-scaffolding.md]
- [Source: .github/workflows/ci.yml — quality-gate job, build-binaries job]
- [Source: CLAUDE.md — Go Quality Rules, Testing Standards]

## Dev Agent Record

### Agent Model Used

### Debug Log References

### Completion Notes List

### CI Findings (Post-Push)

1. **Quality Gate: `errcheck` violation** — `fmt.Fprintf` return value unchecked. golangci-lint v2 with errcheck enabled. Fixed by checking error and wrapping with `%w`.
2. **Docker E2E: missing `go.sum`** — `Dockerfile.test` COPY expects `go.sum` but zero-dependency project doesn't generate one. Fixed by creating empty `go.sum` (valid for no-deps projects).
3. **Coverage regression** — handling `fmt.Fprintf` error added an uncovered branch. Fixed by adding `failWriter` test that exercises the write-error path. Coverage: 81.8%.

### File List
