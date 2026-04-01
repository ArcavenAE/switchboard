---
stepsCompleted: [1, 2, 3, 4, 5, 6]
status: complete
inputDocuments:
  - '_bmad-output/brainstorming/brainstorming-session-2026-03-07-001.md'
  - '_bmad-output/brainstorming/naming-node-type-parking-lot.md'
  - '_bmad-output/brainstorming/session-context-cache.md'
  - '_bmad-output/planning-artifacts/epic-0-project-scaffolding.md'
  - '_bmad-output/implementation-artifacts/story-0.1.md'
date: 2026-03-31
author: maker
---

# Product Brief: Switchboard

## Executive Summary

Switchboard is a session network — a purpose-built switched network for terminal sessions. Not an overlay, not a VPN, not a tunnel. A network whose primitives are sessions, channels, and cryptographic admission, optimized for interactive CLI latency.

Terminals are suddenly mission-critical infrastructure. The explosion of AI coding agents — Claude Code, Codex, Devin — means platform engineers and AI ops professionals now manage fleets of agent sessions across machines worldwide. These fleets run in terminals. The network layer for that doesn't exist. Teams duct-tape SSH tunnels, fight reconnection, and have no visibility into session quality. Switchboard is the infrastructure they're missing.

Switchboard provides low-latency, multi-path, end-to-end encrypted access to remote tmux sessions over virtual switched networks (VSNs). SSH provides the end-to-end trust layer — the same OpenSSH that already secures every server you operate. Switchboard adds latency-optimized multi-path routing, cryptographic network admission, multi-hop failover, and session quality monitoring — all with carrier-grade content separation. Routers see identity, addressing, and traffic patterns — enough to route intelligently — but session content is encrypted end-to-end between nodes. The network operator provides infrastructure; the customer holds the data keys.

---

## Core Vision

### Problem Statement

Terminal sessions are the control plane for modern infrastructure operations and AI agent orchestration. Yet the network layer treats them as an afterthought — just another TCP connection over SSH.

Today's options are inadequate:

- **Raw SSH** provides encryption but is single-path, single-hop, and fragile. A network hiccup kills the session.
- **Mosh** improves reconnection but operates point-to-point with no routing, no multi-path, and no network-level identity.
- **tmate/Upterm** provides relay access but relays can observe session content — no end-to-end encryption through the relay.
- **Overlay networks (Tailscale, Nebula, ZeroTier)** provide connectivity but are general-purpose IP tunnels, not optimized for interactive terminal latency. Their relay nodes may observe traffic depending on configuration.
- **Eternal Terminal** provides reconnection and tmux support but is single-path with no routing intelligence.
- **Browser-based proxies (VibeTunnel, CloudeCode)** put terminals in a browser for AI agent monitoring but are single-machine, web-relayed, with no network-level routing or cryptographic admission.
- **HappyCoder** provides E2E encrypted mobile/web access to Claude Code but is single-agent oriented with no multi-session routing or network-level identity.
- **Claude Code Dispatch/Remote Control** provides official remote access to Claude Code sessions but is vendor-hosted, Claude Code-only, with no self-hosted infrastructure or multi-session routing.
- **Chat integrations (Slack/Discord bots)** route terminal commands through messaging platforms where the platform sees all commands and output — no encryption, no session quality, latency-insensitive.

Every existing tool was designed for the single-user, single-session case. None provides a multi-tenant, multi-session, multi-node network with routing intelligence, cryptographic admission, and terminal-native optimization. No session network exists — because terminals weren't mission-critical infrastructure until AI agents made them so.

### Problem Impact

When remote sessions are unreliable, terminal professionals lose flow state. Keystroke-to-echo latency above 100ms breaks the illusion of local interaction. Session drops force context rebuilding. Lack of quality visibility means users can't distinguish network problems from application problems. For AI ops professionals managing agent fleets across machines, unreliable sessions mean unreliable operations — agents stall, supervisors lose contact, coordination breaks down.

