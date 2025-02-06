package resourceusage

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

var (
	diskIOMutex    sync.Mutex
	prevDiskReads  uint64
	prevDiskWrites uint64

	netMutex    sync.Mutex
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
	diskResult struct {
		used    uint64
		total   uint64
		percent float64
		err     error
	}
	diskIOResult struct {
		reads  uint64
		writes uint64
		err    error
	}
	networkResult struct {
		sent uint64
		recv uint64
		err  error
	}
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
	interval := flag.Int("interval", 1, "Interval in seconds between measurements")
	flag.Parse()

	// Get initial disk IO and network stats
	reads, writes, err := getDiskIO()
	if err != nil {
		log.Fatalf("Initial disk IO error: %v", err) // zatrzymaj program
	}
	prevDiskReads = reads
	prevDiskWrites = writes

	sent, recv, err := getNetworkIO()
	if err != nil {
		log.Fatalf("Initial network error: %v", err) // zatrzymaj program
	}
	prevNetSent = sent
	prevNetRecv = recv

	// Initial CPU reading to avoid 0.0% on first interval
	_, _ = cpu.Percent(0, false)

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now().Format("2006-01-02 15:04:05")

		// Utwórz kanały
		cpuCh := make(chan cpuResult, 1)
		ramCh := make(chan ramResult, 1)
		diskCh := make(chan diskResult, 1)
		diskIOCh := make(chan diskIOResult, 1)
		networkCh := make(chan networkResult, 1)

		// Uruchom równoległe zbieranie metryk
		go func() {
			p, err := getCPUUsage()
			cpuCh <- cpuResult{p, err}
		}()

		go func() {
			u, t, p, err := getRAMUsage()
			ramCh <- ramResult{u, t, p, err}
		}()

		go func() {
			u, t, p, err := getDiskUsage()
			diskCh <- diskResult{u, t, p, err}
		}()

		go func() {
			r, w, err := getDiskIO()
			diskIOCh <- diskIOResult{r, w, err}
		}()

		go func() {
			s, r, err := getNetworkIO()
			networkCh <- networkResult{s, r, err}
		}()

		// Zbierz wyniki
		cpuRes := <-cpuCh
		ramRes := <-ramCh
		diskRes := <-diskCh
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

		if diskRes.err != nil {
			log.Printf("Disk error: %v", diskRes.err)
		}
		diskUsed, diskTotal, diskPercent := diskRes.used, diskRes.total, diskRes.percent

		// Przetwarzanie IO z synchronizacją
		var diskReads, diskWrites uint64
		if diskIORes.err == nil {

			diskReads = diskIORes.reads - prevDiskReads
			diskWrites = diskIORes.writes - prevDiskWrites
			prevDiskReads = diskIORes.reads
			prevDiskWrites = diskIORes.writes

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
