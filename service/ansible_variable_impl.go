package service

import (
	"path/filepath"
	"strings"
	"regexp"
	"io/fs"
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/michaelact/Ansibila.go/model/ansible"
	"github.com/michaelact/Ansibila.go/model/jinja"
	"github.com/michaelact/Ansibila.go/model/config"
	"github.com/michaelact/Ansibila.go/helper"
)

type AnsibleVariableImpl struct {
	Path *config.Path
}

func NewAnsibleVariable(path *config.Path) AnsibleVariable {
	return &AnsibleVariableImpl{
		Path: path, 
	}
}

// func (self *AnsibleVariableImpl) FindAll() ([]ansible.Variable, error) {
func (self *AnsibleVariableImpl) FindAll() ([]string, error) {
	var variables []string
	err := filepath.Walk(self.Path.Directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileExt := filepath.Ext(path)
			if fileExt == ".yml" || fileExt == ".yaml" || fileExt == ".jj" || fileExt == ".j2" {
				lineByLine := helper.Lines{
					FileName:  path,
					Condition: isContainsAnsibleVariable, 
					Parser:    getAnsibleVariable,
				}

				result, err := lineByLine.Extract()
				if err != nil {
					return err
				}

				variables = append(variables, result...)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	variables = slices.Compact(variables)
	return variables, nil
}

func isContainsAnsibleVariable(line string) bool {
	if strings.Contains(line, "}}") && strings.Contains(line, "{{") {
		return true
	}

	return false
}

func getAnsibleVariable(line string) []string {
	regex, _ := regexp.Compile(`{[{%]?[-]?[{]([^{}]*)[{%]?[-]?}`)
	expr := regex.FindAllStringSubmatch(line, -1)

	var result []string
	for _, match := range expr {
		if len(match) == 2 {
			extracted := strings.Trim(match[1], " ")
			for _, word := range strings.Split(extracted, " ") {
				word := strings.Split(word, ".")[0]
				if !(slices.Contains(ansible.BuiltInVariable, word) || slices.Contains(jinja.Symbols, word) || slices.Contains(jinja.Keywords, word)) {
					result = append(result, word)
				}
			}
		}
	}

	return result
}
