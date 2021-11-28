package cmd

/*
   Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

import (
	"fmt"
	"log"

	"github.com/jpoirier55/tri/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo",
	Long: `Add a new line of todo`,
	Run: addNew,
}

func addNew (cmd *cobra.Command, args []string) {
	var items []todo.Item
	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}
	err := todo.SaveItems( ".tridos.json", items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func addRun (cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(".tridos.json")
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
