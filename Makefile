# Makefile for Go Clean Architecture with Kubernetes

.PHONY: help build test docker-build docker-push k8s-setup k8s-deploy k8s-delete clean

APP_NAME := golang-crud
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "latest")
REGISTRY := localhost:5000
ENVIRONMENT := development

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the Go application
	@echo "Building Go application..."
	go build -o bin/$(APP_NAME) .

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@./scripts/build-docker.sh $(VERSION) $(REGISTRY)

docker-push: docker-build ## Build and push Docker image
	@echo "Docker image built and pushed successfully"

minikube-setup: ## Setup Minikube cluster
	@echo "Setting up Minikube..."
	@./scripts/setup-minikube.sh

minikube-docker: ## Configure shell to use minikube's Docker daemon
	@echo "Configuring Docker environment for Minikube..."
	eval $(minikube docker-env)

k8s-deploy-dev: ## Deploy to development environment
	@echo "Deploying to development environment..."
	@./scripts/deploy-k8s.sh development

k8s-deploy-prod: ## Deploy to production environment
	@echo "Deploying to production environment..."
	@./scripts/deploy-k8s.sh production

k8s-delete: ## Delete Kubernetes resources
	@echo "Deleting Kubernetes resources..."
	kubectl delete -k k8s/overlays/$(ENVIRONMENT) || true

k8s-logs: ## View pod logs
	@echo "Viewing pod logs..."
	kubectl logs -l app=golang-crud -n golang-crud --tail=100 -f

k8s-status: ## Check deployment status
	@echo "Checking deployment status..."
	kubectl get all -n golang-crud

k8s-port-forward: ## Port forward to service
	@echo "Port forwarding to service..."
	kubectl port-forward service/golang-crud-service 8080:80 -n golang-crud

helm-install: ## Install using Helm
	@echo "Installing with Helm..."
	helm upgrade --install $(APP_NAME) ./helm/$(APP_NAME) \
		--namespace golang-crud \
		--create-namespace \
		--set image.tag=$(VERSION)

helm-uninstall: ## Uninstall Helm release
	@echo "Uninstalling Helm release..."
	helm uninstall $(APP_NAME) -n golang-crud

clean: ## Clean build artifacts
	@echo "Cleaning up..."
	rm -f bin/$(APP_NAME)
	rm -f coverage.out coverage.html
	docker rmi $(APP_NAME):$(VERSION) 2>/dev/null || true

local-dev: build ## Run application locally for development
	@echo "Starting application locally..."
	./bin/$(APP_NAME)

# Complete deployment pipeline
deploy-pipeline: test docker-build k8s-deploy-dev ## Complete deployment pipeline

# Production deployment
deploy-production: test docker-push k8s-deploy-prod ## Deploy to production
