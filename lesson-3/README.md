# Lesson 3 - Pod

In this lesson we will create the first pod.

Lot of information about pod deployment can be found [here](https://www.mirantis.com/blog/introduction-to-yaml-creating-a-kubernetes-deployment/)

## Creating Pod

### Loading the image
In order for image to be deployed on the Kind cluster, it must be loaded in the Kind nodes. Otherwise pod deployment will fail, as image will not be found.
This process is described [here](https://kind.sigs.k8s.io/docs/user/quick-start/#loading-an-image-into-your-cluster).

```bash
kind load docker-image k8s-tutorial/lesson-1 --name dd-cluster
```

Check if loaded:
```
docker exec -it dd-cluster-control-plane crictl images
```

> N.B. It should be noted that if image is tagget with `:latest`, then Kubernetes will always try to pull the image from the remote Docker registry.
> This is why we are forcing `imagePullPolicy: Never` in [`pod.yaml`](pod.yaml).
> Other way (and probably correct one) is never using `:latest`, i.e. always using versioned images. Then k8s will look for an image locally,
> and try to pull it only if it is not present (more about this can be found [here](https://kubernetes.io/docs/concepts/configuration/overview/#container-images)).

### Creating Pod In The Cluster

```bash
kubectl create -f pod.yaml
> pod/pod-l3 created

kubectl get pods
> NAME     READY   STATUS         RESTARTS   AGE
> pod-l3   0/1     ErrImagePull   0          7s
```

To see more informations about the pod:
```bash
kubectl describe pod pod-l3
```



