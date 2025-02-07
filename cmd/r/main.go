package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/117503445/goutils"
	"github.com/117503445/vscode-lite-ssh/pkg/cli"
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"
)

var rCli struct {
	Target string `arg:"" help:"remote name or path(/remote/**/*)"`
}

func main() {
	goutils.InitZeroLog()
	kong.Parse(&rCli)

	os.Args = []string{os.Args[0]}
	cli.CfgLoad()

	name := ""
	if strings.HasPrefix(rCli.Target, "/") {
		splits := strings.Split(rCli.Target, "/")
		if len(splits) >= 2 {
			name = splits[2]
		} else {
			log.Fatal().Str("target", rCli.Target).Msg("target is invalid")
		}
	} else {
		name = rCli.Target
	}

	sshNode, ok := cli.Cli.Nodes[name]
	if !ok {
		log.Fatal().Str("name", name).Msg("node not found")
	}

	cmd := fmt.Sprintf(`ssh -i %s -p %d -o StrictHostKeyChecking=no %s@%s`, sshNode.Pri, sshNode.Port, sshNode.User, sshNode.Host)
	println(cmd)
	cmds := strings.Split(cmd, " ")
	// log.Debug().Strs("cmds", cmds).Msg("")
	command := exec.Command(cmds[0], cmds[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stdout
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("ssh failed")
	}
}
