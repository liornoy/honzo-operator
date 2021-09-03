# honzo-operator
This is a simple Kubernetes operator for learning purposes.
(The word honzo has no meaning it's just a name of a game character I once created)

## Honzo CR scheme
honzo.yaml file is a template for the honzo CR.
Note the two following things:
1. Under `metadata` there is `finalizers` which defines the finalizer name used for creating pre-deletion logic.
2. The Spec of this CR are:
* `text`: a string that will be printed when creating a new honzo CR.
* `deleteText`: a string that will be printed when deleteing an honzo CR. 



## Getting Started
1. Clone this repo: 
`git clone https://github.com/liornoy/honzo-operator`
2. Start your local kubernetes cluster (with kind or minikube)
3. Install the CRD into the K8s cluster: 
`make install`
5. Create the honzo CR:
`kubectl apply -f honzo.yaml`
6. Delete the honzo CR:
`kubectl delete Honzo example.honzo`

![image](https://user-images.githubusercontent.com/40122521/132027516-8bd88d9e-b4f2-4c8d-a345-a629ae260479.png)

