/*
COMMAND:
	store keys ...
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listKeysGroup bool
var keysGroup string

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage your store keys.",
	Long: `Manage your store keys.

EXAMPLE: 
  store keys add this-is-a-key --value=key-value --group=my-config
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("keys called")
	},
}

/*
COMMAND:
	store keys update ...
*/
var updateKeysCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a specific key from a store group.",
	Long: `Update the value of a key from a specific store group.
The key must exist from the store to be able to be updated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

/*
COMMAND:
	store keys remove ...
*/
var removeKeysCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a key from a store group.",
	Long: `Remove / Delete a key from a store group.
The key must exist from the store to be able to be removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

/*
COMMAND:
	store keys add ...
*/
var addKeysCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a key to the store group.",
	Long: `Add a key to the store group. 
The key should not already exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

	/* sub-functions */
	keysCmd.AddCommand(updateKeysCmd, removeKeysCmd, addKeysCmd)

	/* flags */
	keysCmd.PersistentFlags().StringVarP(&keysGroup, "group", "g", "", "the collection name / group")
	cobra.MarkFlagRequired(updateKeysCmd.PersistentFlags(), "group")
	cobra.MarkFlagRequired(removeKeysCmd.PersistentFlags(), "group")
	cobra.MarkFlagRequired(addKeysCmd.PersistentFlags(), "group")

	colsCmd.Flags().BoolVar(&listKeysGroup, "list", false, "list the current stores created")

}
