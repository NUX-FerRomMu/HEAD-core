# Product context

HEAD organizes development around tickets and the human decisions they need. Neovim is the first frontend; the core remains usable independently. The goal is an MVP capable of developing HEAD itself.

## Essential workflow

- Open a local Git project, maintain its backlog and Kanban, and refine work into tickets with enough context for another agent.
- The human assigns agents. Initial roles are backlog refinement, development, and testing. Backlog agents propose changes for human acceptance; they do not assign work.
- A developer belongs to a ticket, remains available through corrections, and closes when the work ends. Conversations belong to those agents, not a detached global chat.
- The intended implementation cycle is `assigned → doing → committed → fixing → committed → done`; successful validation can proceed from committed directly to human review and integration.
- Testing runs against the committed candidate and reports results and available coverage. Failures return work for correction. Human review approval authorizes integration.
- Review supports side-by-side diffs, selected files, anchored comments, agent replies, forced corrections, and manual edits with their own commits.

## Attention and navigation

A priority queue is central: it groups interventions and review actions, ages waiting items, and offers `next`, quick previews, direct actions, and lower-priority unread requeue. The user can always choose another item.

Switching tickets restores a short factual summary, current state, pending decisions, and links to chat, changes, and tests. Navigation and mappings should fit normal Vim usage.

## Boundaries

Settings belong to the local installation; separate clones or machines have independent configuration. Commands, mappings, branch naming and provider selection are configurable.

Humans can perform the work manually. HEAD prevents its own agents and editor actions from writing the same worktree concurrently; reviewing an older commit remains possible.

Start on Linux while keeping platform-specific code isolated. Other frontends and providers should fit behind clear boundaries. Hosted collaboration, a graphical UI and a general extension marketplace are outside the first MVP.