### Why Now

The AI agent explosion changed what terminals are. A terminal session is no longer one person on one server. It's a fleet of agents, each in a tmux pane, across machines that could be anywhere, managed by supervisors that need reliable real-time access. The number of concurrent terminal sessions per operator went from single digits to dozens or hundreds. The reliability requirements went from "nice to have" to "the agents stop working."

No one built the network layer for this because no one needed it until now.

### Why Existing Solutions Fall Short

Existing tools fall into three categories, none purpose-built for terminal sessions:

**General-purpose transport** (SSH, Mosh, Eternal Terminal) moves encrypted bytes between two endpoints. No routing, no multi-path, no network identity. Single-hop, single-path, fragile.

**General-purpose overlay networks** (Tailscale, Nebula, ZeroTier) create IP connectivity. They tunnel anything — which means they're optimized for nothing. Terminal sessions get the same treatment as file transfers, video calls, and database connections.

**Application-layer remote access** (tmate/Upterm, VibeTunnel/CloudeCode, HappyCoder, Dispatch/Remote Control, Slack/Discord bots) bolts remote access onto specific tools. Each is built for one application (usually Claude Code), not a network primitive. Relay trust varies. No session quality. No infrastructure you own.

Switchboard is none of these. It's a purpose-built session network — a network whose only job is terminal sessions, optimized for exactly that.

**Competitive Comparison:**

| Solution | What It Provides | What It Lacks |
|----------|-----------------|---------------|
| SSH | E2E encryption, universal | Single-path, no failover, no quality monitoring |
| Mosh | Reconnection, state sync | Point-to-point only, no routing, no network admission |
| tmate/Upterm | Relay-based terminal sharing | Relays see content, no E2E through relay, static relay selection |
| Tailscale/Nebula/ZeroTier | Overlay networking, NAT traversal | General-purpose IP tunneling, not terminal-optimized, relay trust varies |
| Eternal Terminal | Reconnection, tmux support | Single-path, no routing intelligence |
| VibeTunnel/CloudeCode | Browser-based terminal access, AI agent monitoring | Web relay sees content, single-machine, no network routing or admission |
| HappyCoder | E2E encrypted mobile/web Claude Code access | Single-agent oriented, no multi-session routing, no network identity |
| Dispatch/Remote Control | Official remote access to Claude Code | Vendor-hosted, Claude Code only, no self-hosted infra, no multi-session routing |
| Chat integrations (Slack/Discord) | Terminal access via messaging platforms | Platform sees all commands/output, no encryption, no session quality |

The gap isn't a missing feature in existing tools. It's a missing category of infrastructure.

### Proposed Solution

Switchboard is a virtual switched network for terminal sessions. Three node types (access nodes, consoles, control nodes) communicate through routers that provide latency-aware multi-path forwarding with carrier-grade content separation.

**Key architectural properties:**

- **Carrier-grade content separation.** Routers have routing intelligence — they authenticate nodes, enforce VSN admission, and make forwarding decisions based on identity, addressing, and traffic patterns. But session content is encrypted end-to-end (SSH) between nodes and opaque to the network operator. The operator sees who communicates with whom, when, and how much. Traffic analysis and contact analysis are within the operator's reach; content inspection and injection are not. The customer holds the data keys.
- **Session-native primitives.** VSNs, channels, half-channels, access nodes, consoles — these aren't IP concepts in disguise. They're abstractions built for how terminal sessions actually work: asymmetric (keystrokes up, output down), latency-sensitive, stateful, and long-lived.
- **Purpose-built for terminal latency.** Timeslice-driven framing, content-type-aware loss recovery (interactive vs. streaming vs. bulk vs. graphics), degradation signaling, and human-perceptible latency budgets. Not a general-purpose IP tunnel.
- **Virtual switched networks.** VSNs provide cryptographic isolation between customers/projects on shared infrastructure. Admission is key-based (OpenSSH keypairs). Multicast presence protocol enables session discovery within a VSN.
- **Multi-path resilience.** Nodes maintain connections to multiple routers. Dual-path forwarding with duplicate-and-race for small payloads, FEC for bursts. Failover in seconds, not minutes.
- **Progressive deployment.** Single E router on a LAN (5-minute setup, the 85% case) scales to multi-hop provider networks without architecture changes. Same binary, different topology.

