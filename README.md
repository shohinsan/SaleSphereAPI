# Sales Backend Web API  

## Install

* https://kustomize.io/
* https://kubernetes.io/docs/tasks/tools/
* https://kind.sigs.k8s.io/

## Upgrade/Increment to Latest Stable Version 

* Dependencies in `go.mod` 

## Steps

Kind creates a cluster named "sale-sphere-cluster" using the specified image (kindest/node:v1.29.2) and configuration file (zarf/k8s/dev/kind-config.yaml).

Various steps are performed during cluster creation:

![Xnapper-2024-03-06-01 13 54](https://github.com/shohinsan/SaleSphereAPI/assets/22685770/7ccf92a7-1e7e-4fdd-9f71-73d8388a6763)

* Ensuring the specified node image is available.
* Preparing nodes for the cluster.
* Writing configuration for the cluster.
* Starting the control plane.
* Installing the Container Network Interface (CNI).
* Installing the StorageClass.
* After successful cluster creation, the kubectl context is set to "kind-sale-sphere-cluster", allowing interaction with the newly created cluster.

![image](https://github.com/shohinsan/SaleSphereAPI/assets/22685770/48093425-4963-412b-9981-c9920da1bfad)

This log shows the execution of the `make all` command, which typically builds and prepares various components of a project. Here's a breakdown of the steps:

Docker builds an image using the provided Dockerfile (zarf/docker/Dockerfile.service).
* The image is tagged as shohinsan/salesphereapi/sales-api:0.0.1.
* During the build process, the following arguments are passed:
* BUILD_REF: Set to 0.0.1.
* BUILD_DATE: Set to the current date and time in UTC format (%Y-%m-%dT%H:%M:%SZ).

The build process includes several stages:
* Loading the base image (docker.io/library/alpine:3.19.1) and build image (docker.io/library/golang:1.22.1).
* Transferring the build context.
* Resolving and extracting dependencies.
* Setting up user and directory permissions.
* Copying files from the local directory to the Docker image.
* Setting the working directory within the image.
* Compiling the Go application (go build).
