package helpers

import (
	"context"
	"os/exec"
)

func ExecuteCommand(name string, commad []string, c context.Context) *exec.Cmd {
	cm := exec.CommandContext(c, name, commad...)
	return cm
}
