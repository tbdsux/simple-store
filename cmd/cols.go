/*
COMMAND:
	store cols ...
*/

package cmd

import (
	"fmt"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

var colsGroup string
var listGroup bool

// colsCmd represents the cols command
var colsCmd = &cobra.Command{
	Use:   "cols",
	Short: "Manage your collections.",
	Args:  cobra.NoArgs,
	Long: `Manage your collections.

EXAMPLE: 
  store cols add https://www.google.com --group=websites
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if listGroup {
			db := internal.DB()
			listCols := db.ListCollections()

			if len(listCols) < 1 {
				fmt.Println("\nYou currently have no collections.")
				return
			}

			fmt.Println("\nCollections:")
			for _, i := range listCols {
				fmt.Printf("  - %s\n", i)
			}
		}
	},
}

/*
COMMAND:
	store cols find ...
*/
var findColsCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a value from the collection.",
	Long: `Find a value from the collection.
It will return matching strings also not only the exact.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find called")
	},
}

/*
COMMAND:
	store cols remove ...
*/
var removeColsCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a value from the collection.",
	Long: `Remove a value from the collection.
The value must exist from the collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

/*
COMMAND:
	store cols add ...
*/
var addColsCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a value to the collection.",
	Long: `Add a value to the collection.
The value must exist from the collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	rootCmd.AddCommand(colsCmd)

	/* sub-functions */
	colsCmd.AddCommand(findColsCmd, removeColsCmd, addColsCmd)

	/* flags */
	colsCmd.PersistentFlags().StringVarP(&colsGroup, "group", "g", "", "the collection name / group")
	cobra.MarkFlagRequired(findColsCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(removeColsCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(addColsCmd.InheritedFlags(), "group")

	colsCmd.Flags().BoolVar(&listGroup, "list", false, "list the current collections created")
}
