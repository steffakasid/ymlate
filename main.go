package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

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
	fmt.Printf("--- m:\n%v\n\n", values)

	content, err := ioutil.ReadFile("template.yml")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template.yml").Parse(string(content))

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
