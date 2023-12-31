package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [name] [MAC]",
	Short:   "Add device data",
	Long:    `Add your device to a local data file by specifying a name and the MAC address of the device.`,
	Args:    cobra.ExactArgs(2),
	Example: `  woled add "My Device" "00:11:22:33:44:55"`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create new type to add JSON data to data file
		type Device struct {
			Name       string `json:"name"`
			MACAddress string `json:"macAddress"`
		}

		// Check if MAC address is valid
		mac := args[1]
		match, _ := regexp.MatchString("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", mac)
		if !match {
			fmt.Println("Invalid MAC address format, use format XX:XX:XX:XX:XX:XX")
			return
		}

		// Assign given arguments to Device type
		newDevice := Device{
			Name:       args[0],
			MACAddress: mac,
		}

		filePath := "data.json"
		// Check if the file doesn't exist
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Create the file
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Failed to create JSON file:", err)
				return
			}
			defer file.Close()
		}

		// Read existing JSON file
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Failed to read JSON file:", err)
			return
		}

		// Unmarshal existing JSON data into slice of type Device
		var deviceList []Device
		err = json.Unmarshal(fileData, &deviceList)
		if err != nil && len(deviceList) > 0 {
			fmt.Println("Failed to Unmarshall JSON data:", err)
		}

		deviceList = append(deviceList, newDevice)

		// Convert device data to JSON
		newDeviceList, err := json.MarshalIndent(deviceList, "", "    ")
		if err != nil {
			fmt.Println("Failed to convert data to JSON:", err)
			return
		}

		// Write JSON data to the data file
		err = ioutil.WriteFile("data.json", newDeviceList, 0644)
		if err != nil {
			fmt.Println("Failed to write data file:", err)
			return
		}

		fmt.Println(newDevice.Name, "added successfully with MAC address", newDevice.MACAddress)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	// Customizing the "usage" display
	addCmd.SetUsageTemplate(`Usage:
  woled add [name] [MAC]

Arguments:
  [name]   string   Name of the device
  [MAC]    string   MAC address of the device

Examples:
  woled add PC-1 00:11:22:33:44:55
  woled add "My computer" 04:AF:32:4B:4C:95
	`)
}
