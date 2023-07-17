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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove saved device",
	Long:    `Remove a saved device by calling this command, with the name of the device as argument. Use command "list" to see list of saved devices.`,
	Args:    cobra.ExactArgs(1),
	Example: `  woled remove PC-2`,
	Run: func(cmd *cobra.Command, args []string) {
		deviceToRemove := args[0]

		// Create new type to save JSON data to config file
		type Device struct {
			Name       string `json:"name"`
			MACAddress string `json:"macAddress"`
		}

		// Read existing JSON file
		filePath := "config.json"
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Failed to read data file:", err)
			return
		}

		// Unmarshal existing JSON data into slice of type Device
		var deviceList []Device
		err = json.Unmarshal(fileData, &deviceList)
		if err != nil {
			fmt.Println("Failed to Unmarshall JSON data:", err)
		}

		// Find the index to remove
		for i, d := range deviceList {
			if d.Name == deviceToRemove {
				// New list is the old list sliced so that the current index is excluded
				deviceList = append(deviceList[:i], deviceList[i+1:]...)

				// Convert device data to JSON
				DeviceList, err := json.MarshalIndent(deviceList, "", "    ")
				if err != nil {
					fmt.Println("Failed to convert data to JSON:", err)
					return
				}

				// Write JSON data to the data file
				err = ioutil.WriteFile("config.json", DeviceList, 0644)
				if err != nil {
					fmt.Println("Failed to write data file:", err)
					return
				}

				fmt.Println(deviceToRemove, "has been successfully deleted.")

				return
			}
		}

		fmt.Println(deviceToRemove, "was not found. Run command 'list' to check existing devices.")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	// Customizing the "usage" display
	removeCmd.SetUsageTemplate(`Usage:
	woled remove [name]

Arguments:
	[name]   string   Name of the device

Examples:
	woled remove PC-1
	woled remove "My computer"
	`)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
