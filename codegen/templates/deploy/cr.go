package deploy

import (
	"github.com/solo-io/autopilot/codegen/model"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func CustomResource(data *model.TemplateData) runtime.Object {
	return customResource(data)
}

func customResource(data *model.TemplateData) *unstructured.Unstructured {
	cr := &unstructured.Unstructured{}
	cr.SetAPIVersion(data.ApiVersion)
	cr.SetKind(data.Kind)
	cr.SetName("example")
	return cr
}