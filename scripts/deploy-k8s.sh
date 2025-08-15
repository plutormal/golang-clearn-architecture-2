#!/bin/bash

# Kubernetes deployment script
set -e

ENVIRONMENT=${1:-development}
NAMESPACE="golang-crud"

echo "Deploying to $ENVIRONMENT environment..."

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "kubectl is not installed or not in PATH"
    exit 1
fi

# Check if kustomize is available
if ! command -v kustomize &> /dev/null; then
    echo "Installing kustomize..."
    kubectl kustomize --help > /dev/null || {
        echo "kustomize is not available"
        exit 1
    }
fi

# Create namespace if it doesn't exist
kubectl create namespace $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

# Deploy using kustomize
echo "Applying Kubernetes manifests for $ENVIRONMENT..."
kubectl apply -k k8s/overlays/$ENVIRONMENT

# Wait for deployment to be ready
echo "Waiting for deployment to be ready..."
kubectl rollout status deployment/$([ "$ENVIRONMENT" = "development" ] && echo "dev-" || echo "prod-")golang-crud-deployment -n $NAMESPACE

# Get service info
echo "Deployment completed successfully!"
echo "Service information:"
kubectl get services -n $NAMESPACE
echo ""
echo "Pod information:"
kubectl get pods -n $NAMESPACE
echo ""
echo "Ingress information:"
kubectl get ingress -n $NAMESPACE
