package main

//操作系统
import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	//"reflect"
)

func main() {
	OsType, family, OsVersion := getOSInfo2()
	fmt.Printf("%s\n", OsType)
	fmt.Printf("%s\n", family)
	fmt.Printf("%s\n", OsVersion)
}

func getOSInfo2() (string, string, string) {
	version, _ := host.KernelVersion()
	platform, family, version, _ := host.PlatformInformation()
	return platform, family, version
}
