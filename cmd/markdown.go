package cmd

import (
	"github.com/spf13/cobra"

	"github.com/michaelact/Ansibila.go/service"
)

func NewMarkdownCommand(ansibleVariable service.AnsibleVariable) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "markdown",
		Short: "Generate Markdown of variable inputs",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := ansibleVariable.FindAll()
			if err != nil{
				return err
			}

			return nil
		},
	}

	return cmd
}
