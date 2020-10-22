package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/dds-sysu/autopilot/cli/pkg/commands"
)

func main() {
	root := commands.AutopilotCli()

	if err := root.Execute(); err != nil {
		log.Fatalf("fatal err: %v", err)
	}
}
