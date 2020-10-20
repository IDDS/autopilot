package metrics

import (
	"fmt"
	"os"

	v1 "gitlab.dds-sysu.tech/691729768/autopilot/api/v1"
)

func GetMetricsServerAddr(meshProvider v1.MeshProvider, controlPlaneNs string) string {
	if metricsServer := os.Getenv("METRICS_SERVER"); metricsServer != "" {
		return metricsServer
	}
	switch meshProvider {
	case v1.MeshProvider_Istio:
		return fmt.Sprintf("http://prometheus.%v:9090", controlPlaneNs)
	}
	panic("currently unsupported: " + meshProvider.String())
}
