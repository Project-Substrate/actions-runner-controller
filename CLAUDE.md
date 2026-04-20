# CLAUDE.md — actions-runner-controller

## Purpose

Actions Runner Controller (ARC) is a Kubernetes operator that orchestrates and autoscales self-hosted GitHub Actions runners. It manages runner scale sets that expand and contract based on workflow queue depth, using ephemeral container-based runners for clean, rapid scaling. Magnon uses this to run CI/CD pipelines on Project-Substrate's own clusters rather than GitHub-hosted runners.

## Tech Stack

- **Language:** Go (1.25+)
- **Framework:** controller-runtime (Kubernetes operator pattern)
- **Key deps:** `github.com/google/go-github/v52`, `github.com/bradleyfalzon/ghinstallation/v2`, `sigs.k8s.io/controller-runtime`, Prometheus client, Ginkgo/Gomega for tests
- **Build:** Make + Docker Buildx (multi-arch: `linux/amd64`, `linux/arm64`)
- **Deployed via:** Helm charts in `charts/`

## Dev Commands

```bash
# Generate manifests and run unit tests
make test

# Run tests with a local kube-apiserver + etcd
make test-with-deps

# Build the controller manager binary
make manager

# Build and push multi-arch Docker image
make docker-buildx

# Generate CRDs and update Helm chart CRDs
make manifests

# Run locally against a kind cluster named "acceptance"
make run
```

## Key Invariants

- Runner pods are **ephemeral**: each job gets a fresh container; never reuse runner state across jobs.
- Scale sets are the current supported mode; the legacy autoscaling modes (`HorizontalRunnerAutoscaler`) are community-maintained only — do not extend them.
- GitHub App authentication is required for runner registration; PAT-based auth is legacy.
- CRD schema changes must remain backward-compatible with at least one prior release — check `CONTRIBUTING.md` before modifying API types in `apis/`.
- The `SYNC_PERIOD` (default 1m) controls how often the controller reconciles; do not reduce it below 30s in production to avoid GitHub API rate limiting.

## What NOT To Do

- Do not use `github.com/google/go-github/v52` GitHub client directly in hot reconcile loops — always go through the cached retryable HTTP client to avoid rate limiting.
- Do not add runner-level persistent volumes; runners must be stateless and disposable.
- Do not modify `go.sum` manually — always use `go mod tidy` after dependency changes.
- Do not deploy changes to the `AutoscalingRunnerSet` CRD without also updating the Helm chart CRDs via `make manifests`.
- Do not skip `make fmt vet` before committing — the CI gate will fail.
