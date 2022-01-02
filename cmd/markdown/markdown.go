/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package markdown

import (
	"github.com/michaelact/Ansibila.go/cmd/markdown/table"
	"github.com/michaelact/Ansibila.go/internal/cli"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for 'markdown' command
func NewCommand(runtime *cli.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.MaximumNArgs(1),
		Use:         "markdown [PATH]",
		Annotations: cli.Annotations("markdown"),
		Aliases:     []string{"md"},
		Short:       "Generate Markdown of inputs and outputs",
	}

	// formatter subcommands
	cmd.AddCommand(table.NewCommand(runtime))
	return cmd
}
