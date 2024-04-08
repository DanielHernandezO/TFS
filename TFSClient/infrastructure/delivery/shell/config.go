package shell

import (
	"github.com/TSF/TFSClient/infrastructure/config"
	"github.com/TSF/TFSClient/infrastructure/delivery"
)

type shellCommands struct{}

func NewShellCommands() delivery.Strategy {
	return &shellCommands{}
}

func (s *shellCommands) Start() {
	config.LoadConfigs()
	dependencies := buildDependencies()
	mapActions(dependencies)
}
