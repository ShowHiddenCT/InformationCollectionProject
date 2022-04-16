package main

import (
	"fmt"
	"log"
	"os/exec"
)

func listAllProcess() {
	cmd := exec.Command("net", "start")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func main() {
	listAllProcess()
}
