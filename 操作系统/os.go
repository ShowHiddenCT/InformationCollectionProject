package main

import (
	"fmt"
	// "runtime"
	"github.com/StackExchange/wmi"
	//"strings"
)

type operatingSystem struct {
	Name    string
	Version string
}

func getOSInfo() {
	var operatingsystem []operatingSystem
	err := wmi.Query("Select * from Win32_OperatingSystem", &operatingsystem)
	if err != nil {
		return
	}
	fmt.Printf("OS info =")
}

func main() {
	getOSInfo()
}
