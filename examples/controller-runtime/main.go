package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/argoapp/pkg/argoapp"
)

func main() {
	var kubeconfig string
	if home, err := os.UserHomeDir(); err == nil {
		flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	ctx := context.Background()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	client, err := client.New(config, client.Options{})
	if err != nil {
		panic(err)
	}

	obj, err := argoapp.NewApplication(argoapp.ApplicationConfig{
		Name: "my-argo-app",

		AppName:                 "dex",
		AppVersion:              "1.2.3",
		AppCatalog:              "control-plane-catalog",
		AppDestinationNamespace: "my-namespace",

		ConfigRef:           "v1",
		DisableForceUpgrade: false,
	})
	if err != nil {
		panic(err)
	}

	err = client.Create(ctx, obj)
	if err != nil {
		panic(err)
	}

}
