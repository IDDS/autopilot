package templates

import (
	"gitlab.dds-sysu.tech/691729768/autopilot/codegen/model"
	"k8s.io/apimachinery/pkg/runtime"
)

type TemplateFunc func(data *model.ProjectData) runtime.Object
