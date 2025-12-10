package main

import (
	"fmt"
	"log"
	"net/http"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

func fatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func parseMetrics(url string) (map[string]*dto.MetricFamily, error) {
	reader, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var parser expfmt.TextParser

	mf, err := parser.TextToMetricFamilies(reader.Body)
	if err != nil {
		return nil, err
	}
	return mf, nil

}

func fetchData() {
	url := "http://localhost:9090/metrics"
	mf, err := parseMetrics(url)
	fatal(err)

	for _, v := range mf {
		// fmt.Println("KEY: ", k)
		fmt.Println("VAL: ", v)
	}

}
