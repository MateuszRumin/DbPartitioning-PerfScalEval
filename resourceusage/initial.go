package main

import (
	"log"

	"github.com/shirou/gopsutil/cpu"
)

func initial() (uint64, uint64, uint64, uint64) {
	// Get initial disk IO and network stats
	reads, writes, err := getDiskIO()
	if err != nil {
		log.Fatalf("Initial disk IO error: %v", err) // zatrzymaj program
	}
	//prevDiskReads = reads
	//prevDiskWrites = writes

	sent, recv, err := getNetworkIO()
	if err != nil {
		log.Fatalf("Initial network error: %v", err) // zatrzymaj program
	}
	//prevNetSent = sent
	//prevNetRecv = recv

	// Initial CPU reading to avoid 0.0% on first interval
	_, _ = cpu.Percent(0, false)

	return reads, writes, sent, recv
}
