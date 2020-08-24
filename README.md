### K8sResourceCloner

This is the kubernetes operator which will clone the given k8s resources in the given namespaces


#### Installation
Update :  
fromNamespace: from which the resources need to be cloned  
toNamespaces: in which namespace it need to be cloned, mutiple namespaces are allowed  
resourceNames: the resource name which need to be copied, mutiple resource name are allowed    
 
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