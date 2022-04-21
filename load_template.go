package golang_ture

import (
	"embed"
	"html/template"
)

var paths = map[string]string{
	"todoHtmlDir":  "internal/templates/html/todo_list",
	"templatesDir": "internal/templates",
}

const (
	extension = "/*.html"
)

var (
	//go:embed internal/templates/*
	files     embed.FS
	Templates map[string]*template.Template
)

func LoadTemplates() error {
	if err := loadTodoListTemplate(); err != nil {
		return err
	}
	return nil
}

func loadTodoListTemplate() error {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}

	funcMap := map[string]interface{}{
		"Increment": func(i int) int {
			return i + 1
		},
	}
	tmpl, err := template.New("todo_list_page.html").Funcs(funcMap).ParseFS(files, paths["todoHtmlDir"]+"/todo_list_page.html", paths["todoHtmlDir"]+extension)
	if err != nil {
		return err
	}

	Templates["todo_list_page.html"] = tmpl
	return nil
}
