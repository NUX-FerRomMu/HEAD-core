# ADR-003: Git isolation and exact approval

Status: accepted.

Allocate one development branch/worktree per implementation ticket and a separate checkout for validation. A writer lease selects the developer or human. Validation and approval reference full candidate/base hashes. Integration uses a dedicated worktree and fast-forwards the target ref with compare-and-swap only when the candidate contains the current approved target base. Rebase/update work happens before approval, never as an invisible integration change.

Alternatives: shared worktrees allow accidental interference; cherry-picking/squashing after approval produces an unreviewed commit; force updating target refs can overwrite concurrent changes. Worktrees consume disk and require careful cleanup, but keep concurrency and audit comprehensible.

Consequences: when target advances, invalidate approval, update the candidate branch, rerun validation, and request review again. A checked-out target branch may only be updated safely when its worktree is clean and has no ADE writer; external changes cause a stop. Keep agent branches and artifacts until explicit cleanup, never delete dirty worktrees automatically.
