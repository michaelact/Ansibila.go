/*
Copyright 2021 The terraform docs Authors.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"github.com/michaelact/Ansibila.go/ansible"
	"github.com/michaelact/Ansibila.go/configs"
)

// Runtime represents the execution runtime for CLI.
type Runtime struct {
	Config    *configs.Config
	Variables Variables
	Modules   Modules
	Playbook  Playbook
	Metadata  ansible.Metadata
}

// NewRuntime returns new instance of Runtime. If `config` is not provided
// default config will be used.
func NewRuntime(config *configs.Config) *Runtime {
	return &Runtime{Config: config}
}
