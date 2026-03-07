---
stepsCompleted: [1, 2-research, 3-first-principles]
inputDocuments: []
session_topic: 'Switchboard - low-latency multi-path E2E encrypted tmux session router architecture with virtual switched networks'
session_goals: 'Naming, architecture design, edge protocol design, failure mode analysis, use cases'
selected_approach: 'ai-recommended'
techniques_used: ['research-synthesis', 'first-principles-thinking']
ideas_generated: []
context_file: ''
next_phase: 'morphological-analysis'
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
