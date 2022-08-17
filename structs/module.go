package structs

import (
	"bytes"
)

type Module struct {
	Name        string
	Prereqs     func() bool
	Enumeration func() ([]string, bool)
	Report      func([]string) *bytes.Buffer
}

func NewModule(name string, prereqs func() bool, enumeration func() ([]string, bool), report func([]string) *bytes.Buffer) *Module {
	return &Module{
		Name:        name,
		Prereqs:     prereqs,
		Enumeration: enumeration,
		Report:      report,
	}
}
