package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the version command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Run a demo application powered by kepler.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be done...")
	},
}

func DemoCmd() *cobra.Command {
	return demoCmd
}

func init() {
	demoCmd.Flags().StringP("listen","L","127.0.0.1:8080","Demo application listen address and port.")
	//rootCmd.AddCommand(newCmd)
}
