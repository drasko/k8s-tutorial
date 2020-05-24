# Lesson 2 - Kubectl and Kind
This lesson covers creation of the first Kubernetes Deployment.

Needed tools:
- Kubectl
- Kind

After that first deployment will be explained.

## Kubectl
Kubectl is Kubernetes CLI, used on local host for remote control of k8s clusters.

### Install
If on Debian:

```bash
sudo apt install kubernetes-client
```

Otherwise, follow the instructions [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Kind
Kind is a tool for running local Kubernetes clusters using Docker container “nodes”.
It helps easy setup of k8s cluster based on containers and does not need VM like in a case of Minikube. 

### Install
```bash
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.8.1/kind-$(uname)-amd64
chmod +x ./kind
mv ./kind /some-dir-in-your-PATH/kind
```

Detailed instructions can be found [here] (https://kind.sigs.k8s.io/docs/user/quick-start/)

### Create cluster
```bash
kind create cluster --name dd-cluster
```
