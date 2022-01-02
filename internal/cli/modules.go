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
	"path/filepath"

	"github.com/michaelact/Ansibila.go/ansible"
	"github.com/michaelact/Ansibila.go/pkg"
	log "github.com/sirupsen/logrus"
)

func (r *Runtime) GetModules() {
	config := r.Config
	dirPath := config.DirectoryPath
	includeDir := []string{"tasks", "handlers"}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if pkg.StringInSlice(path, includeDir) {
			for _, line := range ansible.ExtractModule(path) {
				module := NewModule()
				module.Name = line.Value
				module.FullURL = module.BaseURL + line.Value + "_module.html"
				r.Modules = append(r.Modules, module)
			}
		}

		return nil
	})

	if err != nil {
		log.Error(err)
	}
}

type RuntimeModule struct {
	Name    string
	BaseURL string
	FullURL string
}

func NewModule() RuntimeModule {
	return RuntimeModule{
		Name:    "",
		BaseURL: "https://docs.ansible.com/ansible/latest/collections/ansible/builtin/",
		FullURL: "",
	}
}

type Modules []RuntimeModule
