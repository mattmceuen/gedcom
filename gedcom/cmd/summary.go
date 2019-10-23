/*
Copyright © 2019 Matt McEuen <madgin@madgin.net>

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
    "errors"
	"fmt"

	"github.com/spf13/cobra"
)

var gedFile string

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summarize the data in a gedcom file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Args: func (cmd *cobra.Command, args []string) error {
        if len(args) != 1 {
            return errors.New("summary requires the filename of the .ged to summarize")
        }
        gedFile = args[0]
        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf(".ged file to summarize: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
