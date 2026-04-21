// Copyright (c) 2024-2026 Magnon Compute Corporation. All Rights Reserved.

package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Cmdr struct {
}

func (k Cmdr) CombinedOutput(cmd *exec.Cmd) (string, error) {
	o, err := cmd.CombinedOutput()
	if err != nil {
		args := append([]string{}, cmd.Args...)
		args[0] = cmd.Path

		cs := strings.Join(args, " ")
		s := string(o)
		k.Errorf("%s failed with output:\n%s", cs, s)

		return s, err
	}

	return string(o), nil
}

func (k Cmdr) Errorf(f string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, f+"\n", args...)
}
