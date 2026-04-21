// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package actions

import (
	"context"
	"io"
)

//go:generate mockery --inpackage --name=SessionService
type SessionService interface {
	GetMessage(ctx context.Context, lastMessageId int64, maxCapacity int) (*RunnerScaleSetMessage, error)
	DeleteMessage(ctx context.Context, messageId int64) error
	AcquireJobs(ctx context.Context, requestIds []int64) ([]int64, error)
	io.Closer
}
