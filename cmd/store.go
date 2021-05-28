/*
COMMAND:
	store keys ...
*/

package cmd

import (
	"fmt"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

var listKeysGroup bool
var keysGroup string

// storeValuesCmd represents the keys command
var storeValuesCmd = &cobra.Command{
	Use:   "store",
	Short: "Manage your store keys.",
	Long: `Manage your store keys.

EXAMPLE: 
  store keys add this-is-a-key --value=key-value --group=my-config
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if listKeysGroup {
			db := internal.DB()
			listStores := db.ListStores()

			if len(listStores) < 1 {
				fmt.Println("\nYou currently have no stores.")
				return
			}

			fmt.Println("\nStores:")
			for _, i := range listStores {
				fmt.Printf("  - %s\n", i)
			}
		}
	},
}

/*
COMMAND:
	store keys update ...
*/
var updateStoreCmd = &cobra.Command{
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
var removeStoreCmd = &cobra.Command{
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
var addStoreCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a key to the store group.",
	Long: `Add a key to the store group. 
The key should not already exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(storeValuesCmd)

	/* sub-functions */
	storeValuesCmd.AddCommand(updateStoreCmd, removeStoreCmd, addStoreCmd)

	/* flags */
	storeValuesCmd.PersistentFlags().StringVarP(&keysGroup, "group", "g", "", "the collection name / group")
	cobra.MarkFlagRequired(updateStoreCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(removeStoreCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(addStoreCmd.InheritedFlags(), "group")

	storeValuesCmd.Flags().BoolVar(&listKeysGroup, "list", false, "list the current stores created")

}
