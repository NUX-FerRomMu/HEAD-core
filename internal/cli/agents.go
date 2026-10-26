package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/NUX-FerRomMu/HEAD-core/internal/adapter/codex"
	"github.com/NUX-FerRomMu/HEAD-core/internal/application"
	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
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
	if err := e.JSON(map[string]any{"type": "hea.run.started", "project": e.Project}); err != nil {
		return err
	}
	service := application.Agents{Provider: provider}
	err := service.Run(ctx, domain.AgentTask{Project: e.Project, Prompt: prompt, Model: model}, e.Out, e.Err)
	event := map[string]any{"type": "hea.run.completed"}
	if err != nil {
		event["type"] = "hea.run.failed"
		event["error"] = err.Error()
		if ctx.Err() != nil {
			event["type"] = "hea.run.cancelled"
		}
	}
	if writeErr := e.JSON(event); err == nil {
		err = writeErr
	}
	return err
}
