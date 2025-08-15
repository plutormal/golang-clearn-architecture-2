#!/bin/bash

# Build Docker image script
set -e

APP_NAME="golang-crud"
VERSION=${1:-latest}
REGISTRY=${2:-localhost:5000}

echo "Building Docker image for $APP_NAME:$VERSION"

# Build the image
docker build -t $APP_NAME:$VERSION .

# Tag for registry
docker tag $APP_NAME:$VERSION $REGISTRY/$APP_NAME:$VERSION

echo "Docker image built successfully: $REGISTRY/$APP_NAME:$VERSION"

# Push to registry if registry is provided and not localhost
if [[ "$REGISTRY" != "localhost:5000" ]]; then
    echo "Pushing to registry..."
    docker push $REGISTRY/$APP_NAME:$VERSION
    echo "Image pushed successfully!"
fi
