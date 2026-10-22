# Configuration contract — ade.v1

This is local installation configuration, not repository state. ADE-005 owns decoding/validation; later adapters consume immutable snapshots. Unknown input keys are rejected with field paths to catch typos. Config output redacts sensitive profile arguments and never returns credential values. `config.update` uses an object merge patch: supplied leaves replace values, arrays replace whole arrays, null removes an optional override; removing required global settings restores the documented default or fails if none exists. Result includes monotonically increasing version.

Example of effective config (model intentionally unspecified until Codex probe):

```json
{
  "schema_version": 1,
  "target_branch": "main",
  "branch_template": "{kind}/ade-{ticket_id}-{slug}",
  "profiles": {
    "codex-local": {
      "provider": "codex",
      "executable": "codex",
      "args": [],
      "model": null,
      "environment_allowlist": [],
      "resume": false,
      "max_run_seconds": 3600
    }
  },
  "validation": {
    "checks": [
      {"id":"go-tests","name":"Go tests","argv":["go","test","./..."],"required":true,"timeout_seconds":600}
    ],
    "coverage": null
  },
  "concurrency": {"developers":2,"validations":2},
  "attention": {
    "weights":{"critical":300,"high":200,"normal":100,"low":0},
    "aging_interval_seconds":300,"aging_step":10,
    "requeue_penalty":50,"max_requeue_penalty":200
  },
  "limits": {"log_bytes_per_run":20971520,"provider_input_bytes":65536}
}
```

This example configures Go checks for developing ADE itself. Other projects supply their own checks; do not infer a valid suite or default to Go for arbitrary repositories. Empty checks are allowed in stored config but cannot yield a passing validation. `target_branch` must name an existing local branch; when omitted resolve an unambiguous local default or return a selection-required error. Config creation has version 1 even before a project-specific override exists.

Validation digest = SHA-256 of canonical JSON containing check IDs/names/argv/required/timeouts plus coverage configuration, sorted object keys, preserved array order, UTF-8 compact encoding. Include only validation-affecting settings; changing a keymap or queue weight does not invalidate tests. Check IDs are unique stable nonempty strings. Commands are nonempty arrays of strings, no implicit shell expansion, working directory fixed to the validation workspace. Positive finite timeouts/concurrency/limits are bounded by implementation-documented limits; reject zero/negative values.

Optional coverage object: `{format:"lcov",path:"relative/to/workspace/lcov.info",required:false,min_line_percent:null}`. Reject absolute/traversing/symlink-escaping paths and paths tracked as source in the candidate. A nonnull minimum must be 0..100 and implies required coverage. Coverage files are artifacts of that exact run, not reused results. Only remove a pre-existing declared generated coverage artifact after validating its path and confirming it is not tracked source. Per-project overrides cannot alter daemon transport path or permissions.

Profile arguments are pass-through human configuration only after the adapter validates required tool/version compatibility. An empty `args` list above is not a claimed Codex launch invocation: the adapter supplies its verified protocol launch parameters from ADE-015. `environment_allowlist` refers to names, never values; supply required process runtime variables separately using an explicitly documented minimal baseline (for example PATH and the provider's home/config reference). Model choice and resume support are validated by the real adapter. A missing credential is reported without interactive login by Core.

Neovim `require('ade').setup(opts)` config is frontend-local: executable path, optional auto-start, panel widths, mapping enable flag and action→key map. It cannot change ticket state rules, queue scoring or approval guards. Buffer-local suggested mappings are installed only when enabled; no global mappings by default. ADE-028 documents exact public Lua option names and defaults with backward-compatible tests.
