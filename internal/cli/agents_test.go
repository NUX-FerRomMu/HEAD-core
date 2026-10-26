//go:build linux

package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

func TestAgentOutcomeNotifications(t *testing.T) {
	fixture, err := filepath.Abs("../../tests/fixtures")
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", fixture+string(os.PathListSeparator)+os.Getenv("PATH"))
	project, state := t.TempDir(), t.TempDir()
	run := func(code int, args ...string) string {
		t.Helper()
		var out, stderr bytes.Buffer
		actual := Run(context.Background(), append([]string{"--project", project, "--data-dir", state}, args...), &out, &stderr, "test", map[string]Command{"agent": {Run: AgentCommand}, "notification": {Run: NotificationCommand}})
		if actual != code {
			t.Fatalf("%v: code %d, %s", args, actual, stderr.String())
		}
		return out.String()
	}
	if output := run(0, "agent", "run", "--prompt", "OK"); !strings.Contains(output, "hea.run.completed") || !strings.Contains(output, "fixture response") {
		t.Fatal(output)
	}
	if output := run(1, "agent", "run", "--prompt", "FAIL"); !strings.Contains(output, "hea.run.failed") {
		t.Fatal(output)
	}
	run(1, "agent", "run", "--prompt", " ")
	var items []domain.Notification
	if err := json.Unmarshal([]byte(run(0, "notification", "list")), &items); err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 || !strings.Contains(items[0].Text, "failed") || items[0].Source != "codex" || !strings.Contains(items[1].Text, "completed") {
		t.Fatalf("%+v", items)
	}
}
