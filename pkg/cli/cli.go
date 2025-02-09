package cli

import (
	"maps"
	"regexp"
	"strings"

	"github.com/117503445/goutils"
	"github.com/alecthomas/kong"

	kongtoml "github.com/alecthomas/kong-toml"
	"github.com/rs/zerolog/log"
)

type SshNode struct {
	Host string `help:"node host"`
	Port int    `help:"node port"`
	User string `help:"node user"`
	Path string `help:"node path to mount"`
	Pri  string `help:"node private key path"`
}

var Cli struct {
	CodeServerPassword string             `help:"code-server password"`
	Nodes              map[string]SshNode `help:"nodes"`
}

func cfgCheck() {
	// if Cli.Nodes == nil {
	// 	log.Fatal().Msg("nodes is empty")
	// }

	// name must only contain [a-zA-Z0-9_]
	isValidName := func(name string) bool {
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, name)
		return matched
	}

	for name, node := range Cli.Nodes {
		logger := log.With().Str("node", name).Logger()
		if name == "" {
			logger.Fatal().Msg("node name is empty")
		} else {
			if !isValidName(name) {
				logger.Fatal().Msg("node name is invalid")
			}
		}

		if node.Host == "" {
			logger.Fatal().Msg("node host is empty")
		}
		if node.Port == 0 {
			logger.Warn().Msg("node port is empty, use default 22")
		}
		if node.User == "" {
			logger.Warn().Msg("node user is empty, use default root")
		}

		if node.Pri == "" {
			logger.Fatal().Msg("node pubkey path is empty")
		} else {
			if !goutils.FileExists(node.Pri) {
				logger.Fatal().Msg("node pubkey path is not exists")
			}
		}
	}
}

func cfgSetDefault() {
	if Cli.Nodes == nil {
		Cli.Nodes = make(map[string]SshNode)
	}

	newNodes := make(map[string]SshNode)
	for name, node := range Cli.Nodes {
		newNodes[name] = node
		if node.Port == 0 || node.User == "" {
			port := 22
			if node.Port != 0 {
				port = node.Port
			}

			user := "root"
			if node.User != "" {
				user = node.User
			}

			path := "~"
			if node.Path != "" {
				path = node.Path
			}
			if path != "/" && strings.HasSuffix(path, "/") {
				path = path[:len(path)-1]
			}

			newNode := SshNode{
				Host: node.Host,
				Port: port,
				User: user,
				Pri:  node.Pri,
				Path: path,
			}
			newNodes[name] = newNode
		}
	}
	maps.Copy(Cli.Nodes, newNodes)

}

func CfgLoad() {
	kong.Parse(&Cli, kong.Configuration(kongtoml.Loader, "/workspace/config.toml"))
	cfgCheck()
	cfgSetDefault()
	log.Info().Interface("cli", Cli).Msg("")
}
