package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		monitor()
	}
	defer ticker.Stop()

}

// 开始监控
func monitor() {
	//获取CPU使用率
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		fmt.Println("获取CPU监控信息失败", err)
	} else {
		fmt.Printf("获取cpu运行数据：%.2f%%,", cpuPercent[0])
	}
	//获取内存使用情况
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("获取内存监控信息失败", err)
	} else {
		fmt.Printf("内存使用情况，已用 %.2f%%, 总共%.2f GB，已用 %.2f GB, 可用 %.2f BG \n",
			memInfo.UsedPercent,
			float64(memInfo.Total)/1024/1024/1024,
			float64(memInfo.Used)/1024/1024/1024,
			float64(memInfo.Available)/1024/1024/1024,
		)
	}
	//获取磁盘运行情况
	diskInfo, err := disk.Usage("/")
	if err != nil {
		fmt.Println("获取磁盘监控信息失败", err)
	} else {
		fmt.Printf("磁盘使用情况，已用 %.2f%%, 总共%.2f GB，已用 %.2f GB, 可用 %.2f BG \n",
			diskInfo.UsedPercent,
			float64(diskInfo.Total)/1024/1024/1024,
			float64(diskInfo.Used)/1024/1024/1024,
			float64(diskInfo.Free)/1024/1024/1024,
		)
	}
}
