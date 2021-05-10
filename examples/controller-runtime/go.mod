module github.com/giantswarm/argoapp/examples/controller-runtime

go 1.16

require (
	github.com/giantswarm/argoapp v0.0.0
	k8s.io/client-go v0.21.0
	sigs.k8s.io/controller-runtime v0.8.3
)

replace github.com/giantswarm/argoapp => ../../
