package cg

import (
	"strings"
	"text/template"
)

var FMap = template.FuncMap{
	"lower": func(text string) string {
		return strings.ToLower(text)
	},
	"upper": func(text string) string {
		return strings.ToUpper(text)
	},
}
