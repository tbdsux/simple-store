package cmd

import (
	"fmt"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

// collectionsCmd represents the collections command
var collectionsCmd = &cobra.Command{
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

func init() {
	newCmd.AddCommand(collectionsCmd)
}
