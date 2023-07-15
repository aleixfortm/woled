/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "save",
	Short: "Save device configuration",
	Long:  `Save your device to a local config file by specifying a name and the MAC address of the device.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create new type to save JSON data to config file
		type Device struct {
			Name       string `json:"name"`
			MACAddress string `json:"macAddress"`
		}

		// Check if two arguments have been given in the command, else cancel
		if len(args) != 2 {
			fmt.Println("Provide a name and a MAC address and try again")
			return
		}

		// Assign given arguments to Device type
		device := Device{
			Name:       args[0],
			MACAddress: args[1],
		}

		fmt.Println("Name:", device.Name)
		fmt.Println("MAC:", device.MACAddress)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
