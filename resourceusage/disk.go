// package resourceusage
package main

import (
	"runtime"

	"github.com/shirou/gopsutil/disk"
)

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
