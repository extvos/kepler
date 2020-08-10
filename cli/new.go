package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the version command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project based on kepler.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be done...")
	},
}

func NewCmd() *cobra.Command {
	return newCmd
}

func init() {
	//rootCmd.AddCommand(newCmd)
}
