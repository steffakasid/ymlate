package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type TemplateSettings struct {
	Values map[string]interface{}
}

func main() {
	values := map[string]interface{}{}

	yamlFile, err := ioutil.ReadFile("values.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &values)
	if err != nil {
		panic(err)
	}

	myTemplate := template.New("template.yml")

	myTemplate.Funcs(template.FuncMap{
		"helloWorld": func(feature string) string {
			return "Hello" + feature
		},
		"toyaml": func(yamlObj map[interface{}]interface{}) string {
			out, err := yaml.Marshal(&yamlObj)
			if err != nil {
				panic(err)
			}
			return string(out)
		},
		"indent": func(indent int, str string) string {
			var indentBlanks string
			for i := 0; i < indent; i++ {
				indentBlanks += " "
			}

			var returnString string
			for _, line := range strings.Split(str, "\n") {
				returnString += indentBlanks + line + "\n"
			}
			return returnString
		},
	})

	// TODO: ParseFiles for sure can get multiple templatefiles
	tmpl, err := myTemplate.ParseFiles("template.yml")

	if err != nil {
		panic(err)
	}

	templateSettings := TemplateSettings{
		Values: values,
	}

	err = tmpl.Execute(os.Stdout, templateSettings)
	if err != nil {
		panic(err)
	}
}
