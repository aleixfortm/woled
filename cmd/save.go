/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
		newDevice := Device{
			Name:       args[0],
			MACAddress: args[1],
		}

		// Read existing JSON file
		filePath := "config.json"
		configData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Failed to read JSON file:", err)
			return
		}

		// Unmarshal existing JSON data into slice of type Device
		var deviceList []Device
		err = json.Unmarshal(configData, &deviceList)
		if err != nil {
			fmt.Println("Failed to Unmarshall JSON data:", err)
		}

		deviceList = append(deviceList, newDevice)

		// Convert device data to JSON
		newDeviceList, err := json.MarshalIndent(deviceList, "", "    ")
		if err != nil {
			fmt.Println("Failed to convert data to JSON:", err)
			return
		}

		// Write JSON data to the configuration file
		err = ioutil.WriteFile("config.json", newDeviceList, 0644)
		if err != nil {
			fmt.Println("Failed to write configuration file:", err)
			return
		}

		fmt.Println("Configuration file saved successfully!")

		fmt.Println("Name:", newDevice.Name)
		fmt.Println("MAC:", newDevice.MACAddress)
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
