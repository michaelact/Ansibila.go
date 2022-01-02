/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package main

import (
	"os"

	"github.com/michaelact/Ansibila.go/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
