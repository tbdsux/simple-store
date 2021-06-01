/*
COMMAND:
	store keys ...
*/

package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

var keysGroup string

// storeValuesCmd represents the keys command
var storeValuesCmd = &cobra.Command{
	Use:   "store",
	Short: "Manage your store keys.",
	Long: `Manage your store keys.

EXAMPLE: 
  store keys add this-is-a-key --value=key-value --group=my-config
	`,
}

/*
COMMAND:
	store keys update ...
*/
var updateStoreValue string
var updateStoreValueType string
var updateStoreCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a specific key from a store group.",
	Long: `Update the value of a key from a specific store group.
The key must exist from the store to be able to be updated.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.GetStore(keysGroup)

		var value interface{}
		switch updateStoreValueType {
		case "string":
			value = updateStoreValue
		case "bool":
			if v, err := strconv.ParseBool(updateStoreValue); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		case "int":
			if v, err := strconv.Atoi(updateStoreValue); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		case "float":
			if v, err := strconv.ParseFloat(updateStoreValue, 64); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		default:
			log.Fatalf("Unknown argument type: %s\n", updateStoreValueType)
		}

		if err := db.Update(args[0], value); err != nil {
			log.Fatalln(err)
		}
	},
}

/*
COMMAND:
	store keys remove ...
*/
var removeStoreCmd = &cobra.Command{
	Use:   "remove",
	Args:  cobra.ExactArgs(1),
	Short: "Remove a key from a store group.",
	Long: `Remove / Delete a key from a store group.
The key must exist from the store to be able to be removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.GetStore(keysGroup)

		if err := db.Remove(args[0]); err != nil {
			log.Fatalf("Error trying to remove key: %s from store: %s\n", args[0], keysGroup)
		}

		fmt.Printf("\nSuccessfully removed key: %s from store: %s\n", args[0], keysGroup)
	},
}

/*
COMMAND:
	store keys add ...
*/
var addStoreValue string
var addStoreValueType string
var addStoreCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "Add a key to the store group.",
	Long: `Add a key to the store group. 
The key should not already exist.

EXAMPLE: simple-store store add key --value=123 --type=int --group=my-store`,
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.GetStore(keysGroup)

		var value interface{}
		switch addStoreValueType {
		case "string":
			value = addStoreValue
		case "bool":
			if v, err := strconv.ParseBool(addStoreValue); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		case "int":
			if v, err := strconv.Atoi(addStoreValue); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		case "float":
			if v, err := strconv.ParseFloat(addStoreValue, 64); err != nil {
				log.Fatal(err)
			} else {
				value = v
			}
		default:
			log.Fatalf("Unknown argument type: %s\n", addStoreValueType)
		}

		if err := db.Set(args[0], value); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nSuccessfully set key: %s, value: %s -> store-group: %s", args[0], addStoreValue, keysGroup)
	},
}

/*
COMMAND:
	store keys list ...
*/
var listStoreCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "Remove a key from a store group.",
	Long: `Remove / Delete a key from a store group.
The key must exist from the store to be able to be removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.GetStore(keysGroup)

		fmt.Printf("\nValues in Store: %s\n", keysGroup)

		for key, value := range db.List() {
			fmt.Printf("  - %s : %v\n", key, value)
		}

	},
}

func init() {
	rootCmd.AddCommand(storeValuesCmd)

	/* sub-functions */
	storeValuesCmd.AddCommand(updateStoreCmd, removeStoreCmd, addStoreCmd, listStoreCmd)

	/* flags */
	storeValuesCmd.PersistentFlags().StringVarP(&keysGroup, "group", "g", "", "the collection name / group")
	cobra.MarkFlagRequired(updateStoreCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(removeStoreCmd.InheritedFlags(), "group")
	cobra.MarkFlagRequired(addStoreCmd.InheritedFlags(), "group")

	addStoreCmd.Flags().StringVarP(&addStoreValue, "value", "v", "", "the value of the key")
	addStoreCmd.Flags().StringVarP(&addStoreValueType, "type", "t", "string", "type of the value [string | int | float | bool]")
	addStoreCmd.MarkFlagRequired("value")

	updateStoreCmd.Flags().StringVarP(&updateStoreValue, "value", "v", "", "the value of the key")
	updateStoreCmd.Flags().StringVarP(&updateStoreValueType, "type", "t", "string", "type of the value [string | int | float | bool]")
	updateStoreCmd.MarkFlagRequired("value")
}
