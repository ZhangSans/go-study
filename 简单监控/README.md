使用Go语言实现一个简单的监控工具，可以通过以下步骤完成。这个工具可以监控系统的CPU、内存、磁盘等资源使用情况，并定期输出到控制台或日志文件中。

实现步骤：
导入必要的包：

runtime：获取Go运行时信息（如内存使用）。

time：用于定时任务。

github.com/shirou/gopsutil：第三方库，用于获取系统资源使用情况（如CPU、内存、磁盘等）。

定义监控函数：

监控CPU使用率。

监控内存使用情况。

监控磁盘使用情况。

定时任务：

使用time.Ticker定期执行监控任务。

输出结果：

将监控结果输出到控制台或日志文件。

示例代码

``` go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	// 设置监控间隔时间
	interval := 5 * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 开始监控
	for range ticker.C {
		monitorSystem()
	}
}

func monitorSystem() {
	// 获取CPU使用率
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("获取CPU使用率失败: %v", err)
	} else {
		fmt.Printf("CPU使用率: %.2f%%\n", cpuPercent[0])
	}

	// 获取内存使用情况
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("获取内存使用情况失败: %v", err)
	} else {
		fmt.Printf("内存使用情况: 已用 %.2f%%, 总共 %.2f GB, 已用 %.2f GB, 可用 %.2f GB\n",
			memInfo.UsedPercent,
			float64(memInfo.Total)/1024/1024/1024,
			float64(memInfo.Used)/1024/1024/1024,
			float64(memInfo.Available)/1024/1024/1024,
		)
	}

	// 获取磁盘使用情况
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Printf("获取磁盘使用情况失败: %v", err)
	} else {
		fmt.Printf("磁盘使用情况: 已用 %.2f%%, 总共 %.2f GB, 已用 %.2f GB, 可用 %.2f GB\n",
			diskInfo.UsedPercent,
			float64(diskInfo.Total)/1024/1024/1024,
			float64(diskInfo.Used)/1024/1024/1024,
			float64(diskInfo.Free)/1024/1024/1024,
		)
	}

	fmt.Println("----------------------------------------")
}
```

代码说明：
CPU监控：

使用cpu.Percent获取CPU使用率。

参数time.Second表示采样时间，false表示获取总的CPU使用率。

内存监控：

使用mem.VirtualMemory获取内存使用情况。

UsedPercent表示已用内存百分比，Total表示总内存，Used表示已用内存，Available表示可用内存。

磁盘监控：

使用disk.Usage获取磁盘使用情况。

参数"/"表示监控根目录的磁盘使用情况。

定时任务：

使用time.NewTicker创建一个定时器，每隔5秒执行一次监控任务。

输出结果：

将监控结果格式化输出到控制台。

运行结果示例：
plaintext
复制
CPU使用率: 15.20%
内存使用情况: 已用 45.30%, 总共 16.00 GB, 已用 7.25 GB, 可用 8.75 GB
磁盘使用情况: 已用 60.50%, 总共 500.00 GB, 已用 300.25 GB, 可用 199.75 GB
----------------------------------------
CPU使用率: 12.80%
内存使用情况: 已用 45.50%, 总共 16.00 GB, 已用 7.28 GB, 可用 8.72 GB
磁盘使用情况: 已用 60.50%, 总共 500.00 GB, 已用 300.25 GB, 可用 199.75 GB
----------------------------------------
扩展功能：
日志记录：

将监控结果写入日志文件，而不是输出到控制台。

使用log包或第三方日志库（如logrus）。

报警机制：

当CPU、内存或磁盘使用率超过阈值时，发送报警（如邮件、短信或调用API）。

Web界面：

使用net/http包创建一个简单的Web服务器，实时展示监控数据。

更多监控项：

监控网络流量、进程状态等。

依赖安装：
在运行代码之前，需要安装gopsutil库：

bash
复制
go get github.com/shirou/gopsutil/v3
通过以上代码，你可以实现一个简单的系统监控工具，并根据需求扩展功能。

