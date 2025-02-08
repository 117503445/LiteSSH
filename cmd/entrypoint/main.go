package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "embed"

	"github.com/117503445/goutils"
	"github.com/117503445/vscode-lite-ssh/pkg/cli"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

var codeServerConfig = `auth: password
cert: false`

const dirLogs = "/workspace/logs"

func getLogFileName(dir string) string {
	return fmt.Sprintf("%s/%s.log", dir, goutils.TimeStrSec())
}

func execWithLogs(cmds []string, dirLog string) {
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
		log.Error().Err(err).Msg("Failed to exec")
	}
}

func main() {
	goutils.InitZeroLog()
	goutils.ExecOpt.DumpOutput = true

	err := os.MkdirAll(dirLogs, 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create logs directory")
	}

	cli.CfgLoad()

	codeServerConfigPath := "/root/.config/code-server/config.yaml"
	codeServerPassword := cli.Cli.CodeServerPassword
	if codeServerPassword != "" {
		codeServerConfig += fmt.Sprintf("\npassword: %s", codeServerPassword)
	}
	err = goutils.WriteText(codeServerConfigPath, codeServerConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write code-server config file")
	}

	go func() {
		execWithLogs([]string{"/usr/sbin/code-server"}, fmt.Sprintf("%s/code-server", dirLogs))
	}()

	var isTTY bool
	if isatty.IsTerminal(os.Stdout.Fd()) {
		isTTY = true
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		log.Fatal().Msg("Cygwin terminal is not supported")
	} else {
		isTTY = false
	}

	if isTTY {
		cmd := exec.Command("/bin/fish")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Error().Err(err).Msg("Failed to run fish shell")
		}
	} else {
		goutils.Exec("tail -f /dev/null")
	}
}
