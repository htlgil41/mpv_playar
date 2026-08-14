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

func StartServerUnix(displays int) (StartServerUnixResponse, error) {
	cmd := ExecuteCommand(
		"mpv",
		[]string{
			"--input-ipc-server=/tmp/mpvsocket",
			"--idle=yes",
			"--fullscreen=yes",
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
