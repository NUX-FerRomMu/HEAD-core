package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/adapter/codex"
	"github.com/NUX-FerRomMu/HEAD-core/internal/application"
	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
	"github.com/NUX-FerRomMu/HEAD-core/internal/store"
)

func AgentCommand(ctx context.Context, e *Env, args []string) error {
	if len(args) == 0 || args[0] == "--help" {
		fmt.Fprintln(e.Out, `Usage: hea [OPTIONS] agent COMMAND
  status
  run --prompt TEXT [--model NAME]
Run streams Codex JSONL plus hea.run.started/completed/failed/cancelled events.
Use Ctrl-C to stop the run.`)
		return nil
	}
	op := args[0]
	fs := Flags("agent " + op)
	var prompt, model string
	switch op {
	case "status":
	case "run":
		fs.StringVar(&prompt, "prompt", "", "task instruction")
		fs.StringVar(&model, "model", "", "Codex model")
	default:
		return fmt.Errorf("unknown agent command %q", op)
	}
	if err := Parse(fs, args[1:], 0); err != nil {
		return err
	}
	provider := codex.Provider{}
	if op == "status" {
		statusCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		status, err := provider.Status(statusCtx)
		if err != nil {
			return err
		}
		return e.JSON(status)
	}
	if strings.TrimSpace(prompt) == "" {
		return fmt.Errorf("prompt must not be empty")
	}
	db, err := e.DB()
	if err != nil {
		return err
	}
	notifications, err := store.NewNotifications(ctx, db)
	if err != nil {
		return err
	}
	if err := e.JSON(map[string]any{"type": "hea.run.started", "project": e.Project}); err != nil {
		return err
	}
	service := application.Agents{Provider: provider}
	err = service.Run(ctx, domain.AgentTask{Project: e.Project, Prompt: prompt, Model: model}, e.Out, e.Err)
	event := map[string]any{"type": "hea.run.completed"}
	if err != nil {
		event["type"] = "hea.run.failed"
		event["error"] = err.Error()
		if ctx.Err() != nil {
			event["type"] = "hea.run.cancelled"
		}
	}
	// The run context may already be cancelled; still persist its outcome.
	outcomeCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	text, priority := "Codex task completed", 2
	if err != nil {
		text, priority = "Codex task failed: "+err.Error(), 3
	}
	if ctx.Err() != nil {
		text, priority = "Codex task cancelled", 1
	}
	_, notificationErr := (application.Notifications{Repo: notifications}).Add(outcomeCtx, text, "codex", priority)
	if notificationErr != nil {
		fmt.Fprintln(e.Err, "record Codex outcome:", notificationErr)
		if err == nil {
			err = fmt.Errorf("task completed, but recording its notification failed: %w", notificationErr)
		}
	}
	if writeErr := e.JSON(event); err == nil {
		err = writeErr
	}
	return err
}
