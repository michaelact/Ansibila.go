/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/michaelact/Ansibila.go/cmd/check"
	"github.com/michaelact/Ansibila.go/cmd/markdown"
	"github.com/michaelact/Ansibila.go/configs"
	"github.com/michaelact/Ansibila.go/internal/cli"
	"github.com/michaelact/Ansibila.go/internal/version"
	"github.com/spf13/cobra"
)

func Execute() error {
	err := NewCommand().Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}
	return nil
}

// NewCommand returns a new cobra.Command for 'root' command
func NewCommand() *cobra.Command {
	config := configs.DefaultConfig()
	runtime := cli.NewRuntime(config)
	cmd := &cobra.Command{
		Args:    cobra.MaximumNArgs(1),
		Use:     "ansibila [PATH]",
		Short:   "A utility to generate documentation from Ansible roles in various output formats",
		Version: version.Full(),
	}

	// flag commands
	cmd.PersistentFlags().StringVarP(&config.DirectoryPath, "dirpath", "d", "", "role directory")
	cmd.PersistentFlags().StringVarP(&config.TemplatePath, "tmplpath", "t", "", "template files directory")
	cmd.PersistentFlags().StringVarP(&config.VariableFilename, "varfile", "f", "", "variable filename that containing the `variable: description, type` structure")

	// formatter subcommands
	cmd.AddCommand(markdown.NewCommand(runtime))
	cmd.AddCommand(check.NewCommand(runtime))

	return cmd
}
