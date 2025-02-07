package main

import (
	"github.com/shirou/gopsutil/disk"
)

func getDiskIO() (map[string]uint64, map[string]uint64, map[string]uint64, map[string]uint64, map[string]uint64, map[string]uint64, error) {
	ioCounters, err := disk.IOCounters()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	readsCount := make(map[string]uint64)
	writesCount := make(map[string]uint64)
	readBytes := make(map[string]uint64)
	writeBytes := make(map[string]uint64)
	readTime := make(map[string]uint64)
	writeTime := make(map[string]uint64)

	for device, counter := range ioCounters {
		readsCount[device] = counter.ReadCount
		readBytes[device] = counter.ReadBytes
		readTime[device] = counter.ReadTime
		writesCount[device] = counter.WriteCount
		writeBytes[device] = counter.WriteBytes
		writeTime[device] = counter.WriteTime
	}

	return readsCount, writesCount, readBytes, writeBytes, readTime, writeTime, nil
}

// (1024 * 1024)
