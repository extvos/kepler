package cli

import (
	"fmt"
	"github.com/extvos/kepler/service"
	"log"

	"github.com/spf13/cobra"
)

// newCmd represents the version command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Run a demo application powered by kepler.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be done...")
		listenAddr, _ := cmd.Flags().GetString("listen")
		log.Fatalln(service.Listen(listenAddr))
	},
}

func DemoCmd() *cobra.Command {
	return demoCmd
}

func init() {
	demoCmd.Flags().StringP("listen", "L", "127.0.0.1:8080", "Demo application listen address and port.")
	demoCmd.Flags().StringP("config", "C", "", "Configuration filename.")
	// rootCmd.AddCommand(newCmd)
}
