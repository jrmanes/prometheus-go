#!/bin/bash

echo "Building docker resources..."
make docker_all

echo "Create Kubernetes resources..."
kubectl apply -f ./infra/
