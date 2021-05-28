package cmd

import (
	"fmt"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Create a new store group",
	Args:  internal.ValidateArg,
	Long: `Create a new store group.

If there are multiple args passed, it
will only get the first one.
	
EXAMPLE: store new store my-store`,
	Run: func(cmd *cobra.Command, args []string) {
		cols := args[0]

		db := internal.DB()
		db.Collections(cols)

		fmt.Printf("\nSuccessfully created a new store: \033[1m%s\033[0m\n", cols)
	},
}

func init() {
	newCmd.AddCommand(storeCmd)
}
