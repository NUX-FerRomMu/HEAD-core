# Product context

HEAD supports development with coding agents, with tickets as the main unit of work. The goal is to use HEAD to develop HEAD itself.

## Default workflow

A project opens from a local folder. Its backlog can be maintained manually or refined with an agent. Backlog, development and testing are the initial agent roles.

Agent assignment and review approval are manual by default.

Development agents normally work on a ticket and remain available through corrections. Conversations follow that context by default; shared agents, conversations or a project-wide coordinator remain possible.

The usual cycle is assignment, implementation, commit, testing, review and integration, returning to implementation when corrections are needed. Review includes diffs, line comments, replies and manual edits.

## Attention and navigation

A priority queue is central: it groups interventions and review actions, ages waiting items, and offers `next`, quick previews, direct actions, and lower-priority unread requeue. The user can always choose another item.

Switching tickets restores a short factual summary, current state, pending decisions, and links to chat, changes, and tests. Navigation and mappings should fit normal Vim usage.

Configuration is local to each installation. The first platform is Linux, with other platforms and frontends possible through the core's boundaries.
