package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/common/expfmt"
)

func fetchMetrics() (map[string]float64, error) {
	resp, err := http.Get("http://localhost:9104/metrics")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	parser := expfmt.TextParser{}
	families, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]float64)

	for name, mf := range families {
		for _, m := range mf.GetMetric() {
			if m.Gauge != nil {
				result[name] = m.Gauge.GetValue()
			}
			if m.Counter != nil {
				result[name] = m.Counter.GetValue()
			}
		}
	}

	return result, nil
}

func fetchData() {
	metrics, err := fetchMetrics()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// wybierasz interesujące metryki
	threads := metrics["mysql_global_status_threads_running"]
	queries := metrics["mysql_global_status_queries"]
	uptime := metrics["mysql_global_status_uptime"]

	// print
	fmt.Println("threads_running:", threads)
	fmt.Println("queries:", queries)
	fmt.Println("uptime:", uptime)
}
