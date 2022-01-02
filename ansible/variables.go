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
	"regexp"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	log "github.com/sirupsen/logrus"
)

func ExtractDefaultVariable(filePath string) DefaultVariable {
	content, _ := ioutil.ReadFile(filePath)

	var data DefaultVariable
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		log.Error(err)
	}

	return data
}

func ExtractWrittenVariable(filePath string) WrittenVariable {
	content, _ := ioutil.ReadFile(filePath)

	var data WrittenVariable
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		log.Error(err)
	}

	return data
}

func ExtractActualVariable(filePath string) []reader.Line {
	pattern := `{{ *.* }}`
	re := regexp.MustCompile(pattern)
	lines := reader.Lines{
		FileName: filePath,
		LineNum: 1,
		Condition: func(line reader.Line, lines []reader.Line) bool {
			log.Debug(line)
			status := re.MatchString(line.Value)
			return status
		},
		Parser: func(line reader.Line) ([]string, bool) {
			words := re.FindAllString(line.Value, -1)

			for i := range words {
				words[i] = words[i][len(`{{`) : len(words[i])-len(`}}`)]
		    }

			return words, true
		},
	}

	realVariable := lines.ExtractFromText()
	return realVariable
}
