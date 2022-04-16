package main

import (
	"fmt"
	"log"
	"os/exec"
)

func listAllFW() {
	cmd := exec.Command("net", "start")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func listAllDK() {
	cmd := exec.Command("netstat", "-an")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func main() {
	fmt.Println("后台服务： ")
	listAllFW()
	fmt.Println("开放端口： ")
	listAllDK()

}
