# Shared contract examples

These are design fixtures, not implementation test results. Ticket owners import or translate them into meaningful Go/Lua tests without changing their expected semantics. Mnemonic IDs/hashes in scenarios are symbolic; actual records use valid IDs/full Git object IDs.

- [attention.json](attention.json): exact ranking/default penalties, including old low priority outranking new critical work.
- [workflow.json](workflow.json): current/stale/unauthorized workflow examples and expected observable effects.

Run `python3 specs/check.py` from the repository root to verify the planning package's links, ticket graph, coverage, statuses and queue arithmetic. It does not execute or validate an ADE application. No third-party Python dependencies are needed; Python is only a document-maintenance tool, not part of Core's chosen runtime.
