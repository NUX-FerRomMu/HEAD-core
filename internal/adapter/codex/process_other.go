//go:build !linux

package codex

import "os/exec"

// Linux is the supported initial platform. Other platforms retain the standard
// context cancellation behavior until their process-tree adapters are added.
func configureProcess(cmd *exec.Cmd) {}
