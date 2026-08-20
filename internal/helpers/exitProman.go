package helpers

import (
	"os"
	"playar/internal/libs"
)

func ExirProgram(
	config *libs.ConfigApp,
	asunto string,
	suc string,
	body string,
) {
	SendEmailServerExit(config, asunto, suc, body)
	os.Exit(0)
}
