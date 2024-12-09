package html

import (
	"embed"
	"text/template"
)

var (
	//home        = parse("index.html")
	PageHome       = parseMultiple("template/master/home.html", "template/corpora/search.html", "template/corpora/result.html")
	PageTokenize   = parseMultiple("template/features/tokenize.html")
	PageSignupUser = parseMultiple("template/features/signup.html")
	Pagephonemizer = parseMultiple("template/features/phonemizer.html")
	PageIPA        = parseMultiple("template/features/ipa.html")
	PageCredit     = parseMultiple("template/about/credit.html")
	PageAbout      = parseMultiple("template/about/about.html")
	PageDonate     = parseMultiple("template/about/donate.html")
)

// The configuration below is important and REQUIRED, it is a derective.//

//go:embed * template/*/*
var filesystem embed.FS

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
