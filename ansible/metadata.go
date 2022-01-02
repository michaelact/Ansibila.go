/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package ansible

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	log "github.com/sirupsen/logrus"
)

func ExtractMetadata(filePath string) Metadata {
	content, _ := ioutil.ReadFile(filePath)

	var data Metadata
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		log.Error(err)
	}

	return data
}
