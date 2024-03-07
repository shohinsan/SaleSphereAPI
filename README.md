# Sales Backend Web API  

## Install

* https://kustomize.io/
* https://kubernetes.io/docs/tasks/tools/
* https://kind.sigs.k8s.io/

## Upgrade/Increment to Latest Stable Version 

* Dependencies in `go.mod`
* Makefile

## Steps

Kind creates a cluster named "sale-sphere-cluster" using the specified image (kindest/node:v1.29.2) and configuration file (zarf/k8s/dev/kind-config.yaml).

![Xnapper-2024-03-06-01 13 54](https://github.com/shohinsan/SaleSphereAPI/assets/22685770/7ccf92a7-1e7e-4fdd-9f71-73d8388a6763)

In setting up a Kubernetes cluster, I ensure that the required node image is available, prepare the nodes for the cluster, and configure the cluster accordingly. Once configured, I start the control plane and proceed to install essential components such as the Container Network Interface (CNI) and StorageClass. Upon successful creation of the cluster, I set the kubectl context to "kind-sale-sphere-cluster," enabling seamless interaction with the newly created cluster.

![image](https://github.com/shohinsan/SaleSphereAPI/assets/22685770/48093425-4963-412b-9981-c9920da1bfad)

This log shows the execution of the `make all` command, which typically builds and prepares various components of a project. Here's a breakdown of the steps:

Using Docker, I construct an image based on the provided Dockerfile located at zarf/docker/Dockerfile.service. To mention, to ensure compatibility with Podman, the image is labeled as localhost/shohinsan/salesphereapi/sales-api:0.0.1. Throughout the build process, specific arguments are transmitted: 
* BUILD_REF is designated as 0.0.1
* BUILD_DATE is established as the present date and time in UTC format (%Y-%m-%dT%H:%M:%SZ).

In the build progression, we initiate by establishing the foundational environment, loading essential base and build images such as alpine:3.19.1 and golang:1.22.1 respectively. Following this, we seamlessly integrate required files and configurations by transferring the build context, ensuring all necessary dependencies are resolved and extracted. Security measures are implemented by configuring user and directory permissions within the image. We then copy pertinent local files into the Docker image, facilitating runtime accessibility. Specifying the working directory within the image streamlines operations. Finally, the Go application undergoes compilation using the "go build" command, transforming it into an executable within the Docker environment.

![image](https://github.com/shohinsan/SaleSphereAPI/assets/22685770/d9ed9f6d-1c73-4bc2-83d6-7e4b26f99d4a)

Upon invoking "make dev-status-all" command, the current status of the development environment is queried using various Kubernetes commands. Specifically, "kubectl get pods -o wide --watch --all-namespaces" command retrieves the status of all pods across namespaces. However, to focus solely on the "sales-system" namespace, we need to filter the results. Ensuring that only pods within the "sales-system" namespace are considered, it is observed that the pod named "sales-5bc95c6f54-cs2l9" is indeed running, indicating the successful functioning of the "sales-system" components.