### Key Differentiators

1. **New category: session network.** Not a better tunnel, a smarter relay, or another application hack. A purpose-built network whose only job is terminal sessions. Multi-tenant, multi-session, multi-node by design — built for the fleet-of-agents reality, not the one-user-one-server past.
2. **Carrier-grade content separation.** The network operator sees identity and traffic patterns; the customer controls session content. Routers route intelligently but cannot read or inject payload. This is the trust model that makes shared infrastructure viable — borrowed from telecom, applied to terminal sessions.
3. **Terminal-native optimization.** Every design decision is tuned for keystroke-to-echo latency, not general throughput. Asymmetric half-channels, content-type-aware loss recovery, degradation signaling — built for how terminals actually work.
4. **The illusion of local.** Remote sessions feel like local sessions. The network disappears when it's working; it's honest when it's not (degradation signaling, never mysterious failures).
5. **Progressive complexity.** Two machines, one user, five minutes. Complexity is available but never required. The E router MVP covers 85% of use cases.
6. **Open source, sovereign infrastructure.** Switchboard is open-source software. Run it on your infrastructure. Audit the code. Contribute or fork. No vendor lock-in, no phone-home, no rug-pull. Your sessions, your network, your keys.

---

## Target Users

### Primary Users

**1. Devon — Solo Operator**

One person, two machines. A developer with a workstation at home and a beefy build server in the closet, or a server in the cloud. Runs tmux on the remote machine, SSHes in from wherever. This is the 85% case — the E router MVP user.

**Problem experience:** SSH works until it doesn't. Home wifi hiccups, the session freezes, `~.` to kill the hung connection, reconnect, reattach tmux, find the right pane. Three times a day. The reconnection takes 30 seconds. Getting back to where you were in your head takes longer. Has tried Mosh — it helps with reconnection but the remote machine is behind NAT and Mosh's UDP doesn't punch through reliably. Has tried Eternal Terminal — same NAT problem, and it's another binary to maintain.

**Success:** "I installed Switchboard on both machines. Took five minutes. My sessions don't drop when the wifi blips. I stopped thinking about SSH."

**2. Kai — AI Operations Engineer**

Manages a fleet of 40+ Claude Code agent sessions across 6 machines for a platform engineering team. Each agent runs in a tmux pane, supervised by coordinator agents that also run in tmux. Kai's day is spent monitoring agent work, intervening when agents get stuck, reviewing output, and restarting failed sessions.

**Problem experience:** SSH tunnels to each machine. When the home network hiccups, three sessions drop simultaneously and Kai loses 10 minutes reconnecting and re-finding the right panes. But the reconnection time isn't the real cost — it's the *ambiguity*. When a session feels sluggish, Kai can't tell if it's network degradation or an agent that's stuck. So Kai over-checks, manually polling sessions that are probably fine, because there's no signal saying otherwise. The anxiety tax is higher than the reconnection tax. Uses a patchwork of tmux, SSH, and a Discord bot to monitor. Has tried VibeTunnel but it's single-machine.

**Success:** "I see a yellow indicator — that means network, not agent. I don't need to check. When the indicator is green and the agent isn't producing output, *that's* when I intervene. Switchboard gave me trust in what I'm seeing."

**3. Priya — Platform Engineer**

Operates infrastructure across three cloud providers and a handful of bare-metal machines. Lives in tmux — 15-20 sessions open at any time. Runs ansible, kubectl, terraform from remote sessions. Needs to access the same sessions from her office, her home, and occasionally her phone.

