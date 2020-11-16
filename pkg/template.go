package pkg

import (
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/steffakasid/ymlate/helm/pkg/engine"
	"gopkg.in/yaml.v2"
)

type Template struct {
	TemplateFile string
	ValuesFile   string
	Values       map[string]interface{}
}

func getValues(tpl *Template) {
	values := map[string]interface{}{}

	yamlFile, err := ioutil.ReadFile(tpl.ValuesFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &values)
	if err != nil {
		panic(err)
	}
	tpl.Values = values
}

func Render(tpl Template) {
	getValues(&tpl)

	name := path.Base(tpl.TemplateFile)
	myTemplate := template.New(name)

	myTemplate.Funcs(engine.GetHelmFunction())

	// TODO: ParseFiles for sure can get multiple templatefiles
	tmpl, err := myTemplate.ParseFiles(tpl.TemplateFile)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, tpl)
	if err != nil {
		panic(err)
	}
}
