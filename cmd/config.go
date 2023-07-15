/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var portValue string
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show and tweak default app settings",
	Long:  `Show and tweak default application settings such as broadcast IP address and UDP port.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Port:", portValue)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	rootCmd.Flags().StringVar(&portValue, "port", "p", "Set new default UDP port, recommended 7(Default) or 9")
}
