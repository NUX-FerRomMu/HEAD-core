# Product context

HEAD supports development with coding agents, with tickets as the main unit of work. The goal is to use HEAD to develop HEAD itself.

## Default workflow

A project opens from a local folder. Its backlog can be maintained manually or refined with an agent. Backlog, development and testing are the initial agent roles.

Agent assignment and review approval are manual by default. They are configurable behaviors, not permanent human-only restrictions.

Development agents normally work on a ticket and remain available through corrections. Conversations follow that context by default; shared agents, conversations or a project-wide coordinator remain possible.

The usual cycle is assignment, implementation, commit, testing, review and integration, returning to implementation when corrections are needed. Review includes diffs, line comments, replies and manual edits.

## Attention and navigation

The attention queue collects pending interventions and reviews. Priority and aging help surface waiting work; next, preview and requeue support quick navigation without restricting what the user can select.

Ticket context brings together current state, a short summary, pending actions, conversations, changes and test results. Neovim supplies editing and configurable navigation.

Configuration is local to each installation. The first platform is Linux, with other platforms and frontends possible through the core's boundaries.
