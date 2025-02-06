package resourceusage

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func getCPUUsage() (float64, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	if len(percent) == 0 {
		return 0, fmt.Errorf("no CPU data available")
	}
	return percent[0], nil
}

func getRAMUsage() (uint64, uint64, float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, err
	}
	return memInfo.Used, memInfo.Total, memInfo.UsedPercent, nil
}

func getDiskUsage() (uint64, uint64, float64, error) {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	diskInfo, err := disk.Usage(path)
	if err != nil {
		return 0, 0, 0, err
	}
	return diskInfo.Used, diskInfo.Total, diskInfo.UsedPercent, nil
}

func getDiskIO() (uint64, uint64, error) {
	ioCounters, err := disk.IOCounters()
	if err != nil {
		return 0, 0, err
	}
	var reads, writes uint64
	for _, counter := range ioCounters {
		reads += counter.ReadCount
		writes += counter.WriteCount
	}
	return reads, writes, nil
}

func getNetworkIO() (uint64, uint64, error) {
	ioCounters, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}
	if len(ioCounters) == 0 {
		return 0, 0, fmt.Errorf("no network counters found")
	}
	return ioCounters[0].BytesSent, ioCounters[0].BytesRecv, nil
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func LogUsage() {
	interval := flag.Int("interval", 5, "Interval in seconds between measurements")
	flag.Parse()

	// Initialize previous values for delta calculations
	var (
		prevDiskReads, prevDiskWrites uint64
		prevNetSent, prevNetRecv      uint64
	)

	// Get initial disk IO and network stats
	if reads, writes, err := getDiskIO(); err == nil {
		prevDiskReads = reads
		prevDiskWrites = writes
	} else {
		log.Printf("Error getting initial disk IO: %v", err)
	}

	if sent, recv, err := getNetworkIO(); err == nil {
		prevNetSent = sent
		prevNetRecv = recv
	} else {
		log.Printf("Error getting initial network IO: %v", err)
	}

	// Initial CPU reading to avoid 0.0% on first interval
	_, _ = cpu.Percent(0, false)

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now().Format("2006-01-02 15:04:05")

		// CPU
		cpuPercent, err := getCPUUsage()
		if err != nil {
			log.Printf("Error getting CPU usage: %v", err)
		}

		// RAM
		ramUsed, ramTotal, ramPercent, err := getRAMUsage()
		if err != nil {
			log.Printf("Error getting RAM usage: %v", err)
		}

		// Disk Space
		diskUsed, diskTotal, diskPercent, err := getDiskUsage()
		if err != nil {
			log.Printf("Error getting disk usage: %v", err)
		}

		// Disk IO
		currentReads, currentWrites, err := getDiskIO()
		var diskReads, diskWrites uint64
		if err == nil {
			diskReads = currentReads - prevDiskReads
			diskWrites = currentWrites - prevDiskWrites
			prevDiskReads = currentReads
			prevDiskWrites = currentWrites
		} else {
			log.Printf("Error getting disk IO: %v", err)
		}

		// Network
		currentSent, currentRecv, err := getNetworkIO()
		var netSent, netRecv uint64
		if err == nil {
			netSent = currentSent - prevNetSent
			netRecv = currentRecv - prevNetRecv
			prevNetSent = currentSent
			prevNetRecv = currentRecv
		} else {
			log.Printf("Error getting network IO: %v", err)
		}

		// Display results
		fmt.Printf("[%s]\n", now)
		fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent)
		fmt.Printf("RAM Usage: %s / %s (%.2f%%)\n",
			formatBytes(ramUsed),
			formatBytes(ramTotal),
			ramPercent,
		)
		fmt.Printf("Disk Usage: %s / %s (%.2f%%)\n",
			formatBytes(diskUsed),
			formatBytes(diskTotal),
			diskPercent,
		)
		fmt.Printf("Disk IO: Reads=%d, Writes=%d\n", diskReads, diskWrites)
		fmt.Printf("Network: Sent=%s, Received=%s\n",
			formatBytes(netSent),
			formatBytes(netRecv),
		)
		fmt.Println("----------------------------------------")
	}
}
