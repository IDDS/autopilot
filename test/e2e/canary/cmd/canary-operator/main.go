package main

import (
	"os"

	v1 "github.com/dds-sysu/autopilot/test/e2e/canary/pkg/apis/canarydeployments/v1"
	scheduler "github.com/dds-sysu/autopilot/test/e2e/canary/pkg/scheduler"

	"github.com/dds-sysu/autopilot/pkg/run"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// main is the generated entrypoint to the Autopilot Run function which runs the
// operator
func main() {
	run.RegisterAddToScheme(v1.AddToScheme)
	if err := run.Run(scheduler.AddToManager); err != nil {
		logf.Log.Error(err, "fatal error")
		os.Exit(1)
	}
}
