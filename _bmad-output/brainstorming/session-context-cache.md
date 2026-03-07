# Switchboard Brainstorming Session Context Cache

**Purpose:** Reload this file at the start of a new conversation to restore full context for continuing the brainstorming session.

**Last updated:** 2026-03-07
**Next phase:** Morphological Analysis (Phase 2 of brainstorming)

---

## Project: Switchboard

**What it is:** A low-latency, multi-path, end-to-end encrypted tmux session router architecture that establishes virtual switched networks (VSNs) over switchboard routers. Purpose-built for high-trust, low-latency remote CLI experience.

### Architecture Summary

**Three node types:**
- **Access node** — publishes tmux sessions over the network. Runs on a computer with zero or more tmux sessions.
- **Console** — accesses remote tmux sessions
- **Control** — manages virtual switched networks (VSN configuration)

**One router binary, three deployment modes:**
- **E (Edge-local):** Runs on same device as a node, zero upstream connections. Same-LAN setup for two+ machines. The "5-minute getting started" experience. Graduates to PE when pointed at upstream routers.
- **PE (Provider Edge):** Node-facing + router-facing interfaces. Handles VSN admission, node keep-alives, channel setup. The production router.
- **P (Provider/Core):** Router-facing interfaces only. Pure forwarding. THEORETICAL ONLY — not built until proven needed. Same binary, just zero PE interfaces configured.
- P/PE is an interface-level distinction, not a binary-level distinction.

**Virtual Switched Networks (VSNs):**
- Closed networks similar to ZeroTier networks
- Joined via VSN identifier + private key (public key registered on routers)
- One or more consoles can join; one or more access nodes can join
- Nodes communicate E2E via SSH keypair shared off-net
- Routers cannot peer inside VSNs

**No direct node-to-node communication** — always through a router. Architectural invariant.

**Channels:**
- A channel = end-to-end active tmux session over SSH over VSN over Switchboard
- Two half-channels per link: upstream (console->access, keystrokes) and downstream (access->console, terminal output)
- Each half-channel has its own independent timeslice clock, sequence numbers, FEC strategy

**Timeslice framing principle — "The bus leaves on time, full or not":**
- Each direction has its own independent clock at the node (full duplex, no sync needed)
- Pack whatever bytes are ready when the tick fires, ship them
- Frame size is variable, bounded by buffer contents at tick time
- Tick interval tunable: aggressive (5-10ms) for active typing, relaxed (50-100ms) when idle
- Analogous to WAN/satellite acceleration (SKIPS protocol)

**Multi-path forwarding (post-MVP):**
- Frames sent across the TWO fastest links
- Loop prevention via frame tracking
- STTF (Shortest Time to First Byte) scheduling

**Keep-alive & latency:**
- NTP-style time-in-flight measurement in keep-alive packets
- Unidirectional latency visible to routers and nodes
- Nodes multi-homed to multiple PE routers
- Dynamic path switching to lowest overall latency router

**Encryption model:**
- SSH E2E between nodes is the trust layer — Switchboard does NOT double-encrypt
- Switchboard frame envelope adds auth/integrity/sequencing only
- Lighter techniques (HMAC, auth tags) may suffice for node-to-PE validation
- Inter-router security (encrypted channels between routers) — open question, parked

**Circuit signaling (post-MVP):**
- MPLS-style label switching for multi-hop circuits
- PE routers push/pop labels, P routers switch on labels without parsing frame internals
- Borrow from MPLS LDP/RSVP-TE rather than invent

**Network admission model (open questions, parked):**
- Nodes authenticate to PE via VSN credential (private key signs challenge)
- Forwarding entry with TTL after auth (NAT-table analogy)
- Sideband refresh for entry TTL renewal
- IP migration via re-auth (identity is cryptographic, not IP-based, like QUIC connection IDs)
- PE can challenge nodes at any time (anti-hijacking)

**Deployment context:**
- Routers: containers on AWS (EKS, EC2, k8s, docker-compose)
- Nodes MVP: macOS only
- Nodes short-term: macOS + Linux
- Language: Go (validated — Tailscale, Nebula, WireGuard-go precedent)

**MVP scope: Nodes + E router only.**
- Proves edge protocol, channel model, key-based admission, UX
- Defers: multi-hop, multi-path, FEC, control node, link-state, circuit signaling
- Estimated traffic distribution: ~85% intra-LAN, ~14% single-hop, ~1% multi-path

---

## Research Completed (Phase 0)

Full research is in: `_bmad-output/brainstorming/brainstorming-session-2026-03-07-001.md`

### Key Research Conclusions

1. **Routing:** RON validates overlay routing with active probing. Probing scales to ~50 nodes; needs hierarchy beyond that.
2. **ATM:** Variable-size frames better than fixed cells. Timeslice framing avoids ATM's fatal flaw.
3. **Multi-path:** Lowest-RTT-first (STTF) is optimal. Dual-path forwarding validated. Never round-robin across heterogeneous paths.
4. **One-way delay:** NTP-level clock sync sufficient. Only need relative path comparison.
5. **FEC:** XOR parity for keystrokes, Reed-Solomon for larger bursts. Fountain/Raptor overkill for tiny payloads.
6. **Retransmit/resync:** SRT's TLPKTDROP + selective ARQ. Mosh insight: sync state not streams for display direction.
7. **Encryption overhead:** Target 20-24 bytes per frame. Noise/ChaCha20-Poly1305. But SSH E2E may eliminate need for frame-level encryption entirely.
8. **Comparable projects:** No existing tool combines all of Switchboard's capabilities. Genuine gap.
9. **Language:** Go confirmed. Sub-ms GC pauses invisible at terminal latency scale.
10. **Proposed frame layout:** seq_num(4B) + type(1B) + path_id(1B) + fec_group(2B) + payload + AEAD_tag(16B) = 24 bytes overhead

### Key Go Optimization Notes
- `sync.Pool` for frame buffers
- `net/netip` not `net.IP`
- `GOGC` tuning + `GOMEMLIMIT`
- `runtime/metrics` for GC monitoring
- Never allocate per frame

---

## First Principles Established (Phase 1)

1. Human perception is the latency budget (~100ms feels instant)
2. Terminal sessions are asymmetric — two half-channels with different characteristics
3. The bus leaves on time, full or not — timeslice framing, not packet framing
4. No direct node-to-node — always through a router (architectural invariant)
5. One router binary, topology determines role — E, PE, P (P is theoretical only)
6. SSH E2E is the trust layer — Switchboard adds routing and admission, not encryption
7. Circuit signaling is a solved problem — borrow from MPLS when needed
8. The MVP is nodes + E router — prove edge protocol and UX first

---

## Brainstorming Technique Queue

| Phase | Technique | Status |
|-------|-----------|--------|
| 0 | Research Synthesis | COMPLETE |
| 1 | First Principles Thinking | COMPLETE |
| 2 | Morphological Analysis | NEXT |
| 3 | Chaos Engineering | Pending |
| 4 | Analogical Thinking (naming) | Deferred to separate session |

---

## Naming Parking Lot

See: `_bmad-output/brainstorming/naming-node-type-parking-lot.md`

Current favorite: **access node** (clean trio with console node, control node)
Router modes: E (edge-local), PE (provider edge), P (provider core)

---

## How to Continue

1. Load this file for context
2. Load the BMAD analyst agent: `/bmad-agent-bmm-analyst`
3. Select [BP] Brainstorm Project
4. Tell Mary: "Continue from Phase 2: Morphological Analysis. Context is in `_bmad-output/brainstorming/session-context-cache.md`"
