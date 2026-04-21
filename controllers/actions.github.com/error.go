// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package actionsgithubcom

type controllerError string

func (e controllerError) Error() string {
	return string(e)
}

const (
	retryableError = controllerError("retryable error")
	fatalError     = controllerError("fatal error")
)
