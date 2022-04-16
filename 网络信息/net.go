package main

import (
	"fmt"
	// "reflect"
	"github.com/shirou/gopsutil/net"
	//"net"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

// 获取网络信息 ： 网络地址、网关、防火墙状态及规则等信息
type NetworkDetails struct {
	Name      string                  // 接口名
	IpAddress syscall.IpAddressString //ip地址
	IpMask    syscall.IpMaskString    //子网掩码
	Gateway   syscall.IpAddressString //默认网关
}

func GetLocalIP() {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return
	}
	fmt.Println("index\tname\thardwareaddr\taddr")
	for i := 0; i < len(netInterfaces); i++ {
		fmt.Println("接口名：", netInterfaces[i].Name)
		fmt.Println("硬件地址：", netInterfaces[i].HardwareAddr)
		fmt.Println("ip4：", netInterfaces[i].Addrs[1].Addr)
		fmt.Println("ip6：", netInterfaces[i].Addrs[0].Addr)
		fmt.Println("===========")
	}
}

// 获取适配器列表
func getAdapterList() (*syscall.IpAdapterInfo, error) {
	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}

func localAddresses() []NetworkDetails {
	// 存储网络信息  数组
	networkDetails := []NetworkDetails{}

	// 返回系统网络接口的列表
	ifaces, err := net.Interfaces()
	if err != nil {
		return networkDetails
	}

	//返回一个*syscall.IpAdapterInfo链表，通过Next指向下一个网卡。
	aList, err := getAdapterList()
	if err != nil {
		return networkDetails
	}

	// 通过index匹配适配器和接口
	for _, ifi := range ifaces {
		for ai := aList; ai != nil; ai = ai.Next {
			index := ai.Index

			if ifi.Index == int(index) {
				// 获取适配器的IP链表
				ipl := &ai.IpAddressList
				// 获取适配器的网关链表
				gwl := &ai.GatewayList

				// fmt.Println(reflect.TypeOf(ipl.IpAddress))
				for ; ipl != nil; ipl = ipl.Next {
					netDetail := NetworkDetails{ifi.Name, ipl.IpAddress, ipl.IpMask, gwl.IpAddress}
					networkDetails = append(networkDetails, netDetail)
					//fmt.Printf("%s: ip地址:%s,  子网掩码:(%s),  默认网关:%s  \n", ifi.Name, ipl.IpAddress, ipl.IpMask, gwl.IpAddress)
				}

			}
		}
	}

	return networkDetails
}

func getState() {
	cmd := exec.Command("netsh", "advfirewall", "show", "allprofiles")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	//fmt.Println(cmd.Stdout,reflect.TypeOf(cmd.Stdout))
	//fmt.Println(cmd.Stderr,reflect.TypeOf(cmd.Stderr))

}

func main() {
	//GetLocalIP()
	Net := localAddresses()
	for i := 0; i < len(Net); i++ {
		fmt.Printf("%s : ip地址 : %s , 子网掩码:(%s),默认网关:%s \n -------\n", Net[i].Name, Net[i].IpAddress, Net[i].IpMask, Net[i].Gateway)
	}
	//getState()
}
