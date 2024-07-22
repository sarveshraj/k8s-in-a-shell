# K8s in a shell

### Begin with

1. Read about the origins of K8s
2. Understand declarative vs imperative programming
3. Understand nodes, clusters and manifests in K8s
4. Install `kubectl`, `minikube` and `k9s`
5. Start `minikube` with
```bash
minikube start
```
6. Run `k9s` and get familiar navigating the minikube cluster with it
```bash
k9s
```
7. Try to `shell` into a pod, switch namespaces, deleting a pod etc

### Namespace

1. A namespace in K8s is an organizational construct
2. The manifest of our namespace is in `namespace/namespace.yaml`
3. Apply the namespace by running
```bash
kubectl apply -f namespace/namespace.yaml
```

### Pod

1. Navigate to `deployment/main.go`
2. This is the app we'll be deploying to K8s - we'll call it server
3. Go through the pod manifest `deployment/pod.yaml`
4. Understand that a pod is a group of containers and is assumed to be stateless
5. Apply the pod by running
```bash
kubectl apply -f deployment/pod.yaml
```
6. Kill the pod

### Replica set

1. Go through the replica set manifest `deployment/replicaset.yaml`
2. Note and understand `replicas` and `template` in the manifest
3. A replica set will always try to maintain the number of pods specified to it
4. All pods in an replica set are considered interchangeable
5. Apply the replica set by running
```bash
kubectl apply -f deployment/replicaset.yaml
```
6. Delete a pod in the replica set and see K8s create another one
7. Update the image in `deployment/replicaset.yaml` to a non-existent one and re-apply the replica set
8. Notice how the pods were immediately updated and are now in a failing state
9. Delete the replica set

### Deployment

1. Go through the deployment manifest `deployment/deployment.yaml`
2. Notice how similar it is to the replica set's manifest
3. Apply the deployment by running
```bash
kubectl apply -f deployment/deployment.yaml
```
4. Verify that applying a deployment created a replica set internally
5. Update the image in `deployment/deployment.yaml` to a non-existent one and re-apply the deployment
6. Notice how the deployment did not terminate older pods till the new one is healthy
7. Rollback the faulty deployment with
```bash
kubectl rollout undo deployment/server-deployment -n k8s-in-a-shell
```
8. Notice how the failing pod gets terminated

### Volume

1. Containers in a pod don't have access to persistent storage that exists beyond the pod's lifecycle by default
2. Containers in a pod don't share storage by default
3. Volumes solve both these problems
4. We will deploy redis to understand volumes
5. Go through the persistent volume claim manifest `volume/volume.yaml`
6. Understand that in the claim we are only requesting for storage of the specified configuration
7. The storage itself is dynamically allocated
8. Go through redis's deployment manifest `volume/deployment.yaml`
9. Notice and understand the relationship between `volumeMounts`, `volumes` and the volume manifest `volume/volume.yaml`
10. Apply the volume and the deployment
```bash
kubectl apply -f volume/volume.yaml
kubectl apply -f volume/deployment.yaml
```
11. Shell into redis pod and set some data in redis
```bash
redis-cli
set mykey myvalue
```
12. Delete this pod and wait for K8s to create a new pod
13. Shell into the new pod and attempt to get the data
```bash
redis-cli
get mykey
```
14. Verify that the response is `"myvalue"`

### Service

1. A service is a way to expose workloads within and outside a K8s cluster
2. Go through the frontend application code `service/index.js` - specifically its `/ping` and `/` APIs
3. Go through the frontend's service manifest `service/service.yaml`
4. This is of type `LoadBalancer` which is used to expose workloads outside the cluster
5. Apply the service
```bash
kubectl apply -f service/service.yaml
```
6. Expose the service external IP directly to the host operating system (your machine)
```bash
# in a new terminal window
minikube tunnel
```
7. Open `localhost:3000/ping` on your browser - you should see `pong`
8. You can now access the frontend app outside the cluster!
9. Notice how redis is being used in `index.js`
10. This is called a FQDN - read about it
11. Open `volume/service.yaml` and go through the service manifest
12. This service is of type `ClusterIP` which is used to expose workloads within the cluster
13. Apply the service to expose redis
```bash
kubectl apply -f volume/service.yaml
```

### Cron job

1. A cron job as the name suggests is used to run recurring workloads
2. Go through the manifest at `cronjob/cronjob.yaml` - we call it worker
3. Go through the cron job code `cronjob/main.py`

### Stitching it all together

1. Open `localhost:3000` on your browser
2. Try to submit a wage - it should fail - can you guess why?
3. If you guessed that our server is not exposed, you are right!
4. Read then apply the server's service
```bash
kubectl apply -f deployment/service.yaml
```
5. Now try submitting a wage again - it should succeed now
6. Open worker logs and check the 30% of your submitted wage was paid as tax

