# argoapp

A convenience library managing Giant Swarm Management Cluster Argo CD Application CRs.

It is used in:

- [architect](https://github.com/giantswarm/architect/)
- [opsctl](https://github.com/giantswarm/opsctl/)
- [release-operator](https://github.com/giantswarm/release-operator/)

## FAQ

#### Why not using upstream Argo CD types?

Argo CD has a mono-repo approach with huge dependency tree (https://github.com/argoproj/argo-cd/). This repository has only `k8s.io/apimachinery` dependency. Smaller dependency footprint makes the library easier to vendor and maintain (thinking of nancy security reports).

#### Why not copying upstream Argo CD types?

Upstream type definitions are intertwined with business logic and protobuf generated code and OpenAPI generated code. Removing all of that off is an effort to repeated every single time something changes upstream. When using undefined that should be way easier.

For the reference upstream types are defined here https://github.com/argoproj/argo-cd/tree/master/pkg/apis/application/v1alpha1.
