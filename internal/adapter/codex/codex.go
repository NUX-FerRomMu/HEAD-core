// Package codex connects to the installed Codex CLI using its existing login
// and configuration. It never invokes a shell or changes approval policy.
package codex

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

type Provider struct{ Binary string }

func (p Provider) binary() string {
	if p.Binary == "" {
		return "codex"
	}
	return p.Binary
}
func (p Provider) command(ctx context.Context, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, p.binary(), args...)
	configureProcess(cmd)
	cmd.WaitDelay = 2 * time.Second
	return cmd
}
func (p Provider) Status(ctx context.Context) (domain.ProviderStatus, error) {
	status := domain.ProviderStatus{}
	cmd := p.command(ctx, "--version")
	output, err := cmd.Output()
	if err != nil {
		if ctx.Err() != nil {
			return status, ctx.Err()
		}
		status.Detail = "Codex unavailable; install the Codex CLI and put it on PATH"
		return status, nil
	}
	status.Available = true
	status.Version = strings.TrimSpace(string(output))
	if err = p.command(ctx, "login", "status").Run(); err != nil {
		if ctx.Err() != nil {
			return status, ctx.Err()
		}
		status.Detail = "Codex login is not ready; run codex login"
		return status, nil
	}
	status.LoginPresent = true
	status.Detail = "Local login found; a run verifies server access"
	return status, nil
}
func (p Provider) Run(ctx context.Context, task domain.AgentTask, out, stderr io.Writer) error {
	args := []string{"exec", "--json", "--color", "never", "--skip-git-repo-check", "-C", task.Project}
	if task.Model != "" {
		args = append(args, "--model", task.Model)
	}
	args = append(args, "-")
	cmd := p.command(ctx, args...)
	cmd.Dir = task.Project
	cmd.Stdin = strings.NewReader(task.Prompt)
	cmd.Stdout = out
	cmd.Stderr = stderr
	err := cmd.Run()
	if ctx.Err() != nil {
		return fmt.Errorf("Codex run cancelled: %w", ctx.Err())
	}
	if err != nil {
		return fmt.Errorf("Codex run failed: %w", err)
	}
	return nil
}
