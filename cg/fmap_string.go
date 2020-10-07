package cg

import (
	"strings"
	"text/template"
)

var fmap = template.FuncMap{
	"lower": func(text string) string {
		return strings.ToLower(text)
	},
	"upper": func(text string) string {
		return strings.ToUpper(text)
	},
	"firstUp": func(text string) string {
		return strings.ToUpper(string(text[0])) + text[1:]
	},
	"firstLow": func(text string) string {
		return strings.ToLower(string(text[0])) + text[1:]
	},
}
