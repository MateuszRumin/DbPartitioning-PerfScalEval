package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv" 
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Zmienna globalna do przechowywania metryk
var (
	// Rejestracja jednej metryki z etykietą `source`
	customMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "custom_exporter_metrics",
			Help: "Metrics from mysql_exporter and windows_exporter combined",
		},
		[]string{"source", "metric_name"},
	)
)

// Funkcja do pobierania metryk z eksportera
func fetchMetrics(url string, source string) {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Błąd pobierania metryk z", url, ":", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Błąd odczytu danych:", err)
		return
	}

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		metricName := parts[0]
		valueStr := parts[1]

		// Tworzymy pełną nazwę metryki z etykietą `source`
		fullMetricName := fmt.Sprintf("%s_%s", source, metricName)

		// Parsowanie wartości metryki
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println("Błąd parsowania wartości dla metryki:", fullMetricName)
			continue
		}

		// Zapis metryki z etykietą `source`
		customMetrics.WithLabelValues(source, metricName).Set(value)
	}
}

// Funkcja, która uruchamia pobieranie metryk w tle
func startMetricCollection() {
	for {
		// Pobieranie metryk z obu exporterów
		fetchMetrics("http://localhost:9104/metrics", "mysql_exporter")
		fetchMetrics("http://localhost:9182/metrics", "windows_exporter")
		time.Sleep(5 * time.Second) // Co 5 sekund
	}
}

func main() {
	// Rejestracja metryk w Prometheusie
	prometheus.MustRegister(customMetrics)

	// Uruchomienie procesu zbierania metryk
	go startMetricCollection()

	// Udostępnienie metryk na porcie 9105
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Custom exporter running on :9105")
	http.ListenAndServe(":9105", nil)
}
