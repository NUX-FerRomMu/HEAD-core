# ADR-004: Tool-based Codex and bounded sessions

Status: accepted; live compatibility gated by ADE-015.

Implement a provider-neutral lifecycle port with Codex as the first real adapter, using a supported structured interface of the installed Codex tool. Do not depend exclusively on the OpenAI API, scrape a terminal UI, or invent command flags. ADE-015 records actual executable/version, supported launch/events/approval/chat/resume behavior and fixtures before the adapter is built.

A logical ADE agent belongs to one role/ticket; provider processes may end between turns. Developer agents remain addressable until completion/cancellation even if the subprocess exits. Chat is delivered only through a supported capability: a live turn when supported, or a new/resumed turn with explicit UI status. A new context includes the durable handoff. Unknown session continuity is never presented as preserved memory.

If installed Codex lacks a suitable structured interface, mark the real-provider path blocked with evidence and an exact supported-version recommendation. Continue with the deterministic fake provider and Core. Do not declare the real-provider MVP complete or switch to API-only integration silently.
