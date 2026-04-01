---
stepsCompleted: [1, 2-research, 3-first-principles, 4-morphological-analysis-complete, 5-loss-recovery-qos-research, 6-downstream-strategy-research, 7-switchboard-for-mcp-exploration, 8-tmux-control-mode-depth, 9-router-control-plane]
inputDocuments: []
session_topic: 'Switchboard - low-latency multi-path E2E encrypted tmux session router architecture with virtual switched networks'
session_goals: 'Naming, architecture design, edge protocol design, failure mode analysis, use cases'
selected_approach: 'ai-recommended'
techniques_used: ['research-synthesis', 'first-principles-thinking', 'morphological-analysis', 'values-exploration', 'research-synthesis-loss-recovery', 'cross-pollination', 'constraint-mapping', 'research-synthesis-downstream-strategy', 'exploration-mcp', 'research-synthesis-tmux-control-mode', 'design-router-control-plane']
ideas_generated: []
context_file: ''
next_phase: 'product-brief'
morphological_parameters_completed: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
morphological_parameters_next: null
queued_sessions:
  - 'COMPLETED: technical-research: loss recovery and QoS techniques for interactive overlays (X.25, interleaving, hybrid ARQ+FEC)'
  - 'COMPLETED: technical-research: downstream strategy (scrollback, graphics/Sixel/Kitty, TUI, tmux control mode, Claude Code use case)'
  - 'EXPLORED-PARKED: exploration: Switchboard for MCP - infrastructure alignment real, gaps are MCP HTTP-inherited limitations, undisclosed context pending'
  - 'EXPLORED-PARKED: exploration: Switchboard for MCP - agent-to-agent and agent-to-tool overlay network'
  - 'COMPLETED: technical-research: tmux control mode depth, console-side integration options'
  - 'COMPLETED: design: router control plane (topology, link-state, forwarding computation, distributed database)'
  - 'DEFERRED-TO-ARCHITECTURE: design: router management plane (config, monitoring, operations, upgrades)'
  - 'DEFERRED-TO-ARCHITECTURE: design: admission keying particulars (authentication, reauthentication, revocation propagation)'
