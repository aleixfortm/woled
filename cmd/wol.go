/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

// wolCmd represents the wol command
var wolCmd = &cobra.Command{
	Use:     "wol",
	Short:   "Broadcast a WOL packet to local network",
	Long:    `Send a WOL packet to the local network, broadcasted to IP 255.255.255.255, although this can be tweaked using 'config' command.`,
	Args:    cobra.ExactArgs(1),
	Example: `  woled wol PC-1`,
	Run: func(cmd *cobra.Command, args []string) {

		requestedDevice := args[0]

		// Create new type to save JSON data to config file
		type Device struct {
			Name       string `json:"name"`
			MACAddress string `json:"macAddress"`
		}

		// Read existing JSON file
		filePath := "config.json"
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Failed to read JSON file:", err)
			return
		}
		// Unmarshal existing JSON data into slice of type Device
		var deviceList []Device
		err = json.Unmarshal(fileData, &deviceList)
		if err != nil {
			fmt.Println("Failed to Unmarshall JSON data:", err)
		}

		var foundDevice Device
		for _, device := range deviceList {
			if device.Name == requestedDevice {
				foundDevice = device
				break
			}
		}

		if foundDevice == (Device{}) {
			fmt.Println("Error:", requestedDevice, "does not exist. Run 'list' command to see available devices.")
			return
		}

		// Create the magic packet by concatenating the MAC address 16 times
		magicPacket := strings.Repeat("\xff", 6) + strings.Repeat(foundDevice.MACAddress, 16)

		// The broadcast IP address and port used for WOL packets
		broadcastAddr := "255.255.255.255:9"

		// Resolve the UDP address for the broadcast IP and port
		udpAddr, err := net.ResolveUDPAddr("udp", broadcastAddr)
		if err != nil {
			fmt.Println("Failed to resolve UDP address:", err)
			return
		}

		// Create a UDP connection to the broadcast address
		conn, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			fmt.Println("Failed to dial UDP:", err)
			return
		}
		defer conn.Close()

		// Send the magic packet as a UDP broadcast
		_, err = conn.Write([]byte(magicPacket))
		if err != nil {
			fmt.Println("Failed to send WOL packet:", err)
			return
		}

		fmt.Println("WOL packet sent successfully!")
	},
}

func init() {
	rootCmd.AddCommand(wolCmd)

}
