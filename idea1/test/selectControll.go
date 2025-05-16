package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"perfscaleval/config"
	"perfscaleval/selectqueries"
	"sync"
	"time"
)

// Struktura do przechowywania wyników profilowania i dodatkowych informacji
type TestResult struct {
	TestName      string        `json:"test_name"`       // Nazwa testu
	ThreadNum     int           `json:"thread_num"`      // Numer wątku
	QueryResults  []QueryResult `json:"query_results"`   // Wyniki zapytań
	Timeout       float64       `json:"timeout"`         // Czas timeoutu w sekundach
	TotalTestTime float64       `json:"total_test_time"` // Całkowity czas testu w sekundach
}

// Struktura do przechowywania wyników pojedynczego zapytania
type QueryResult struct {
	QueryID  int            `json:"query_id"` // ID zapytania
	Query    string         `json:"query"`    // Treść zapytania
	Duration float64        `json:"duration"` // Całkowity czas wykonania zapytania w sekundach
	Stages   []ProfileStage `json:"stages"`   // Etapy wykonania zapytania
}

// Struktura do przechowywania etapów profilowania
type ProfileStage struct {
	Status   string  `json:"status"`   // Status etapu
	Duration float64 `json:"duration"` // Czas trwania etapu w sekundach
}

// Struktura do przechowywania statystyk dla wątku
type ThreadStats struct {
	TestName             string             `json:"test_name"`              // Nazwa testu
	ThreadNum            int                `json:"thread_num"`             // Numer wątku
	TotalDuration        float64            `json:"total_duration"`         // Suma czasów wszystkich zapytań w sekundach
	TotalStatusDurations map[string]float64 `json:"total_status_durations"` // Suma czasów dla każdego statusu w sekundach
	AvgQueryDuration     float64            `json:"avg_query_duration"`     // Średni czas wykonania zapytania w sekundach
	AvgStatusDurations   map[string]float64 `json:"avg_status_durations"`   // Średni czas dla każdego statusu w sekundach
	TotalAppTime         float64            `json:"total_app_time"`         // Całkowity czas w aplikacji dla wątku w sekundach
	AvgAppTimePerQuery   float64            `json:"avg_app_time_per_query"` // Średni czas w aplikacji na zapytanie w sekundach
}

// Struktura do przechowywania statystyk ogólnych
type OverallStats struct {
	TestName             string             `json:"test_name"`              // Nazwa testu
	TotalDuration        float64            `json:"total_duration"`         // Suma czasów wszystkich zapytań w sekundach
	TotalStatusDurations map[string]float64 `json:"total_status_durations"` // Suma czasów dla każdego statusu w sekundach
	AvgQueryDuration     float64            `json:"avg_query_duration"`     // Średni czas wykonania zapytania w sekundach
	AvgStatusDurations   map[string]float64 `json:"avg_status_durations"`   // Średni czas dla każdego statusu w sekundach
	TotalAppTime         float64            `json:"total_app_time"`         // Całkowity czas w aplikacji w sekundach
	AvgAppTimePerQuery   float64            `json:"avg_app_time_per_query"` // Średni czas w aplikacji na zapytanie w sekundach
	StartTimestamp       string             `json:"start_timestamp"`        // Timestamp rozpoczęcia testu
	EndTimestamp         string             `json:"end_timestamp"`          // Timestamp zakończenia testu
}

func selectControl() {

	startTime := time.Now() // Rozpocznij pomiar czasu całego testu

	val, ok := selectqueries.SelectQueriesStack[config.TestName]
	if !ok || val == nil {
		fmt.Println("Błąd flagi")
		return
	}

	var wg sync.WaitGroup
	wg.Add(config.NumConnections)

	// Slice do przechowywania statystyk z wszystkich wątków
	allStats := make([]ThreadStats, config.NumConnections)

	for i := 0; i < config.NumConnections; i++ {
		go func(connIndex int) {
			defer wg.Done()
			// Użyj połączenia z config.connections[connIndex]
			results := executeQueries(val, config.Connections[connIndex], connIndex)

			// Oblicz statystyki dla wątku
			allStats[connIndex] = calculateThreadStats(results, startTime, connIndex)
		}(i)
	}

	wg.Wait()

	// Wyświetl statystyki z wszystkich wątków
	// for _, stats := range allStats {
	// 	printThreadStats(stats)
	// }

	// Oblicz i wyświetl ogólne statystyki
	endTime := time.Now()
	overallStats := calculateOverallStats(allStats, startTime, endTime)
	printOverallStats(overallStats)

	totalTestTime := time.Since(startTime).Seconds() // Oblicz całkowity czas testu w sekundach
	fmt.Printf("Całkowity czas testu: %.6f sekund\n", totalTestTime)

}

