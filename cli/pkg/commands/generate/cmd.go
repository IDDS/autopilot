package generate

import (
	"github.com/dds-sysu/autopilot/codegen"
	"github.com/dds-sysu/autopilot/codegen/util"
	"github.com/spf13/cobra"
)

var (
	forceOverwrite bool
	deepcopyOnly   bool
)

func NewCmd() *cobra.Command {
	genCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates code, build, and deployment files for the Operator",
		Long: `The autopilot generate command (re-)generates all assets for coding, building and deploying the Operator.

ap generate should be run whenever the autopilot.yaml, autopilot-operator.yaml, or the spec.go for the top-level CRD have changed.

Re-run this command when you have updated your autopilot.yaml or your api's spec.go.

Note: Kubernetes YAML Manifests generated by ap generate have fields that must be replaced using sed.
The ap deploy command will perform this automatically; however, to apply the manifests by hand:

cat deploy/<resource>.yaml | \
	sed 's/REPLACE_IMAGE/<your-docker-image>/' | \
	sed 's/REPLACE_NAMESPACE/<operator deployment namespace>/' | \
	kubectl apply -f- 

`,
		RunE: genFunc,
	}
	genCmd.PersistentFlags().BoolVarP(&forceOverwrite, "overwrite", "f", false, "force overwriting files that are meant to be modified by the user (spec.go, worker.go, etc.)")
	genCmd.PersistentFlags().BoolVarP(&deepcopyOnly, "deepcopy-only", "d", false, "Only update the zz_generated.deepcopy.go file.")
	return genCmd
}

func genFunc(cmd *cobra.Command, args []string) error {
	util.MustInProjectRoot()

	return codegen.Run(".", forceOverwrite, deepcopyOnly)
}
