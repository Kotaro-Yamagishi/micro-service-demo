#!/bin/bash

# Docker Hub username
DOCKER_HUB_USERNAME=DOCKER_HUB_USERNAME

# Build and push book_service
cd book_service
docker build -t $DOCKER_HUB_USERNAME/book_service .
docker push $DOCKER_HUB_USERNAME/book_service
kubectl apply -f book_service.yaml
cd ..

# Build and push user_service
cd user_service
docker build -t $DOCKER_HUB_USERNAME/user_service .
docker push $DOCKER_HUB_USERNAME/user_service
kubectl apply -f user_service.yaml
cd ..

# Build and push order_service
cd order_service
docker build -t $DOCKER_HUB_USERNAME/order_service .
docker push $DOCKER_HUB_USERNAME/order_service
kubectl apply -f order_service.yaml
cd ..

# Apply ingress
kubectl apply -f ingress.yaml

# Confirm deployments
kubectl get pods
kubectl get services
kubectl get ingress
