package testingutil

import (
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

type FakeK8sClient struct {
	DevWorkspaceResources map[string]v1alpha2.DevWorkspaceTemplate
	Errors                map[string]string
}
