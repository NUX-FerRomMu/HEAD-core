package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

func TestTicketWorkflow(t *testing.T) {
	project, state := t.TempDir(), t.TempDir()
	run := func(want int, args ...string) []byte {
		t.Helper()
		var out, stderr bytes.Buffer
		full := append([]string{"--project", project, "--data-dir", state, "ticket"}, args...)
		code := Run(context.Background(), full, &out, &stderr, "test", map[string]Command{"ticket": {Run: TicketCommand}})
		if code != want {
			t.Fatalf("%v: exit %d: %s", args, code, stderr.String())
		}
		return out.Bytes()
	}
	var first, second domain.Ticket
	if err := json.Unmarshal(run(0, "create", "--title", "  Café 🚀  ", "--description", "line 1\nline 2", "--criterion", "uno", "--criterion", "dos"), &first); err != nil {
		t.Fatal(err)
	}
	if first.Title != "Café 🚀" || first.Description != "line 1\nline 2" || !reflect.DeepEqual(first.Criteria, []string{"uno", "dos"}) || first.State != "todo" {
		t.Fatalf("%+v", first)
	}
	json.Unmarshal(run(0, "create", "--title", "second"), &second)
	if second.ID == first.ID {
		t.Fatal("duplicate ID")
	}
	id := fmt.Sprint(first.ID)
	run(0, "move", id, "review")
	run(0, "edit", id, "--title", "changed", "--description", "", "--clear-criteria")
	var got domain.Ticket
	json.Unmarshal(run(0, "get", id), &got)
	if got.Title != "changed" || got.Description != "" || len(got.Criteria) != 0 || got.State != "review" {
		t.Fatalf("%+v", got)
	}
	run(1, "edit", id, "--title", " ")
	run(1, "move", id, "invalid")
	run(1, "move", id, "done", "unexpected")
	run(1, "get", "99999")
	run(1, "create", "--title", " ")
	run(1, "create", "--title", "x", "--unknown")
	run(1, "edit", id)
	var board map[string][]domain.Ticket
	json.Unmarshal(run(0, "board"), &board)
	if len(board) != 4 || len(board["review"]) != 1 || len(board["todo"]) != 1 || len(board["done"]) != 0 {
		t.Fatalf("%+v", board)
	}
	var list []domain.Ticket
	json.Unmarshal(run(0, "list"), &list)
	if len(list) != 2 || list[0].ID != first.ID || list[1].ID != second.ID {
		t.Fatalf("%+v", list)
	}
	// A new project under the same state root cannot see the original tickets.
	project = t.TempDir()
	if string(run(0, "list")) != "[]\n" {
		t.Fatal("project leak")
	}
}
