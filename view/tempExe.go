package view

import (
	"html/template"
	"net/http"
)
type Template struct {
	htmlTpl *template.Template
}

func ParseTemplate(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, err
	}
	return Template{htmlTpl: tpl}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) error {
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}