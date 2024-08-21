/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/ddoffy/clipb/handlers"
	"github.com/spf13/cobra"
)

// retrieveCmd represents the retrieve command
var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve clipboard content from Redis",
	Long: `
		Retrieve clipboard content from Redis
		`,
	Run: func(cmd *cobra.Command, args []string) {
		rdb := handlers.GetRedisClient()

		// get clipboard content
		content, err := rdb.LPop("clipboard").Result()
		if err != nil {
			log.Fatalf("Error getting clipboard content from Redis: %v", err)
		}

		fmt.Println(content)

		// set clipboard content
		err = clipboard.WriteAll(content)
		if err != nil {
			log.Fatalf("Error setting clipboard content: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(retrieveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// retrieveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// retrieveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
