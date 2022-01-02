/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package ansible

type DefaultVariable map[string]interface{}

func (data DefaultVariable) GetKeys() []string {
	keys := make([]string, len(data))

	i := 0
	for k := range data {
		keys[i] = k
		i++
	}

	return keys
}

type WrittenVariable map[string]struct {
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

func (data WrittenVariable) GetKeys() []string {
	keys := make([]string, len(data))

	i := 0
	for k := range data {
		keys[i] = k
		i++
	}

	return keys
}

type Metadata struct {
	Dependencies []interface{} `yaml:"dependencies"`
	GalaxyInfo   struct {
		RoleName          string  `yaml:"role_name"`
		Author            string  `yaml:"author"`
		Description       string  `yaml:"description"`
		Company           string  `yaml:"company"`
		License           string  `yaml:"license"`
		MinAnsibleVersion float64 `yaml:"min_ansible_version"`
		Platforms         []struct {
			Name     string    `yaml:"name"`
			Versions []string  `yaml:"versions"`
		} `yaml:"platforms"`
		GalaxyTags []string `yaml:"galaxy_tags"`
	} `yaml:"galaxy_info"`
}
