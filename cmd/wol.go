/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

// wolCmd represents the wol command
var wolCmd = &cobra.Command{
	Use:   "wol",
	Short: "Broadcast a WOL packet to local network",
	Long:  `Send a WOL packet to the local network, broadcasted to IP 192.168.1.255, although this can be tweaked using 'config' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// The MAC address of the target device in the format "XX:XX:XX:XX:XX:XX"
		macAddress := "00:11:22:33:44:55"

		// Create the magic packet by concatenating the MAC address 16 times
		magicPacket := strings.Repeat("\xff", 6) + strings.Repeat(macAddress, 16)

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
