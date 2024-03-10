package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const FuncTemplate = `package bytesize_icons_templ
{{- range .}}
templ {{.Name}}() {
  {{.SVG | indent 2 }}
}
{{- end}}

`
type Icon struct {
  SVG string
  Name string
}

// Source: https://github.com/Masterminds/sprig/blob/581758eb7d96ae4d113649668fa96acc74d46e7f/functions.go#L50
func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}
func main() {
  files, err := os.ReadDir("bytesize-icons/dist/icons")
  if err != nil {
    log.Fatal(err)
  }

  templFile, err := os.OpenFile("src.templ", os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    log.Fatal(err)
  }
  t, err := template.New("x").Funcs(map[string]any{"indent":indent}).Parse(FuncTemplate)
  if err != nil {
    log.Fatal(err)
  }
  
  
  icons := make([]Icon, len(files))
  for i := range files {
    basename := filepath.Base(files[i].Name())
    extIndex := strings.Index(basename, ".")
    funcName := strcase.ToCamel(basename[:extIndex])

    f, err := os.ReadFile("bytesize-icons/dist/icons/" + files[i].Name())
    if err != nil {
      log.Fatal(err)
    }
    
    icons[i] = Icon{
      SVG: string(f),
      Name: funcName,
    }
  }
  err = t.Execute(templFile, icons)
  if err != nil {
    log.Fatal(err)
  }



}
