package cmd

import (
	"fmt"

	"github.com/TheBoringDude/simple-store/cmd/internal"
	"github.com/spf13/cobra"
)

var showKeysGroup bool
var listGroup bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simple-store",
	Short: "Store Stuff from the CLI.",
	Long: `Manage collections, configs, stores from your CLI.
	
Maybe you are comfortable in using your terminal than a gui.
Store and manage stuff from your cli or terminal.

`,
	Run: func(cmd *cobra.Command, args []string) {
		// show help if the --show* flag is not defined
		if !listGroup && !showKeysGroup {
			cmd.Help()
			return
		}

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
		if showKeysGroup {
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

// run
func Execute() {
	cobra.CheckErr(rootCmd.Execute())

}

func init() {
	rootCmd.Flags().BoolVar(&listGroup, "show-cols", false, "list the current collections created")
	rootCmd.Flags().BoolVar(&showKeysGroup, "show-stores", false, "list the current stores created")
}
