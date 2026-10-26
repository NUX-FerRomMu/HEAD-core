package cli

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/NUX-FerRomMu/HEAD-core/internal/store"
)

type Command struct {
	Summary string
	Run     func(context.Context, *Env, []string) error
}

type Env struct {
	Project  string
	StateDir string
	Out      io.Writer
	Err      io.Writer
	db       *sql.DB
}

func (e *Env) DB() (*sql.DB, error) {
	if e.db == nil {
		key := fmt.Sprintf("%x", sha256.Sum256([]byte(e.Project)))
		db, err := store.Open(filepath.Join(e.StateDir, "projects", key, "state.sqlite"))
		if err != nil {
			return nil, err
		}
		e.db = db
	}
	return e.db, nil
}
func (e *Env) JSON(value any) error { return json.NewEncoder(e.Out).Encode(value) }

func Run(ctx context.Context, args []string, out, errOut io.Writer, version string, commands map[string]Command) int {
	fs := Flags("hea")
	project := fs.String("project", ".", "project directory")
	stateDir := fs.String("data-dir", "", "local state directory")
	help := fs.Bool("help", false, "show help")
	fs.BoolVar(help, "h", false, "show help")
	showVersion := fs.Bool("version", false, "show version")
	fail := func(err error) int { fmt.Fprintln(errOut, "hea:", err); return 1 }
	if err := fs.Parse(args); err != nil {
		return fail(err)
	}
	if *help || len(args) == 0 {
		fmt.Fprintln(out, "Usage: hea [--project DIR] [--data-dir DIR] COMMAND [ARGS]\nFeature commands return JSON. Global options go before the command.\n\nOptions: --help, --version")
		names := make([]string, 0, len(commands))
		for name := range commands {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Fprintf(out, "  %-14s %s\n", name, commands[name].Summary)
		}
		return 0
	}
	if *showVersion {
		fmt.Fprintln(out, "hea", version)
		return 0
	}
	if fs.NArg() == 0 {
		return fail(errors.New("missing command; use --help"))
	}
	command, ok := commands[fs.Arg(0)]
	if !ok {
		return fail(fmt.Errorf("unknown command %q; use --help", fs.Arg(0)))
	}
	root, err := ProjectRoot(*project)
	if err != nil {
		return fail(err)
	}
	dir, err := StateDir(*stateDir)
	if err != nil {
		return fail(err)
	}
	env := &Env{Project: root, StateDir: dir, Out: out, Err: errOut}
	defer func() {
		if env.db != nil {
			env.db.Close()
		}
	}()
	if err = command.Run(ctx, env, fs.Args()[1:]); err != nil {
		return fail(err)
	}
	return 0
}

// ProjectRoot treats each Git worktree as its own project. Outside Git, it uses
// the selected directory. Symlink aliases resolve to the same local state.
func ProjectRoot(path string) (string, error) {
	root, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		return "", fmt.Errorf("project: %w", err)
	}
	info, err := os.Stat(root)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("project is not a directory: %s", root)
	}
	for dir := root; ; dir = filepath.Dir(dir) {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}
		if filepath.Dir(dir) == dir {
			break
		}
	}
	return root, nil
}
func StateDir(override string) (string, error) {
	if override != "" {
		return filepath.Abs(override)
	}
	root := os.Getenv("XDG_STATE_HOME")
	if root == "" || !filepath.IsAbs(root) {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		root = filepath.Join(home, ".local", "state")
	}
	return filepath.Join(root, "hea"), nil
}
func Flags(name string) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs
}
func Parse(fs *flag.FlagSet, args []string, positionals int) error {
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != positionals {
		return fmt.Errorf("%s: expected %d positional arguments, got %d", fs.Name(), positionals, fs.NArg())
	}
	return nil
}

type Strings []string

func (s *Strings) String() string     { return fmt.Sprint([]string(*s)) }
func (s *Strings) Set(v string) error { *s = append(*s, v); return nil }
