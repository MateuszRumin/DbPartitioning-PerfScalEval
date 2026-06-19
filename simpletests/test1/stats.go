package main

import (
	"sort"
	"time"
)

func percentile(durations []time.Duration, p float64) time.Duration {
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	idx := int(float64(len(durations)) * p)
	if idx >= len(durations) {
		idx = len(durations) - 1
	}
	return durations[idx]
}

func summarize(results []QueryResult) map[string]any {
	group := make(map[string][]time.Duration)

	for _, r := range results {
		group[r.Query] = append(group[r.Query], r.Duration)
	}

	out := make(map[string]any)

	for q, arr := range group {
		out[q] = map[string]any{
			"p50": percentile(arr, 0.50),
			"p95": percentile(arr, 0.95),
			"p99": percentile(arr, 0.99),
		}
	}

	return out
}
