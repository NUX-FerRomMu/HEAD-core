#!/usr/bin/env python3
"""Validate this planning package. Does not test an ADE implementation."""
from datetime import datetime
import json
from pathlib import Path
import re
import sys

ROOT = Path(__file__).resolve().parent.parent
errors = []


def require(condition, message):
    if not condition:
        errors.append(message)


def main():
    data = json.loads((ROOT / "specs/backlog.json").read_text())
    tickets = data["tickets"]
    by_id = {t["id"]: t for t in tickets}
    require(len(by_id) == len(tickets), "Duplicate ticket IDs")
    require(data["language"] == {"core": "Go", "neovim": "Lua"}, "Language decision drift")
    epic_ids = {e["id"] for e in data["epics"]}
    md_backlog = (ROOT / "specs/backlog.md").read_text()
    requirements = set()
    for ticket in tickets:
        tid = ticket["id"]
        path = ROOT / ticket["path"]
        require(path.is_file(), f"{tid}: missing spec {path}")
        require(ticket["epic"] in epic_ids, f"{tid}: unknown epic")
        require(ticket["status"] in {"ready", "blocked", "in_progress", "review", "done"}, f"{tid}: unknown status")
        require(len(set(ticket["dependencies"])) == len(ticket["dependencies"]), f"{tid}: repeated dependency")
        require(ticket["gate"] is None or ticket["gate"] in data["gates"], f"{tid}: missing gate")
        for dep in ticket["dependencies"]:
            require(dep in by_id and dep != tid, f"{tid}: invalid dependency {dep}")
        if ticket["status"] == "ready":
            require(all(by_id.get(dep, {}).get("status") == "done" for dep in ticket["dependencies"]), f"{tid}: ready before dependencies done")
            if ticket["gate"]:
                require(data["gates"][ticket["gate"]]["status"] == "satisfied", f"{tid}: ready before gate satisfied")
        if path.is_file():
            text = path.read_text()
            require(f"- Status: {ticket['status']};" in text, f"{tid}: spec status differs from JSON")
            expected = ", ".join(ticket["dependencies"]) or "none"
            require(f"- Dependencies: {expected}\n" in text, f"{tid}: dependency text differs from JSON")
            gate = ticket["gate"] or "none"
            require(f"- Evidence gate: {gate}\n" in text, f"{tid}: gate differs from JSON")
            for heading in ["Outcome", "Required reading", "Scope and implementation boundaries", "Acceptance criteria", "Verification", "Handoff"]:
                require(f"## {heading}" in text, f"{tid}: missing {heading}")
            require(len(re.findall(r"\*\*AC\d+\*\*", text)) >= 3, f"{tid}: insufficient acceptance examples")
        line = next((line for line in md_backlog.splitlines() if line.startswith(f"| [{tid}]")), "")
        require(bool(line), f"{tid}: absent from backlog table")
        if line:
            cells = [c.strip() for c in line.split("|")][1:-1]
            require(cells[4] == ticket["status"], f"{tid}: backlog status differs from JSON")
            require(cells[5] == (", ".join(ticket["dependencies"]) or "—"), f"{tid}: backlog dependencies differ from JSON")
            require(cells[6] == (ticket["gate"] or "—"), f"{tid}: backlog gate differs from JSON")
        requirements.update(ticket["requirements"])
    require(requirements == {f"R{i:02}" for i in range(1, 17)}, "Requirement coverage missing or unknown IDs")
    covered = []
    for epic in data["epics"]:
        require((ROOT / epic["path"]).is_file(), f"Missing epic {epic['id']}")
        covered.extend(epic["tickets"])
        for tid in epic["tickets"]:
            require(tid in by_id and by_id[tid]["epic"] == epic["id"], f"Epic membership mismatch: {tid}")
    require(sorted(covered) == sorted(by_id), "Epic ticket coverage differs from manifest")

    visiting, visited, order = set(), set(), []

    def visit(tid):
        if tid in visiting:
            errors.append(f"Dependency cycle at {tid}")
            return
        if tid in visited or tid not in by_id:
            return
        visiting.add(tid)
        for dep in by_id[tid]["dependencies"]:
            visit(dep)
        visiting.remove(tid)
        visited.add(tid)
        order.append(tid)

    for tid in by_id:
        visit(tid)
    files = [p for directory in [ROOT / "Docs", ROOT / "specs"] for p in directory.rglob("*.md")]
    link_count = 0
    for path in files:
        for target in re.findall(r"\[[^\]]*\]\(([^)]+)\)", path.read_text()):
            if re.match(r"^[a-zA-Z]+://", target) or target.startswith("#"):
                continue
            target = target.split("#", 1)[0].strip("<>")
            resolved = (path.parent / target).resolve()
            require(resolved.exists(), f"Broken link in {path.relative_to(ROOT)}: {target}")
            link_count += 1
    fixture = json.loads((ROOT / "specs/fixtures/attention.json").read_text())
    policy = fixture["policy"]
    now = datetime.fromisoformat(fixture["evaluated_at"].replace("Z", "+00:00"))
    scores = {}
    for item in fixture["items"]:
        created = datetime.fromisoformat(item["created_at"].replace("Z", "+00:00"))
        age = max(0, int((now - created).total_seconds() // policy["aging_interval_seconds"]))
        score = policy["weights"][item["severity"]] + age * policy["aging_step"] - min(item["requeue_count"] * policy["requeue_penalty"], policy["max_requeue_penalty"])
        require(score == item["expected_score"], f"Queue score mismatch: {item['id']}")
        scores[item["id"]] = score
    ordered = sorted(fixture["items"], key=lambda i: (-scores[i["id"]], i["created_at"], i["id"]))
    require([i["id"] for i in ordered] == fixture["expected_list_order"], "Queue order mismatch")
    unread = [i["id"] for i in ordered if i["state"] == "unread"]
    require(unread == fixture["expected_unread_order"], "Unread order mismatch")
    require(unread[0] == fixture["expected_next"], "Next item mismatch")
    workflows = json.loads((ROOT / "specs/fixtures/workflow.json").read_text())["scenarios"]
    require(len({s["id"] for s in workflows}) == len(workflows), "Duplicate workflow examples")
    for scenario in workflows:
        require(all(scenario.get(k) for k in ["given", "when", "then"]), f"Incomplete workflow example: {scenario['id']}")
    if errors:
        print("\n".join(f"ERROR: {error}" for error in errors), file=sys.stderr)
        return 1
    ready = ", ".join(t["id"] for t in tickets if t["status"] == "ready")
    print(f"PASS: {len(tickets)} tickets, {len(data['epics'])} epics, 16 covered requirements, acyclic dependencies.")
    print(f"PASS: {len(files)} Markdown files, {link_count} local links, statuses/gates consistent.")
    print(f"PASS: {len(fixture['items'])} queue examples and {len(workflows)} workflow examples are internally valid.")
    print(f"Ready to assign: {ready}.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
