package deploy

import (
	"github.com/sirupsen/logrus"
	"github.com/dds-sysu/autopilot/codegen/model"
	"github.com/dds-sysu/autopilot/codegen/util"
	"github.com/dds-sysu/autopilot/pkg/defaults"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func ConfigMap(data *model.ProjectData) runtime.Object {
	return configMap(data)
}

func configMap(data *model.ProjectData) *v1.ConfigMap {
	yam, err := util.MarshalYaml(&data.AutopilotOperator)
	if err != nil {
		logrus.Fatalf("failed to marshal autopilot operator config")
	}
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: data.OperatorName,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.SchemeGroupVersion.String(),
			Kind:       "ConfigMap",
		},
		Data: map[string]string{
			defaults.OperatorFile: string(yam),
		},
	}
}
