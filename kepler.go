/*
Copyright © 2020 Mingcai SHEN
*/
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/extvos/kepler/cli"
)

//var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kepler",
	Short: "A quick framework of Golang...",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("verbose", "V", false, "Toggle debug messages")
}


func init() {
	//rootCmd.AddCommand(cmd.InitCmd(), cmd.VersionCmd(), cmd.ServeCmd())
	rootCmd.AddCommand(cli.NewCmd())
}

func main() {
	Execute()
}
