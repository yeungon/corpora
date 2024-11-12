package html

import (
	"embed"
	"strings"
	"text/template"
)

var (
	//home        = parse("index.html")
	home        = parseMultiple("template/master/home.html", "template/corpora/search.html")
	profileShow = parseMultiple("template/profile/show.html")
	signupUser  = parseMultiple("template/profile/signup.html")
	profileEdit = parseMultiple("template/profile/edit.html")
)

// The configuration below is important and REQUIRED, it is a derective.//

//go:embed * template/profile/* template/master/* template/corpora/*
var filesystem embed.FS

var funcs = template.FuncMap{
	"uppercase": func(v string) string {
		return strings.ToUpper(v)
	},
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("template/master/layout.html").Funcs(funcs).ParseFS(filesystem, "template/master/layout.html", file))
}

// Helper function to parse multiple template files, always including template/master/layout.html
func parseMultiple(files ...string) *template.Template {
	// Default template files loaded by default
	allFiles := append([]string{"template/master/layout.html", "template/master/nav.html", "template/master/header.html", "template/master/footer.html"}, files...)
	return template.Must(
		template.New("template/master/layout.html").Funcs(funcs).ParseFS(filesystem, allFiles...))
}

//NOTE: We might create many more version of parseMultiple for specific purposes. <----------
// ref: https://philipptanlak.com/web-frontends-in-go/#implementing-the-template-renderers
// future work: https://github.com/dstpierre/tpl/tree/main
