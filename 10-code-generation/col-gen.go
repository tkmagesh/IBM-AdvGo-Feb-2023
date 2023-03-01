//go:build exclude

package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var collectionUtilTemplate = `
/* Auto generated code. Do not modify it */
package {{.PkgName}}
type {{.TypeName}}s []{{.TypeName}}

func (items *{{.TypeName}}s) IndexOf(item {{.TypeName}}) int {
for idx, p := range *items {
if p == item {
return idx
}
}
return -1
}

func (items *{{.TypeName}}s) Includes(item {{.TypeName}}) bool {
return items.IndexOf(item) != -1
}

func (items *{{.TypeName}}s) Any(predicate func({{.TypeName}}) bool) bool {
for _, item := range *items {
if criteria(item) {
return true
}
}
return false
}


func (items {{.TypeName}}s) Filter(predicate func(item {{.TypeName}}) bool) {{.TypeName}}s {
	result := {{.TypeName}}s{}
	for _, p := range items {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}
`

type TemplateData struct {
	TypeName string
	PkgName  string
}

func main() {
	typeName := flag.String("N", "Test", "Type Name")
	pkgName := flag.String("P", "TestPackage", "Package Name")
	flag.Parse()
	templateData := TemplateData{*typeName, *pkgName}
	fileName := templateData.TypeName + "s.go"
	file, _ := os.Create(fileName)
	defer file.Close()

	parsedT := template.Must(template.New("").Parse(collectionUtilTemplate))
	parsedT.Execute(file, templateData)
	fmt.Println(fileName, " created!")
}
