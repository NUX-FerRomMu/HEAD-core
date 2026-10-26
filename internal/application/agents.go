package application

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/NUX-FerRomMu/HEAD-core/internal/domain"
)

type AgentProvider interface {
	Status(context.Context) (domain.ProviderStatus, error)
	Run(context.Context, domain.AgentTask, io.Writer, io.Writer) error
}
type Agents struct{ Provider AgentProvider }

func (a Agents) Run(ctx context.Context, task domain.AgentTask, out, stderr io.Writer) error {
	if strings.TrimSpace(task.Prompt) == "" {
		return fmt.Errorf("prompt must not be empty")
	}
	return a.Provider.Run(ctx, task, out, stderr)
}
