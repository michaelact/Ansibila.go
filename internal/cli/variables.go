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
	"path/filepath"
	"strings"

	"github.com/michaelact/Ansibila.go/ansible"
	"github.com/michaelact/Ansibila.go/pkg"
	log "github.com/sirupsen/logrus"
)

func (r *Runtime) GetVariables() {
	config := r.Config
	dirPath := config.DirectoryPath
	varFile := config.VariableFilename
	includeDir := []string{"tasks", "meta", "defaults", "vars", "handlers", "templates"}

	writtenVar := ansible.ExtractWrittenVariable(dirPath + "/" + varFile)
	defaultVar := ansible.ExtractDefaultVariable(dirPath + "/defaults/main.yml")

	// Return all variables used in ansible roles
	var actualVar []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if pkg.StringInSlice(path, includeDir) {
			for _, line := range ansible.ExtractActualVariable(path) {
				actualVar = append(actualVar, line.Value)
			}
		}

		return nil
	})

	if err != nil {
		log.Error(err)
	}

	for _, varName := range writtenVar.GetKeys() {
		var variable RuntimeVariable
		variable.Name = varName
		variable.Type = writtenVar[varName].Type
		variable.Description = writtenVar[varName].Description
		variable.Value = defaultVar[varName]
		variable.Declared = true

		variable.CheckUsed(actualVar)
		variable.Required = variable.InUse && variable.Value == nil

		if variable.Value != nil {
			variable.TypeCheck()
		}

		r.Variables = append(r.Variables, variable)
	}

	storedVarName := r.Variables.GetNames()
	// Cross Check existence between written variable and default value
	for _, varName := range append(defaultVar.GetKeys()) {
		if pkg.StringInSlice(varName, storedVarName) {
			continue
		} else {
			var variable RuntimeVariable
			variable.Name = varName
			variable.Value = defaultVar[varName]
			variable.Declared = false

			variable.CheckUsed(actualVar)
			variable.Required = variable.InUse && variable.Value == nil
			r.Variables = append(r.Variables, variable)
		}
	}
}

type RuntimeVariable struct {
	Name        string
	Type        string
	Description string
	Value       interface{}
	ValueType   string
	InUse       bool
	CorrectType bool
	Declared    bool
	Required    bool
}

func (r *RuntimeVariable) TypeCheck() bool {
	if r.Value == "" {
		log.Debug("Variable ", r.Name, " doesn't have default value.")
	}

	types := []string{"int", "str", "bool", "dict", "list"}

	r.CorrectType = pkg.StringInSlice(r.Type, types)
	if !r.CorrectType {
		log.Info("Datatype of ", r.Name, " variable is not found: ", r.Type)
	}

	r.ValueType = pkg.TypeOf(r.Value)
	if strings.Contains(r.ValueType, "map") {
		r.ValueType = "dict"
	} else if strings.Contains(r.ValueType, "[]") {
		r.ValueType = "list"
	} else if strings.Contains(r.ValueType, "string") {
		r.ValueType = "str"
	}

	r.CorrectType = r.ValueType == r.Type
	if !r.CorrectType {
		log.Warn("Datatype of ", r.Name, " is not match with the default value")
	}

	return r.CorrectType
}

func (r *RuntimeVariable) UnDeclare() {
	r.Declared = false
	log.Warn("Variable ", r.Name, " is not declared")
}

func (r *RuntimeVariable) CheckUsed(actualVars []string) bool {
	for _, wordInSlice := range actualVars {
		if strings.Contains(wordInSlice, r.Name) {
			r.InUse = true
			break
		} else {
			r.InUse = false
		}
	}

	if !r.InUse {
		log.Warn("Variable ", r.Name, " is not in use")
	}

	return r.InUse
}

type Variables []RuntimeVariable

func (v Variables) GetNames() []string {
	keys := make([]string, len(v))

	i := 0
	for _, k := range v {
		keys[i] = k.Name
		i++
	}

	return keys
}
