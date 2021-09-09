package argoapp

import (
	"log"
	"testing"

	"github.com/giantswarm/microerror"
)

func Test_UnstructuredToArgoApplication(t *testing.T) {
	obj, err := NewUnstructuredApplication(ApplicationConfig{
		Name: "my-argo-app",

		AppName:                 "dex",
		AppVersion:              "1.2.3",
		AppCatalog:              "control-plane-catalog",
		AppDestinationNamespace: "my-namespace",

		ConfigRef:           "v1",
		DisableForceUpgrade: false,
	})
	if err != nil {
		log.Fatalf("Test failed:\n%s", microerror.Mask(err))
	}

	app, err := UnstructuredToArgoApplication(obj)
	if err != nil {
		log.Fatalf("Test failed:\n%s", microerror.Mask(err))
	}

	match := true
	match = match && app.APIVersion == "argoproj.io/v1alpha1"
	match = match && app.Kind == "Application"
	match = match && app.Name == "my-argo-app"
	match = match && app.Namespace == "argocd"
	if !match {
		log.Fatalf("Argo Application does not match unstructured:\n%+v\n\n%+v", obj, app)
	}
}