session_bootstrap: |
  ## Session Bootstrap — Resume Instructions

  This brainstorming session is run under the analyst agent (Mary) using the
  brainstorming workflow. The user goes by "maker".

  ### Where We Are
  - **Phases complete:** Research Synthesis (10 topics), First Principles (8 truths),
    Values/Philosophy (10 values + death conditions), Morphological Analysis (12 parameters — ALL COMPLETE),
    Loss Recovery & QoS Technical Research (queued session #1 — COMPLETE),
    Downstream Strategy Technical Research (queued session #2 — COMPLETE)
  - **Next phase:** Product brief, then PRD. Remaining queued sessions (#6 management plane, #7 admission keying) deferred to architecture work.
  - **Session #3 (Switchboard for MCP): explored, parked** — undisclosed context pending.
  - **Session #4 (tmux control mode depth): complete** — console-side decision: CN-E configurable (CN-B/D MVP, CN-C post-MVP, CN-A long-term).
  - **Session #5 (router control plane): complete** — one mechanism (reliable flooding of signed messages) serves topology, admission, and membership.
  - **All 12 parameters have working directions chosen.**
  - **Parameter 2 now has concrete technique selections** — see Queued Session 1 Results.
  - **Parameter 6 downstream now decided: D-CE** (hybrid via content-type tiering) — see Queued Session 2 Results.

  ### Parameters Summary (1-12)

  | # | Parameter | Working Direction |
  |---|-----------|------------------|
  | 1 | Timer Regimes | Idle/Negotiation/Active; independent half-channel clocks; fixed tick MVP |
  | 2 | Loss Recovery | E/F hybrid duplicate+ARQ, adaptive (techniques need research) |
  | 3 | Degradation Signaling | Visual + clock step-down + advisory + graceful hold |
  | 4 | Channel Establishment | Multicast presence protocol, signed, 3 trigger types |
  | 5 | VSN Admission | Two-tier: admission keys (VSN) + session auth (access node), signed multicast |
  | 6 | Upstream/Downstream | Upstream: U-C idempotent replay. Downstream: ALL OPEN (research flagged) |
  | 7 | Router Forwarding | A (address) MVP → F (address+label) post-MVP; 3 multicast addresses per VSN |
  | 8 | tmux Integration | AN-E (control mode + PTY fallback); console: all options open, configurable |
  | 9 | Node Multi-homing | MH-C active/active, PS-D policy+latency, FT-D layered failover, ID-D crypto identity, IF-D upstream free/downstream TBD |
  | 10 | Frame Envelope | EL-B two-layer, router HMAC auth, ADDR-C VSN-scoped hash (8 byte), 44-byte outer + ~16-byte channel header, EXT-D versioned outer/TLV channel |
  | 11 | Router-to-Router | TD-E seed+gossip, SE-A link-state + SE-E control observer, latency+loss metrics, RS-D Noise protocol auth. Design space mapped, details are router control plane (undescribed) |
  | 12 | Control Node (VSN Mgmt) | Control node is a USER-FACING NODE TYPE, not infrastructure. Manages VSN lifecycle and key registration. VR-C first key bootstrapped OOB. KS-D key carries VSN+role, permissions at Tier 2. CM-A single control node MVP, CM-B peers later. CP-C submits changes as network service, distribution is the network's problem. |

  ### Critical Correction Applied in Parameter 12
  - **Control node is a node type that uses the network.** It does NOT manage routers.
  - **Three distinct planes exist:**
    1. User/data plane — nodes using the network (described)
    2. Router control plane — topology, link-state, forwarding, distributed database (NOT YET DESCRIBED)
    3. Router management plane — config, monitoring, operations (NOT YET DESCRIBED)
  - Console nodes CAN revoke/edit/expire keys for VSNs they have keys for
  - How routers maintain a network-distributed database for admission state, VSN membership,
    and revocation propagation is a foundational undescribed piece that multiple parameters depend on
  - Authentication and reauthentication particulars have not been described

  ### Open Questions from Parameter 12
  1. Can consoles register NEW keys, or only revoke/edit/expire?
  2. Can access nodes do any key management? (Instinct: no)
  3. Is there a permission hierarchy among keys? Can a console revoke a control node's key?

  ### Open Research Needs
  1. Loss recovery techniques (X.25, interleaving, hybrid ARQ+FEC)
  2. Downstream strategy (scrollback, graphics, tmux control mode depth)
  3. Console-side tmux integration options
  4. Switchboard for MCP exploration (agent-to-agent overlay)
  5. Router control plane design (distributed database is key dependency)
  6. Router management plane design
  7. Admission keying particulars (auth, reauth, revocation propagation)

  ### The Soul (Value #4)
  Simple, reliable sessions over unreliable, complex networks.

  ### Naming Still Open
  The third node type (tmux-publishing node) is called "access node" as working name.
  See: _bmad-output/brainstorming/naming-node-type-parking-lot.md
---

# Brainstorming Session Results

**Facilitator:** maker
**Date:** 2026-03-07

## Session Overview

**Topic:** Switchboard - a low-latency, multi-path, end-to-end encrypted tmux session routing architecture that establishes virtual switched networks (VSNs) over switchboard routers. Purpose-built for high-trust, low-latency remote CLI experience with persistent tmux session orchestration.

**Goals:** Naming (especially the tmux-publishing node type), architecture refinement, edge protocol design (ECC/fragmentation/reassembly/envelope optimization), failure mode analysis (jitter, flaky routes, retransmit/resync), and use cases/positioning.

### Architecture Context

**Key Components:**
- Three node types: (needs-a-name) publishers, consoles, controls
- Virtual Switched Networks (VSNs) with public/private key admission
- Routers: blind relays with link-state routing, latency-aware
- ATM-style timeslice frame assembly
- Dual fastest-path forwarding with loop prevention
- NTP-style unidirectional latency probes in keep-alives
- Multi-path burst + ECC at edges
- Multi-homed nodes with dynamic router switching
- Edge protocol handler for ECC, fragmentation, sequencing, reassembly, encryption envelope optimization

### Session Setup

_Session initialized with AI-recommended technique approach. Rich architectural context provided by maker covering networking, security, protocol design, and naming challenges. Research phase added per maker's request before brainstorming techniques begin._

## Phase 0: Research Synthesis

### 1. Low-Latency Overlay Routing Protocols

**Key Research:**
- **RON (Resilient Overlay Networks)** — Andersen, Balakrishnan, Kaashoek, Morris (MIT, SOSP 2001). The foundational paper on overlay routing for fault detection and latency optimization. RON nodes exchange path quality metrics via a routing protocol and build forwarding tables based on latency, loss rate, and throughput. Recovered from outages in <20 seconds vs minutes for BGP. ~5% of transfers doubled TCP throughput.
- **ShorTor** — Improved Tor latency via multi-hop overlay routing, demonstrating that even anonymity networks benefit from latency-aware relay selection.
- **Hop-by-hop multipath overlay routing** — Recent research on leveraging intermediate node routing decisions in WANs for resource optimization.

**Directly applicable to Switchboard:**
- RON's active probing + passive observation model maps directly to Switchboard's keep-alive latency measurement
- Link-state exchange between routers is a proven approach; RON validated it at scale
- Path selection based on measured latency (not topology) is the right approach

**Pitfalls:** RON found that probing overhead scales O(n^2) with nodes. Switchboard routers need to be smart about which paths to probe — full mesh probing won't scale to hundreds of routers.

---

### 2. ATM Lessons Learned

**Key Research:**
- **Kalmanek, "A Retrospective View of ATM"** (AT&T Labs, ACM CCR 2002) — authoritative postmortem
- **rule11.tech, "Technologies that Didn't: ATM"** — practical retrospective

**What worked (borrow these):**
- Cell-based switching with fixed timeslice assembly — predictable, low-jitter forwarding
- Virtual circuits (analogous to Switchboard's VSN channels) — path state established before data flows
- QoS guarantees through traffic shaping

**What failed (avoid these):**
- **48-byte cell size was a political compromise**, not a technical optimum. Not a power of 2. Led to massive overhead: providers saw 80% link utilization but <40% goodput due to cell tax and reassembly failures
- **Complexity killed it** — hardware/software implementation cost was enormous
- **Reassembly fragility** — out-of-order cells caused entire packets to be discarded
- **Rigidity** — couldn't adapt to variable payload sizes efficiently

**Critical lesson for Switchboard:** Variable-size frames are better than fixed cells. ATM's fixed 48-byte cell was its fatal flaw. Switchboard's timeslice-driven frame assembly (pack whatever's in the buffer) avoids this — frames naturally vary in size based on terminal activity. This is the right instinct.

---

### 3. Multi-Path Routing & Packet Scheduling

**Key Research:**
- **MPTCP (RFC 6824/8684)** — Linux kernel default scheduler uses lowest-RTT-first. Key insight: the scheduler is the most critical component for latency.
- **BLEST scheduler** — estimates buffer blocking to reduce head-of-line blocking across paths
- **STTF (Shortest Time to First Byte)** — sends via path expected to deliver first byte earliest. Best for latency-sensitive workloads.
- **LLHD scheduler** — achieved 25% higher throughput and 45% lower delay than Linux default
- **QUIC (RFC 9000)** — multiplexed streams over UDP, eliminates head-of-line blocking, connection migration via connection IDs, 0-RTT resumption

**Directly applicable:**
- Switchboard's dual-fastest-path forwarding aligns with MPTCP research showing that sending over 2 paths with intelligent scheduling beats single-path
- Receiver-side reordering buffer is mandatory — packets from different paths arrive out of order
- QUIC's connection ID concept maps to Switchboard's VSN channel identifiers for connection migration
- STTF scheduling philosophy matches Switchboard's "lowest latency path" approach

**Pitfalls:** MPTCP research consistently shows that naive round-robin across heterogeneous paths causes severe head-of-line blocking. Path selection MUST be latency-weighted.

---

### 4. One-Way Delay Measurement

**Key Research:**
- **OWAMP (RFC 4656)** — One-Way Active Measurement Protocol. Requires synchronized clocks via GPS/NTP at endpoints. Consists of OWAMP-Control (TCP) and OWAMP-Test (UDP).
- **PTP (IEEE 1588)** — Precision Time Protocol, sub-microsecond accuracy
- **NTP** — typically 1-10ms accuracy; sufficient for Switchboard's use case

**Directly applicable:**
- OWAMP's approach of embedding timestamps in probe packets is exactly what Switchboard's keep-alive packets should do
- Unidirectional measurement requires *some* clock synchronization — but NTP-level accuracy (milliseconds) is sufficient for terminal session routing decisions
- The keep-alive packet should carry: sender timestamp, receiver's last-known offset, and path identifier

**Key insight:** Perfect clock sync isn't needed. Switchboard only needs *relative* latency comparison between paths. If Path A consistently shows 20ms and Path B shows 45ms, you pick A — even if the absolute measurements are off by a few ms due to clock skew.

---

### 5. Forward Error Correction for Small Payloads

**Key Research:**
- **Simple XOR FEC (RFC 5109)** — XOR parity across packet groups. (k+1, k) code: one redundancy packet per k data packets. Minimal overhead, recovers single loss per group.
- **Reed-Solomon (RFC 5510/6865)** — Optimal erasure recovery, but computationally heavier. Source symbols are part of encoding (systematic code), simplifying decoding.
- **Fountain/Raptor codes (RFC 5053/6330)** — Near-optimal for large blocks, but overhead is poor for very small payloads (<1500 bytes). Designed for bulk transfer, not interactive.
- **LDPC** — Good for large blocks, poor for tiny payloads.

**Recommendation for Switchboard:**
- **XOR parity is the right starting point** for Switchboard's use case. For dual-path forwarding with tiny payloads (keystrokes = 1-100 bytes), simple XOR across the two path copies provides single-loss recovery with minimal overhead and near-zero computational cost.
- **Reed-Solomon** is the upgrade path if more than 2 paths are used or if loss patterns require recovering from multiple simultaneous losses.
- Fountain/Raptor codes are **overkill** — they're optimized for large data blocks, not keystroke-sized frames.

---

### 6. Retransmit & Resync Protocols

**Key Research:**
- **SRT (Secure Reliable Transport)** — Open-source by Haivision. Uses selective repeat ARQ + optional FEC. Key innovation: **Too-Late Packet Drop (TLPKTDROP)** — both sender and receiver drop packets that can't arrive in time. Retransmission is immediate upon loss detection, not timer-based.
- **Mosh SSP (State Synchronization Protocol)** — Synchronizes *objects* (terminal state), not byte streams. Idempotent datagrams with AES-OCB encryption. Frame rate controlled. Roaming via sequence-number-based IP address migration.
- **Eternal Terminal** — Uses BackedReader/BackedWriter with sequence numbers and buffered resend on reconnection. Supports tmux -CC control mode.

**Critical design insight from Mosh:** Mosh doesn't synchronize a stream — it synchronizes *state*. Every datagram is a diff between numbered states. This means lost packets don't cause gaps — the next datagram just carries a newer diff. This is brilliant for terminal display (server->client) but doesn't work for keystrokes (client->server), which are sequential.

**Recommendation for Switchboard:**
- **Server->client (terminal output):** Consider Mosh-style state synchronization for display updates — diffs of terminal state, not byte streams. Lost frames are naturally superseded.
- **Client->server (keystrokes):** Must be reliable and ordered. SRT's selective ARQ with TLPKTDROP is the right model — immediate retransmit on loss, drop frames that are too late.
- **Resync:** On sequence gap detection, trigger fast resync by having the receiver request a state snapshot (terminal state) rather than replaying lost frames.

---

### 7. Encryption Envelope Overhead

**Key Research:**
- **WireGuard** — 32 bytes total per-packet overhead (16-byte header + 16-byte Poly1305 auth tag). Over IPv4/UDP: 60 bytes outer header. Uses ChaCha20-Poly1305 (RFC 7539).
- **Noise Protocol Framework** — 16 bytes authentication overhead per transport message (Poly1305 tag). Max message 65535 bytes. Used by WireGuard, WhatsApp, Slack.
- **DTLS** — Heavier overhead than Noise/WireGuard due to TLS record layer headers.

**Recommendation for Switchboard:**
- **Target: 16-32 bytes crypto overhead per frame.** Noise/ChaCha20-Poly1305 is the gold standard for lightweight authenticated encryption.
- For Switchboard's tiny payloads (1-byte keystroke), 32 bytes of overhead on a 1-byte payload is 97% overhead. This argues strongly for **batching within the timeslice** — even a few milliseconds of batching can aggregate multiple keystrokes or terminal bytes to improve the payload-to-overhead ratio.
- The E2E SSH encryption is already present; Switchboard's frame envelope should add minimal additional overhead. Consider whether the SSH layer can serve as the sole encryption, with the Switchboard frame adding only authentication + sequencing (not double-encrypting).

---

### 8. Similar Projects Analysis

| Project | Architecture | What Switchboard Borrows | What Switchboard Improves |
|---------|-------------|------------------------|--------------------------|
| **Mosh** | UDP, SSP state sync, AES-OCB, single server | State sync model for display, roaming via seq numbers | Multi-path, multi-hop routing, network-level resilience |
| **Tailscale** | WireGuard overlay, DERP relays, coordination server | DERP blind-relay model (routers can't see content), NAT traversal, latency-based server selection | Purpose-built for terminal sessions, no general IP tunneling overhead |
| **ZeroTier** | VL1 (crypto peer-to-peer) + VL2 (ethernet emulation), planet/moon/leaf | VSN admission model (cryptographic identity + network ID), moon concept for private infrastructure | No ethernet emulation overhead, terminal-specific optimization |
| **Nebula** | Noise protocol, lighthouse discovery, certificate-based auth, UDP hole punching | Certificate-based VSN admission, lighthouse = Switchboard control nodes | Multi-path active forwarding vs single-path, latency-optimized routing |
| **Eternal Terminal** | SSH + custom reconnection, BackedReader/BackedWriter, tmux -CC support | Sequence-numbered buffered resend, tmux control mode integration | Network-level routing, multi-path, not just reconnection |
| **tmate** | tmux fork, relay servers (SF/NY/London/Singapore), msgpack+gzip over SSH | Relay server geographic distribution, jailed tmux server isolation | True multi-hop routing vs static relay, E2E encryption (tmate relays can see content) |
| **QUIC** | UDP, multiplexed streams, connection migration, 0-RTT, TLS 1.3 built-in | Connection migration via IDs, stream multiplexing, 0-RTT concepts | Terminal-specific framing, multi-path forwarding (QUIC is single-path) |

**Key gap Switchboard fills:** No existing tool combines multi-path latency-optimized routing + E2E encrypted terminal sessions + virtual network admission + purpose-built framing for interactive CLI. They all solve *parts* of this problem.

---

### 9. Language Assessment: Go vs Rust

**Go (current project language):**

| Factor | Assessment |
|--------|-----------|
| **GC pauses** | Go 1.24: 15-25% pause reduction. Go 1.25: experimental 10-40% GC overhead reduction. Sub-millisecond pauses typical. For terminal sessions (human-perceptible latency ~50-100ms), GC pauses are well within tolerance. |
| **Concurrency** | Goroutines excel at managing thousands of concurrent connections (router use case). Scheduling overhead exists but is negligible for this workload. |
| **Crypto** | Go stdlib crypto is well-optimized. ChaCha20-Poly1305 available. |
| **Containers** | Static binaries, small images (scratch base), instant startup. Excellent for EKS/k8s. |
| **Cross-compilation** | `GOOS=darwin GOARCH=arm64` just works. macOS + Linux from single build. |
| **Precedent** | Tailscale, WireGuard-go, Nebula — all production overlay networking in Go. Tailscale recently achieved >10Gbps with wireguard-go, surpassing kernel WireGuard. |

**Rust:**

| Factor | Assessment |
|--------|-----------|
| **Latency** | Zero GC, deterministic memory. Eliminates tail-latency jitter entirely. |
| **Performance** | Cloudflare chose Rust (Pingora) for proxy infrastructure, citing memory safety + performance. |
| **Downsides** | Steeper learning curve, slower iteration, smaller contributor pool. Compile times. async ecosystem complexity (tokio runtime choices). |
| **Containers** | Even smaller binaries than Go. musl static linking. |

**Recommendation: Go is the right choice. Here's why:**

1. **The latency budget is human-perceptible (>10ms round-trip is fine).** Go's sub-ms GC pauses are invisible at this scale. Rust's zero-GC advantage matters for microsecond-sensitive workloads (HFT, game servers at 120fps tick rate), not terminal sessions.
2. **Tailscale, Nebula, and WireGuard-go prove Go works** for exactly this class of problem — encrypted overlay networking with latency sensitivity.
3. **Development velocity matters for MVP.** Go's simplicity means faster iteration. The project already has Go scaffolding.
4. **Operational simplicity for containers.** Go's single static binary with no runtime dependencies is ideal for EKS/k8s deployment.
5. **macOS support is trivial.** Go cross-compiles to darwin/arm64 and darwin/amd64 with zero friction.

**Consider Rust only if:** Profiling reveals that GC pauses cause perceptible jitter in production, OR if the edge protocol handler's ECC computation becomes a bottleneck. In that case, a Go program calling a Rust shared library (via CGo/FFI) for the hot path is a viable hybrid.

---

### 10. Deployment Context Established

**Routers (server-side):**
- Container images: Go static binary on `scratch` or `distroless` base
- Target platforms: AWS EKS, EC2 (Docker), k8s, docker-compose
- Single binary per container, no dependencies
- Health checks via keep-alive metrics endpoint
- Horizontal scaling: add more router containers to increase network capacity

**Nodes (client-side MVP):**
- macOS only (darwin/arm64 + darwin/amd64)
- Short-term: add Linux (amd64 + arm64)
- Distributed as single binary (Homebrew tap, direct download)
- Integrates with local tmux via tmux control mode or socket

---

## Phase 1: First Principles Thinking

### Fundamental Truths Established

1. **Human perception is the latency budget.** Keystroke-to-echo under ~100ms feels instant, under ~200ms feels responsive, over ~500ms feels broken. This is neuroscience, not negotiable.

2. **Terminal sessions are asymmetric.** Two half-channels with fundamentally different characteristics:
   - **Upstream (console to access node):** Keystrokes. Tiny, ordered, loss-intolerant.
   - **Downstream (access node to console):** Terminal output. Bursty, larger, state-syncable (Mosh-style diffs viable).

3. **The bus leaves on time, full or not.** Timeslice-driven framing, not packet-driven. Each direction has its own independent clock at the node. No clock synchronization needed between sender and receiver (full duplex). Pack whatever bytes are ready when the tick fires and ship them. Analogous to WAN/satellite acceleration (SKIPS protocol) — minimize idle time across the latency chasm. Frame size is variable, bounded by buffer contents at tick time. Tick interval is tunable: aggressive (5-10ms) for active typing, relaxed (50-100ms) when idle.

4. **No direct node-to-node communication.** Always through a router. This is what makes it a switched network, not a mesh. This is an architectural invariant.

5. **One router binary, topology determines role.** Three deployment modes, not three binaries:
   - **E (Edge-local):** Runs on the same device as a node, zero upstream router connections. Same-LAN quick setup between co-located machines. The "getting started in 5 minutes" experience.
   - **PE (Provider Edge):** Node-facing + router-facing interfaces. Handles VSN admission, node keep-alives, channel setup. The production router.
   - **P (Provider/Core):** Router-facing interfaces only. Pure forwarding, no node awareness. **Theoretical only** — not built until simulation or production data proves it earns its keep.
   - An E router that gets pointed at an upstream PE router graduates to PE. Same binary, different config.
   - The P/PE distinction is an **interface-level concept**, not a binary-level concept. A router has PE interfaces (node-facing) and P interfaces (router-facing). A "P router" is simply one with zero PE interfaces configured.

6. **SSH E2E is the trust layer.** Switchboard adds routing and admission, not encryption. The payload is already SSH-encrypted between nodes. Switchboard's frame envelope should add auth/integrity/sequencing, not double-encrypt. Lighter techniques (HMAC, auth tags) may suffice for frame-level validation between node and PE router.

7. **Circuit signaling is a solved problem.** MPLS LDP/RSVP-TE provides decades of production-proven label distribution and path establishment. When multi-hop circuits are needed, study and adapt rather than invent. Label-switching model for core forwarding (PE pushes/pops labels, P routers switch on labels without parsing frame internals) is the direction — but deferred until post-MVP.

8. **The MVP is nodes + E router.** Proves out the dominant use case (intra-LAN, single-hop) and the core value:
   - Edge protocol (framing, timeslice clock, sequencing)
   - Channel model (half-channels, asymmetric handling)
   - Key-based VSN admission (local)
   - Access node + console interaction model
   - The actual user experience — does this feel better than raw SSH?
   - Deliberately defers: multi-hop routing, multi-path forwarding, FEC/ECC, control node, router-to-router link-state, MPLS-inspired circuit establishment.

### Architectural Insights

**Router scaling:** The P/PE interface distinction has implications for cloud scaling. PE interfaces carry node protocol overhead (admission, keep-alives, edge framing). P interfaces carry only forwarding load. In a k8s environment, scaling out P-interface-only containers to handle traffic bursts is lighter than scaling full PE routers. This optimization is deferred but the interface separation keeps the door open.

**Network admission model (open questions, parked for later):**
- Nodes authenticate to PE routers via VSN credential (private key signs challenge)
- After auth, PE router creates a forwarding entry with TTL (NAT-table analogy)
- Sideband refresh: periodic re-auth to keep the entry alive
- IP migration: re-auth on sideband with new source, channel continues (identity is cryptographic, not network-based, similar to QUIC connection IDs)
- Challenge-resync: PE can challenge a node at any time to guard against session hijacking
- Inter-router security: P routers should speak only to known routers. IP filtering alone is insufficient (MITM). Encrypted channels between routers likely necessary — but parked.

**Traffic engineering (parked for later):**
- Flow control, jitter detection, latency measurement, packet drop detection
- Estimated traffic distribution: ~85% intra-LAN (E router), ~14% single-hop/single-PE bridging, ~1% multi-path/multi-router. The MVP covers the 85%.
- P routers at major IaaS POPs/AZs (AWS, etc.) is a future deployment pattern for the 1%.

### Channel Model

- **Channel:** An end-to-end active tmux session over SSH over VSN over Switchboard, between a console and an access node.
- **Half-channel (upstream):** Console to access node. Keystrokes. Own independent timeslice clock, sequence numbers, FEC strategy.
- **Half-channel (downstream):** Access node to console. Terminal output. Own independent timeslice clock, sequence numbers, FEC strategy.
- **Circuit identity:** When multi-hop is introduced, MPLS-style label handles negotiated via signaling. The node+PE path on each side acts as an address for circuit endpoints. Deferred to post-MVP.

### Naming Parking Lot Update

- Node types: access node (publisher), console, control — established in Phase 0
- Router modes: E (edge-local), PE (provider edge), P (provider core) — established in Phase 1
- See: `_bmad-output/brainstorming/naming-node-type-parking-lot.md`

---

## Philosophy: The Soul of Switchboard

_Values exploration session — what Switchboard believes, who it serves, what it refuses to become._

### The One-Line Soul

**Simple, reliable sessions over unreliable, complex networks.**

### Values

1. **Secure as OpenSSH.** End-to-end encrypted. The network cannot observe, inject, or tamper. Non-negotiable.

2. **Rock solid.** Self-healing, set-and-forget. No mystery failures. Concrete under your feet.

3. **The network is a fact, not a mystery.** When physics intrudes, we're transparent. Never mysterious.

4. **Simple, reliable sessions over unreliable, complex networks.**

5. **Simple things are simple.** Complexity is available, never required. Two machines, one user, five minutes.

6. **Function over form.** We carry keystrokes and terminal output. We optimize for exactly that, better than anyone.

7. **The illusion of local.** Your session feels like it's on your machine, no matter where it runs. The network disappears. Like a fly-by-wire pilot — hands on controls, aircraft responds, the distance is invisible. Switchboard extends that illusion around the world to any CLI, TUI, or shell, without thinking about which computer the session runs on.

8. **Terminal professionals are who we serve.** People who live in shells. Who think in CLI. Whose work happens at a prompt. Platform engineers, network admins, AI ops professionals managing fleets of supervisor-agents across machines worldwide.

9. **Sessions are sovereign, infrastructure is shared.** Your control channel belongs to you. Not Slack, not Discord, not AWS, not anyone. Multi-tenant routers, cryptographically isolated sessions. Even root on the router gives you nothing — no MITM, no keystroke injection, no observation. Without the data keys, the only attack is availability or quality.

10. **Resilience is the architecture.** A router going down is a routing event, not an outage. Rolling updates, per-VSN isolation. Never everyone's problem.

### Project Decisions (not values, but true)

- tmux-first, depth before breadth — be great at tmux before considering screen, Zellij, raw shell
- Implementation earns the protocol — battle-test through generations before standardizing
- OpenSSH key pairs as the credential system
- Go as the implementation language

### The Experience

Freedom to breathe and think. Will becoming action with low friction. Focus on what to manifest, not how to connect. The feeling of a remote control operator — hands on controls, the system responds, the distance is invisible.

### Death Conditions (Switchboard has failed if)

1. **The UX breaks flow state.** Stutters, hangs, freezes. If the user notices the network, Switchboard has failed.
2. **Session security is compromised** by anything other than stolen SSH keys.
3. **Operational fragility.** Can't do rolling updates, vulnerability requires taking the whole network down, affects all users simultaneously.
4. **Complexity barrier.** The solo operator with two machines can't set up and forget in minutes.

### Who We Serve

Terminal professionals who need to access multiple remote sessions reliably, regardless of where they are and where their keyboard is. Specifically:
- Platform engineers operating infrastructure
- Network admins operating complex networks
- AI-assisted devops professionals managing dozens of supervisor-agents and teams of agents across machines that could be anywhere in the world
- Anyone who lives in tmux and needs it to work everywhere, always

### Feature Surface Discovered: Access Modes

| Mode | Direction | Use Case |
|------|-----------|----------|
| **Full access** | Bidirectional (read-write) | Normal interactive session |
| **Read-only console** | Downstream only (view, no keystrokes) | "Come look at this", monitoring |
| **Read-only per-session** | Downstream only, scoped to specific tmux session | Publish one session without exposing the node |
| **Read-only per-node** | Downstream only, all sessions on a node | Monitor everything on a machine |
| **Published view** | Read-only, flagged for display | SOC video wall, team dashboard, shared TUI |

Connects to Tier 2 session auth (Parameter 5) — authorized key list needs a permission level (read-write vs. read-only), not just allow/deny. Access mode is bound to the key.

### Performance Philosophy

- Keystrokes punch through fast — no unnecessary milliseconds of latency
- Bulk data downstream doesn't feel like 2400 baud — if we need data dictionaries, zcompression at the nodes, fast-path routing for small bytes, we do that
- We're not a video streaming service — we know what we carry and we optimize for exactly that
- Copy/paste, LSP, image uploads over terminal are bulk operations with different characteristics than keystrokes — handle them differently

---

## Phase 2: Morphological Analysis

_Systematic exploration of Switchboard's key architectural parameters and their design options. Each parameter maps a design dimension with viable options, a working direction, and notes on what's parked for future study._

### Parameter 1: Timer Regimes

The system operates in distinct timing modes, not a single clock:

| Regime | Trigger | Timer Behavior | Purpose |
|--------|---------|---------------|---------|
| **Idle** | No attached tmux sessions | NAT keepalive / heartbeat / latency probe cadence | Stay alive through NATs, measure path quality, detect failures |
| **Negotiation** | Session inventory request, pre-attach | Request/response cadence for listing remote tmux sessions, channel setup | Discovery & establishment |
| **Active Session** | Attached to tmux session | Core loop — tunable tick, independent per half-channel | The main event — frame assembly & shipping |

**Active Session sub-parameters (parked for future study):**
- Tick interval — tunable per network conditions (fixed for now)
- Multi-path link count — how many parallel paths
- Overlap % — how much redundant data across paths
- Offset — stagger cut points across paths so frames don't fragment at the same byte boundary
- Timeshift for multi-path — possibly stagger *when* each path's frame fires, not just *where* it cuts

**Working direction:** Independent half-channel clocks (E), fixed tick for MVP (A). Options B (activity-adaptive), C (latency-adaptive), D (payload-adaptive) remain in design space for future study.

---

### Parameter 2: Loss Recovery Strategy

**Working direction:** E/F hybrid — duplicate-and-race on fastest paths with ARQ fallback, adaptive switching based on measured conditions.

**Options explored:**
- A. Duplicate-and-race (send identical frame on all paths, first arrival wins)
- B. Staggered redundancy (overlap % with offset across paths)
- C. XOR parity path (k data + 1 parity)
- D. Selective retransmit / ARQ (SRT-style with TLPKTDROP)
- E. Hybrid: duplicate + ARQ fallback
- F. Adaptive (switch strategy based on measured loss rate)

**Key insight — the reassembly window:** The clock creates a known time budget on the receive side. Because you know what was sent, when, and how long paths take, you get a window between expected arrival and the MUST-deliver deadline. Within that window: compare path copies, request retransmit if needed, decide whether to deliver partial data and signal degradation.

**Philosophy:** Guaranteed level of service over unreliable service (the internet). When quality can't be maintained, be honest — signal degradation rather than silently degrade. Options include stepping down clocks, warning users to reduce refresh rate or window size.

**Research flagged:** Specific techniques TBD. Need targeted research on X.25 guaranteed delivery (LAP-B, sliding window, selective reject), interleaving vs. duplication (digital radio/satellite techniques), and hybrid ARQ+FEC decision models (SRT, QUIC, WebRTC).

---

### Parameter 3: Degradation Signaling

_Emerged from Parameter 2 discussion — elevated to its own parameter._

| Signal | Description |
|--------|-------------|
| **Visual indicator** | Status bar / tmux status line showing connection quality (green/yellow/red) |
| **Clock step-down** | Automatically reduce tick rate to give more reassembly headroom |
| **Advisory** | Suggest user actions: smaller window, lower refresh, switch to batch mode |
| **Graceful hold** | Freeze display briefly rather than show corrupt/partial output |

---

### Parameter 4: Channel Establishment — Multicast Presence Protocol

**Working direction:** VSN-scoped multicast presence protocol over signed advertisements (Option B from original exploration).

**Who advertises what:**

| Node Type | Advertises |
|-----------|-----------|
| **Access node** | Available tmux sessions, which consoles are attached to which sessions, connection quality per console |
| **Console** | Its own presence on the VSN, which access node/session it's attached to, connection quality |
| **Control** | Listens and aggregates; may advertise VSN-level state |

**When they advertise:**

| Trigger | Description |
|---------|-------------|
| **On change** | New tmux session created, session closed, console attaches/detaches, quality state change |
| **Scheduled** | Periodic heartbeat — full state refresh on a cadence |
| **On request** | Multicast ping → all nodes respond with current state |

**Properties:** No central dependency. Eventual consistency (self-corrects on heartbeat or ping). Connection quality is first-class data visible to all VSN members. Natural fit for E router MVP (LAN multicast is trivial); scaling across PE routers to WAN is a known-hard problem for post-MVP.

**Open questions parked:** Multicast scope across router boundaries. Advertisement payload size (full state vs. deltas). Security of session metadata visible to all VSN members.

---

### Parameter 5: VSN Admission & Session Identity

**Working direction:** Two-tier key model.

**Tier 1 — VSN Admission (network level):**
- Control node registers public keys against a VSN identifier
- Node presents: private key signature + VSN identifier → if public key is registered, node is admitted
- Key flexibility: one keypair can serve both access node and console roles, OR separate keys per role, OR multiple nodes sharing a VSN with different keys
- Admission gets you on the network — you can see multicast traffic and exist as a presence

**Tier 2 — Session Authorization (tmux level):**
- Access node maintains a list of authorized console public keys
- Console requests tmux session → access node checks console's public key
- Granularity TBD (per-access-node? per-session? per-tmux-session-name?)

**Flow:**
```
Console → [Tier 1: prove VSN membership to PE/E router] → on the VSN
       → [multicast: discover available sessions]
       → [Tier 2: request tmux attach, access node checks console's pubkey] → attached
```

**Multicast envelope:** Signed (Option B) for MVP — if you've passed Tier 1 admission, you're trusted to see network state, but nobody can forge advertisements. Full envelope approach (what's signed vs. encrypted vs. clear) connects to philosophy session — what is the trust boundary?

---

### Parameter 6: Upstream/Downstream Asymmetry

**Upstream — Working direction: U-C (idempotent replay)**
Each upstream frame carries the last N keystrokes, not just new ones. Receiver deduplicates. Loss is self-healing without retransmit. Overhead is negligible — a 10-keystroke sliding window is ~10-20 extra bytes on a frame that already has 16-32 bytes of crypto envelope.

**Downstream — ALL OPTIONS OPEN, research flagged:**
- D-A. Reliable ordered stream
- D-B. Mosh-style state sync (diffs between numbered terminal states)
- D-C. Hybrid: state sync + reliable for scrollback
- D-D. Snapshot + delta (keyframes + diffs, like video encoding)
- D-E. Tiered by content type (detect interactive prompt vs. log flood vs. file transfer)

**Why downstream is hard — content diversity:**

| Content Type | Characteristics | Implications |
|-------------|-----------------|--------------|
| Interactive prompt | Small, frequent updates | State sync works well |
| Log flood | Massive burst, sequential | User needs scrollback — can't discard intermediate states |
| Claude Code output | Long streaming, scrollback critical | Every line matters — closer to upstream semantics |
| TUI applications | Full screen redraws | Snapshot+delta natural fit |
| Image display | Sixel, iTerm2 inline, Kitty graphics protocol | Binary, potentially large, loss-intolerant |
| Overlays | tmux popups, fzf, floating panes | Layered rendering — diffs get complex |

**Research flagged:** tmux control mode internals, scrollback buffer semantics, Sixel/Kitty graphics interaction with state diffing, content-type detection feasibility.

---

### Parameter 7: Router Forwarding Model

**Working direction:** A (address-based) for MVP → F (A for control/signaling + B/E label/channel-switching for data) post-MVP. Revisit flag set.

**Options explored:**
- A. Address-based forwarding (destination node ID → next-hop lookup)
- B. Label switching (MPLS-inspired, PE pushes/pops labels, P routers swap only)
- C. Source-routed (full path in header, each router pops and forwards)
- D. VSN-scoped flooding (NOT NEEDED — address-based with multicast addresses covers this)
- E. Channel-ID switching (per-channel granularity, functionally similar to B)
- F. Hybrid: A for control, B/E for data

**Key insight: D (flooding) eliminated.** Address-based forwarding with reserved multicast addresses replaces flooding entirely.

**Multicast address model (three per VSN):**

| Address | Who Sends | Who Listens |
|---------|-----------|-------------|
| `[VSN-ID]:[access-nodes]` | Access nodes (sessions, quality) | Consoles, control |
| `[VSN-ID]:[consoles]` | Consoles (presence, attachments, quality) | Access nodes, control |
| `[VSN-ID]:[control]` | Control (policy, admission updates, network state) | Access nodes, consoles |

**Control node special behavior:** Listens on all three multicast address types. Can join multiple VSNs simultaneously. Only node type with full cross-VSN visibility.

Router treats multicast as "forward to all interfaces with subscribers to this address" — just multiple entries in the same forwarding table. No special flooding logic needed.

---

### Morphological Matrix Summary (Parameters 1-8, inline details below; Parameters 9-12 at end of document)

| # | Parameter | Working Direction | Status |
|---|-----------|------------------|--------|
| 1 | Timer Regimes | Idle / Negotiation / Active; independent half-channel clocks; fixed tick MVP | Decided (adaptive parked) |
| 2 | Loss Recovery | E/F hybrid duplicate+ARQ, adaptive switching | Direction set, techniques TBD (research flagged) |
| 3 | Degradation Signaling | Visual + clock step-down + advisory + graceful hold | Direction set |
| 4 | Channel Establishment | Multicast presence protocol, signed, three trigger types | Decided |
| 5 | VSN Admission | Two-tier: admission keys (VSN) + session auth (access node), signed multicast | Decided (envelope details TBD) |
| 6 | Upstream/Downstream | Upstream: U-C idempotent replay. Downstream: D-CE hybrid via content-type tiering (MVP: D-A reliable stream) | Decided |
| 7 | Router Forwarding | A (address) MVP → F (address+label) post-MVP; 3 multicast addresses per VSN | Decided (revisit flag set) |
| 8 | tmux Integration | AN-E (control mode + PTY fallback). Console: CN-E configurable (CN-B/D MVP, CN-C post-MVP, CN-A long-term) | Decided |
| 9 | Node Multi-homing | MH-C active/active, PS-D policy+latency, FT-D layered failover, ID-D crypto identity, IF-D upstream free/downstream TBD | Decided |
| 10 | Frame Envelope | EL-B two-layer, router HMAC, ADDR-C 8-byte VSN-scoped hash, 44-byte outer + ~16-byte channel, EXT-D versioned outer/TLV channel | Decided |
| 11 | Router-to-Router | TD-E seed+gossip, SE-A link-state + SE-E observer, latency+loss metrics, RS-D Noise auth | Design space mapped, router control plane TBD |
| 12 | Control Node (VSN Mgmt) | User-facing node type. VR-C OOB bootstrap, KS-D VSN+role, CM-A→CM-B, CP-C network service | Decided, open questions parked |

### Parameter 8: tmux Integration Model

**Access Node — Working direction: AN-E (hybrid: control mode + PTY fallback)**
Use tmux control mode (`-CC`) when available (tmux 1.8+), fall back to PTY proxy for older tmux or non-tmux sessions. Control mode provides structured, machine-readable output (`%begin`/`%end` blocks, `%output` notifications, programmatic pane access). PTY proxy captures raw terminal byte stream — simple, zero tmux-specific knowledge required.

**Console Side — ALL OPTIONS OPEN, configurable, research flagged:**

| Option | Description |
|--------|-------------|
| **CN-A. Control mode consumer** | Parse control mode output, reconstruct terminal state locally. Enables local scrollback, local resize, client-side intelligence. |
| **CN-B. PTY injection** | Pipe received bytes into local PTY. Terminal emulator renders. Switchboard invisible. Graphics pass through naturally. |
| **CN-C. Local tmux session** | Local tmux mirrors remote. Native scrollback, copy-paste, mouse. "Illusion of local" — strongest version. Requires tmux on both ends. |
| **CN-D. Direct terminal write** | Write bytes to stdout. Simplest. No local intelligence. |

Working direction: likely a configurable subset (CN-E), auto-detected based on local environment. Needs study.

**Key connection:** Console-side choice directly affects downstream strategy (Parameter 6). Control mode enables structured state sync (Mosh-style). PTY proxy means raw byte stream — state sync much harder.

**Potential angle discovered:** Switchboard's console as a general-purpose tmux control mode consumer — not locked to iTerm2 like today's only implementation. Significant engineering investment but strongly aligned with "the illusion of local" value.

**Research flagged:**
- Control mode protocol depth — what structured data is available, completeness
- Graphics passthrough: control mode vs. PTY — does control mode preserve Sixel/Kitty or strip it
- Local tmux mirroring feasibility — bidirectional state sync complexity
- Console auto-detection of best mode based on local environment
- Eternal Terminal's experience with tmux `-CC` support

---

### Parameter 9: Node Multi-homing

**Working direction:** MH-C (active/active), PS-D (policy+latency), FT-D (layered failover), ID-D (crypto identity), IF-D (upstream free via idempotent replay, downstream TBD).

A node maintains live connections to multiple routers simultaneously. It sends data on multiple paths (duplicate-and-race originates at the node, not inside the router mesh). Policy sets router preferences (prefer local E, prefer specific PE for specific VSNs), latency measurements break ties and drive real-time path selection.

**Connection cardinality (MH-C active/active):** The node is a multi-path sender. This is a meaningful complexity decision — the node isn't just "connect to a router and let the network handle it." The edge protocol handler at the node needs multi-path awareness.

**Path selection (PS-D policy+latency):** Policy sets preferences (prefer local E router, fall back to PE, prefer router X for VSN Y). Latency measurements (keep-alive RTT) break ties and trigger failover.

**Failover (FT-D layered):** Three layers, from most graceful to last resort:
1. Router-initiated drain (FT-C) — router signals "going away," node migrates before disconnect. Enables zero-downtime rolling updates.
2. Degradation threshold (FT-B) — latency or loss exceeds threshold, preemptive switch.
3. Keep-alive timeout (FT-A) — N missed keep-alives, router declared dead.

**Identity (ID-D crypto identity):** Node identity IS its keypair. Any router serving the VSN recognizes the node by its public key. No migration ceremony — node proves key ownership via challenge-response. Pre-authentication to standby routers (ID-B optimization on ID-D) completes the challenge-response in advance so switchover is instant.

**In-flight frames (IF-D):** Upstream switchover is free — idempotent replay (Parameter 6 U-C) means the next frame on any path carries the full sliding window of last N keystrokes. Downstream switchover inherits whatever strategy emerges from the downstream research (Parameter 6, still open).

**Key implications:**
1. The node is an active participant in path diversity — edge protocol handler must be multi-path aware
2. Pre-auth to standby routers is worth doing as MVP optimization
3. Router drain signal needs a wire format (small addition, huge operational value)
4. Downstream switchover is coupled to Parameter 6 downstream strategy

---

### Parameter 10: Frame Envelope Structure

**Working direction:** EL-B (two-layer), router HMAC auth, ADDR-C (VSN-scoped hash, 8 bytes), 44-byte outer + ~16-byte channel header, EXT-D (versioned outer, TLV channel).

**Layering (EL-B two-layer):**
- **Outer layer (router-visible):** Addressing, frame type, length, HMAC auth tag. Routers parse this to forward.
- **Inner layer (endpoint-only):** Channel header + SSH-encrypted payload. Routers cannot read this.

SSH is the E2E trust layer (First Principle #6). Switchboard adds auth/integrity/sequencing, not double-encryption.

**Router HMAC auth tag:** The router verifies frames came from a legitimate VSN member. A VSN is a cryptographic trust boundary — non-members' frames are rejected at the first router, not forwarded to be rejected at the endpoint. Lightweight HMAC, not a full signature.

**Address format (ADDR-C VSN-scoped hash):** `hash(VSN-ID || pubkey)` truncated to 8 bytes. Self-derived from cryptographic identity (ID-D). No assignment authority needed. 2^64 collision resistance per VSN — sufficient for any realistic VSN size. ADDR-D (hash + short label) deferred to post-MVP label-switching.

**Outer header (44 bytes):**

| Field | Bytes | Purpose |
|-------|-------|---------|
| Version | 1 | Protocol evolution |
| Frame type | 1 | Data, keep-alive, control, drain signal, multicast advertisement |
| VSN ID | 8 | Hash-derived, scopes forwarding table |
| Destination | 8 | ADDR-C |
| Source | 8 | Return path + router auth verification |
| Frame length | 2 | Variable frames need explicit length |
| HMAC tag | 16 | Router-verifiable admission proof |
| **Total** | **44** | |

No TTL/hop-limit — unnecessary for E-router MVP (single hop). Added when multi-hop arrives via version bump.

**Channel header (~16 bytes, endpoint-only):**

| Field | Size (est.) | Purpose |
|-------|-------------|---------|
| Channel ID | 2-4 bytes | Which half-channel |
| Sequence number | 4 bytes | Ordering and loss detection |
| Timestamp | 4-8 bytes | Sender's tick timestamp |
| FEC metadata | 1-2 bytes | Parity group, position, type |
| Flags | 1 byte | Degradation signal, priority, fragmentation |

Replay window (upstream idempotent keystrokes) and TLV extensions follow the fixed core.

**Extensibility (EXT-D):** Outer header is versioned — fixed format, fast parsing, version bump for changes. Channel header allows TLV extensions — endpoints are smarter, updated more frequently, can experiment without router upgrades.

**Frame size budget (single keystroke worst case):** ~96 bytes total (44 outer + 16 channel + 36 SSH-encrypted keystroke). At 100 keystrokes/second peak typing: 9.6KB/s. Bandwidth is trivial; latency is what matters, and a 96-byte UDP frame traverses any link in microseconds.

---

### Parameter 11: Router-to-Router Protocol

_Design space mapped. Details are router control plane — not yet fully described. This parameter establishes the direction; implementation is a queued design session._

**Working direction:** TD-E (control-node seed + gossip), SE-A (link-state) + SE-E (control node as observer), latency+loss minimum viable metrics, triggered updates + dampening for convergence, RS-D (Noise protocol) for router-to-router auth.

**Topology discovery (TD-E):** Bootstrap with known routers (control node is the natural seed — routers register with it, it tells them about each other). Gossip for runtime discovery. Static config (TD-A) as first step for initial multi-router deployments.

**State exchange (SE-A link-state + SE-E observer):** Each router floods local link costs (latency, loss) to all routers. Every router computes full topology independently. Control node listens and aggregates for monitoring/dashboards but does NOT compute forwarding — routers do that locally. At 2-20 routers (realistic trajectory), full link-state is proven and well-understood.

**Minimum viable metrics:** One-way latency + loss rate. Jitter derived from latency measurements. Capacity and admin cost are future additions. All metrics observable and reportable — "the network is a fact, not a mystery" (Value #3).

**Convergence:** Triggered updates (not periodic-only) for fast reconvergence. Dampening for flapping links. Target: alternative paths found within ~1-2 seconds of router failure. Link-state achieves this easily at scale of 2-20 routers.

**Security (RS-D Noise protocol):** Static keypairs, Noise handshake. Same crypto framework as node-to-router. No PKI overhead. Consistent cryptographic model across the whole system.

**Convergence scenarios:**
- Router comes up → registers, establishes adjacencies, exchanges link-state, begins forwarding
- Router dies hard → neighbors detect via keep-alive timeout, flood updated link-state, reconverge
- Router drains gracefully → drain signal to neighbors (FT-C from Parameter 9), neighbors update before disappearance
- Link degrades → measured metrics cross threshold, updated link-state, network reconverges
- Flapping → dampening suppresses rapid state changes

---

### Parameter 12: Control Node — VSN Management

_Re-scoped after critical correction. Control node is a user-facing node type, not infrastructure. It uses the network; it does not manage it._

**Critical distinction — three planes:**
1. **User/data plane** — nodes (access, console, control) using the network. Sessions, VSN registration, key management. **This is what we've been designing.**
2. **Router control plane** — topology, link-state, forwarding computation, distributed database. **Not yet described.**
3. **Router management plane** — config, monitoring, operations, upgrades. **Not yet described.**

**What the control node does:**
- Register VSNs (if authorized)
- Register keys against VSNs — granting access to other control nodes, consoles, and access nodes
- Revoke/edit/expire keys
- Deregister VSNs

**Console nodes can also** revoke/edit/expire keys for VSNs they have keys for. Key management is not exclusive to control nodes.

**Working direction:**

**VSN registration (VR-C):** First control key bootstrapped out of band (at the router management plane). After that, control nodes self-manage — they register additional keys, other control nodes, etc. The network doesn't decide policy about who creates VSNs; that's an operational decision at the management plane.

**Key scope (KS-D):** Key carries VSN membership + role (control, console, access node). Access modes (read-only, read-write, etc. from the Feature Surface) are Tier 2 — negotiated between console and access node. Tier 1 stays clean: are you on this network, and as what type of node?

**Multiplicity (CM-A → CM-B):** Single control node per VSN for MVP. Multiple equal peers as natural growth path for multi-operator VSNs.

**How changes reach the network (CP-C):** Control node submits key changes as a network service request. How routers distribute and store that change is a router control plane problem — the undescribed distributed database. The control node's contract: "I'm authorized to make this change. Here's the change. You figure out distribution."

**Open questions (parked):**
1. Can consoles register NEW keys, or only revoke/edit/expire existing ones?
2. Can access nodes do any key management? (Instinct: no)
3. Is there a permission hierarchy among keys? Can a console revoke a control node's key?

---

### Morphological Matrix — Complete Summary

| # | Parameter | Working Direction | Status |
|---|-----------|------------------|--------|
| 1 | Timer Regimes | Idle / Negotiation / Active; independent half-channel clocks; fixed tick MVP | Decided (adaptive parked) |
| 2 | Loss Recovery | E/F hybrid duplicate+ARQ, adaptive switching | Direction set, techniques TBD (research flagged) |
| 3 | Degradation Signaling | Visual + clock step-down + advisory + graceful hold | Direction set |
| 4 | Channel Establishment | Multicast presence protocol, signed, three trigger types | Decided |
| 5 | VSN Admission | Two-tier: admission keys (VSN) + session auth (access node), signed multicast | Decided (envelope details TBD) |
| 6 | Upstream/Downstream | Upstream: U-C idempotent replay. Downstream: D-CE hybrid via content-type tiering (MVP: D-A reliable stream) | Decided |
| 7 | Router Forwarding | A (address) MVP → F (address+label) post-MVP; 3 multicast addresses per VSN | Decided (revisit flag set) |
| 8 | tmux Integration | AN-E (control mode + PTY fallback). Console: CN-E configurable (CN-B/D MVP, CN-C post-MVP, CN-A long-term) | Decided |
| 9 | Node Multi-homing | MH-C active/active, PS-D policy+latency, FT-D layered failover, ID-D crypto identity, IF-D upstream free/downstream TBD | Decided |
| 10 | Frame Envelope | EL-B two-layer, router HMAC, ADDR-C 8-byte VSN-scoped hash, 44-byte outer + ~16-byte channel, EXT-D versioned outer/TLV channel | Decided |
| 11 | Router-to-Router | TD-E seed+gossip, SE-A link-state + SE-E observer, latency+loss metrics, RS-D Noise auth | Design space mapped, router control plane TBD |
| 12 | Control Node (VSN Mgmt) | User-facing node type. VR-C OOB bootstrap, KS-D VSN+role, CM-A→CM-B, CP-C network service | Decided, open questions parked |

---

### Queued Sessions

1. ~~**Philosophy/Values Exploration**~~ — **COMPLETED** (see Philosophy section above)

2. ~~**Technical Research: Loss Recovery & QoS**~~ — **COMPLETED** (see Queued Session 1 Results below)

3. ~~**Technical Research: Downstream Strategy**~~ — **COMPLETED** (see Queued Session 2 Results below)

4. ~~**Exploration: Switchboard for MCP**~~ — **EXPLORED, PARKED** (see Queued Session 3 Results below). Infrastructure alignment is real. Key insight: the apparent gaps are limitations of MCP's HTTP-inherited design, not inherent requirements of agent communication. Undisclosed context prevents full assessment. Reopens when human discloses or director design matures.

5. ~~**Design: Router Control Plane**~~ — **COMPLETED** (see Queued Session 5 Results below).

   ~~**Technical Research: tmux control mode depth, console-side integration options**~~ — **COMPLETED** (see Queued Session 4 Results below). Merged into session #4 since it was originally queued separately but covers the same Parameter 8 scope.

6. **Design: Router Management Plane** — Configuration, monitoring, operations, upgrades. How routers are provisioned, how the first control key is bootstrapped (VR-C), rolling updates, observability. **Deferred to architecture work.**

7. **Design: Admission Keying Particulars** — Authentication and reauthentication flows, revocation propagation mechanics, key lifecycle (creation, distribution, rotation, expiry, revocation). Depends on router distributed database design (#5). **Deferred to architecture work.** (Note: #5 designed the distributed database — this session can proceed when needed.)

---

## Queued Session 1 Results: Loss Recovery & QoS Technical Research

_Completed 2026-03-31. Techniques: research-synthesis, cross-pollination, constraint-mapping._
_Feeds: Parameter 2 (Loss Recovery Strategy) technique selection._

### Research Synthesis — Three Targets

#### Target 1: X.25 / LAP-B — Guaranteed Delivery Over Unreliable Links

**LAP-B (Link Access Procedure, Balanced) — ITU-T X.25 Layer 2:**

Key mechanisms from this bit-oriented, full-duplex data link protocol:

1. **Sliding Window with Go-Back-N and Selective Reject (SREJ)**
   - Window size: 1-7 (mod 8) or 1-127 (extended, mod 128)
   - Each frame carries N(S) (send sequence) and N(R) (receive sequence, ACKs all < N(R))
   - **Go-Back-N (REJ):** "Missing frame 3 — resend from 3 onward." Simple but wasteful.
   - **Selective Reject (SREJ):** "Missing frame 3 — resend *only* 3." Bandwidth-efficient, requires receiver buffering.
   - **For Switchboard:** SREJ is the right choice — bandwidth is cheap, latency matters. A frame lost on path A may already have arrived on path B.

2. **Timer-Driven Retransmit (T1) + Retry Limit (N2)**
   - T1 ≈ slightly more than one RTT. After N2 failures, link declared dead.
   - **For Switchboard:** T1 maps to timeslice clock + one RTT. Multi-path advantage: retransmit via *different* path.

3. **Piggyback ACKs**
   - N(R) carried free in every data frame — ACKs ride on reverse-direction traffic.
   - **For Switchboard:** Half-channel structure is a natural piggyback vehicle.

4. **Frame Check Sequence (16-bit CRC)**
   - Corrupt frames silently discarded, retransmit handles recovery.
   - **For Switchboard:** HMAC in outer envelope already serves this purpose, but stronger.

**What transfers:** Sliding window with SREJ, piggybacked ACKs, timer-driven retransmit with clean failure escalation, window size as memory/flow-control bound.

**What doesn't:** Go-Back-N is wasteful. Single-path assumption. Mandatory in-order delivery.

#### Target 2: Interleaving vs. Duplication — Digital Radio & Satellite

**Duplication (send it twice):**

Used in DVB-S2 (satellite critical streams), dual-homed financial feeds (NYSE/NASDAQ), MPTCP redundant scheduler, Pro-MPEG (critical audio).

| Attribute | Value |
|-----------|-------|
| Bandwidth cost | 2x (100% overhead) |
| Recovery latency | Zero |
| Burst tolerance | Survives total loss of one path |
| Complexity | Trivial |

**For Switchboard upstream:** 2× on ~96 bytes = 192 bytes/tick. At 100 ticks/sec = 19.2 KB/s. Bandwidth is irrelevant. Duplication is essentially free.

**Interleaving (scatter to survive):**

Rearranges data so burst errors spread across multiple FEC recovery groups when de-interleaved.

Used in GSM/GPRS (8-frame depth, ~40ms), DAB/DAB+ (~384ms depth), DVB-T, CCSDS (deep space), CD audio (CIRC, survives 4mm scratches).

| Attribute | Value |
|-----------|-------|
| Bandwidth cost | Zero (rearrangement only, but requires FEC underneath) |
| Recovery latency | Depth × frame interval |
| Burst tolerance | Proportional to depth |
| Complexity | Moderate — depth tuning and buffering |

**Critical tradeoff:** Interleaving trades latency for burst tolerance.
- Depth 3, tick 10ms = 30ms added. Marginal.
- Depth 8, tick 10ms = 80ms. Dangerous.
- Depth 3, tick 50ms = 150ms. **Dead.** Exceeds perception threshold.

**Head-to-head for Switchboard:**

| Dimension | Duplication | Interleaving |
|-----------|-------------|--------------|
| Upstream (keystrokes) | Perfect. Free. | Overkill. U-C already handles loss. |
| Downstream (interactive) | Fine. Cheap. | Unnecessary for small frames. |
| Downstream (burst/bulk) | Expensive at 10KB+. | Natural fit with latency budget. |
| Downstream (graphics) | Very expensive at 100KB+. | Loss-intolerant — needs reliable retransmit instead. |

**Conclusion:** Not competing strategies — complementary, stratified by content type.

#### Target 3: Hybrid ARQ+FEC Decision Models — SRT, QUIC, WebRTC

**SRT (Secure Reliable Transport) — Haivision:**

Live video over public internet. Sub-second latency. Bloomberg, NHL, Twitch ingest.

- **ARQ-first with Too-Late Packet Drop (TLPKTDROP):** Sender buffer sized to latency budget. Immediate NAK on sequence gap. Retransmit if within budget. If too late → both sender and receiver drop. No wasted bandwidth on stale data.
- **Optional FEC (since v1.5):** Column+row XOR parity. FEC recovers before ARQ fires (saves one RTT). ARQ handles what FEC can't.
- **Transfers to Switchboard:** TLPKTDROP is philosophically aligned — late keystrokes are worse than dropped ones. Immediate NAK beats timer-based detection. Latency budget as first-class parameter.

**QUIC (RFC 9000 + RFC 9002) — Google/IETF:**

~30% of global web traffic. Deliberately no FEC — multiplexed streams eliminate head-of-line blocking, reducing loss impact enough for ARQ alone.

- **Smart ARQ:** Packet threshold (3 subsequent ACKs = loss), time threshold (RTT × 9/8), Probe Timeout with exponential backoff.
- **Retransmit frames, not packets:** Never resend same packet. New packet carries old frames. Avoids retransmission ambiguity.
- **Per-path RTT tracking:** All timeouts derived from measured RTT.
- **Transfers to Switchboard:** "Retransmit frames not packets" = timeslice model by construction. Per-path RTT tracking feeds path selection.

**WebRTC (RFC 8854 + RFC 5109) — Google Meet, Zoom, Discord:**

Closest regime to Switchboard: small payloads (20ms voice frames), strict latency, bidirectional.

- **Four-tier recovery hierarchy:**
  1. Redundant coding (RFC 2198) — previous frame's payload at lower quality in every packet
  2. XOR FEC (ulpfec) — parity groups, no RTT cost
  3. NACK retransmit — if RTT allows within playout deadline
  4. Concealment / skip — if too late
- **Adaptive FEC rate (GCC):** 0% at <1% loss → 10-20% at 2-5% → 30%+ at >5%.
- **Transfers to Switchboard:** The four-tier hierarchy is the answer. Redundant coding validates U-C idempotent replay. Adaptive FEC is directly applicable. Uneven Level Protection — upstream gets max protection, downstream gets tiered.

**Synthesis across all three:**

| Mechanism | SRT | QUIC | WebRTC | Switchboard Fit |
|-----------|-----|------|--------|-----------------|
| FEC | Optional add-on | Excluded | Core layer | Core for downstream burst |
| ARQ trigger | Immediate NAK | Packet+time threshold | NACK with RTT check | Immediate NAK + RTT budget check |
| Too-late policy | TLPKTDROP | Implicit PTO backoff | Playout deadline | TLPKTDROP — explicit, first-class |
| Retransmit unit | Original packet | New packet/old frames | Original packet | New frame/old content (QUIC model) |
| Adaptation | Connection-time budget | RTT-derived timeouts | Real-time FEC rate | Per-regime + measured loss |
| Multi-path | No | No | No | **Switchboard's unique advantage** |

### Cross-Pollination — Five Mechanism Transfers

#### CP-1: Reassembly Window as Decision Engine

The receiver knows: when the frame was sent (timestamp), how long each path takes (per-path RTT), when the next tick fires (local clock). This creates a decision window:

```
Frame sent ──── Expected arrival ──── Next tick deadline
                                  ↑
                    Decision window: all recovery here
```

Four-step cascade within the window:

| Step | Mechanism | Source | Cost |
|------|-----------|--------|------|
| 1 | Check duplicate from alternate path | Satellite/financial feeds | Zero |
| 2 | Check FEC parity group | WebRTC ulpfec / SMPTE 2022-1 | Zero latency |
| 3 | SREJ via alternate path | X.25 LAP-B + multi-path | One alternate-path RTT |
| 4 | TLPKTDROP + degradation signal | SRT | Zero — decision |

Step 3 decision logic:
```
remaining_window = next_tick_deadline - now
if remaining_window > min(alternate_path_RTTs):
    send SREJ via fastest alternate path
else:
    TLPKTDROP → deliver partial, signal degradation
```

Multi-path advantage: SREJ goes via the path that *didn't* lose the frame — different failure domain than the lossy path.

#### CP-2: No Retransmit, Only New Frames (QUIC × Timeslice)

QUIC's "never resend same packet, send new packet with old frames" and Switchboard's "the bus leaves on time" are the same idea discovered independently.

- **Upstream:** Already solved. U-C idempotent replay = every frame carries last N keystrokes. No retransmit mechanism needed.
- **Downstream with state sync (Mosh-style):** Lost frame superseded by next diff covering larger delta. No retransmit needed.
- **Downstream with sequential content:** Lost content included in next tick's frame alongside new content. Frame gets bigger, sequence stays clean.

**Implication:** For state-syncable downstream content, ARQ may not be needed at all. ARQ reserved for sequential, non-supersedable content (log streams, file transfers, scrollback).

#### CP-3: Regime-Aware Recovery Profiles (WebRTC GCC × Timer Regimes)

| Regime | Recovery Profile |
|--------|-----------------|
| **Idle** | No FEC, no duplication. Keep-alive only. ARQ on keep-alive loss. |
| **Negotiation** | Duplication only. Small critical messages. |
| **Active — low loss** (<1%) | Upstream: duplication. Downstream: no FEC, ARQ if needed. |
| **Active — moderate** (1-5%) | Upstream: duplication. Downstream: XOR parity K=4. ARQ backstop. |
| **Active — high** (>5%) | Duplication. Aggressive FEC K=2-3 + interleaving depth 2-3. Degradation signal. |
| **Active — critical** (>15%) | State-sync-only mode. Clock step-down. Prominent warning. |

Adaptation loop: measure loss every K ticks. Step up protection immediately. Step down slowly (hysteresis — confirm improvement before reducing protection).

Recovery profile transitions connect directly to Parameter 3 degradation signaling — same loss measurement drives both.

#### CP-4: Piggybacked ACKs + SACK Bitmap (X.25 × Half-Channels)

Each half-channel's data frames carry ACKs for the reverse half-channel. 6 bytes added to channel header:

| Field | Size | Purpose |
|-------|------|---------|
| `ack_seq` | 4 bytes | Highest contiguous reverse-channel sequence received |
| `ack_bitmap` | 2 bytes | SACK bitmap beyond ack_seq for selective reject |

Example: `ack_seq=40, bitmap=0b1101` → "Have 40, 41, 43, 44. Missing 42." Sender includes 42's content in next frame.

Standalone ACK frames fire only when reverse half-channel is idle (no data frames to piggyback on).

#### CP-5: Content-Type Stratified Interleaving

| Content Type | Interleaving | Rationale |
|-------------|-------------|-----------|
| Interactive prompt | None. Duplication + ARQ. | Every ms matters. |
| Streaming output | Depth 2-3 if moderate+ loss. | User reading, not reacting. 20-30ms invisible. |
| Bulk transfer | Depth 3-5. | Throughput > per-frame latency. |
| Graphics (Sixel/Kitty) | None. Reliable retransmit. | Loss-intolerant binary. Partial = corruption. |

Content-type detection options: heuristic (output rate), tmux control mode metadata, explicit signaling, conservative default (treat as interactive, switch after sustained bulk).

### Constraint Mapping — Mechanism Scoring

All mechanisms scored against 9 hard constraints (C1: <100ms latency, C2: independent half-channels, C3: upstream U-C, C4: dual-path, C5: 44+16 byte frame budget, C6: E router MVP, C7: variable frames, C8: SSH trust layer, C9: degradation signaled).

| Mechanism | Constraints Passed | Key Limitation | Verdict |
|-----------|-------------------|----------------|---------|
| **Duplication** | All | Single-path E router = no redundancy (degrades gracefully) | Always on |
| **XOR Parity FEC K=4** | All | Variable frames need padding+length metadata. 40ms group window at 10ms ticks. | Downstream, post-MVP |
| **SREJ via alternate path** | All | Requires dual-path. One RTT cost. | Downstream backstop, post-MVP |
| **TLPKTDROP** | All | Must signal degradation (C9) | Always on |
| **Piggybacked ACK+SACK** | All | 6 bytes added to channel header (~16→~22 bytes) | Always on |
| **Adaptive FEC rate** | All | Meta-controller, not a mechanism itself | Always on |
| **Interleaving depth 2-3** | Conditional C1 | Depth × tick = latency cost. Limited value on E router MVP. | Downstream bulk only, post-MVP |

**Variable-frame FEC solution:**
```
Parity payload = XOR(pad(frame1, MAX), ..., pad(frameK, MAX))
Parity metadata = [len1, len2, ..., lenK]  (K × 2 bytes)
Recovery: XOR parity with received frames → missing frame (padded). Trim to stored length.
```

### Parameter 2: Updated Specification — Concrete Techniques Selected

**MVP (E router, single or dual path):**

```
Upstream:
  - Duplication across available paths (always)
  - U-C idempotent replay: sliding window of last N keystrokes (always)
  - Piggybacked downstream ACK + SACK bitmap in every frame (always)
  - TLPKTDROP: keystroke older than perception budget → discard (always)

Downstream:
  - Duplication across available paths when frame size < threshold (always)
  - Piggybacked upstream ACK + SACK bitmap in every frame (always)
  - TLPKTDROP: frame past tick deadline → skip, deliver next state,
    signal degradation via Parameter 3 (always)
  - ARQ: on SACK gap detection, include missed content in next tick's
    frame (QUIC model — new frame with old content, not retransmit)
  - Two recovery regimes: normal (duplication + ARQ) and degraded
    (ARQ + clock step-down + degradation signal)
```

**Post-MVP (multi-path, multi-hop):**

```
Add to downstream:
  - XOR parity FEC: K=4 groups, variable-frame with length metadata
    Activated at moderate loss (1-5%), regime-aware
  - SREJ via alternate path: triggered by SACK bitmap, only when
    remaining decision window > alternate path RTT
  - Interleaving: depth 2-3, bulk content only, activated at high
    loss (>5%)
  - Four-regime adaptive profile (low/moderate/high/critical loss)
    with fast step-up, slow step-down (hysteresis)

Recovery cascade (within reassembly decision window):
  1. Check duplicate arrival from alternate path → use if present
  2. Check FEC parity group → recover if complete
  3. SREJ if window > min(alternate RTTs) → retransmit via fastest alt
  4. TLPKTDROP → skip, deliver available state, signal degradation
```

**Channel header additions (to existing ~16 byte budget):**

| Field | Size | Purpose |
|-------|------|---------|
| `ack_seq` | 4 bytes | Highest contiguous reverse-channel sequence received |
| `ack_bitmap` | 2 bytes | SACK bitmap for selective reject |

**FEC channel extension (parity frames only, TLV):**

| Field | Size | Purpose |
|-------|------|---------|
| `fec_group_id` | 2 bytes | Parity group identifier |
| `fec_group_size` | 1 byte | K value |
| `fec_position` | 1 byte | Position in group (0=parity) |
| `frame_lengths` | K × 2 bytes | Original frame lengths for variable-size XOR recovery |

### Key Research Sources Referenced

- **X.25/LAP-B:** ITU-T X.25, HDLC/LAP-B sliding window, selective reject (SREJ)
- **Interleaving:** GSM (3GPP TS 05.03), DAB (ETSI EN 300 401), DVB-T (ETSI EN 300 744), CCSDS (131.0-B), CD CIRC (IEC 60908)
- **SRT:** Haivision SRT Protocol Technical Overview, TLPKTDROP, SMPTE 2022-1/7 FEC
- **QUIC:** RFC 9000 (transport), RFC 9002 (loss detection and congestion control)
- **WebRTC:** RFC 8854 (media transport), RFC 5109 (ulpfec), RFC 2198 (redundant audio), RFC 4585 (RTCP feedback/NACK), Google Congestion Control (GCC)
- **Duplication:** DVB-S2 PLP, MPTCP redundant scheduler, Pro-MPEG/SMPTE 2022-1
- **Multi-path:** MPTCP (RFC 8684), BLEST scheduler, STTF scheduler

### Connections to Other Parameters

- **Parameter 1 (Timer Regimes):** Recovery profiles map to regimes. Idle = no FEC. Active = regime-aware cascade.
- **Parameter 3 (Degradation Signaling):** Profile transitions trigger degradation signals. Same loss measurement drives both.
- **Parameter 6 (Upstream/Downstream):** Upstream fully resolved (U-C + duplication). Downstream technique selection now concrete but content-type stratification depends on downstream strategy decision (still open — queued session #2).
- **Parameter 9 (Multi-homing):** SREJ via alternate path leverages multi-path. Pre-auth to standby routers enables instant switchover.
- **Parameter 10 (Frame Envelope):** Channel header grows by 6 bytes (ACK+SACK). FEC extension uses TLV in parity frames.
- **Parameter 11 (Router-to-Router):** Hysteresis/dampening pattern shared between FEC adaptation and link-state convergence.

### Open Questions Remaining

1. ~~**Downstream content-type detection**~~ — **RESOLVED** in queued session #2. Three-layer detector: escape sequence scan + pane state cache + byte-rate heuristic.
2. **FEC group size tuning** — K=4 is the starting point. May need per-content-type K values (smaller K for interactive, larger K for bulk).
3. **SACK bitmap width** — 2 bytes = 16 frames of selective ACK history. Sufficient? Or does high-loss regime need wider bitmap?
4. **Interleaving activation threshold** — "high loss >5%" is a starting point. Needs simulation or production data.
5. **Upstream sliding window depth N** — how many keystrokes in the idempotent replay window? 10? 20? Depends on typical keystroke rate vs. tick interval.

---

## Queued Session 2 Results: Downstream Strategy Technical Research

_Completed 2026-03-31. Technique: research-synthesis._
_Feeds: Parameter 6 (Upstream/Downstream Asymmetry) — downstream decision._
_Decision: **D-CE** — hybrid (state sync + reliable stream) via content-type tiering._

### Research Synthesis — Five Targets

#### Target 1: tmux Control Mode Internals

**Protocol:** Control mode (`tmux -CC`) emits structured text instead of rendering to a PTY.

**Key message types:**
- `%output %<pane-id> <escaped-data>` — raw terminal bytes, tagged by pane
- `%begin` / `%end` — structured command output blocks
- `%session-changed`, `%window-add`, `%window-close`, `%pane-mode-changed`, `%layout-change` — lifecycle events
- `%extended-output` (tmux 3.4+) — richer metadata per output chunk

**What control mode gives:**
- Pane-multiplexed output without parsing tmux's own rendering
- Event-driven state tracking (window/pane create/destroy, layout changes)
- Programmatic command execution with structured responses
- Clean separation of data (`%output`) and control (`%begin`/`%end`)

**What control mode does NOT give:**
- Terminal state (no `%screen-state` snapshot — only the byte stream)
- Scrollback contents (requires `capture-pane -p -S -N` command, not streamed)
- Content type (tmux doesn't know interactive from bulk from graphics)
- Cursor position / screen dimensions (queryable, not streamed)

**Finding:** Control mode gives tagged byte streams per pane, not state snapshots. State sync (Mosh-style) requires a terminal state tracker (headless terminal emulator) on the access node consuming `%output` and maintaining a screen buffer. This is proven feasible engineering — a solved class of problem — but requires purpose-built, minimal implementation.

#### Target 2: Scrollback Buffer Semantics

**Two different things called "scrollback":**
1. **tmux's scrollback** — history buffer per pane, on the access node. Authoritative. Accessible via `capture-pane`.
2. **Terminal emulator's scrollback** — local terminal's buffer of received bytes. Only as complete as what arrived.

**Scrollback fetch is a different protocol mode than live streaming:**

| Attribute | Live streaming | Scrollback fetch |
|-----------|---------------|-----------------|
| Latency sensitivity | Critical (<100ms) | Tolerant (200-500ms, user-initiated) |
| Loss tolerance | Content-type dependent | Zero — must be complete and ordered |
| Direction | Push | Pull (console requests) |
| Recovery | Cascade (dup → FEC → SREJ → drop) | Reliable retransmit only, no TLPKTDROP |

**Finding:** Scrollback is a pull-based reliable fetch, separate from live streaming. State sync handles live display. Reliable fetch handles scrollback. The two coexist naturally — they map to tmux's existing live/copy-mode split. Console enters scroll mode on scroll-up, returns to live mode on scroll-to-bottom.

#### Target 3: Sixel/Kitty Graphics Protocols

**Three competing terminal graphics protocols:**

| Protocol | Origin | Encoding | Chunked? | tmux support |
|----------|--------|----------|----------|-------------|
| Sixel | DEC 1984 | ASCII (~3-4x overhead) | No (continuous blob) | tmux 3.4+ passthrough |
| Kitty | 2017 | Base64 (~1.33x) | Yes (`m=1`/`m=0`) | tmux 3.4+ passthrough |
| iTerm2 inline | 2010s | Base64 (~1.33x) | No (single blob) | tmux 3.4+ passthrough |

**Common property: all three have zero loss tolerance.** A single missing byte corrupts the image or crashes the parser. Partial graphics data is worse than no graphics data.

**Detection is feasible via escape sequence delimiters:**
- Sixel: `ESC P ... ESC \`
- Kitty: `ESC _G ... ESC \`
- iTerm2: `ESC ]1337;File= ... BEL`

**Kitty's chunked model maps naturally to timeslice frames** — one chunk per tick, reassembled atomically on the console side. Sixel can be chunked artificially at row boundaries (`-` characters).

**Finding:** Graphics need reliable delivery (no TLPKTDROP), atomic reassembly, and escape sequence-based detection. MVP can pass through without special handling (E router, low loss). Post-MVP adds detection + reliable-mode flag in channel header.

#### Target 4: Content-Type Detection Feasibility

**Three-layer detection model:**

**Layer 1 — Escape Sequence Detection (per-frame, highest priority):**
- Graphics: Sixel/Kitty/iTerm2 delimiter detection → reliable mode
- TUI: Alternate screen buffer `?1049h`/`?1049l` → state sync mode
- Cost: Linear scan, same work terminal state tracker already does

**Layer 2 — Pane State Cache (event-driven):**
- `alternate_on` flag (cached on `%pane-mode-changed`)
- `pane_current_command` (cached on first output, updated on change)
- Cost: One tmux query per state change, cached

**Layer 3 — Byte-Rate Heuristic (per-tick, fallback):**
- < ~200 bytes/tick → interactive
- < ~4KB/tick → streaming
- \> ~4KB/tick → bulk flood
- Cost: Counter, negligible

**Combined output: content_type flag (3 bits) in channel header flags byte:**

| Value | Type | Delivery Path | TLPKTDROP |
|-------|------|---------------|-----------|
| 0 | Interactive | State sync | Acceptable |
| 1 | Streaming | Reliable stream | Reluctant |
| 2 | Bulk | Reliable stream | Acceptable |
| 3 | TUI | State sync | Acceptable |
| 4 | Graphics | Reliable stream (atomic) | Never |

**Finding:** Content-type detection is feasible, reliable for graphics and TUI (escape sequences), adequate for interactive/streaming/bulk (byte rate). Cost is negligible if access node runs a terminal state tracker.

#### Target 5: Claude Code Use Case

**Claude Code is the hardest downstream case** — it interleaves content types within a single pane within seconds:
- Interactive (prompt, confirmation) → streaming (reasoning, response) → bulk (code diffs, tool output) → TUI (status line) → back to interactive

**Key findings:**
1. **Streaming text where every line matters.** Claude's reasoning is the work product. TLPKTDROP is dangerous — a dropped line loses rationale. Streaming class, reluctant TLPKTDROP.
2. **Tool execution output is variable.** Transitions from streaming to bulk naturally. Byte-rate heuristic handles this.
3. **Status line / input box.** Not full alternate-screen TUI. Small, infrequent. Falls to interactive by byte rate. No special treatment.
4. **Scrollback is first-class.** Users will scroll back through reasoning, diffs, tool output. Strongest argument for reliable scrollback fetch (D-C hybrid).
5. **Mixed stream demands per-frame classification.** No single strategy works for all of Claude Code's output types.

**Finding:** Claude Code is the strongest argument for D-CE. Its mixed-type output validates per-frame content-type tiering over any uniform downstream strategy.

### Parameter 6: Updated Specification — Downstream Decision

**Decision: D-CE — Hybrid via Content-Type Tiering**

D-C (state sync + reliable for scrollback) and D-E (tiered by content type) converge: D-C describes the two delivery paths, D-E describes the per-frame routing decision between them.

**Access node architecture:**

```
tmux control mode (%output per pane)
        │
        ↓
┌─────────────────────────┐
│ Content-Type Detector    │
│ L1: escape sequence scan │
│ L2: pane state cache     │
│ L3: byte-rate heuristic  │
└──────┬──────┬────────────┘
       │      │
       ↓      ↓
┌────────────┐ ┌──────────────────┐
│ State Sync │ │ Reliable Stream  │
│ Path       │ │ Path             │
│            │ │                  │
│ Screen buf │ │ Ordered bytes    │
│ tracker,   │ │ with seq nums,   │
│ diff vs    │ │ ARQ recovery,    │
│ last-ack'd │ │ feeds scrollback │
│ state,     │ │                  │
│ TLPKTDROP  │ │ Graphics submode:│
│ ok         │ │ atomic chunks,   │
│            │ │ no TLPKTDROP     │
└─────┬──────┘ └────────┬─────────┘
      └────────┬─────────┘
               ↓
      Frame assembly (tick)
      content_type in header
```

**Console side:**
- State sync frames → update local screen buffer, render
- Reliable stream frames → append local scrollback, render if live mode
- Graphics frames → buffer until complete, render atomically
- Scroll-up → scrollback fetch request to access node (pull, reliable)
- Scroll-to-bottom → return to live mode

**MVP (simplified):**

```
Downstream MVP: D-A reliable ordered stream.
  All content treated uniformly.
  Piggybacked ACK + SACK in upstream frames.
  ARQ on gap (QUIC model: new frame with old content).
  TLPKTDROP as last resort, degradation signaled.
  No terminal state tracker. No content-type detection.
  No state sync. No scrollback fetch protocol.
```

**Post-MVP (full D-CE):**

```
Downstream post-MVP: D-CE hybrid via content-type tiering.
  Terminal state tracker on access node (control mode consumer).
  Three-layer content-type detector (escape + cache + byte rate).
  Two delivery paths: state sync + reliable stream.
  Five content types: interactive (0), streaming (1), bulk (2), TUI (3), graphics (4).
  Per-type recovery profiles (from loss recovery session):
    Interactive/TUI: state sync, duplication, TLPKTDROP acceptable
    Streaming: reliable stream, duplication + ARQ, reluctant TLPKTDROP
    Bulk: reliable stream, FEC + interleaving if lossy, TLPKTDROP acceptable
    Graphics: reliable stream atomic, no TLPKTDROP, chunked delivery
  Scrollback fetch: pull-based, reliable, separate protocol mode.
```

### Connections to Other Parameters and Sessions

- **Parameter 2 (Loss Recovery):** Content-type flag drives which recovery profile applies per frame. The four-tier cascade from session #1 applies differently to each content type.
- **Parameter 3 (Degradation Signaling):** TLPKTDROP policy varies by content type. Degradation signal includes which content types are being affected.
- **Parameter 8 (tmux Integration):** AN-E control mode is prerequisite for content-type detection and terminal state tracking. PTY fallback limits detection to byte-rate heuristic only.
- **Parameter 10 (Frame Envelope):** Content-type field uses 3 bits in existing flags byte. No additional header space needed.
- **Session #1 open question #1 (content-type detection):** Resolved — three-layer detector.

### Open Questions from Downstream Research

1. **Terminal state tracker implementation** — purpose-built minimal screen buffer, or adapt existing library (e.g., Go terminal emulator library)? Engineering scope assessment needed.
2. **State sync diff format** — character-level diffs, line-level diffs, or screen-region diffs? Tradeoff between diff size and compute cost.
3. **Scrollback fetch protocol** — request format (line range? byte offset?), flow control, interaction with live mode. Separate channel or flag on existing channel?
4. **Content-type transition latency** — one tick lag when content type changes (byte-rate heuristic lags by one tick). Is this acceptable for all transitions?
5. **Sixel chunking boundaries** — row boundaries (`-`) are natural chunk points. Are they sufficient, or do large Sixel images need sub-row chunking?
6. **State sync acknowledgment** — how does the console ACK screen state? Per-diff sequence number? Last-seen state hash? Needed for the access node to know what to diff against.
7. **iTerm2 is not a reference implementation** — terminal state tracking is proven feasible but Switchboard must build its own minimal, purpose-built tracker. No dependency on bloated consumer software.

---

## Queued Session 3 Results: Switchboard for MCP — Exploration

_Completed 2026-03-31. Technique: exploration._
_Status: **Explored, parked.** Infrastructure alignment real. Incomplete — undisclosed context prevents full assessment._

### What MCP Is (Current State)

Model Context Protocol (MCP) — Anthropic, 2024-2025. Client-server protocol for connecting AI agents to tools and data sources via JSON-RPC.

**Two defined transports:**
1. **stdio** — MCP server as subprocess, stdin/stdout, local only.
2. **Streamable HTTP (SSE)** — HTTP endpoint, POST for requests, SSE for streaming responses, session management via headers. Remote-capable.

**Traffic characteristics:**
- Tool discovery: 1-10 KB
- Tool invocation: 100 bytes - 5 KB
- Tool results: 100 bytes - 1 MB (variable)
- Progress notifications: 50-200 bytes

**Multi-agent use case:** Agents expose MCP server interfaces to each other. Currently requires per-agent HTTP endpoint, TLS, per-connection auth, SSE for streaming (unidirectional).

### Infrastructure Alignment with Switchboard

| MCP Remote Need | Current (HTTPS) | Switchboard Equivalent |
|----------------|-----------------|----------------------|
| Agent connectivity | HTTP endpoint per agent | VSN membership |
| Authentication | OAuth/Bearer per connection | VSN admission + session auth — prove once |
| Encryption | TLS per connection | SSH E2E |
| Bidirectional comms | POST + SSE (asymmetric) | Full duplex half-channels |
| Service discovery | DNS / hardcoded URLs | Multicast presence protocol |
| Failover | Application-level retry | Multi-path + SREJ + path switching |
| Multi-tenancy | Separate endpoints / auth | VSN cryptographic isolation |

**The structural argument:** Switchboard already solves network-level concerns that HTTPS re-solves per-connection at the application level. MCP over Switchboard would look like stdio-style transport (JSON-RPC over a bidirectional byte stream) with network-level routing, admission, and failover underneath. The agent doesn't need an HTTP server — it reads/writes a channel.

### Apparent Gaps — And Their True Nature

Several properties of MCP appear mismatched with Switchboard's architecture:

1. **Request-response vs. stream-oriented** — MCP is JSON-RPC request/response. Switchboard's timeslice framing is optimized for continuous streams.
2. **No latency sensitivity** — MCP tool calls take 100ms-10s. Switchboard's sub-100ms optimization is irrelevant.
3. **No terminal substrate** — MCP endpoints aren't terminal sessions. The access node / console model assumes terminal producers and consumers.
4. **Variable payload sizes** — 1 MB tool results don't benefit from timeslice framing or idempotent replay.

**Critical reframing (from maker):** These gaps are properties of MCP's *current transport design*, not inherent requirements of agent communication. MCP was designed by people from a web-services world. The protocol's shape — HTTP transport, request-response orientation, no quality-of-service, no built-in failover, no network-level identity — reflects its authors' assumptions inherited from web development, not the problem's actual requirements.

**The real question is not "does Switchboard fit MCP?" but "does Switchboard fit what agent communication *should be* — which MCP approximates badly because it inherited web-services assumptions?"**

This reframing changes the assessment fundamentally:
- Request-response may be the wrong pattern for agent communication (agents are persistent processes that *stream* work to each other, not stateless HTTP handlers)
- Latency sensitivity may matter more than current MCP acknowledges (real-time agent coordination, not just tool invocation)
- Terminal substrate may be more natural than HTTP for agents that already live in terminals
- The apparent mismatch between Switchboard and MCP may indicate that MCP is misshapen, not that Switchboard is wrong for the job

### What's Undisclosed

The human has indicated there are aspects of this use case not yet on the table. The exploration is intentionally incomplete. The fit assessment — particularly around what agent communication *should be* vs. what MCP currently provides — cannot be completed with missing inputs.

### Parking Conditions

**Status: Explored, parked. Not graveyard, not frontier.**

**Reopens when:**
- The human discloses the missing context
- Marvel's agent communication protocol (charter F3) matures
- Director's design (charter F1) clarifies agent-to-agent communication needs
- Evidence emerges about what agent communication *should be* beyond MCP's HTTP-inherited model

**Would graveyard if:**
- MCP transport evolves to solve remote well (HTTP/3 + QUIC would close most gaps)
- The workload mismatch proves fundamental to agent communication itself, not just to MCP's HTTP bias
- Second use case dilutes switchboard's terminal-session focus in a way that harms the primary use case

**Would promote to frontier if:**
- Undisclosed context reveals that the alignment is deeper than infrastructure overlap
- Director design shows that agent communication needs connection-oriented, quality-managed, terminal-native channels — exactly what Switchboard provides
- Evidence that MCP's HTTP assumptions are actively holding back agent capabilities, and Switchboard's model enables capabilities that HTTP-based MCP cannot

---

## Queued Session 4 Results: tmux Control Mode Depth & Console-Side Integration

_Completed 2026-03-31. Technique: research-synthesis._
_Feeds: Parameter 8 (tmux Integration) — console-side decision._
_Decision: **CN-E configurable** — CN-B/CN-D for MVP, CN-C for tmux users post-MVP, CN-A as long-term foundation._

### Control Mode Protocol — Full Depth

**Key discoveries beyond session #2 baseline:**

**`%subscription-changed` (tmux 3.4+):** Subscribe to any tmux format variable, get pushed updates. Eliminates polling for content-type detection:
```
refresh-client -B "alt:#{alternate_on}"
→ %subscription-changed alt 1
```

Relevant subscribable variables: `alternate_on` (TUI detection), `cursor_x`/`cursor_y`, `pane_current_command`, `pane_width`/`pane_height`, `pane_dead`, `scroll_position`, `pane_in_mode`.

**`%pause` / `%continue` (tmux 3.4+):** Per-pane output flow control. Access node can pause output when downstream channel is congested — tmux buffers at source. Alternative to TLPKTDROP for streaming content where completeness matters more than freshness.

**`%extended-output` (tmux 3.4+):** Richer metadata per output chunk. Newer, less documented, but indicates tmux is moving toward more structured output metadata.

**Connection to session #2 findings:** `%subscription-changed` makes the content-type detector's layer 2 (pane state cache) fully event-driven with zero polling. `%pause`/`%continue` adds a sixth option to the recovery cascade — buffer at source instead of TLPKTDROP, applicable to streaming content types.

### Console-Side Integration — Four Options Assessed

**CN-A: Control Mode Consumer**
- Console parses control mode, reconstructs terminal state locally
- Local scrollback (owns the buffer), local resize (re-renders), client-side intelligence
- Engineering scope: building a minimal terminal emulator
- Compatible with state sync downstream (D-CE)
- Go terminal emulator libraries are sparse — may need to build or port

**CN-B: PTY Injection**
- Pipe received bytes into local PTY, terminal emulator renders
- Simplest useful implementation — graphics pass through naturally
- No local state tracking, no re-render on resize
- Works with reliable ordered stream (D-A), NOT with state sync (D-CE)

**CN-C: Local tmux Mirroring**
- Local tmux mirrors remote topology — windows, panes, layout
- Native scrollback, copy-paste, mouse, splits — strongest "illusion of local"
- Requires tmux on console side (acceptable for target audience)
- Significant complexity: bidirectional sync, keystroke routing, pane mapping
- Remote tmux is authoritative; local pane creation intercepted or prevented

**CN-D: Direct Terminal Write**
- `write(stdout, received_bytes)` — simplest possible
- Terminal emulator handles everything
- Adequate for MVP when loss is rare (E router LAN)

### Console-Side Decision

**CN-E: Configurable, with opinionated default progression:**

| Phase | Console Mode | Downstream Mode | Target |
|-------|-------------|----------------|--------|
| **MVP** | CN-D or CN-B | D-A reliable ordered stream | Works now, simple |
| **Post-MVP tier 1** | CN-C (local tmux) | D-A or partial D-CE | Premium for tmux users |
| **Post-MVP tier 2** | CN-A (control mode consumer) | Full D-CE state sync | Client-side intelligence |

**Key architectural insight:** The access node and console can be at different capability levels. A smart access node running full D-CE state sync can serve a dumb CN-B/CN-D console by sending diff-applied output as a byte stream. Console capability doesn't constrain access node capability.

### Updated Parameter 8 Specification

```
Access node: AN-E (control mode + PTY fallback) — unchanged.
  tmux 3.4+ features used:
    %subscription-changed for event-driven pane state monitoring
    %pause/%continue for source-side flow control
    %extended-output for richer metadata

Console: CN-E configurable.
  MVP: CN-D (direct write) or CN-B (PTY injection).
    Reliable ordered stream downstream. Switchboard invisible.
  Post-MVP tier 1: CN-C (local tmux mirroring).
    Optional. Falls back to CN-B if local tmux unavailable.
    Strongest "illusion of local" for target audience.
  Post-MVP tier 2: CN-A (control mode consumer).
    Terminal state tracker on console side.
    Enables full D-CE state sync at both ends.
    Local scrollback, resize, search, highlighting.
    Long-term foundation for client-side intelligence.
```

### Connections to Other Parameters and Sessions

- **Parameter 6 / Session #2 (Downstream Strategy):** CN-A enables full D-CE on console side. CN-B/CN-D work with D-A only. Access node D-CE works regardless of console capability.
- **Parameter 3 (Degradation Signaling):** `%pause`/`%continue` adds source-side buffering as alternative to TLPKTDROP for streaming content.
- **Session #2 Target 4 (Content-Type Detection):** `%subscription-changed` makes layer 2 fully event-driven.
- **Session #1 (Loss Recovery):** Source-side pause via `%pause` could be a fifth step in the recovery cascade, between SREJ and TLPKTDROP — for content types where completeness matters.

### Open Questions

1. **Go terminal emulator library** — build vs. port for CN-A? Scope assessment needed.
2. **CN-C local tmux mirroring** — bidirectional sync protocol. How does the console prevent local pane creation from conflicting with remote authority?
3. **CN-C keystroke routing** — mapping local pane IDs to remote pane IDs. Needs a pane registry.
4. **`%pause` integration with recovery cascade** — should source-side pause be step 4.5 (between SREJ and TLPKTDROP) or a content-type-specific alternative to TLPKTDROP?
5. **tmux version requirements** — `%subscription-changed`, `%pause`, and `%extended-output` require tmux 3.4+. What's the fallback for older tmux? Polling + no pause + basic `%output` only.

---

## Queued Session 5 Results: Router Control Plane Design

_Completed 2026-03-31. Technique: design._
_Feeds: Parameters 5, 7, 9, 11, 12 — the foundational undescribed piece._
_Key insight: **One mechanism (reliable flooding of signed messages) serves all three control plane functions.**_

### Three Functions of the Control Plane

| Function | What It Does | Consistency Requirement |
|----------|-------------|------------------------|
| **Topology & routing** | Routers discover each other, exchange link-state, compute forwarding tables | Reliable flood, fast convergence (~1-2s) |
| **Admission state distribution** | Distributed database of which keys are admitted to which VSNs | Reliable flood, revocation wins over admission |
| **VSN membership propagation** | Which nodes are currently connected to which routers, for multicast forwarding | Reliable flood, eventual consistency ok |

### Function 1: Topology and Routing

**Bootstrap (TD-E):** Router starts with seed list (known routers). Control node's router is the natural seed. Gossip discovers additional routers within a few rounds.

**Link-State Advertisement (SE-A):**
```
LSA:
  router_id: <hash of router public key, 8 bytes>
  sequence: <u64, monotonically increasing>
  timestamp: <u64, UTC nanos>
  links:
    - neighbor: <router_id>
      latency_ms: <measured one-way>
      loss_rate: <measured>
      jitter_ms: <derived from latency variance>
```

**Flooding:** Router receives LSA with higher sequence than stored → updates link-state database → forwards to all neighbors except source. Classic reliable flooding (OSPF model).

**Forwarding table computation:** Each router runs Dijkstra independently over its link-state database with latency as metric. Result: next-hop for each destination. At 2-20 routers, Dijkstra is microseconds.

**Convergence:**
- Triggered updates on link quality change beyond threshold → immediate LSA flood
- Dampening for flapping links → suppress rapid oscillation
- Periodic full refresh every T seconds → catch any missed updates
- Target: alternative paths found within ~1-2 seconds of router failure

### Function 2: Admission State Distribution

**The distributed database.** Stores which public keys are admitted to which VSNs, with role (control/console/access_node), expiry, and revocation state.

**Consistency model: reliable flooding with confirmation.**

Key registration flow:
1. Control node sends signed key registration to its connected router
2. Router validates signature (control node must have `role: control` for target VSN)
3. Router adds key to local admission database
4. Router floods registration to all other routers (same mechanism as LSAs)
5. Each receiving router validates signature + adds key
6. Confirmation propagates back → control node gets acknowledgment

**Revocation:** Same mechanism. Signed revocation floods to all routers. **Revocation wins over admission** — if registration and revocation are in flight simultaneously, revocation takes precedence.

**Forwarding-time admission check:**
```
Frame arrives with: VSN_ID (8 bytes) + source_addr (8 bytes) + HMAC (16 bytes)
Lookup: admission_db[vsn_id][source_addr] → admitted? → forward or drop
```
Hash table lookup, O(1) per frame. Database size is small (admitted keys across VSNs). Memory and CPU: negligible.

### Function 3: VSN Membership Propagation

**Tracks which nodes are currently connected to which routers.** Needed for multicast forwarding — "which routers have subscribers for this VSN multicast address?"

**Join/leave protocol:**
- Node connects to router, joins VSN → router floods "join" to all routers
- Node disconnects → router floods "leave" to all routers
- Receiving routers update their membership tables

**Multicast forwarding:** When a multicast frame arrives for VSN X group Y, forward to all routers listed as having subscribers for that group. PIM-like but simpler — full mesh of state at 2-20 routers, no rendezvous points, no tree switching.

**Eventual consistency is fine:** A join taking 1-2 seconds to propagate means the new node misses one or two multicast heartbeats, then catches up.

### Unified Control Plane Protocol

All three functions share:
- **Same flooding mechanism** — reliable flooding to all authenticated peers
- **Same authentication** — Noise handshake between routers (RS-D), signed messages
- **Same wire format envelope:**

```
Control Message Envelope:
  type: u8           # lsa, key_register, key_revoke, membership_join, membership_leave
  router_id: [u8;8]  # originating router
  sequence: u64      # per-router monotonic
  timestamp: u64     # UTC nanos
  signature: [u8;64] # Ed25519 or Noise static key
  payload: [u8;...]  # type-specific content
```

**Control node messages are double-signed:** by the control node (proving authority) and by the forwarding router (proving legitimate receipt).

### MVP Control Plane

E router MVP (single router, single hop):
```
Topology: N/A (one router)
Admission: Local database only, no flooding needed
Membership: Local table only, multicast is local delivery
```

**Trivially simple.** All three functions collapse to local operations. Distributed protocol activates when the second router appears.

### Connections to Other Parameters

- **Parameter 5 (VSN Admission):** Admission state distribution is the mechanism that makes VSN admission work across routers.
- **Parameter 7 (Forwarding):** Membership propagation enables multicast forwarding across routers.
- **Parameter 9 (Multi-homing):** When a node switches routers, the new router already has admission state (via flooding). Membership join/leave handles the path update.
- **Parameter 11 (Router-to-Router):** This session concretized the working directions from P11 into a protocol design.
- **Parameter 12 (Control Node):** "Control node submits changes as a network service" — the service is the control plane's admission flooding mechanism.

### Open Questions (Deferred to Architecture)

1. **LSA format optimization** — is the full LSA format above sufficient, or do we need compact LSAs for frequent updates?
2. **Admission database persistence** — should routers persist admission state to disk, or rebuild from flooding on restart? Persistence means faster restart. Flooding means simpler code and guaranteed freshness.
3. **Split-brain scenarios** — what happens if the router network partitions? Each partition has stale admission state for nodes on the other side. Heals on reconnect via sequence-number-based reconciliation.
4. **Control message prioritization** — should revocations be prioritized over registrations in the flooding queue? Revocation-wins semantics suggest yes.
5. **Router management plane** — how routers are configured, monitored, updated. Deferred to queued session #6 (now architecture work).
6. **Admission keying particulars** — authentication flows, key lifecycle. Deferred to queued session #7 (now architecture work). The distributed database from this session unblocks that design.
