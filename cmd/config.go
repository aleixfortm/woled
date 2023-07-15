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

type Config struct {
	Port int `json:"port"`
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show and tweak default app settings",
	Long:  `Show and tweak default application settings such as broadcast IP address and UDP port.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read the current configuration from the file
		config, err := readConfigFromFile()
		if err != nil {
			fmt.Println("Failed to read config file:", err)
			return
		}

		// Check if the --port flag is provided
		portFlag, _ := cmd.Flags().GetInt("port")
		if portFlag != 0 {
			// Update the port value in the configuration
			config.Port = portFlag

			// Save the updated configuration to the file
			err = saveConfigToFile(config)
			if err != nil {
				fmt.Println("Failed to save config file:", err)
				return
			}

			fmt.Println("Configuration updated successfully!")
		} else {
			// Display the current configuration
			fmt.Println("Current Configuration:")
			fmt.Println("Port:", config.Port)
		}
	},
}

func readConfigFromFile() (*Config, error) {
	filePath := "config.json"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfigToFile(config *Config) error {
	filePath := "config.json"
	fileData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().IntP("port", "p", 0, "Set new default UDP port, recommended 7(Default) or 9")
}
