package main

import (
	"github.com/117503445/goutils"
	"github.com/117503445/vscode-lite-ssh/pkg/cli"
	"github.com/117503445/vscode-lite-ssh/pkg/mount"
)

func main() {
	goutils.InitZeroLog()
	cli.CfgLoad()
	mountManager := mount.NewMountManager(cli.Cli.Nodes)
	mountManager.Start()
}
