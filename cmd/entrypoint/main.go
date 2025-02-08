package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "embed"

	"github.com/117503445/goutils"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

//go:embed code-server-config-template.yaml
var codeServerConfigTemplate string

func main() {
	goutils.InitZeroLog()
	goutils.ExecOpt.DumpOutput = true

	codeServerConfigPath := "/root/.config/code-server/config.yaml"
	codeServerPassword := os.Getenv("CODE_SERVER_PASSWORD")
	if codeServerPassword == "" {
		log.Warn().Msg("CODE_SERVER_PASSWORD is not set, use default password")
		codeServerPassword = "123456"
	}
	codeServerConfigText := fmt.Sprintf(codeServerConfigTemplate, codeServerPassword)
	err := goutils.WriteText(codeServerConfigPath, codeServerConfigText)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write code-server config file")
	}

	go func() {
		file, err := os.OpenFile("/docker-dev/logs/goreman.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		cmd := exec.Command("goreman", "start")
		// cmd.Stdin = os.Stdin
		cmd.Stdout = file
		cmd.Stderr = file
		cmd.Dir = "/docker-dev"
		log.Info().Str("log", "/docker-dev/logs/goreman.log").Msg("Starting goreman")
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("Failed to run goreman")
		}
	}()

	fileCustomEntrypoint := "/entrypoint"
	if goutils.PathExists(fileCustomEntrypoint) {
		goutils.Exec(fileCustomEntrypoint)
	}

	var isTTY bool
	if isatty.IsTerminal(os.Stdout.Fd()) {
		isTTY = true
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		log.Fatal().Msg("Cygwin terminal is not supported")
	} else {
		isTTY = false
	}

	// log.Debug().Bool("isTTY", isTTY).Msg("Check if TTY")

	if isTTY {
		cmd := exec.Command("/bin/fish")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		// log.Debug().Msg("Enter fish shell")
		err := cmd.Run()
		if err != nil {
			log.Error().Err(err).Msg("Failed to run fish shell")
		}
		// log.Debug().Msg("Exit fish shell")
	} else {
		goutils.Exec("tail -f /dev/null")
	}
}
