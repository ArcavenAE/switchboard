# Naming Parking Lot: The Third Node Type

**Status:** Deferred to future brainstorming session (Analogical Thinking technique recommended)

## The Problem

Switchboard has three node types. Two are named:
- **Console** — accesses remote tmux sessions
- **Control** — manages virtual switched networks

The third type — the node that *publishes* tmux sessions over the network, running on a computer with zero or more tmux sessions — needs a name.

## Candidates So Far

| Name | Pros | Cons |
|------|------|------|
| **Anchor** (anker) | Tmux sessions "anchor" to the VSN. Conveys stability, permanence. | Doesn't convey the multi-session/access aspect. Slightly passive. |
| **Fan** | Implies multiple access points, fan-out. | Too generic. Doesn't convey the tmux publishing role. |
| **Exchange** | Telephony heritage (telephone exchange). Fits the switchboard metaphor well. | Could be confused with message exchange patterns. |
| **CO (Central Office)** | Deep telephony analog. | Too telco-jargon for a modern tool. |
| **Switch** | Switch connecting to router feels natural. Extends the switchboard metaphor. | Could be confused with network switches. "Switch node" on a "switchboard" is redundant. |
| **Access Node** | Clear, descriptive. "Access node, console node, control node" is a clean trio. | Generic. Doesn't evoke the switchboard metaphor. |

## Abbreviation Concerns

- "Switchboard Access Node" = SAN (conflicts with Storage Area Network)
- "Switchboard Console Node" = SCN
- "Switchboard Control Node" = SCN (collision!)
- Need to differentiate console and control in abbreviated form

## Maker's Current Favorite

**Access Node** — clean trio with console node and control node. Descriptive. Professional.

## Future Session Notes

- Use Analogical Thinking technique: what is this node *like*? Telephony (operator, station, trunk), postal (office, depot), radio (transmitter, beacon), computing (host, server, daemon, agent)
- The switchboard metaphor is rich — explore: operator, jack, plug, cord, trunk, line, extension, position, board
- Consider: does the name need to be a single word or can it be compound?
- Consider: will users type this name frequently (CLI commands)?
