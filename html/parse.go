package html

import (
	"embed"
	"strings"
	"text/template"
)

// ref: https://philipptanlak.com/web-frontends-in-go/#implementing-the-template-renderers

//go:embed *
var files embed.FS

var funcs = template.FuncMap{
	"uppercase": func(v string) string {
		return strings.ToUpper(v)
	},
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").Funcs(funcs).ParseFS(files, "layout.html", file))
}
