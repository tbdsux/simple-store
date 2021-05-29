package cmd

import (
	"fmt"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new group.",
	Long:  `Create a new group.`,
}

// collectionsCmd represents the collections command
var collectionsNewCmd = &cobra.Command{
	Use:   "collections",
	Short: "Create a new collections group",
	Args:  internal.ValidateArg,
	Long: `Create a new collections group.

If there are multiple args passed, it
will only get the first one.
	
EXAMPLE: store new collections my-collection`,
	Run: func(cmd *cobra.Command, args []string) {
		cols := args[0]

		db := internal.DB()
		db.Collections(cols)

		fmt.Printf("\nSuccessfully created a new collection: \033[1m%s\033[0m\n", cols)
	},
}

// storeCmd represents the store command
var storeNewCmd = &cobra.Command{
	Use:   "store",
	Short: "Create a new store group",
	Args:  internal.ValidateArg,
	Long: `Create a new store group.

If there are multiple args passed, it
will only get the first one.
	
EXAMPLE: store new store my-store`,
	Run: func(cmd *cobra.Command, args []string) {
		store := args[0]

		db := internal.DB()
		db.Store(store)

		fmt.Printf("\nSuccessfully created a new store: \033[1m%s\033[0m\n", store)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	/* sub-commands */
	newCmd.AddCommand(collectionsNewCmd, storeNewCmd)
}
