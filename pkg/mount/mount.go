package mount

import (
	"os"
	"os/exec"
	"sync"

	"github.com/117503445/vscode-lite-ssh/pkg/cli"
	"github.com/rs/zerolog/log"
)

const dirRemote = "/remote"

type MountManager struct {
	nodes map[string]cli.SshNode
}

func NewMountManager(nodes map[string]cli.SshNode) *MountManager {
	return &MountManager{nodes: nodes}
}

func (m *MountManager) Start() {
	log.Info().Msg("Mount Manager Start")

	err := os.MkdirAll(dirRemote, 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("create remote dir failed")
	}

	writeRcloneCfg := func() {
		// TODO
	}
	writeRcloneCfg()

	mount := func(name string, node cli.SshNode) {
		logger := log.With().Str("node", name).Logger()
		dirNode := dirRemote + "/" + name
		err := os.MkdirAll(dirNode, 0755)
		if err != nil {
			logger.Fatal().Err(err).Msg("create node dir failed")
		}
		cmds := []string{"rclone", "mount", name + ":", dirNode, "--allow-non-empty", "--allow-other", "--vfs-cache-mode", "full", "-vvv"}
		cmd := exec.Command(cmds[0], cmds[1:]...)
		logger.Info().Str("cmd", cmd.String()).Msg("")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		err = cmd.Run()
		if err != nil {
			logger.Fatal().Err(err).Msg("mount failed")
		}
	}

	var sg sync.WaitGroup
	for name, node := range m.nodes {
		sg.Add(1)
		go func(name string, node cli.SshNode) {
			defer sg.Done()
			mount(name, node)
		}(name, node)
	}
	sg.Wait()

	log.Info().Msg("Mount Manager Stop")

}
