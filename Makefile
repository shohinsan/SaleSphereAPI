# ==============================================================================
# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.22.1
KIND            := kindest/node:v1.29.2
KIND_CLUSTER    := sale-sphere-cluster
POSTGRES        := postgres:16.2
BASE_IMAGE_NAME := shohinsan/salesphereapi
SERVICE_NAME    := sales-api
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)
DRIVER 			:= app/services/sales-api/main.go
PRETTYLOG 		:= app/tooling/logfmt/main.go	

# example:  VERSION  	 := "0.0.1-$(shell git rev-parse --short HEAD)" tied to repository


# ==============================================================================
# Define targets

run: 
	go run $(DRIVER) | go run $(PRETTYLOG)

tidy-up:
	go mod tidy

# ==============================================================================
# Running from within k8s/kind

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

# ==============================================================================
# Building containers 

all: service

service:
	docker build \
		-f zarf/docker/Dockerfile.service \
		-t $(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# metrics:
# 	docker build \
# 		-f zarf/docker/dockerfile.metrics \
# 		-t $(METRICS_IMAGE) \
# 		--build-arg BUILD_REF=$(VERSION) \
# 		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
# 		.

