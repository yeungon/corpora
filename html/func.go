package html

import (
	"strings"
	"text/template"
)

// Add function to increment a number
func add(a, b int) int {
	return a + b
}

// Sub function to decrement a number
func sub(a, b int) int {
	return a - b
}

var funcs = template.FuncMap{
	"uppercase": func(v string) string {
		return strings.ToUpper(v)
	},
	"add": add,
	"sub": sub,
}
