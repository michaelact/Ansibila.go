package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/michaelact/Ansibila.go/model/config"
	"github.com/michaelact/Ansibila.go/service"
)

var rootCmd = &cobra.Command{
	Use:   "Ansibila",
	Short: "Generate documentation from Ansible roles in various output formats.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	path := config.Path{}

	ansibleVariable := service.NewAnsibleVariable(&path)

	rootCmd.AddCommand(NewMarkdownCommand(ansibleVariable))

	flags := rootCmd.Flags()
	flags.StringVarP(&path.Directory, "directory", "d", ".", "Root directory of ansible role")
	flags.StringVarP(&path.Template, "template", "t", "", "Directory of template files")
	flags.StringVarP(&path.Variable, "variable", "v", "variables.yaml", "Variable filepath")
}
