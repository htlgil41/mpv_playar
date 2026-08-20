package helpers

import (
	"context"
	"fmt"
	"playar/internal/libs"
	"time"
)

type StartServerUnixResponse struct {
	Pid  int
	Path string
}

func StartServerUnix(
	displays int,
	config *libs.ConfigApp,
) (StartServerUnixResponse, error) {
	cmd := ExecuteCommand(
		"mpv",
		[]string{
			fmt.Sprintf("--input-ipc-server=%s", config.Paths.Path_servermpv),
			"--idle=yes",
			"--fullscreen=yes",
			"--keepaspect=no",
			"--no-osc",
			"--osd-level=0",
			"--cursor-autohide=0",
			fmt.Sprintf("--vf=scale=%d:%d", config.App.Scale_x, config.App.Scale_y),
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
