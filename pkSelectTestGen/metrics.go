package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

type Sample struct {
	Time     time.Time
	Duration time.Duration
}

type Metrics struct {
	mu      sync.Mutex
	samples []Sample
}

func (m *Metrics) Add(duration time.Duration) {
	m.mu.Lock()
	m.samples = append(m.samples, Sample{
		Time:     time.Now(),
		Duration: duration,
	})
	m.mu.Unlock()
}

var metrics = &Metrics{}

func percentile(values []float64, p float64) float64 {
	if len(values) == 0 {
		return 0
	}

	sort.Float64s(values)

	idx := int(math.Ceil(float64(len(values))*p/100.0)) - 1

	if idx < 0 {
		idx = 0
	}

	if idx >= len(values) {
		idx = len(values) - 1
	}

	return values[idx]
}

func PrintMinuteReport() {

	type Bucket struct {
		Durations []float64
		Count     int
	}

	metrics.mu.Lock()
	defer metrics.mu.Unlock()

	if len(metrics.samples) == 0 {
		return
	}

	startTime := metrics.samples[0].Time

	buckets := make(map[int]*Bucket)

	for _, sample := range metrics.samples {

		minute := int(sample.Time.Sub(startTime).Minutes())

		if buckets[minute] == nil {
			buckets[minute] = &Bucket{}
		}

		buckets[minute].Count++

		buckets[minute].Durations = append(
			buckets[minute].Durations,
			float64(sample.Duration.Microseconds())/1000.0,
		)
	}

	fmt.Printf(
		"%-8s %-10s %-10s %-10s %-10s %-10s %-10s\n",
		"Minute",
		"Queries",
		"QPS",
		"P90",
		"P95",
		"P99",
		"MAX",
	)

	for minute := 0; minute < len(buckets); minute++ {

		b := buckets[minute]

		if b == nil {
			continue
		}

		max := 0.0
		for _, d := range b.Durations {
			if d > max {
				max = d
			}
		}

		qps := float64(b.Count) / 60.0

		fmt.Printf(
			"%-8d %-10d %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f\n",
			minute,
			b.Count,
			qps,
			percentile(append([]float64(nil), b.Durations...), 90),
			percentile(append([]float64(nil), b.Durations...), 95),
			percentile(append([]float64(nil), b.Durations...), 99),
			max,
		)
	}
}
