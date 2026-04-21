# actions-runner-controller

ARC self-hosted runner fleet configuration for `magnon-enterprise-runners`. Manages the ARC k8s runner pods across ash and hel1 clusters. YAML/Kustomize manifests. All CI across 33 project orgs uses these runners — never `ubuntu-latest`. No deployable service beyond ARC itself.
