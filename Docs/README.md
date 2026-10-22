# ADE implementation design

Start here. This is the **design and implementation backlog** for [NUX-FerRomMu/HEAD-core](https://github.com/NUX-FerRomMu/HEAD-core), not an implemented application. `Docs/` contains architecture and supporting decisions; `specs/` contains the contracts and agent-ready work. The product name ADE is retained from the original context. Tickets are tracked as repository files; no external tracker issues have been created.

The product source is [the original Spanish context](source/ADE_CONTEXTO.md). The user's subsequent instruction fixes **Go for Core and Lua for Neovim**. The other technical choices below are the architecture selected for implementation; they are not claims that a prototype has already validated them.

## Read in this order

1. [Product and MVP requirements](product.md).
2. [Architecture and package boundaries](architecture.md).
3. [Decision records](adr/README.md).
4. [Domain and invariants](../specs/contracts/domain.md), [RPC](../specs/contracts/rpc.md), [Git](../specs/contracts/git.md), [providers](../specs/contracts/provider.md), [attention](../specs/contracts/attention.md), and [configuration](../specs/contracts/configuration.md).
5. [Delivery roadmap](delivery.md), [backlog](../specs/backlog.md), and [agent execution instructions](../specs/AGENTS.md).

## Authority and change control

Order of authority: explicit user instructions; product requirements; accepted decision records; shared contracts; individual tickets; examples. A later user instruction can supersede this baseline. Record the resulting change in the affected docs before implementing it. Conflicting lower-level documents are defects to fix, not permission to guess. Implementation agents must not silently redesign shared contracts to make a ticket easier.

Contract version is `ade.v1`; this design baseline is `2026-10-22`. Tickets are implementation-ready when their dependencies and explicit gates are satisfied. Only ADE-001 and ADE-015 start ready. All other work is deliberately blocked until its prerequisites land. `ready` means adequately specified, not assigned: the user still assigns every ticket.

## Navigation

- [Testing and release strategy](testing.md)
- [Operational recovery and local trust boundary](operations.md)
- [Open risks, chosen defaults, and research gate](decisions-and-risks.md)
- [Requirement-to-ticket coverage](../specs/traceability.md)
- [Machine-readable execution graph](../specs/backlog.json)
- [Reusable ticket format](../specs/TEMPLATE.md)
- [Planning-package validation report](package-validation.md)

The initial production deliverable is a Linux local development tool. Portability is a package-boundary requirement; Windows/macOS support is not claimed until their platform adapters and CI pass.
