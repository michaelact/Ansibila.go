/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"github.com/michaelact/Ansibila.go/ansible"
)

func (r *Runtime) GetMetadata() {
	config := r.Config
	dirPath := config.DirectoryPath
	r.Metadata = ansible.ExtractMetadata(dirPath + "/meta/main.yml")
}
