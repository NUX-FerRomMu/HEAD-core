//go:build linux

package codex

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

func executable(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "codex")
	if err := os.WriteFile(path, []byte("#!/bin/sh\nset -eu\n"+body), 0700); err != nil {
		t.Fatal(err)
	}
	return path
}
func TestArgumentsStdinAndFailure(t *testing.T) {
	dir := t.TempDir()
	argsFile := filepath.Join(dir, "args")
	promptFile := filepath.Join(dir, "prompt")
	t.Setenv("HEAD_TEST_ARGS", argsFile)
	t.Setenv("HEAD_TEST_PROMPT", promptFile)
	p := Provider{Binary: executable(t, `printf '%s\n' "$@" > "$HEAD_TEST_ARGS"
cat > "$HEAD_TEST_PROMPT"
printf '{"type":"turn.completed"}\n'
printf 'provider diagnostic\n' >&2
`)}
	var out, stderr bytes.Buffer
	prompt := "Keep $(touch /tmp/not-executed) literal\nCafé"
	if err := p.Run(context.Background(), domain.AgentTask{Project: dir, Prompt: prompt, Model: "selected-model"}, &out, &stderr); err != nil {
		t.Fatal(err)
	}
	args, _ := os.ReadFile(argsFile)
	want := "exec\n--json\n--color\nnever\n--skip-git-repo-check\n-C\n" + dir + "\n--model\nselected-model\n-\n"
	if string(args) != want {
		t.Fatalf("arguments: %s", args)
	}
	got, _ := os.ReadFile(promptFile)
	if string(got) != prompt {
		t.Fatalf("prompt: %q", got)
	}
	var event map[string]any
	if err := json.Unmarshal(out.Bytes(), &event); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stderr.String(), "provider diagnostic") {
		t.Fatal(stderr.String())
	}
	p.Binary = executable(t, "exit 7\n")
	if err := p.Run(context.Background(), domain.AgentTask{Project: dir, Prompt: "x"}, io.Discard, io.Discard); err == nil || !strings.Contains(err.Error(), "exit status 7") {
		t.Fatalf("%v", err)
	}
	p.Binary = filepath.Join(t.TempDir(), "missing")
	if err := p.Run(context.Background(), domain.AgentTask{Project: dir, Prompt: "x"}, io.Discard, io.Discard); err == nil {
		t.Fatal("missing executable succeeded")
	}
}
func TestStatus(t *testing.T) {
	p := Provider{Binary: executable(t, `if [ "$1" = "--version" ]; then echo codex-test; exit 0; fi
if [ "$1" = login ] && [ "$2" = status ]; then exit 0; fi
exit 1
`)}
	s, err := p.Status(context.Background())
	if err != nil || !s.Available || !s.LoginPresent || s.Version != "codex-test" {
		t.Fatalf("%+v %v", s, err)
	}
	p.Binary = executable(t, `if [ "$1" = "--version" ]; then echo codex-test; exit 0; fi
exit 1
`)
	s, err = p.Status(context.Background())
	if err != nil || !s.Available || s.LoginPresent {
		t.Fatalf("%+v %v", s, err)
	}
	p.Binary = "/does-not-exist"
	s, err = p.Status(context.Background())
	if err != nil || s.Available {
		t.Fatalf("%+v %v", s, err)
	}
}

type signalWriter struct {
	once  sync.Once
	ready chan struct{}
}

func (w *signalWriter) Write(p []byte) (int, error) {
	w.once.Do(func() { close(w.ready) })
	return len(p), nil
}
func TestStreamingAndProcessTreeCancellation(t *testing.T) {
	dir := t.TempDir()
	pidFile := filepath.Join(dir, "child")
	t.Setenv("HEAD_TEST_CHILD", pidFile)
	p := Provider{Binary: executable(t, `sleep 30 &
echo $! > "$HEAD_TEST_CHILD"
printf '{"type":"turn.started"}\n'
wait
`)}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	writer := &signalWriter{ready: make(chan struct{})}
	done := make(chan error, 1)
	go func() { done <- p.Run(ctx, domain.AgentTask{Project: dir, Prompt: "x"}, writer, io.Discard) }()
	select {
	case <-writer.ready:
	case <-time.After(3 * time.Second):
		t.Fatal("output did not stream before completion")
	}
	select {
	case err := <-done:
		t.Fatalf("ended before cancellation: %v", err)
	default:
	}
	cancel()
	select {
	case err := <-done:
		if err == nil || !strings.Contains(err.Error(), "cancelled") {
			t.Fatalf("%v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("run did not stop")
	}
	pid, err := os.ReadFile(pidFile)
	if err != nil {
		t.Fatal(err)
	}
	statPath := "/proc/" + strings.TrimSpace(string(pid)) + "/stat"
	deadline := time.Now().Add(time.Second)
	for {
		stat, err := os.ReadFile(statPath)
		if os.IsNotExist(err) {
			break
		}
		// A zombie has exited; the container's init owns reaping adopted children.
		if err == nil && strings.Contains(string(stat), ") Z ") {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("descendant still running: %s (%v)", stat, err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
