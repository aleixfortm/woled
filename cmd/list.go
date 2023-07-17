package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of saved devices",
	Long:  `list command displays a list of saved devices from data.json file`,
	Run: func(cmd *cobra.Command, args []string) {

		type Device struct {
			Name       string `json:"name"`
			MACAddress string `json:"macAddress"`
		}

		// Check whether file exists
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

		fmt.Println(len(fileData))

		// Stop function if there is no data to show
		if len(fileData) == 0 {
			fmt.Println("Your device list is empty. Run 'add' command to add a device.")
			return
		}

		// Unmarshal existing JSON data into slice of type Device
		var deviceList []Device
		err = json.Unmarshal(fileData, &deviceList)
		if err != nil {
			fmt.Println("Failed to Unmarshall JSON data:", err)
		}

		fmt.Println("Device list:")
		for i, deviceData := range deviceList {
			fmt.Println(" ", i, "", deviceData.Name, "", deviceData.MACAddress)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Customizing the "usage" display
	listCmd.SetUsageTemplate(`Usage:
	woled list

Arguments:
	None

Examples:
	woled list
	`)
}
