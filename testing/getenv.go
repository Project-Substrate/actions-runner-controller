// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package testing

import (
	"os"
	"testing"
)

func Getenv(t *testing.T, name string, opts ...string) string {
	t.Helper()

	v := os.Getenv(name)
	if v == "" {
		if len(opts) == 0 {
			t.Fatal(name + " must be set")
		}
		v = opts[0]
	}
	return v
}
