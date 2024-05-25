package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

var moduleElements = []string{"dhttp", "repo", "types"}

func GenerateModule(name string) error {
	module := NewModule(name)

	err := os.Mkdir(module.ServiceName, 0777)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	for idx := range moduleElements {
		fileName := filepath.Join(
			wd,
			"internal",
			"templates",
			fmt.Sprintf("%s.tmpl", moduleElements[idx]),
		)

		folderName := fmt.Sprintf(
			"%s/%s",
			module.ServiceName,
			moduleElements[idx],
		)

		err = os.Mkdir(folderName, 0777)
		if err != nil {
			return err
		}

		err = writeFile(fileName, module, moduleElements[idx])
		if err != nil {
			return err
		}
	}

	fileName := filepath.Join(wd, "internal", "templates", "usecase.tmpl")
	return writeFile(fileName, module, "")
}

func writeFile(templateName string, module *Module, element string) error {
	data, err := os.ReadFile(templateName)
	if err != nil {
		return err
	}
	tmpl, err := template.New(element).Parse(string(data))
	if err != nil {
		return err
	}

	filename := filepath.Join(
		module.ServiceName,
		element,
		fmt.Sprintf("%s.go", module.ServiceName),
	)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, module)
	if err != nil {
		return err
	}

	return nil
}
