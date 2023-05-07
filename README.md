# 🛠️ Sandkube

A Kubernetes sandbox cluster configured to test different infrastructural tools and components.

About the cluster:

- a multi-node cluster powered by k3d
- deploys a few connected microservice via Tilt
- supports WriteReadOnce volumes via local-path provisioner

## Use cases

WriteReadMany (WRX) volumes or an NFS server:

```sh
make nfs-server-install
```
