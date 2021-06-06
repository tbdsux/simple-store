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
package cmd

import (
	"fmt"
	"log"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a group.",
	Long:  `Remove a group.`,
}

// collectionsCmd represents the collections command
var collectionsRemoveCmd = &cobra.Command{
	Use:   "collections",
	Short: "Remove collections group",
	Args:  cobra.ExactArgs(1),
	Long: `Remove collections group.

This will remove the key and its corresponding file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cols := args[0]

		internal.GetCols(cols)
		db := internal.DB()

		// remove the collection
		if err := db.RemoveCollection(cols); err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("\nSuccessfully removed collection: \033[1m%s\033[0m\n", cols)
	},
}

// storeCmd represents the store command
var storeRemoveCmd = &cobra.Command{
	Use:   "store",
	Short: "Remove store group",
	Args:  cobra.ExactArgs(1),
	Long: `Remove store group.

This will remove the key and its corresponding file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		store := args[0]

		internal.GetStore(store)
		db := internal.DB()

		// remove the store
		if err := db.RemoveStore(store); err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("\nSuccessfully removed store: \033[1m%s\033[0m\n", store)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.AddCommand(collectionsRemoveCmd, storeRemoveCmd)
}
