package main

import "go-wol/cmd"

// Create new type to save JSON data to config file
type Device struct {
	Name       string `json:"name"`
	MACAddress string `json:"macAddress"`
}

func main() {
	cmd.Execute()
}
