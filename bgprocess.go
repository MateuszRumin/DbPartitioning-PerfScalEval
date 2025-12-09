package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	scraperCancel  context.CancelFunc
	scraperMutex   sync.Mutex
	scraperRunning bool
)

func startBackgroundScraper() {
	scraperMutex.Lock()
	if scraperRunning {
		scraperMutex.Unlock()
		return
	}
	scraperRunning = true
	scraperMutex.Unlock()

	var ctx context.Context
	ctx, scraperCancel = context.WithCancel(context.Background())

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fetchData()
			case <-ctx.Done():
				fmt.Println("Scraper stopped")
				return
			}
		}
	}()
}

func stopBackgroundScraper() {
	scraperMutex.Lock()
	defer scraperMutex.Unlock()

	if !scraperRunning {
		return
	}

	scraperCancel()
	scraperRunning = false
}
