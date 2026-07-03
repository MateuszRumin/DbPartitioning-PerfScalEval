package main

import (
	"github.com/shirou/gopsutil/mem"
)

func getRAMUsage() (uint64, uint64, float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, err
	}
	return memInfo.Used, memInfo.Total, memInfo.UsedPercent, nil
}
