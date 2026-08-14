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

func StartServerUnix() (StartServerUnixResponse, error) {
	cmd := ExecuteCommand(
		"mpv",
		[]string{
			"--input-ipc-server=/tmp/mpvsocket",
			"--idle=yes",
			"--vo=null",
			"--fullscreen=yes",
			"--screen=0",
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