**Problem experience:** SSH with Mosh for the flaky connections. Works most of the time but has no multi-path — when the primary route to the Singapore datacenter degrades, she waits or bounces through a jump host manually. No session quality visibility — "is it slow because of the network or because kubectl is thinking?" Cannot hand a session to a colleague without sharing SSH keys or setting up tmate. In her SOC2-audited environment, sharing credentials is a compliance violation — so she simply can't share sessions at all. Pair debugging means "I'll share my screen on Zoom."

**Success:** "I hand a colleague a read-only key to one specific session. No credential sharing, no compliance violation, no Zoom screen share. They see exactly what I see. When I switch from office to home, the session doesn't even blink."

**4. Marcus — Network / Infrastructure Admin**

Manages datacenter infrastructure, network gear, and monitoring systems. Needs persistent, reliable access to console sessions on remote equipment. Operates across sites connected by sometimes-unreliable WAN links. Deeply skeptical of adding another network layer on top of his carefully managed infrastructure. His skepticism isn't personality — he's deployed "simple" overlay networks that turned into operational nightmares. The E router earns trust by being reversible.

**Problem experience:** SSH over VPN. When the VPN link degrades, sessions freeze. Uses Eternal Terminal for some hosts, raw SSH for others, a jump host topology he maintains by hand. No unified view. Has built his career on tmux and expects tools to respect that. His first reaction to Switchboard: "I don't need another overlay network." His second reaction, after seeing the E router: "This doesn't touch my infrastructure. It's just two machines on the LAN."

**Success:** "I started with the E router between my laptop and the console server. No infrastructure changes. When I added a second site, the sessions failed over to the backup path automatically. I see link quality in real time. And the router infrastructure can't read what I'm typing into those console sessions — carrier-grade separation, not marketing."

### Secondary Users

**Team leads and supervisors** who need visibility into what their operators are doing. Today: they ask on Slack, they walk over and look at a screen, or they have no visibility at all. In distributed teams, "walk over" doesn't exist. They're managing by asking "how's it going?" and trusting the answer.

With Switchboard: read-only access to specific sessions via key-based access modes. Published views for SOC dashboards. A team lead can see what Kai's agent fleet is doing without interrupting Kai, without needing Kai's credentials, without Kai needing to share their screen on a video call. Their adoption is driven by the primary users already being on Switchboard.

### User Journey

**Two adoption tracks, both valid:**

- **Scale track:** Devon → Kai → Priya. More sessions, more machines, more paths. The product grows with the user's fleet.
- **Trust track:** Marcus. Skeptic → E router trial → production infrastructure. The product earns trust progressively.

Some users enter on scale (Devon needs two machines to just work). Others on trust (Marcus needs to see it before he believes it). Priya is on both — she needs scale *and* trust (SOC2 compliance, credential separation).

**Discovery:** Through the problem. Devon hits the "session dropped again" moment. Kai hits the "three sessions dropped on a network hiccup" pain for the Nth time and searches for alternatives. Priya hears about it from another platform engineer. Marcus encounters it in a conference talk and is skeptical until he sees the E router demo.

**Onboarding (target experience):** Install Switchboard. Run an E router on the LAN. Two machines, five minutes. The first session feels like SSH but doesn't drop. Devon's "aha" comes when the wifi blips and the session stays connected. Kai's "aha" comes when the degradation indicator turns yellow instead of the session freezing — network problem, not agent problem, no need to check. Priya's "aha" comes when she hands a colleague a read-only key and it just works.

**Core usage:** Switchboard replaces the SSH-to-remote-tmux workflow entirely. Operators connect to their VSN, see their sessions, attach. It becomes the substrate — invisible when working, honest when degrading.

**Success moment:** The first network outage where sessions survive. The first time a colleague gets read-only access with a key instead of shared credentials. The first time Kai sees 40 agent sessions across 6 machines with quality indicators and realizes the anxiety is gone. The first time Marcus's WAN link fails over and he finds out from the indicator, not from a frozen screen.

