package main

import (
	`flag`
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	hostname := flag.String("hostname", "", "测试的域名")
	startPort := flag.Int("start-port", 80, "开始扫描端口号")
	endPort := flag.Int("end-port", 100, "结束扫描端口号")
	timeout := flag.Duration("timeout", time.Millisecond*200, "超时时间")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			open := isOpen(*hostname, p, *timeout)
			if open {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Printf("打开的端口 : %v\n", ports)
}

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		conn.Close()
		return true
	} else {
		return false
	}
}
