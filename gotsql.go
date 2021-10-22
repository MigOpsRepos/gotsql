package gotsql

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type GoTSQL struct {
	Templates map[string]*template.Template
}

func (g *GoTSQL) Load(filedir string) {
	if g.Templates == nil {
		g.Templates = make(map[string]*template.Template)
	}

	err := filepath.Walk(filedir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".gotsql" {
			return nil
		}
		t, er := template.ParseFiles(path)
		if er != nil {
			return fmt.Errorf("error parsing template %s", path, er)
		}
		g.Templates[path[:len(path)-len(filepath.Ext(path))]] = t
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func (g *GoTSQL) Get(commandName string, data interface{}) (string, error) {
	namespace := filepath.Dir(commandName)

	if t, ok := g.Templates[namespace]; ok {
		var b bytes.Buffer

		templateName := filepath.Base(commandName)
		tmpl_err := t.ExecuteTemplate(&b, templateName, data)
		if tmpl_err != nil {
			return "", tmpl_err
		}

		return b.String(), nil
	} else {
		return "", fmt.Errorf("template %s not found or loaded", namespace)
	}

}
