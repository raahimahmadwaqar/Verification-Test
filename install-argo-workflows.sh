#!/bin/bash

# Check if version is supplied
if [ -z "$1" ]; then
  echo "Please provide the Argo Workflows version (e.g., v3.5.6)"
  exit 1
fi

ARGO_WORKFLOWS_VERSION=$1

# Create namespace for Argo Workflows
kubectl create namespace argo-workflows

# Apply the quick-start manifest
kubectl apply -n argo-workflows -f "https://github.com/argoproj/argo-workflows/releases/download/${ARGO_WORKFLOWS_VERSION}/quick-start-minimal.yaml"

echo "Argo Workflows version ${ARGO_WORKFLOWS_VERSION} installed successfully in namespace argo-workflows"