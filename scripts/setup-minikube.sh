#!/bin/bash

# Setup minikube for local development
set -e

echo "Setting up Minikube for local Kubernetes development..."

# Start minikube
echo "Starting Minikube..."
minikube start --driver=docker --cpus=2 --memory=4096

# Enable addons
echo "Enabling Minikube addons..."
minikube addons enable ingress
minikube addons enable metrics-server
minikube addons enable dashboard

# Configure docker environment
echo "Configuring Docker environment..."
eval $(minikube docker-env)

# Verify setup
echo "Verifying Minikube setup..."
kubectl cluster-info
kubectl get nodes

echo ""
echo "Minikube setup completed successfully!"
echo "To use minikube's Docker daemon, run: eval \$(minikube docker-env)"
echo "To access dashboard, run: minikube dashboard"
echo "To get minikube IP, run: minikube ip"
