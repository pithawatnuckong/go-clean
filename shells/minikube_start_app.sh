#!/bin/bash

MINIKUBE_CN="minikube"
NAMESPACE="go-app"

if [ ! "$(docker ps -aq -f name=$MINIKUBE_CN)" ]; then
  echo "minikube container not found"
  if ! minikube start; then echo "minikube not found"; fi
else
  echo "minikube is already running."
fi

# Access minikube's container registry
eval $(minikube docker-env)

docker images

# Delete image
if [ "$(docker images -q -f 'reference=go-app')" ]; then docker rmi $(docker images -q -f 'reference=go-app'); fi

# Create container inside minikube registry
docker build ../. -t go-app:v0.0.1

docker images

# Apply configs to kube
cd ../ && make apply && cd shells || exit

kubectl -n $NAMESPACE get pod
