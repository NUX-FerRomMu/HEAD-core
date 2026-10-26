package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
	"testing"
)

func TestNotificationCommands(t *testing.T) {
	project, state := t.TempDir(), t.TempDir()
	run := func(code int, args ...string) []byte {
		t.Helper()
		var out, stderr bytes.Buffer
		actual := Run(context.Background(), append([]string{"--project", project, "--data-dir", state, "notification"}, args...), &out, &stderr, "test", map[string]Command{"notification": {Run: NotificationCommand}})
		if code != actual {
			t.Fatalf("%v: %d %s", args, actual, stderr.String())
		}
		return out.Bytes()
	}
	if string(run(0, "next")) != "null\n" {
		t.Fatal("expected null")
	}
	var n domain.Notification
	if err := json.Unmarshal(run(0, "add", "--text", "Needs review", "--priority", "3"), &n); err != nil {
		t.Fatal(err)
	}
	id := fmt.Sprint(n.ID)
	run(0, "requeue", id)
	run(0, "dismiss", id)
	if string(run(0, "list")) != "[]\n" {
		t.Fatal("dismissed notification listed")
	}
	run(1, "dismiss", id)
	run(1, "add", "--text", "x", "--priority", "4")
	run(1, "add", "--text", " ")
	run(1, "next", "unexpected")
	run(1, "dismiss", "invalid")
}
