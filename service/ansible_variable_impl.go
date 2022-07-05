package service

import (
	"fmt"

	"github.com/michaelact/Ansibila.go/model/ansible"
	"github.com/michaelact/Ansibila.go/model/config"
)

type AnsibleVariableImpl struct {
	Path *config.Path
}

func NewAnsibleVariable(path *config.Path) AnsibleVariable {
	return &AnsibleVariableImpl{
		Path: path, 
	}
}

func (self *AnsibleVariableImpl) FindAll() []ansible.Variable {
	fmt.Println("Masuk sini")
	return nil
}
