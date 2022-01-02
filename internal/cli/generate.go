/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"os"
	"strings"
	"text/template"

	"github.com/michaelact/Ansibila.go/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func (r *Runtime) Generate(cmd *cobra.Command, args []string) {
	config := r.Config
	tmplPath := config.TemplatePath + "/"

	r.GetVariables()
	r.GetModules()
	r.GetPlaybook()
	r.GetMetadata()

	command := cmd.Annotations["command"]
	tFilename := strings.Replace(command, " ", `_`, 1) + ".tmpl"
	tFilepath := tmplPath + tFilename
	t, err := template.New(tFilename).Funcs(pkg.TemplateFuncs()).ParseFiles(tFilepath)
	if err != nil {
		log.Error(err)
	}

	err = t.Execute(os.Stdout, r)
	if err != nil {
		log.Error(err)
	}
}