**Long-term:** Switchboard is the session infrastructure. Devon never thinks about SSH again. Kai's agent fleet monitoring is solved infrastructure, not daily firefighting. Priya onboards new team members with keys, not credential sharing. Marcus's E router graduates to a PE router connected to the wider network. Progressive deployment, no rearchitecture.

---

## Success Metrics

### User Success — Measurable in Testing

| Metric | Target | Measurement |
|--------|--------|-------------|
| **Session survivability** | Sessions survive network transitions (wifi switch, path failover) without user intervention | CI test suite: simulate network transitions, assert session continuity |
| **Keystroke-to-echo latency** | p50 < 50ms, p99 < 100ms on E router (LAN); p50 < 100ms, p99 < 200ms on PE (WAN) | OTEL telemetry when enabled; integration test harness |
| **Time to first session** | < 5 minutes from install to first working session (Devon's E router case) | Timed onboarding walkthrough; step count ≤ 3 commands |
| **Path failover time** | < 2 seconds from path loss to successful failover | Integration test: kill primary path, measure time to session resumption |
| **Session sharing** | Read-only access to a specific session via key, no credential sharing | Functional test: issue key, verify read-only access, verify no write capability |

### User Success — Measurable via Opt-In Telemetry

| Metric | What It Reveals | Requires |
|--------|----------------|----------|
| **Degradation indicator accuracy** | Does the indicator match actual measured path quality? | OTEL enabled; compare indicator state vs. measured latency/loss |
| **Recovery cascade effectiveness** | How often does each recovery tier fire? (duplication, FEC, SREJ, TLPKTDROP) | OTEL enabled; per-tier counters |

### Project Health — Community Signals

| Signal | What It Means |
|--------|---------------|
| **Issues filed by production users** | People are using it for real work and caring enough to report problems |
| **PRs from non-maintainers** | The project is approachable and worth contributing to |
| **Ecosystem packaging** | Homebrew, distro repos, container images maintained by others |
| **Conference/blog mentions** | People are talking about it in infrastructure contexts |

### Leading Indicators (Rumble Strips)

Warning signs that we're heading toward a death condition before users report it:

| Death Condition | Leading Indicator | Threshold |
|----------------|-------------------|-----------|
| UX breaks flow state | p99 keystroke-to-echo latency | Creeping above 150ms in integration tests |
| Complexity barrier | Onboarding step count | Getting-started guide exceeds 1 page or 3 commands |
| Complexity barrier | Binary dependency count | Requires anything beyond itself and tmux |
| Operational fragility | Single points of failure | Any component failure takes down all sessions |
| Operational fragility | Rolling update capability | Can't update a router without dropping all its sessions |
| Security compromise | Payload visibility at router | Any test showing router can access decrypted content |
| Security compromise | Outer header byte count | Growth beyond 44 bytes requires explicit security justification |
| *All death conditions* | Dependency count (software + service) | No proprietary service dependency, ever. Software dependencies minimized — every dependency is an attack surface and a sovereignty risk |
| *All death conditions* | Bus factor | Starting at 1 (honest). Goal: grow. Project must be approachable enough that others can contribute meaningfully |

### Death Conditions (We've Failed If)

1. The UX breaks flow state — stutters, hangs, freezes that make users notice the network
2. Session security is compromised by anything other than stolen SSH keys
3. Operational fragility — can't do rolling updates, a vulnerability requires taking the whole network down
4. Complexity barrier — Devon can't set up and forget in five minutes

### What We Don't Measure

Switchboard is open-source software. Its success is measured by adoption and utility, not revenue. Sustainability comes from the community it serves and the value it provides — the same model that sustains OpenSSH, tmux, and the infrastructure this project builds on.

Not metrics:
- Revenue, MRR, conversion rates — this is free software, not a product with a sales funnel
- Download counts — downloads don't mean usage
- Feature count — more features is not more success
- Usage telemetry without consent — no phone-home, ever (SOUL.md §6)

---

## MVP Scope

### Core Features (The E Router Release)

The MVP is Devon's case: one person, two machines, one E router, five minutes. It proves the core value proposition — does a session network feel better than raw SSH?

**Implementation note:** The wire protocol carries all fields from day one — frame envelope, admission, session authorization — even before all features are fully implemented. Permissive defaults (accept all keys, no real validation) in early builds, not missing fields. The protocol shape must not change. The PRD will sequence the build into milestones; this brief describes the target MVP.

**What ships:**

1. **E router** — single binary, single hop. Runs on the same LAN as the nodes. No multi-hop, no inter-router protocol. Same binary that later becomes PE/P — just with no upstream router connections configured.

2. **Access node** — connects to E router, publishes tmux sessions via control mode (AN-E: control mode + PTY fallback). Advertises available sessions.

3. **Console** — connects to E router, discovers available sessions, attaches. MVP console is CN-D (direct terminal write) or CN-B (PTY injection). Simplest thing that works.

4. **VSN admission (Tier 1)** — key-based. Control node registers public keys against a VSN. Node presents signed challenge to E router. Admitted or rejected. Local database on the E router (no distributed flooding — single router).

5. **Session authorization (Tier 2)** — access node maintains authorized console keys. Per-session access modes: full access (read-write), read-only. Key carries permission level.

6. **Edge protocol — upstream** — timeslice-driven framing, U-C idempotent replay (sliding window of last N keystrokes), piggybacked ACK + SACK bitmap, TLPKTDROP for keystrokes past perception budget.

7. **Edge protocol — downstream** — D-A reliable ordered stream. All downstream content treated uniformly. ARQ on SACK gap (QUIC model: new frame with old content). TLPKTDROP as last resort with degradation signal.

8. **Frame envelope** — 44-byte outer header (version, frame type, VSN ID, destination, source, length, HMAC) + ~22-byte channel header (channel ID, sequence, timestamp, FEC metadata, flags, ack_seq, ack_bitmap). Carrier-grade content separation — router sees outer, endpoints see inner.

9. **Degradation signaling** — quality indicator visible to the console. Green/yellow/red based on measured latency and loss. When the network degrades, the user knows.

10. **Single-path operation** — E router MVP may have only one path between node and router. Duplication across paths activates when two paths exist (multi-homed node to two E routers on the same LAN, or E router with upstream PE). Not required for MVP to be useful.

### Out of Scope for MVP

The PRD will sequence the MVP build into milestones. Not all MVP features ship simultaneously — the wire protocol and basic session connectivity come first, admission and degradation signaling layer on top. The following features are out of scope for the *entire* MVP, not just the first milestone:

| Feature | Why Deferred | When It Comes |
|---------|-------------|---------------|
| Multi-hop routing | E router is single hop. Inter-router protocol (link-state, Dijkstra) not needed until PE/P routers exist. | Second router deployed |
| XOR parity FEC | MVP downstream is reliable ordered stream. FEC adds complexity for a loss scenario (multi-hop) that MVP doesn't have. | Post-MVP with multi-path |
| SREJ via alternate path | Requires multiple paths. MVP may be single-path. | Post-MVP with multi-homing |
| Interleaving | Post-MVP loss recovery mechanism for bulk downstream under high loss. | Post-MVP with multi-hop |
| Adaptive FEC rate / four-regime profiles | MVP has two regimes (normal + degraded). Full four-regime profiles come with FEC. | Post-MVP |
| Content-type detection | Requires terminal state tracker on access node. MVP sends all downstream uniformly. | Post-MVP (D-CE) |
| Terminal state tracker / state sync | Significant engineering. MVP uses reliable ordered stream (D-A). | Post-MVP (CN-A) |
| Scrollback fetch protocol | Console-side scrollback is the local terminal's scrollback in MVP. | Post-MVP (D-CE + CN-A) |
| Local tmux mirroring (CN-C) | Premium console experience. MVP is CN-D/CN-B. | Post-MVP |
| Control node as separate node type | MVP: control operations happen at the E router directly (CLI commands). | When multi-router requires distributed key management |
| Multi-host scheduling | Marvel concern, not switchboard. | When marvel integrates |
| Router-to-router protocol (Noise, LSA flooding) | Single router. No peers to talk to. | Second router |
| Distributed admission database | Single router. Local admission DB only. | Second router |
| Membership propagation (join/leave flooding) | Single router. All nodes are local. | Second router |
| OTEL telemetry export | Available but not required. Build the hooks, don't require the collector. | Always optional |

### MVP Success Criteria

The MVP succeeds if:

1. **Devon's test:** Install on two machines. Run E router. Attach to remote tmux session. Switch wifi networks. Session survives. Time from install to working session: under 5 minutes.
2. **Latency test:** Keystroke-to-echo p99 < 100ms on LAN. The session feels like local tmux.
3. **Degradation test:** Introduce packet loss on the path. Console shows degradation indicator. User knows it's the network, not the application.
4. **Sharing test:** Issue a read-only key to a second console. Second console can view the session but not send keystrokes.
5. **Security test:** Capture traffic at the E router. Verify payload is opaque — router sees addressing and HMAC, not session content.
6. **Protocol conformance test:** Wire format matches specification. Frame envelope, channel header, and admission messages conform to defined byte layouts. No undocumented fields, no version mismatches between node and router.

### Future Vision

The MVP is the foundation. What grows from it:

**Near-term (post-MVP, same architecture):**
- Multi-path: dual-path forwarding, duplicate-and-race, FEC, SREJ, full recovery cascade
- Content-type-aware downstream (D-CE): terminal state tracker, per-type recovery profiles
- Console upgrades: CN-C (local tmux mirroring), CN-A (control mode consumer)
- Multi-router: PE routers, link-state routing, distributed admission, membership propagation
- `%pause`/`%continue` integration for source-side flow control

**Medium-term (architecture extensions):**
- P routers (core forwarding, no node awareness) — when scale demands it
- Multi-host scheduling via marvel integration
- Director integration for inter-agent communication
- Content packs and spectacle integration

**Long-term (the full vision):**
- A global session network. PE routers at major cloud POPs. Any terminal professional, anywhere, connects to their sessions with the illusion of local. The network is a fact, not a mystery. Simple, reliable sessions over unreliable, complex networks.
- The MCP exploration reopens — Switchboard as session infrastructure for agent-to-agent communication, if the undisclosed context supports it.

---

## Risks and Assumptions

**Risks:**

1. **The Mosh problem — "good enough" inertia.** SSH is familiar. Dropped sessions are annoying but routine. The real competitor isn't any tool in the competitive table — it's the habit of tolerating what exists. The five-minute E router onboarding and progressive deployment are our answers, but adoption requires the pain to be acute enough to try something new.

2. **tmux dependency.** Switchboard is tmux-first by design decision ("tmux-first, depth before breadth" — project decision from the values session). If tmux development stalls or a successor displaces it, Switchboard's value proposition is coupled to that dependency. Mitigation: the access node's PTY fallback (AN-E) works with any terminal session, not just tmux. tmux is the optimized path, not the only path.

3. **Agent fleet growth assumption.** The "Why Now" argument rests on AI agents running in terminals at scale. If the industry moves toward browser-based or API-only agents, Kai's persona weakens. Devon, Priya, and Marcus remain valid regardless — they predate the agent explosion. Switchboard is useful without the agent use case; the agent use case makes it urgent.

**Assumptions:**

4. **Terminal professionals will self-serve.** The target audience installs software, reads man pages, and configures infrastructure. Switchboard doesn't need an onboarding wizard or a GUI. A clear README and a working binary are sufficient.

5. **OpenSSH remains the standard.** The E2E trust layer is SSH. If a post-SSH encrypted terminal protocol emerges, the session network architecture survives but the trust layer needs rework.
