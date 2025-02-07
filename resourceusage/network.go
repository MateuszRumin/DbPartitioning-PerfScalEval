// package resourceusage
package main

import (
	"fmt"

	"github.com/shirou/gopsutil/net"
)

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
