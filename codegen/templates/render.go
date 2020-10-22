package templates

import (
	"github.com/dds-sysu/autopilot/codegen/model"
	"k8s.io/apimachinery/pkg/runtime"
)

type TemplateFunc func(data *model.ProjectData) runtime.Object
