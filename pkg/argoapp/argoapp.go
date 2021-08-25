package argoapp

import (
	"github.com/giantswarm/microerror"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Unstructured example
// https://github.com/kubernetes/client-go/blob/master/examples/dynamic-create-update-delete-deployment/main.go

const (
	argoNamespace       = "argocd"
	argoAPIVersion      = "argoproj.io/v1alpha1"
	argoApplicationKind = "Application"

	argoProjectName = "collections"

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
	// AppDestinationNamespace is the namespace where the application's
	// manifests are created.
	AppDestinationNamespace string

	// ConfigRef is the valid git ref of giantswarm/config repository used
	// to configure the application. Usually the desired value is the major
	// tag, e.g.: v1, v2, etc.
	ConfigRef string
	// DisableForceUpgrade sets appropriate annotation to prevent helm
	// force upgrades.
	DisableForceUpgrade bool
}

func NewApplication(config ApplicationConfig) (*unstructured.Unstructured, error) {
	if config.Name == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Name must not be empty", config)
	}
	if config.AppName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.AppName must not be empty", config)
	}
	if config.AppVersion == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.AppVersion must not be empty", config)
	}
	if config.AppCatalog == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.AppCatalog must not be empty", config)
	}
	if config.AppDestinationNamespace == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.AppDestinationNamespace must not be empty", config)
	}
	if config.ConfigRef == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ConfigRef must not be empty", config)
	}

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
				"targetRevision": config.ConfigRef,
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
					// If set to true allows deleting all application resources during automatic syncing (false by default).
					"allowEmpty": false,
					"selfHeal":   true,
				},
			},
		},
	}

	return &unstructured.Unstructured{Object: obj}, nil
}
