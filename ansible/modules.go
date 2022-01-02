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
	"strings"
	log "github.com/sirupsen/logrus"
)

func ExtractModule(filePath string) []reader.Line {
	lines := reader.Lines{
		FileName: filePath,
		LineNum: 1,
		Condition: func(line reader.Line, lines []reader.Line) bool {
			log.Debug(line)
			unique := true
			for _, module := range lines {
				if strings.Contains(line.Value, module.Value) {
					unique = false
					break
				}
			}

			status := unique && strings.HasPrefix(line.Parent, "- name:")
			return status
		},
		Parser: func(line reader.Line) ([]string, bool) {
			word := line.Value
			word = strings.TrimSpace(word)
			word = word[: strings.Index(word, ":")]

			return []string{ word }, true
		},
	}

	realModule := lines.ExtractFromText()
	return realModule
}
