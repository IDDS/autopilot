package commands

import (
	log "github.com/sirupsen/logrus"
	"gitlab.dds-sysu.tech/691729768/autopilot/cli/pkg/commands/build"
	"gitlab.dds-sysu.tech/691729768/autopilot/cli/pkg/commands/deploy"
	"gitlab.dds-sysu.tech/691729768/autopilot/cli/pkg/commands/generate"
	"gitlab.dds-sysu.tech/691729768/autopilot/cli/pkg/commands/initialize"
	"gitlab.dds-sysu.tech/691729768/autopilot/cli/pkg/commands/logs"
	"gitlab.dds-sysu.tech/691729768/autopilot/codegen/util"
	"gitlab.dds-sysu.tech/691729768/autopilot/pkg/version"
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
