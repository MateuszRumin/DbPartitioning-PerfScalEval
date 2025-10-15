// package resourceusage
package main

import (
	"flag"
	"fmt"
	"log"

	"time"
)

var (
	prevReadsCount, prevWritesCount map[string]uint64
	prevReadBytes, prevWriteBytes   map[string]uint64
	prevReadTime, prevWriteTime     map[string]uint64

	prevNetSent, prevNetRecv uint64
)

type (
	cpuResult struct {
		percent float64
		err     error
	}
	ramResult struct {
		used, total uint64
		percent     float64
		err         error
	}
	diskIOResult struct {
		readCount, writeCount map[string]uint64
		readBytes, writeBytes map[string]uint64
		readTime, writeTime   map[string]uint64
		err                   error
	}
	networkResult struct {
		sent, recv uint64
		err        error
	}
)

func main() {
	db, err := SetConnection()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	}
	defer db.Close()
	fmt.Println("Chose test sets")
	err = db.Ping()
	if err != nil {
		fmt.Println("Połączenie z bazą danych nie działa.")
	}

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
		readsCount, writesCount, readBytes, writeBytes, readTime, writeTime := make(map[string]uint64), make(map[string]uint64), make(map[string]uint64), make(map[string]uint64), make(map[string]uint64), make(map[string]uint64)
		if diskIORes.err == nil {

			for device := range diskIORes.readCount {
				// Dla odczytów
				readsCount[device] = diskIORes.readCount[device] - prevReadsCount[device]
				writesCount[device] = diskIORes.writeCount[device] - prevWritesCount[device]

				// Dla bajtów
				readBytes[device] = diskIORes.readBytes[device] - prevReadBytes[device]
				writeBytes[device] = diskIORes.writeBytes[device] - prevWriteBytes[device]

				// Dla czasu
				readTime[device] = diskIORes.readTime[device] - prevReadTime[device]
				writeTime[device] = diskIORes.writeTime[device] - prevWriteTime[device]
			}

			prevReadsCount, prevWritesCount = diskIORes.readCount, diskIORes.writeCount
			prevReadBytes, prevWriteBytes = diskIORes.readBytes, diskIORes.writeBytes
			prevReadTime, prevWriteTime = diskIORes.readTime, diskIORes.writeTime
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

		disk1Name, disk2Name := "C:", "F:"
		// Pobierz wartości dla dysku C:
		disk1Reads, disk1Writes := getDiskValue(readsCount, disk1Name), getDiskValue(writesCount, disk1Name)
		disk1RBytes, disk1WBytes := getDiskValue(readBytes, disk1Name), getDiskValue(writeBytes, disk1Name)
		disk1RTime, disk1WTime := getDiskValue(readTime, disk1Name), getDiskValue(writeTime, disk1Name)
		// Pobierz wartości dla dysku F:
		disk2Reads, disk2Writes := getDiskValue(readsCount, disk2Name), getDiskValue(writesCount, disk2Name)
		disk2RBytes, disk2WBytes := getDiskValue(readBytes, disk2Name), getDiskValue(writeBytes, disk2Name)
		disk2RTime, disk2WTime := getDiskValue(readTime, disk2Name), getDiskValue(writeTime, disk2Name)

		query := `
		INSERT INTO resource_usages (
    		timestamp, cpu_percent, ram_used, ram_percent,
    		disk1_name, disk1_reads, disk1_writes, disk1_read_bytes, disk1_write_bytes, disk1_read_time, disk1_write_time,
    		disk2_name, disk2_reads, disk2_writes, disk2_read_bytes, disk2_write_bytes, disk2_read_time, disk2_write_time,
    		network_sent, network_received) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		_, err := db.Exec(query,
			now, // timestamp
			cpuPercent,
			ramUsed,
			ramPercent,

			disk1Name, // disk1_name (C:)
			disk1Reads,
			disk1Writes,
			disk1RBytes,
			disk1WBytes,
			disk1RTime,
			disk1WTime,

			disk2Name, // disk2_name (F:)
			disk2Reads,
			disk2Writes,
			disk2RBytes,
			disk2WBytes,
			disk2RTime,
			disk2WTime,

			netSent,
			netRecv,
		)

		if err != nil {
			log.Printf("Błąd zapisu do bazy: %v", err)
		}
		// Display results
		fmt.Printf("[%s]\n",
			now)

		fmt.Printf("CPU Usage: %.10f %%\n", cpuPercent)

		fmt.Printf("RAM Usage: %d / %d (%.2f  %%)\n", ramUsed, ramTotal, ramPercent)

		log.Printf("Disk IO - Device %s || ReadsCount=%d, WritesCount=%d || ReadsBytes=%d, WritesBytes=%d || Read Time=%dms, Write Time=%dms",
			disk1Name, disk1Reads, disk1Writes, disk1RBytes, disk1WBytes, disk1RTime, disk1WTime)

		log.Printf("Disk IO - Device %s || ReadsCount=%d, WritesCount=%d || ReadsBytes=%d, WritesBytes=%d || Read Time=%dms, Write Time=%dms",
			disk2Name, disk2Reads, disk2Writes, disk2RBytes, disk2WBytes, disk2RTime, disk2WTime)

		fmt.Printf("Network: Sent=%d, Received=%d\n", netSent, netRecv)

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
