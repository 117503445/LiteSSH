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
	subs := []string{}
	if strings.HasPrefix(rCli.Target, "/") {
		splits := strings.Split(rCli.Target, "/")
		if len(splits) >= 2 {
			name = splits[2]
			if goutils.FileExists(rCli.Target) {
				if len(splits) < 4 {
					log.Fatal().Str("target", rCli.Target).Msg("target is invalid")
				}
				// remove last item (file)
				subs = splits[3 : len(splits)-1]
			} else if goutils.DirExists(rCli.Target) {
				subs = splits[3:]
			} else {
				log.Fatal().Str("target", rCli.Target).Msg("target is invalid")
			}
		} else {
			log.Fatal().Str("target", rCli.Target).Msg("target is invalid")
		}
	} else {
		name = rCli.Target
	}
	subPath := strings.Join(subs, "/")
	log.Debug().Str("name", name).Str("subPath", subPath).Msg("")

	sshNode, ok := cli.Cli.Nodes[name]
	if !ok {
		log.Fatal().Str("name", name).Msg("node not found")
	}
	
	p := ""
	if strings.HasPrefix(sshNode.Path, "/") {
		p = sshNode.Path + "/" + subPath
	} else if sshNode.Path == "~" {
		p = "~/" + subPath
	} else {
		p = "~/" + sshNode.Path + "/" + subPath
	}

	cmds := []string{"ssh", "-i", sshNode.Pri, "-p", fmt.Sprintf("%d", sshNode.Port), "-o", "StrictHostKeyChecking=no", fmt.Sprintf("%s@%s", sshNode.User, sshNode.Host), "-t", fmt.Sprintf("cd %s && $SHELL -l", p)}
	command := exec.Command(cmds[0], cmds[1:]...)
	log.Debug().Str("command", command.String()).Msg("")
	command.Stdout = os.Stdout
	command.Stderr = os.Stdout
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("ssh failed")
	}
}
