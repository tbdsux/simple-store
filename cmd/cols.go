/*
COMMAND:
	store cols ...
*/

package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

var colsGroup string

// colsCmd represents the cols command
var colsCmd = &cobra.Command{
	Use:   "cols",
	Short: "Manage your collections.",
	Args:  cobra.NoArgs,
	Long: `Manage your collections.

EXAMPLE: 
  store cols add https://www.google.com --group=websites
	`,
}

/*
COMMAND:
	store cols find ...
*/
var findColsCmd = &cobra.Command{
	Use:   "find",
	Args:  cobra.ExactArgs(1),
	Short: "Find a value from the collection.",
	Long: `Find a value from the collection.
It will return matching strings also not only the exact.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.GetCols(colsGroup)

		result := db.FindAll(args[0])

		fmt.Printf("\nSearch for: %s -> group: %s\n", args[0], colsGroup)

		for _, i := range result {
			fmt.Printf(" > %s \n", i)
		}
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
The value must exist from the collection and should be exact.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		// THIS IS A TODO:
		// `minidb` has no function Collections.Remove() yet
	},
}

/*
COMMAND:
	store cols add ...
*/
var addColsArgType string
var addColsCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "Add a value to the collection.",
	Long: `Add a value to the collection.
The value must exist from the collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.GetCols(colsGroup)

		var value interface{}
		// get check value type
		switch addColsArgType {
		case "string":
			value = args[0]
		case "bool":
			if v, err := strconv.ParseBool(args[0]); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		case "int":
			if v, err := strconv.Atoi(args[0]); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		case "float":
			if v, err := strconv.ParseFloat(args[0], 64); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		default:
			log.Fatalf("Unknown argument type: %s\n", addColsArgType)
		}

		// push the first arg
		db.Push(value)

		fmt.Printf("\nSuccessfully added item: `%s` -> group: `%s`\n", args[0], colsGroup)
	},
}

/*
COMMAND:
	store cols list ...
*/
var listColsOneline bool
var listColsCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "Add a value to the collection.",
	Long: `Add a value to the collection.
The value must exist from the collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.GetCols(colsGroup)

		lists := db.List()

		fmt.Printf("\nValues in Collection: %s\n", colsGroup)
		if listColsOneline {
			for _, i := range lists {
				fmt.Printf("  - %s\n", i)
			}
			return
		}

		fmt.Println(lists)
	},
}

func init() {
	rootCmd.AddCommand(colsCmd)

	/* sub-functions */
	colsCmd.AddCommand(findColsCmd, removeColsCmd, addColsCmd, listColsCmd)

	/* flags */
	colsCmd.PersistentFlags().StringVarP(&colsGroup, "group", "g", "", "the collection name / group")
	cobra.MarkFlagRequired(findColsCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(removeColsCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(addColsCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(listColsCmd.InheritedFlags(), "group")

	listColsCmd.Flags().BoolVar(&listColsOneline, "oneline", false, "print each values in each line")
	addColsCmd.Flags().StringVarP(&addColsArgType, "type", "t", "string", "the type of the argument [string | int | float | bool]")
}
