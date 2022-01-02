/*
Copyright 2021 The Michael Act Author.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package ansible

import (
	"github.com/michaelact/Ansibila.go/internal/reader"
)

func ExtractPlaybook(filePath string) []reader.Line {
	lines := reader.Lines{
		FileName: filePath,
		LineNum: 1,
		Condition: func(line reader.Line, lines []reader.Line) bool { return true },
		Parser: func(line reader.Line) ([]string, bool) { return []string{ line.Value }, true },
	}

	realPlaybook := lines.ExtractFromText()
	return realPlaybook
}
