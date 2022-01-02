/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package configs

// Config represents all the available config options that can be accessed and
// passed through CLI.
type Config struct {
	DirectoryPath 	 string
	TemplatePath     string
	VariableFilename string
}

func DefaultConfig() *Config {
	return &Config {
		DirectoryPath:	  ".", 
		VariableFilename: "variables.yml", 
	}
}
