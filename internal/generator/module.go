package generator

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"ray/internal/config"
	"strings"
)

type Module struct {
	PackageName string
	Name        string
	ServiceName string
	Prefix      string
}

func NewModule(name string) *Module {
	return &Module{
		PackageName: config.Get().Package,
		Name:        cases.Title(language.English).String(name),
		ServiceName: strings.ToLower(name),
		Prefix:      prefix(strings.ToLower(name)),
	}
}
