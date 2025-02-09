package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "embed"

	"github.com/117503445/goutils"
	"github.com/117503445/vscode-lite-ssh/pkg/cli"
	"github.com/117503445/vscode-lite-ssh/pkg/common"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

var codeServerConfig = `cert: false
bind-addr: 0.0.0.0:4444`

func main() {
	goutils.InitZeroLog()
	goutils.ExecOpt.DumpOutput = true

	cli.CfgLoad()

	codeServerConfigPath := "/root/.config/code-server/config.yaml"
	if cli.Cli.CodeServerPassword != "" {
		codeServerConfig += fmt.Sprintf("\nauth: password\npassword: %s", cli.Cli.CodeServerPassword)
	} else {
		codeServerConfig += "\nauth: none"
	}
	err := goutils.WriteText(codeServerConfigPath, codeServerConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write code-server config file")
	}

	go func() {
		common.ExecWithLogs([]string{"/usr/sbin/code-server"}, "code-server")
	}()

	go func() {
		common.ExecWithLogs([]string{"litessh"}, "litessh-mount")
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
		if err != nil && err.Error() != "exit status 127" {
			log.Error().Err(err).Msg("Failed to run fish shell")
		}
	} else {
		goutils.Exec("tail -f /dev/null")
	}
}
