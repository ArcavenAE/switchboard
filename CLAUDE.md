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

@.claude/rules/_index.md

## How to Work Here (kos Process)

### Re-introduction
Read charter.md before any substantive work. It contains:
- Current bedrock (what's committed)
- Current frontier (what's under exploration)
- Current graveyard (what's been ruled out)

### Session Protocol
1. Read charter.md (orient)
2. Identify the highest-value open question — or capture new ideas in _kos/ideas/
3. Write an Exploration Brief in _kos/probes/
4. Do the probe work
5. Write a finding in _kos/findings/
6. Harvest: update affected nodes, move files if confidence changed
7. Update charter.md if bedrock changed

Cross-repo questions belong in the orchestrator's _kos/, not here.

### Ideas (pre-hypothesis brainstorming)
Ideas live in _kos/ideas/ as markdown files. Generative, possibly contradictory,
no commitment. When an idea crystallizes, extract into a frontier question + brief.

### Node Files
Nodes live in _kos/nodes/[confidence]/[id].yaml
Schema follows kos schema v0.3.
One node per file. Filename = node id.

### Confidence Changes
Moving a file between confidence directories IS the promotion.
Always accompany with a commit message explaining the evidence.

### Harvest Verification
Before starting the next cycle, verify:
- [ ] Finding written and committed
- [ ] Charter updated if bedrock changed
- [ ] Frontier questions updated (closed, opened, or revised)
- [ ] Exploration briefs marked complete or carried forward
