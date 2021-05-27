package internal

import (
	"errors"

	"github.com/spf13/cobra"
)

// a custom arg validator for command
func ValidateArg(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires one arg")
	}
	return nil
}
