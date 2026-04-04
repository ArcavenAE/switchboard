# CLAUDE.md — Switchboard

## Project Overview

Switchboard is a Go project.

- **Language:** Go 1.25+
- **Build:** `just build` · `just test` · `just lint` · `just fmt`

## Project Structure

```
cmd/switchboard/      # Entry point
internal/             # Internal packages
```

## Development Workflow

**Git workflow:** Gitflow. `develop` is the default branch. Branch from and
PR into `develop`. Alpha releases are cut automatically from `develop`.
Stable releases are cut from `main` via version tags (`v*`). Do not push
directly to `main`.

```bash
just fmt              # gofumpt formatting (run before every commit)
just lint             # golangci-lint — must pass with zero warnings
just test             # go test ./... -v
just test-race        # Race detector — run before pushing
```

## Go Quality Rules

### Idiomatic Go — MUST Follow

These rules prevent the most common AI-generated Go anti-patterns.

**1. Use `fmt.Fprintf` — never `WriteString` + `Sprintf`**
```go
// WRONG — allocates intermediate string
s.WriteString(fmt.Sprintf("Task: %s", name))

// RIGHT — writes directly to the writer
fmt.Fprintf(&s, "Task: %s", name)
```

**2. Never nil-check before `len`**
```go
// WRONG — len handles nil slices/maps (returns 0)
if tasks != nil && len(tasks) > 0 { ... }

// RIGHT
if len(tasks) > 0 { ... }
```

**3. Always check error returns**
```go
// WRONG — silently ignoring error
data, _ := json.Marshal(task)

// RIGHT — handle or propagate every error
data, err := json.Marshal(task)
if err != nil {
    return fmt.Errorf("marshal task %s: %w", task.ID, err)
}
```

**4. Wrap errors with context using `%w`**
```go
// WRONG — loses error chain
return fmt.Errorf("failed to save: %v", err)

// RIGHT — preserves chain for errors.Is/errors.As
return fmt.Errorf("save task %s: %w", id, err)
```

**5. Accept interfaces, return concrete types**
```go
// WRONG — returning interface hides implementation
func NewProvider() Provider { ... }

// RIGHT — return the concrete type
func NewProvider(path string) *FileProvider { ... }
```

**6. `context.Context` is always the first parameter**
```go
// WRONG
func Load(path string, ctx context.Context) error

// RIGHT
func Load(ctx context.Context, path string) error
```

**7. Don't use `interface{}`/`any` without justification**
- Prefer specific types or generics over `any`
- If `any` is needed, document why in a comment

**8. Prefer value receivers unless mutation is needed**
```go
// Use pointer receiver only when:
// - The method mutates the receiver
// - The struct is large (>~64 bytes) and copying is expensive
// - Consistency: if one method needs pointer, all should use pointer
```

**9. No `init()` functions**
- Pass dependencies explicitly via constructors
- Configuration belongs in `main()` or factory functions

**10. Timestamps always in UTC**
```go
// WRONG
time.Now()

// RIGHT
time.Now().UTC()
```

### Error Handling

- Every exported function that can fail returns `error` as last return value
- Use `errors.Is()` and `errors.As()` for error inspection — never string matching
- Define sentinel errors as package-level `var` with documentation
- No panics in library code

### Testing Standards

- **Table-driven tests** for any function with >2 test cases
- **Use stdlib `testing`** — no testify. Use `t.Fatal`, `t.Errorf`, `t.Helper()`
- **Use `t.Helper()`** in test helper functions so failures report the caller's line
- **Use `t.Cleanup()`** instead of `defer` for test resource cleanup
- **Test files** live alongside source: `foo.go` → `foo_test.go`
- **Test fixtures** in `testdata/` directories
- Mark independent tests with `t.Parallel()` where safe

### Code Organization

- **Package naming:** lowercase, single word — no underscores, no camelCase
- **File naming:** lowercase snake_case (`task_pool.go`, `handler.go`)
- **One primary type per file**
- **Import order:** stdlib → external → internal (gofumpt enforces this)
- **Keep packages small** — split when a package exceeds ~10 files

### Common AI Mistakes to Avoid

1. **Don't create unnecessary abstractions** — three similar lines are better than a premature helper
2. **Don't add unused parameters** "for future use" — YAGNI
3. **Don't shadow imports** — `var errors = ...` shadows the `errors` package
4. **Don't use `log.Fatal`/`os.Exit` outside `main()`** — let errors propagate
5. **Don't buffer channels without justification** — unbuffered is the default for a reason
6. **Don't use `sync.Mutex` when `atomic` suffices** for simple counters/flags
7. **Don't create `utils` or `helpers` packages** — put functions where they're used
8. **Don't add comments that restate the code** — only comment the "why", not the "what"
9. **Don't use `strings.Builder` then call `Sprintf` into it** — use `fmt.Fprintf` directly
10. **Don't return `bool, error` as a substitute for `error`**

### Formatting & Linting

- **Formatter:** `gofumpt` (stricter than `gofmt`) — run via `just fmt`
- **Linter:** `golangci-lint run ./...` — must pass with zero warnings
- **Vet:** `go vet ./...` — runs as part of `golangci-lint`
- Never disable linter rules with `//nolint` without a justifying comment

### Go Proverbs to Follow

> The bigger the interface, the weaker the abstraction.

> Make the zero value useful.

> A little copying is better than a little dependency.

> Don't communicate by sharing memory; share memory by communicating.

> Errors are values — program with them.

> Don't just check errors, handle them gracefully.
