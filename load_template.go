package golang_ture

import (
	"embed"
	"html/template"
)

var paths = map[string]string{
	"todoHtmlDir":  "internal/templates/html/todo_list",
	"tmplDir":      "internal/templates/tmpl",
	"templatesDir": "internal/templates",
}

const (
	extensionHtml = "/*.html"
	extensionTmpl = "/*.tmpl"
)

var (
	//go:embed internal/templates/*
	files     embed.FS
	Templates map[string]*template.Template
)

func LoadTemplates() error {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}
	if err := loadTodoListTemplate(); err != nil {
		return err
	}
	return nil
}

func loadTodoListTemplate() error {
	funcMap := map[string]interface{}{
		"Increment": func(i int) int {
			return i + 1
		},
	}
	tmpl, err := template.New("todo_list_page.html").Funcs(funcMap).ParseFS(files, paths["todoHtmlDir"]+"/todo_list_page.html", paths["todoHtmlDir"]+extensionHtml)
	if err != nil {
		return err
	}

	Templates["todo_list_page.html"] = tmpl
	return nil
}

func LoadGetSetGenTemplate() error {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}

	tmpl, err := template.New("get_set_gen.tmpl").ParseFS(files, paths["tmplDir"]+"/get_set_gen.tmpl", paths["tmplDir"]+extensionTmpl)
	if err != nil {
		return err
	}

	Templates["get_set_gen.tmpl"] = tmpl
	return nil
}
