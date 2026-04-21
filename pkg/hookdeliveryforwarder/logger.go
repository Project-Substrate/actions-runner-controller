// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package hookdeliveryforwarder

import (
	"fmt"
	"os"
)

type logger struct {
}

func (f logger) Logf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", args...)
}

func (f logger) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
