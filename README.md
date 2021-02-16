## K8sResourceCloner

K8sResourceCloner is the kubernetes operator which will clone the given k8s resources from one namespace to another

### How to run
There are diffrent ways to run the application 
1) Run locally outside the cluster
```sh
make install run
```
2) Run as a Deployment inside the cluster
```sh
export USERNAME=<quay-namespace>
make docker-build docker-push IMG=quay.io/$USERNAME/K8sResourceCloner-operator:v0.0.1
make deploy IMG=quay.io/$USERNAME/K8sResourceCloner-operator:v0.0.1
```

Create a K8sResourceCloner CR
```sh
kubectl apply -f config/samples/cache_v1alpha1_k8sresourcecloner.yaml
```

### Installation
Update :  
**fromNamespace**: from which the resources need to be cloned  
**toNamespaces**: in which namespace it need to be cloned, mutiple namespaces are allowed  
**resourceNames**: the resource name which need to be copied, mutiple resource name are allowed    
 
```
apiVersion: cache.example.org/v1alpha1
kind: K8sResourceCloner
metadata:
  name: k8sresourcecloner-sample
spec:
  fromNamespace : default
  toNamespaces: 
  - test
  - demo
  resourceNames:  
  - mysecret
  - test-configMap
   ```

After deploying the above Custom resource, it will create the give resouces in specified namespace

# Resources:
https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/
