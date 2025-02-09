package common

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/117503445/goutils"
	"github.com/rs/zerolog/log"
)

func getLogFileName(dir string) string {
	return fmt.Sprintf("%s/%s.log", dir, goutils.TimeStrSec())
}
const dirLogs = "/workspace/logs"

func ExecWithLogs(cmds []string, dirsubLog string) {
	dirLog := fmt.Sprintf("%s/%s", dirLogs, dirsubLog)
	
	err := os.MkdirAll(dirLog, 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create code-server logs directory")
	}

	file, err := os.OpenFile(getLogFileName(dirLog), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open goreman.log")
	}
	defer file.Close()

	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdout = file
	cmd.Stderr = file
	if err := cmd.Run(); err != nil {
		log.Error().Str("cmd", cmd.String()).Err(err).Msg("Failed to exec")
	}
}
