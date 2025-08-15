# Kubernetes Deployment Guide

This guide explains how to deploy the Go Clean Architecture application to Kubernetes using various Infrastructure as Code (IaC) approaches.

## Prerequisites

- Docker
- Kubernetes cluster (Minikube for local development)
- kubectl
- Helm (optional)
- make

## Quick Start

### 1. Setup Local Kubernetes Cluster (Minikube)

```bash
# Setup Minikube
make minikube-setup

# Configure Docker to use Minikube's daemon
eval $(minikube docker-env)
```

### 2. Build and Deploy

```bash
# Complete deployment pipeline (test, build, deploy)
make deploy-pipeline
```

### 3. Access the Application

```bash
# Port forward to access locally
make k8s-port-forward

# Or get minikube IP and use ingress
minikube ip
# Add to /etc/hosts: <minikube-ip> golang-crud.local
# Access: http://golang-crud.local
```

## Detailed Deployment Options

### Option 1: Using Kustomize (Recommended)

#### Development Environment
```bash
# Deploy to development
kubectl apply -k k8s/overlays/development

# Check status
kubectl get all -n golang-crud
```

#### Production Environment
```bash
# Deploy to production
kubectl apply -k k8s/overlays/production

# Check status
kubectl get all -n golang-crud
```

### Option 2: Using Helm Charts

```bash
# Install using Helm
make helm-install

# Upgrade
helm upgrade golang-crud ./helm/golang-crud --namespace golang-crud

# Uninstall
make helm-uninstall
```

### Option 3: Using Scripts

```bash
# Build Docker image
./scripts/build-docker.sh v1.0.0

# Deploy to Kubernetes
./scripts/deploy-k8s.sh development
```

## Infrastructure Components

### Kubernetes Resources

1. **Namespace**: `golang-crud`
2. **Deployment**: Application pods with rolling update strategy
3. **Service**: ClusterIP service for internal communication
4. **Ingress**: External access routing
5. **ConfigMap**: Application configuration
6. **HPA**: Horizontal Pod Autoscaler for scaling
7. **SecurityContext**: Security constraints

### Resource Limits

| Environment | Replicas | CPU Request | CPU Limit | Memory Request | Memory Limit |
|-------------|----------|-------------|-----------|----------------|--------------|
| Development | 1        | 25m         | 50m       | 32Mi           | 64Mi         |
| Production  | 5        | 100m        | 200m      | 128Mi          | 256Mi        |

### Auto Scaling

- **Min Replicas**: 2 (production), 1 (development)
- **Max Replicas**: 10
- **CPU Threshold**: 70%
- **Memory Threshold**: 80%

## Environment Management

### Configuration

Different environments are managed using Kustomize overlays:

- **Base**: Common configuration in `k8s/`
- **Development**: Overrides in `k8s/overlays/development/`
- **Production**: Overrides in `k8s/overlays/production/`

### Environment Variables

| Variable  | Development | Production |
|-----------|-------------|------------|
| APP_ENV   | development | production |
| LOG_LEVEL | debug       | warn       |
| PORT      | 8080        | 8080       |

## Health Checks

### Readiness Probe
- **Path**: `/health`
- **Initial Delay**: 10s
- **Period**: 5s

### Liveness Probe
- **Path**: `/health`
- **Initial Delay**: 30s
- **Period**: 10s

## Security

### Pod Security Context
- **Run as non-root user**: ✅
- **Read-only root filesystem**: ✅
- **Drop all capabilities**: ✅
- **No privilege escalation**: ✅

### Network Policies
```bash
# Apply network policies (optional)
kubectl apply -f k8s/security/network-policy.yaml
```

## Monitoring and Logging

### View Logs
```bash
# Application logs
make k8s-logs

# Specific pod logs
kubectl logs <pod-name> -n golang-crud

# Follow logs
kubectl logs -f deployment/golang-crud-deployment -n golang-crud
```

### Metrics
```bash
# Pod metrics (requires metrics-server)
kubectl top pods -n golang-crud

# Node metrics
kubectl top nodes
```

## Troubleshooting

### Common Issues

1. **Image Pull Errors**
   ```bash
   # Check if using Minikube Docker daemon
   eval $(minikube docker-env)
   docker images | grep golang-crud
   ```

2. **Service Not Accessible**
   ```bash
   # Check service endpoints
   kubectl get endpoints -n golang-crud
   
   # Port forward for testing
   kubectl port-forward service/golang-crud-service 8080:80 -n golang-crud
   ```

3. **Pod Not Starting**
   ```bash
   # Describe pod for details
   kubectl describe pod <pod-name> -n golang-crud
   
   # Check events
   kubectl get events -n golang-crud --sort-by=.metadata.creationTimestamp
   ```

### Debug Commands

```bash
# Check cluster status
kubectl cluster-info

# Check node status
kubectl get nodes

# Check all resources in namespace
kubectl get all -n golang-crud

# Describe deployment
kubectl describe deployment golang-crud-deployment -n golang-crud

# Check ingress
kubectl get ingress -n golang-crud
kubectl describe ingress golang-crud-ingress -n golang-crud
```

## CI/CD Integration

### GitLab CI Example
```yaml
stages:
  - test
  - build
  - deploy

test:
  stage: test
  script:
    - make test

build:
  stage: build
  script:
    - make docker-build

deploy:
  stage: deploy
  script:
    - make k8s-deploy-dev
  only:
    - main
```

### GitHub Actions Example
```yaml
name: Deploy to Kubernetes

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Deploy
      run: |
        make test
        make docker-build
        make k8s-deploy-dev
```

## Production Checklist

- [ ] Security scanning of Docker images
- [ ] Resource limits configured
- [ ] Health checks implemented
- [ ] Monitoring and alerting setup
- [ ] Backup and recovery plan
- [ ] SSL/TLS certificates configured
- [ ] Network policies applied
- [ ] Secrets management
- [ ] Rolling update strategy tested
- [ ] Disaster recovery plan

## Useful Commands

```bash
# Build everything
make build

# Run tests
make test

# Complete deployment
make deploy-pipeline

# Check status
make k8s-status

# View logs
make k8s-logs

# Port forward
make k8s-port-forward

# Clean up
make k8s-delete
make clean
```
