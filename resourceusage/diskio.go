package main

import (
	"github.com/shirou/gopsutil/disk"
)

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
