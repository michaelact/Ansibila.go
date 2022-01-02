package table

import (
	"github.com/michaelact/Ansibila.go/internal/cli"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra. Command for 'table' command
func NewCommand(runtime *cli.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.MaximumNArgs(1),
		Use:         "table",
		Aliases:     []string{"tb"},
		Annotations: cli.Annotations("markdown table"),
		Short:       "Generate Markdown tables of inputs and outputs",
		Run:         runtime.Generate,
	}

	return cmd
}
