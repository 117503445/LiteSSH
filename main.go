package main

import (
	"github.com/117503445/goutils"
	"github.com/117503445/vscode-lite-ssh/pkg/cli"
)

func main() {
	goutils.InitZeroLog()
	cli.CfgLoad()
}
