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

	"github.com/michaelact/Ansibila.go/ansible"
	"github.com/michaelact/Ansibila.go/internal/reader"
	log "github.com/sirupsen/logrus"
)

func (r *Runtime) GetPlaybook() {
	config := r.Config
	dirPath := config.DirectoryPath
	playbookPath := dirPath + "/molecule/default/playbook.yml"

	_, err := os.Stat(playbookPath)
	if err == nil {
		r.Playbook = ansible.ExtractPlaybook(playbookPath)
	} else {
		log.Error(err)
	}
}

type Playbook []reader.Line
