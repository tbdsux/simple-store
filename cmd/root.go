package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simple-store",
	Short: "Store Stuff from the CLI.",
	Long: `Manage collections, configs, stores from your CLI.
	
Maybe you are comfortable in using your terminal than a gui.
Store and manage stuff from your cli or terminal.

`,
}

// run
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
