/*
Copyright © 2020 Steffen Rumpf <ymlate@steffen-rumpf.de>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type RenderSettings struct {
	TemplateFile string
	ValuesFile   string
	Values       map[string]interface{}
}

var renderSettings = RenderSettings{}

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "render the given template and it's values",
	Long: `This command will render the given template,
	merge values and directly output the result on the
	command line.`,
	Run: func(cmd *cobra.Command, args []string) {

		values := map[string]interface{}{}

		yamlFile, err := ioutil.ReadFile(renderSettings.ValuesFile)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(yamlFile, &values)
		if err != nil {
			panic(err)
		}

		name := path.Base(renderSettings.TemplateFile)
		myTemplate := template.New(name)

		myTemplate.Funcs(template.FuncMap{
			"trim": func(p string) string {
				return strings.TrimSpace(p)
			},
			"trimAll": func(c, p string) string {
				return strings.Trim(p, c)
			},
			"trimPrefix": func(prefix, p string) string {
				return strings.TrimPrefix(p, prefix)
			},
			"trimSuffix": func(suffix, p string) string {
				return strings.TrimSuffix(p, suffix)
			},
			"lowwer": func(p string) string {
				return strings.ToLower(p)
			},
			"upper": func(p string) string {
				return strings.ToUpper(p)
			},
			"title": func(p string) string {
				return strings.Title(p)
			},
			"untitle": func(p string) string {
				return strings.ToLower(p)
			},
			"repeat": func(n int, p string) string {
				return strings.Repeat(p, n)
			},
			"substr": func(start, end int, p string) string {
				return string(p[start:end])
			},
			"nospace": func(p string) string {
				return strings.ReplaceAll(p, " ", "")
			},
			"trunc": func(n int, p string) string {
				// TODO: implement me
				return p
			},
			"abbrev": func(maxlen int, p string) string {
				// TODO: implement me
				return p
			},
			"abbrevboth": func(loffset, maxlen int, p string) string {
				// TODO: implement me
				return p
			},
			"initials": func(p string) string {
				// TODO: implement me
				return p
			},
			"randAlphaNum": func(len int) string {
				return ""
			},
			"randAlpha": func(len int) string {
				return ""
			},
			"randNummeric": func(len int) string {
				return ""
			},
			"randAscii": func(len int) string {
				return ""
			},
			"wrap": func(col int, p string) string {
				return p
			},
			"toyaml": func(yamlObj map[interface{}]interface{}) string {
				out, err := yaml.Marshal(yamlObj)
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
			"nindent": func(p string) string { return "not implemented" },
		})

		// TODO: ParseFiles for sure can get multiple templatefiles
		tmpl, err := myTemplate.ParseFiles(renderSettings.TemplateFile)

		if err != nil {
			panic(err)
		}

		renderSettings.Values = values

		err = tmpl.Execute(os.Stdout, renderSettings)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	renderCmd.Flags().StringVarP(&renderSettings.TemplateFile, "template", "t", "", "The template file to be rendered")
	renderCmd.Flags().StringVarP(&renderSettings.ValuesFile, "values", "f", "", "A values file to be rendered into the template")
}
