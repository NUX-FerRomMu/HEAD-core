# HEAD-003 — Create and list tickets

Dependency: HEAD-002. Scope: two application operations with persistence.

Add a ticket record with generated ID, title, description and acceptance criteria. Create accepts a nonempty trimmed title; description and criteria may be empty while drafting. List returns stored tickets in insertion order. Add the required migration and application-facing repository interface.

## Acceptance

- Creating two tickets gives distinct IDs; listing after reopening the database returns their original content and order.
- A blank title fails without inserting a row. A failed write returns an error, not a successful ticket.
- Application/domain code does not import the SQLite driver or CLI package.

Verification: validation tests and temporary-database integration tests. Editing, lifecycle states, project relationships and agent assignment are outside this ticket.
