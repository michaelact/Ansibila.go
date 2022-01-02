/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package check

import (
	"github.com/michaelact/Ansibila.go/internal/cli"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for 'check' command
func NewCommand(runtime *cli.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.MaximumNArgs(1),
		Use:         "check [/path/to/roles/name/]",
		Aliases:     []string{"ch"},
		Annotations: cli.Annotations("check"),
		Short:       "Check Compability between Real Variable and Written Variable",
		Run:         runtime.Check,
	}

	return cmd
}
