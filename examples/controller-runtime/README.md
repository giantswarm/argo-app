# controller-runtime example

This example demonstrates how to use this library using controller-runtime
client.

1. Setup a kind cluster:

```
kind create cluster
```

2. Install Argo CD:

```
kubectl --context kind-kind create namespace argocd
kubectl --context kind-kind apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

3. Run the example:

```
kubectl config use-context kind-kind && go run main.go --kubeconfig=$HOME/.kube/config
```

4. Verify the CR is created:

```
$ kubectl --context kind-kind get application -n argocd my-argo-app
NAME          SYNC STATUS   HEALTH STATUS
my-argo-app   Unknown       Unknown
```
