# ADR-001: Go Core and Lua client

Status: accepted. Source: explicit user instruction plus the independent-Core requirement.

Choose a Go modular monolith and one binary for daemon, CLI, and Neovim bridge. Choose Lua for the Neovim plugin. Go owns all domain state and rules. Lua renders server projections and forwards user intent. Use standard Go tools and a small set of pinned dependencies.

Alternatives: TypeScript Core was explicitly rejected by the user; embedding the business logic in Neovim would violate frontend independence; microservices would complicate local installation and recovery without an MVP benefit. Rust is viable but not selected given the user's Go choice.

Consequences: maintain JSON compatibility across Go/Lua, implement process orchestration explicitly, and avoid runtime Go plugins. A future GUI uses the same protocol. This decision does not prescribe implementation of every possible extension before MVP.
