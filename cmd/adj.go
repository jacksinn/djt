// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/jacksinn/djt/text"
)

// adjCmd represents the adj command
var adjCmd = &cobra.Command{
	Use:   "adj",
	Short: "Add Adjectives",
	Long: `Adds adjectives to the adjectives.json`,
	Run: adjRun,
}

func adjRun(cmd *cobra.Command, args []string) {
	adjectives := []text.Adjective{}
	for _, x := range args {
		adjectives = append(adjectives, text.Adjective{Text:x})
	}
	err := text.SaveAdjectives("adjectives.json", adjectives)

	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println(adjectives)
}

func init() {
	RootCmd.AddCommand(adjCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adjCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adjCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
