# Lesson 3 - Pod

In this lesson we will create the first pod.

## Creating Pod

```bash
kubectl create -f pod.yaml
> pod/pod-l3 created

kubectl get pods
> NAME     READY   STATUS         RESTARTS   AGE
> pod-l3   0/1     ErrImagePull   0          7s
```


