package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
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
