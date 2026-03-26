---
stepsCompleted: [1, 2-research, 3-first-principles, 4-morphological-analysis-complete]
inputDocuments: []
session_topic: 'Switchboard - low-latency multi-path E2E encrypted tmux session router architecture with virtual switched networks'
session_goals: 'Naming, architecture design, edge protocol design, failure mode analysis, use cases'
selected_approach: 'ai-recommended'
techniques_used: ['research-synthesis', 'first-principles-thinking', 'morphological-analysis', 'values-exploration']
ideas_generated: []
context_file: ''
next_phase: 'post-morphological-synthesis-or-queued-sessions'
morphological_parameters_completed: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
morphological_parameters_next: null
queued_sessions:
  - 'technical-research: loss recovery and QoS techniques for interactive overlays (X.25, interleaving, hybrid ARQ+FEC)'
  - 'technical-research: downstream strategy (scrollback, graphics/Sixel/Kitty, TUI, tmux control mode, Claude Code use case)'
  - 'exploration: Switchboard for MCP - agent-to-agent and agent-to-tool overlay network'
  - 'technical-research: tmux control mode depth, console-side integration options'
  - 'design: router control plane (topology, link-state, forwarding computation, distributed database)'
  - 'design: router management plane (config, monitoring, operations, upgrades)'
  - 'design: admission keying particulars (authentication, reauthentication, revocation propagation)'
session_bootstrap: |
  ## Session Bootstrap — Resume Instructions

  This brainstorming session is run under the analyst agent (Mary) using the
  brainstorming workflow. The user goes by "maker".

  ### Where We Are
  - **Phases complete:** Research Synthesis (10 topics), First Principles (8 truths),
    Values/Philosophy (10 values + death conditions), Morphological Analysis (12 parameters — ALL COMPLETE)
  - **Next phase:** Post-morphological synthesis, or pivot to one of the queued sessions
  - **All 12 parameters have working directions chosen.**

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
| 6 | Upstream/Downstream | Upstream: U-C idempotent replay. Downstream: ALL OPEN | Upstream decided, downstream research flagged |
| 7 | Router Forwarding | A (address) MVP → F (address+label) post-MVP; 3 multicast addresses per VSN | Decided (revisit flag set) |
| 8 | tmux Integration | AN-E (control mode + PTY fallback). Console: all options open, configurable | Access node decided, console research flagged |
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
| 6 | Upstream/Downstream | Upstream: U-C idempotent replay. Downstream: ALL OPEN | Upstream decided, downstream research flagged |
| 7 | Router Forwarding | A (address) MVP → F (address+label) post-MVP; 3 multicast addresses per VSN | Decided (revisit flag set) |
| 8 | tmux Integration | AN-E (control mode + PTY fallback). Console: all options open, configurable | Access node decided, console research flagged |
| 9 | Node Multi-homing | MH-C active/active, PS-D policy+latency, FT-D layered failover, ID-D crypto identity, IF-D upstream free/downstream TBD | Decided |
| 10 | Frame Envelope | EL-B two-layer, router HMAC, ADDR-C 8-byte VSN-scoped hash, 44-byte outer + ~16-byte channel, EXT-D versioned outer/TLV channel | Decided |
| 11 | Router-to-Router | TD-E seed+gossip, SE-A link-state + SE-E observer, latency+loss metrics, RS-D Noise auth | Design space mapped, router control plane TBD |
| 12 | Control Node (VSN Mgmt) | User-facing node type. VR-C OOB bootstrap, KS-D VSN+role, CM-A→CM-B, CP-C network service | Decided, open questions parked |

---

### Queued Sessions

1. ~~**Philosophy/Values Exploration**~~ — **COMPLETED** (see Philosophy section above)

2. **Technical Research: Loss Recovery & QoS** — X.25 guaranteed delivery (LAP-B, sliding window, selective reject), interleaving vs. duplication (digital radio/satellite), hybrid ARQ+FEC decision models (SRT, QUIC, WebRTC). Feeds Parameter 2 technique selection.

3. **Technical Research: Downstream Strategy** — tmux control mode internals, scrollback buffer semantics, Sixel/Kitty graphics protocol, content-type detection feasibility, Claude Code scrollback use case. Feeds Parameter 6 downstream decision.

4. **Exploration: Switchboard for MCP** — Switchboard as an MPLS-like overlay network for MCP, primarily for agent-teaming and agent-tooling. A replacement for agent-to-agent and agent-to-tool virtual networks over reliable, untrusted infrastructure. Originated from exploring remote Claude Code director-to-multiclaude supervisor communication channels. Discovered Claude Relay and the prevalence of MCP as the serious multi-agent communication method. Core question: why WebSocket+MCP+SSH? Could connection-oriented, terminal-native, text-oriented communication be more effective than streaming WebSockets and abused HTTP? This extends Switchboard's values (sovereign sessions, untrusted network, simple reliable communication) into the agent-to-agent domain. **Not a pivot — a potential second use case that shares the same infrastructure and values.**

5. **Design: Router Control Plane** — Topology discovery, link-state exchange, forwarding computation, network-distributed database for admission state/VSN membership/revocation propagation. Foundational piece that multiple parameters depend on.

6. **Design: Router Management Plane** — Configuration, monitoring, operations, upgrades. How routers are provisioned, how the first control key is bootstrapped (VR-C), rolling updates, observability.

7. **Design: Admission Keying Particulars** — Authentication and reauthentication flows, revocation propagation mechanics, key lifecycle (creation, distribution, rotation, expiry, revocation). Depends on router distributed database design (#5).
