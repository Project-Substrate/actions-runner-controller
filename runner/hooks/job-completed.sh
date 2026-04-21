#!/usr/bin/env bash
# Copyright (c) Magnon Compute Corporation. All rights reserved.
set -Eeuo pipefail

# shellcheck source=runner/logger.sh
source logger.sh

log.debug "Running ARC Job Completed Hooks"

for hook in /etc/arc/hooks/job-completed.d/*; do
  log.debug "Running hook: $hook"
  "$hook" "$@"
done
