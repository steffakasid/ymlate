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
	"github.com/spf13/cobra"
	"github.com/steffakasid/ymlate/pkg"
)

var tpl = pkg.Template{}

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "render the given template and it's values",
	Long: `This command will render the given template,
	merge values and directly output the result on the
	command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Render(tpl)
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
	renderCmd.Flags().StringVarP(&tpl.TemplateFile, "template", "t", "", "The template file to be rendered")
	renderCmd.Flags().StringVarP(&tpl.ValuesFile, "values", "f", "", "A values file to be rendered into the template")
}
