package {{.pkg.name}}

import "fmt"

func {{.pkg.func}}()  {
	{{- range .pkg.message.words}}
	fmt.Println("{{.}}")
	{{- end}}
}