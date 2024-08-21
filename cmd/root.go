/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ddoffy/clipb/handlers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clipb",
	Short: "Store cliipboard content in Redis",
	Long: `
		clipb is a command line tool to store clipboard content in Redis.
		It can be used to store clipboard content in Redis and retrieve it later.
		`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		rdb := handlers.GetRedisClient()

		// get clipboard content
		content, err := handlers.GetClipboardContent()
		if err != nil {
			log.Fatalf("Error getting clipboard content: %v", err)
		}

		fmt.Println(content)

		// set clipboard content
		err = rdb.LPush("clipboard", content).Err()
		if err != nil {
			log.Fatalf("Error setting clipboard content in Redis: %v", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(retrieveCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clipb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
