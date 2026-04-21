// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package actionssummerwindnet

func filterLabels(labels map[string]string, filter string) map[string]string {
	filtered := map[string]string{}

	for k, v := range labels {
		if k != filter {
			filtered[k] = v
		}
	}

	return filtered
}
