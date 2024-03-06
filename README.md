# SaleSphereAPI

## Install

* https://kustomize.io/
* https://kubernetes.io/docs/tasks/tools/
* https://kind.sigs.k8s.io/

## Upgrade/Increment to Latest Stable Version 

* Dependencies in `go.mod` 

## Steps I have taken

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
