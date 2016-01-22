package template_testing

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestGeneratingModel(t *testing.T) {
	b, err := ioutil.ReadFile("./model.go.template")
	if err != nil {
		t.Fatal(err)
	}

	funcs := template.FuncMap{
		"capitolize": strings.Title,
		"lowercase":  strings.ToLower,
	}

	moduleTemplate := struct {
		Package      string
		Name         string
		Fields       map[string]string
		FieldArgList string
	}{
		Package: "TESTMODEL",
		Name:    "User",
		Fields: map[string]string{
			"name": "string",
			"age":  "int",
		},
		FieldArgList: "name string, age int",
	}

	t.Log(string(b))
	tmpl, err := template.New("Test template").Funcs(funcs).Parse(string(b))
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("user_model.go")
	if err != nil {
		t.Fatal(err)
	}
	err = tmpl.Execute(f, moduleTemplate)
	if err != nil {
		t.Fatal(err)
	}
}
