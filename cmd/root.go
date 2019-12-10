// Package CMD provides CLI utility
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	// homedir "github.com/mitchellh/go-homedir"
	// "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "syntax-helper",
	Short: "CLI to provide code syntax",
	Long: `This project is designed to help coders get access to quick and proper code syntax

The first language to be integrated is Golang (https://golang.org/).`,
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
