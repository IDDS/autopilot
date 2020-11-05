package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/dds-sysu/autopilot/cli/pkg/commands/build"
	"github.com/dds-sysu/autopilot/cli/pkg/commands/deploy"
	"github.com/dds-sysu/autopilot/cli/pkg/commands/generate"
	"github.com/dds-sysu/autopilot/cli/pkg/commands/initialize"
	"github.com/dds-sysu/autopilot/cli/pkg/commands/logs"
	"github.com/dds-sysu/autopilot/codegen/util"
	"github.com/dds-sysu/autopilot/pkg/version"
	"github.com/spf13/cobra"
)

var verbose bool

func AutopilotCli() *cobra.Command {
	root := &cobra.Command{
		Use:     "autopilot",
		Aliases: []string{"ap"},
		Short:   "An SDK for building Service Mesh Operators with ease",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				if err := util.SetGoVerbose(); err != nil {
					log.Fatalf("Could not set GOFLAGS: (%v)", err)
				}
				log.SetLevel(log.DebugLevel)
				log.Debug("Debug logging is set")
			}
		},
		Version: version.Version,
	}
	root.AddCommand(initialize.NewCmd())
	root.AddCommand(generate.NewCmd())
	root.AddCommand(build.NewCmd())
	root.AddCommand(deploy.NewCmd())
	root.AddCommand(logs.NewCmd())

	root.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")

	return root
}
