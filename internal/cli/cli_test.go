package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestHelpVersionAndErrorsDoNotCreateState(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "not-created")
	for _, tc := range []struct {
		args     []string
		code     int
		contains string
	}{
		{[]string{"--help"}, 0, "Usage:"}, {[]string{"--version"}, 0, "hea dev"}, {[]string{"--invalid"}, 1, ""}, {[]string{"missing"}, 1, ""},
	} {
		var out, stderr bytes.Buffer
		args := append([]string{"--data-dir", dir, "--project", "/does-not-exist"}, tc.args...)
		code := Run(context.Background(), args, &out, &stderr, "dev", nil)
		if code != tc.code {
			t.Fatalf("%v: code %d stderr %s", tc.args, code, stderr.String())
		}
		if tc.code == 0 && !bytes.Contains(out.Bytes(), []byte(tc.contains)) {
			t.Fatal(out.String())
		}
		if tc.code != 0 && stderr.Len() == 0 {
			t.Fatal("missing diagnostic")
		}
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Fatalf("created state: %v", err)
	}
}
func TestProjectRoot(t *testing.T) {
	for _, worktree := range []bool{false, true} {
		root := t.TempDir()
		sub := filepath.Join(root, "src")
		os.Mkdir(sub, 0700)
		if worktree {
			os.WriteFile(filepath.Join(root, ".git"), []byte("gitdir: elsewhere"), 0600)
		} else {
			os.Mkdir(filepath.Join(root, ".git"), 0700)
		}
		alias := filepath.Join(t.TempDir(), "alias")
		if err := os.Symlink(sub, alias); err != nil {
			t.Fatal(err)
		}
		got, err := ProjectRoot(alias)
		if err != nil || got != root {
			t.Fatalf("%q %v", got, err)
		}
	}
}
func TestProjectIsolation(t *testing.T) {
	state := t.TempDir()
	first := &Env{Project: t.TempDir(), StateDir: state}
	second := &Env{Project: t.TempDir(), StateDir: state}
	a, err := first.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer a.Close()
	b, err := second.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()
	if _, err = a.Exec("CREATE TABLE example (value TEXT)"); err != nil {
		t.Fatal(err)
	}
	if _, err = b.Exec("SELECT * FROM example"); err == nil {
		t.Fatal("project data leaked")
	}
}