func executeQueries(queries *[]string, conn *sql.DB, threadNum int) []QueryResult {
	// Włącz profilowanie (jeśli baza danych to obsługuje)
	_, err := conn.Exec("SET profiling = 1")
	if err != nil {
		fmt.Printf("Błąd podczas włączania profilowania (wątek %d): %v\n", threadNum, err)
		return nil
	}

	var queryResults []QueryResult

	for _, query := range *queries {
		// Wykonaj zapytanie
		_, err := conn.Exec(query)
		if err != nil {
			fmt.Printf("Wątek %d: Błąd podczas wykonywania zapytania: %v\n", threadNum, err)
			continue
		}

		// Pobierz szczegółowe informacje o profilowaniu dla ostatniego zapytania
		profileRows, err := conn.Query("SHOW PROFILE")
		if err != nil {
			fmt.Printf("Wątek %d: Błąd podczas pobierania profilowania: %v\n", threadNum, err)
			continue
		}

		// Przechowuj etapy profilowania
		var stages []ProfileStage
		for profileRows.Next() {
			var (
				status   string
				duration float64
			)
			if err := profileRows.Scan(&status, &duration); err != nil {
				fmt.Printf("Wątek %d: Błąd podczas parsowania profilowania: %v\n", threadNum, err)
				continue
			}
			stages = append(stages, ProfileStage{Status: status, Duration: duration})
		}
		profileRows.Close()

		// Pobierz całkowity czas wykonania zapytania
		var totalDuration float64
		row := conn.QueryRow("SELECT SUM(DURATION) FROM INFORMATION_SCHEMA.PROFILING WHERE QUERY_ID = (SELECT MAX(QUERY_ID) FROM INFORMATION_SCHEMA.PROFILING)")
		if err := row.Scan(&totalDuration); err != nil {
			fmt.Printf("Wątek %d: Błąd podczas pobierania całkowitego czasu zapytania: %v\n", threadNum, err)
			continue
		}

		// Dodaj wyniki do listy
		queryResults = append(queryResults, QueryResult{
			QueryID:  len(queryResults) + 1, // Unikalny identyfikator zapytania
			Query:    query,
			Duration: totalDuration,
			Stages:   stages,
		})
	}

	return queryResults
}

func calculateThreadStats(results []QueryResult, startTime time.Time, threadNum int) ThreadStats {
	stats := ThreadStats{
		TestName:             config.TestName,
		ThreadNum:            threadNum,
		TotalStatusDurations: make(map[string]float64),
		AvgStatusDurations:   make(map[string]float64),
	}

	// Oblicz sumy i średnie
	totalQueries := len(results)
	if totalQueries == 0 {
		return stats
	}

	// Suma czasów wszystkich zapytań
	for _, result := range results {
		stats.TotalDuration += result.Duration
		for _, stage := range result.Stages {
			stats.TotalStatusDurations[stage.Status] += stage.Duration
		}
	}

	// Średni czas wykonania zapytania
	stats.AvgQueryDuration = stats.TotalDuration / float64(totalQueries)

	// Średni czas dla każdego statusu
	for status, total := range stats.TotalStatusDurations {
		stats.AvgStatusDurations[status] = total / float64(totalQueries)
	}

	// Całkowity czas w aplikacji dla wątku w sekundach
	stats.TotalAppTime = time.Since(startTime).Seconds()

	// Średni czas w aplikacji na zapytanie w sekundach
	stats.AvgAppTimePerQuery = stats.TotalAppTime / float64(totalQueries)

	return stats
}

func calculateOverallStats(allStats []ThreadStats, startTime, endTime time.Time) OverallStats {
	overallStats := OverallStats{
		TestName:             config.TestName,
		TotalStatusDurations: make(map[string]float64),
		AvgStatusDurations:   make(map[string]float64),
		StartTimestamp:       startTime.Format(time.RFC3339), // Timestamp rozpoczęcia testu
		EndTimestamp:         endTime.Format(time.RFC3339),   // Timestamp zakończenia testu
	}

	totalThreads := len(allStats)
	if totalThreads == 0 {
		return overallStats
	}

	// Suma czasów wszystkich wątków
	for _, stats := range allStats {
		overallStats.TotalDuration += stats.TotalDuration
		overallStats.TotalAppTime += stats.TotalAppTime

		// Suma czasów dla każdego statusu
		for status, duration := range stats.TotalStatusDurations {
			overallStats.TotalStatusDurations[status] += duration
		}
	}

	// Średni czas wykonania zapytania
	overallStats.AvgQueryDuration = overallStats.TotalDuration / float64(totalThreads)

	// Średni czas dla każdego statusu
	for status, total := range overallStats.TotalStatusDurations {
		overallStats.AvgStatusDurations[status] = total / float64(totalThreads)
	}

	// Średni czas w aplikacji na zapytanie
	overallStats.AvgAppTimePerQuery = overallStats.TotalAppTime / float64(totalThreads)

	return overallStats
}

func printOverallStats(stats OverallStats) {
	// Konwertuj strukturę na JSON
	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		fmt.Println("Błąd podczas konwersji do JSON:", err)
		return
	}

	// Wyświetl JSON
	fmt.Println(string(jsonData))
}
