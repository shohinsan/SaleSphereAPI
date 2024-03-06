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
SERVICE_NAME    := sales-api
DRIVER 			:= app/services/sales-api/main.go
PRETTYLOG 		:= app/tooling/logfmt/main.go	


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
