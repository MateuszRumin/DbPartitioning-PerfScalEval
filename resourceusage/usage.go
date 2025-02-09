// package resourceusage
package main

import (
	"flag"
	"fmt"
	"log"

	"time"
)

var (
	prevReadsCount  map[string]uint64
	prevWritesCount map[string]uint64
	prevReadBytes   map[string]uint64
	prevWriteBytes  map[string]uint64
	prevReadTime    map[string]uint64
	prevWriteTime   map[string]uint64

	prevNetSent uint64
	prevNetRecv uint64
)

type (
	cpuResult struct {
		percent float64
		err     error
	}
	ramResult struct {
		used    uint64
		total   uint64
		percent float64
		err     error
	}
	diskIOResult struct {
		readCount  map[string]uint64
		writeCount map[string]uint64
		readBytes  map[string]uint64
		writeBytes map[string]uint64
		readTime   map[string]uint64
		writeTime  map[string]uint64
		err        error
	}
	networkResult struct {
		sent uint64
		recv uint64
		err  error
	}
)

func main() {
	interval := flag.Int("interval", 1, "Interval in seconds between measurements")
	flag.Parse()

	prevReadsCount, prevWritesCount, prevReadBytes, prevWriteBytes, prevReadTime, prevWriteTime, prevNetSent, prevNetRecv = initial()

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now().Format("2006-01-02 15:04:05")

		// Utwórz kanały
		cpuCh := make(chan cpuResult, 1)
		ramCh := make(chan ramResult, 1)
		diskIOCh := make(chan diskIOResult, 1)
		networkCh := make(chan networkResult, 1)

		// Uruchom równoległe zbieranie metryk
		go func() { p, err := getCPUUsage(); cpuCh <- cpuResult{p, err} }()
		go func() { u, t, p, err := getRAMUsage(); ramCh <- ramResult{u, t, p, err} }()
		go func() {
			rc, wc, rm, wm, rt, wt, err := getDiskIO()
			diskIOCh <- diskIOResult{rc, wc, rm, wm, rt, wt, err}
		}()
		go func() { s, r, err := getNetworkIO(); networkCh <- networkResult{s, r, err} }()

		// Zbierz wyniki
		cpuRes := <-cpuCh
		ramRes := <-ramCh
		diskIORes := <-diskIOCh
		networkRes := <-networkCh

		// Przetwórz wyniki
		if cpuRes.err != nil {
			log.Printf("CPU error: %v", cpuRes.err)
		}
		cpuPercent := cpuRes.percent

		if ramRes.err != nil {
			log.Printf("RAM error: %v", ramRes.err)
		}
		ramUsed, ramTotal, ramPercent := ramRes.used, ramRes.total, ramRes.percent

		// Przetwarzanie IO z synchronizacją
		var readsCount, writesCount, readBytes, writeBytes, readTime, writeTime map[string]uint64
		if diskIORes.err == nil {
			readsCount = make(map[string]uint64)
			writesCount = make(map[string]uint64)
			readBytes = make(map[string]uint64)
			writeBytes = make(map[string]uint64)
			readTime = make(map[string]uint64)
			writeTime = make(map[string]uint64)

			for device, current := range diskIORes.readCount {
				prev := prevReadsCount[device]
				readsCount[device] = current - prev
			}
			for device, current := range diskIORes.writeCount {
				prev := prevWritesCount[device]
				writesCount[device] = current - prev
			}
			for device, current := range diskIORes.readBytes {
				prev := prevReadBytes[device]
				readBytes[device] = current - prev
			}
			for device, current := range diskIORes.writeBytes {
				prev := prevWriteBytes[device]
				writeBytes[device] = current - prev
			}
			for device, current := range diskIORes.readTime {
				prev := prevReadTime[device]
				readTime[device] = current - prev
			}
			for device, current := range diskIORes.writeTime {
				prev := prevWriteTime[device]
				writeTime[device] = current - prev
			}

			prevReadsCount = diskIORes.readCount
			prevWritesCount = diskIORes.writeCount
			prevReadBytes = diskIORes.readBytes
			prevWriteBytes = diskIORes.writeBytes
			prevReadTime = diskIORes.readTime
			prevWriteTime = diskIORes.writeTime
		} else {
			log.Printf("Disk IO error: %v", diskIORes.err)
		}

		var netSent, netRecv uint64
		if networkRes.err == nil {

			netSent = networkRes.sent - prevNetSent
			netRecv = networkRes.recv - prevNetRecv
			prevNetSent = networkRes.sent
			prevNetRecv = networkRes.recv

		} else {
			log.Printf("Network error: %v", networkRes.err)
		}

		disk1Name := "C:"
		disk2Name := "F:"

		// Pobierz wartości dla dysku C:
		disk1Reads := getDiskValue(readsCount, disk1Name)
		disk1Writes := getDiskValue(writesCount, disk1Name)
		disk1RBytes := getDiskValue(readBytes, disk1Name)
		disk1WBytes := getDiskValue(writeBytes, disk1Name)
		disk1RTime := getDiskValue(readTime, disk1Name)
		disk1WTime := getDiskValue(writeTime, disk1Name)

		// Pobierz wartości dla dysku F:
		disk2Reads := getDiskValue(readsCount, disk2Name)
		disk2Writes := getDiskValue(writesCount, disk2Name)
		disk2RBytes := getDiskValue(readBytes, disk2Name)
		disk2WBytes := getDiskValue(writeBytes, disk2Name)
		disk2RTime := getDiskValue(readTime, disk2Name)
		disk2WTime := getDiskValue(writeTime, disk2Name)

		// Display results
		fmt.Printf("[%s]\n",
			now)

		fmt.Printf("CPU Usage: %.10f %%\n",
			cpuPercent)

		fmt.Printf("RAM Usage: %d / %d (%.2f  %%)\n",
			ramUsed, ramTotal, ramPercent)

		log.Printf("Disk IO - Device %s || ReadsCount=%d, WritesCount=%d || ReadsBytes=%d, WritesBytes=%d || Read Time=%dms, Write Time=%dms",
			disk1Name, disk1Reads, disk1Writes, disk1RBytes, disk1WBytes, disk1RTime, disk1WTime)

		log.Printf("Disk IO - Device %s || ReadsCount=%d, WritesCount=%d || ReadsBytes=%d, WritesBytes=%d || Read Time=%dms, Write Time=%dms",
			disk2Name, disk2Reads, disk2Writes, disk2RBytes, disk2WBytes, disk2RTime, disk2WTime)

		fmt.Printf("Network: Sent=%d, Received=%d\n",
			netSent, netRecv)

		fmt.Println("----------------------------------------")
	}

}

//formatBytes(readBytes[device]),  %s

func getDiskValue(m map[string]uint64, device string) uint64 {
	if val, ok := m[device]; ok {
		return val
	}
	return 0
}
