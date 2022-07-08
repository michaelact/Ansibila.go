package service

// import (
// 	"github.com/michaelact/Ansibila.go/model/ansible"
// )

type AnsibleVariable interface {
	// FindAll() ([]ansible.Variable, error)
	FindAll() ([]string, error)
}
