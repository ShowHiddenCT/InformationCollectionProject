package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"strconv"
	//	"reflect"
)

func main() {
	pi, _ := process.Pids()

	// fmt.Printf("进程号\t\t进程名\t\t占用资源\n")
	for i := 0; i < len(pi); i++ {
		p, _ := process.NewProcess(pi[i])
		strPid := strconv.FormatInt(int64(pi[i]), 10)
		pm, _ := p.MemoryPercent()
		pn, _ := p.Name()
		//	fmt.Println(pn)
		fmt.Printf("进程号：%s,  进程名：%s,占用资源：%f%c\n", strPid, pn, pm, '%')

	}

}
