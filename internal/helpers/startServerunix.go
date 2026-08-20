package helpers

import (
	"context"
	"fmt"
	"time"
)

type StartServerUnixResponse struct {
	Pid  int
	Path string
}

func StartServerUnix(displays int, pipe string) (StartServerUnixResponse, error) {
	cmd := ExecuteCommand(
		"mpv",
		[]string{
			fmt.Sprintf("--input-ipc-server=%s", pipe),
			"--idle=yes",
			"--fullscreen=yes",
			"--keepaspect=no",
			"--no-osc",
			"--osd-level=0",
			"--cursor-autohide=0",
			"--vf=scale=3072:256",
			"--no-keepaspect-window",
			fmt.Sprintf("--screen=%d", displays),
			"--loop-playlist=inf",
		},
		context.Background(),
	)

	err_start := cmd.Start()
	if err_start != nil {
		fmt.Println(err_start.Error())
		return StartServerUnixResponse{}, err_start
	}

	time.Sleep(5 * time.Second)
	return StartServerUnixResponse{
		Pid:  cmd.Process.Pid,
		Path: cmd.Path,
	}, nil
}
