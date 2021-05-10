package argoapp

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Unstructured example
// https://github.com/kubernetes/client-go/blob/master/examples/dynamic-create-update-delete-deployment/main.go

const (
	argoNamespace       = "argocd"
	argoAPIVersion      = "argoproj.io/v1alpha1"
	argoApplicationKind = "Application"

	argoProjectName = "draughtsman2"

	configRepoURL = "https://github.com/giantswarm/config.git"
)

type ApplicationConfig struct {
	// Name of the Argo CD Application CR to be created in the argocd
	// namespace.
	Name string

	// AppName as defined in the App Catalog.
	AppName string
	// AppVersion as defined in the App Catalog.
	AppVersion string
	// AppCatalog name.
	AppCatalog string
	// AppConfigVersion is the valid git ref of giantswarm/config
	// repository. Usually the desired value is the major tag, e.g.: v1,
	// v2, etc.
	AppConfigVersion string
	// AppDestinationNamespace is the namespace where the application's
	// manifests are created.
	AppDestinationNamespace string
	// DisableForceUpgrade sets appropriate annotation to prevent helm
	// force upgrades.
	DisableForceUpgrade bool
}

func NewApplication(config ApplicationConfig) (*unstructured.Unstructured, error) {
	// See the argo-cd source for detailed object structure:
	// https://github.com/argoproj/argo-cd/blob/master/pkg/apis/application/v1alpha1/types.go
	obj := map[string]interface{}{
		"apiVersion": argoAPIVersion,
		"kind":       argoApplicationKind,
		"metadata": map[string]interface{}{
			"name":      config.Name,
			"namespace": argoNamespace,
		},
		"spec": map[string]interface{}{
			"project": argoProjectName,
			"source": map[string]interface{}{
				"repoURL":        configRepoURL,
				"targetRevision": config.AppConfigVersion,
				"path":           ".",
				"plugin": map[string]interface{}{
					"name": "konfigure",
					"env": []map[string]interface{}{
						{
							"name":  "KONFIGURE_APP_NAME",
							"value": config.AppName,
						},
						{
							"name":  "KONFIGURE_APP_VERSION",
							"value": config.AppVersion,
						},
						{
							"name":  "KONFIGURE_APP_CATALOG",
							"value": config.AppCatalog,
						},
					},
				},
			},
			"destination": map[string]interface{}{
				"namespace": config.AppDestinationNamespace,
				"server":    "https://kubernetes.default.svc",
			},
			"syncPolicy": map[string]interface{}{
				"automated": map[string]interface{}{
					"prune": true,
				},
			},
		},
	}

	return &unstructured.Unstructured{Object: obj}, nil
}
