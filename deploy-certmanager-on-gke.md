Deploy cert-manager on Google Kubernetes Engine (GKE) and create SSL certificates for Ingress using Let's Encrypt


1. To quickly install gcloud 

gcloud components install kubectl

2. Configure gCloud with Google cloud project

gcloud init 

3. export PROJECT=your-project  # Your Google Cloud project ID.
export REGION=europe-west1   # Your Google Cloud region.


We have completed with the initialisation part 

Part 1. 

Create a Kubernetes Cluster 

export CLUSTER=test-cluster-1
gcloud container clusters create $CLUSTER --preemptible --num-nodes=1
gcloud components install gke-gcloud-auth-plugin
export USE_GKE_GCLOUD_AUTH_PLUGIN=True
gcloud container clusters get-credentials $CLUSTER
kubectl get nodes -o wide


Part 2. 

Deploy a simple webserver 

kubectl create deployment web --image=gcr.io/google-samples/hello-app:1.0
kubectl expose deployment web --port=8080


Part 3. 

Create a Static external ip address

gcloud compute addresses create web-ip --global
gcloud compute addresses list

Finally, we will save the IP address into an environment variable for later use. Display the IP address with the following command:


gcloud compute addresses describe web-ip --format='value(address)' --global
export IP_ADDRESS=198.51.100.1  # Replace with your IP address

Part 4. 

Create a domain name for your web server -- Using namecheap or aws-certificate-manager. 

Part 5. 
```
# ingress.yaml


apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: web-ingress
  annotations:
    # This tells Google Cloud to create an External Load Balancer to realize this Ingress
    kubernetes.io/ingress.class: gce
    # This enables HTTP connections from Internet clients
    kubernetes.io/ingress.allow-http: "true"
    # This tells Google Cloud to associate the External Load Balancer with the static IP which we created earlier
    kubernetes.io/ingress.global-static-ip-name: web-ip
spec:
  defaultBackend:
    service:
      name: web
      port:
        number: 8080

```


