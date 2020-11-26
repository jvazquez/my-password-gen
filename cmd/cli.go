/*
Copyright © 2020 Jorge Omar Vazquez

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
	"fmt"
	"github.com/my-password-gen/internal"
	"github.com/spf13/cobra"
	"log"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Generate a password from command line",
	Long: `Generate a password from command line with
multiple options. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		useSymbols, _ := cmd.Flags().GetBool("symbols")
		seed, _ := cmd.Flags().GetInt("words")
		fixed, _ := cmd.Flags().GetBool("fixed")
		defaultSymbol, _ := cmd.Flags().GetString("defaultsymbol")

		passwordGenerator := internal.PasswordGenerator{UseSymbols: useSymbols, Seed: seed, UseFixedSymbol: fixed}

		if fixed == true && len(defaultSymbol) == 0 {
			log.Fatal("You are using a fixed symbol but did not provide the default")
		} else if fixed == true && len(defaultSymbol) != 0 {
			passwordGenerator.Separator = defaultSymbol
		}

		fmt.Printf("%s\n", passwordGenerator.Execute())
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cliCmd.Flags().BoolP("symbols", "s", false, "Use symbols instead of spaces")
	cliCmd.Flags().Int("words", internal.DefaultWords, "How much words we will use")
	cliCmd.Flags().BoolP("fixed", "f", false, "Use fixed symbol")
	cliCmd.Flags().String("defaultsymbol", "", "Use the provided default symbol")
}
