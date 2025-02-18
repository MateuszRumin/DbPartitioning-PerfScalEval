package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"perfscaleval/config"
	"perfscaleval/dbconects"
	"perfscaleval/selectqueries"
	"sync"
	"time"
)

// Struktura do przechowywania wyników profilowania i dodatkowych informacji
type TestResult struct {
	TestName      string        `json:"test_name"`       // Nazwa testu
	ThreadNum     int           `json:"thread_num"`      // Numer wątku
	QueryResults  []QueryResult `json:"query_results"`   // Wyniki zapytań
	Timeout       time.Duration `json:"timeout"`         // Czas timeoutu
	TotalTestTime time.Duration `json:"total_test_time"` // Całkowity czas testu
}

// Struktura do przechowywania wyników pojedynczego zapytania
type QueryResult struct {
	QueryID  int            `json:"query_id"` // ID zapytania
	Query    string         `json:"query"`    // Treść zapytania
	Duration float64        `json:"duration"` // Całkowity czas wykonania zapytania
	Stages   []ProfileStage `json:"stages"`   // Etapy wykonania zapytania
}

// Struktura do przechowywania etapów profilowania
type ProfileStage struct {
	Status   string  `json:"status"`   // Status etapu
	Duration float64 `json:"duration"` // Czas trwania etapu
}

// Struktura do przechowywania statystyk dla wątku
type ThreadStats struct {
	TestName             string             `json:"test_name"`              // Nazwa testu
	ThreadNum            int                `json:"thread_num"`             // Numer wątku
	TotalDuration        float64            `json:"total_duration"`         // Suma czasów wszystkich zapytań
	TotalStatusDurations map[string]float64 `json:"total_status_durations"` // Suma czasów dla każdego statusu
	AvgQueryDuration     float64            `json:"avg_query_duration"`     // Średni czas wykonania zapytania
	AvgStatusDurations   map[string]float64 `json:"avg_status_durations"`   // Średni czas dla każdego statusu
	TotalAppTime         time.Duration      `json:"total_app_time"`         // Całkowity czas w aplikacji dla wątku
	AvgAppTimePerQuery   float64            `json:"avg_app_time_per_query"` // Średni czas w aplikacji na zapytanie
}

func selectAnalizeControl() {
	if config.NumConnections <= 0 {
		fmt.Println("Błąd flagi: nie można przeprowadzić testów bez połączenia z bazą danych")
		return
	}

	startTime := time.Now() // Rozpocznij pomiar czasu całego testu

	if config.NumConnections == 1 {
		oneConnectionSelect()
	} else {
		multipleConnectionSelect(startTime)
	}

	totalTestTime := time.Since(startTime) // Oblicz całkowity czas testu
	fmt.Printf("Całkowity czas testu: %v\n", totalTestTime)
}

func oneConnectionSelect() {
	dbconects.SetConnections()
	if !config.ConectionWorking {
		fmt.Println("Brak połączenia")
		return
	}

	val, ok := selectqueries.SelectQueriesStack[config.TestName]
	if !ok || val == nil {
		fmt.Println("Błąd flagi")
		return
	}

	startTime := time.Now() // Rozpocznij pomiar czasu dla tego wątku

	// Użyj pierwszego połączenia z config.connections
	results := executeQueries(val, config.Connections[0], 0) // 0 to numer wątku

	// Oblicz statystyki dla wątku
	threadStats := calculateThreadStats(results, startTime, 0)

	// Wyświetl statystyki w formacie JSON
	printThreadStats(threadStats)
}

func multipleConnectionSelect(startTime time.Time) {
	dbconects.SetConnections()
	if !config.ConectionWorking {
		fmt.Println("Brak połączenia")
		return
	}

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
	for _, stats := range allStats {
		printThreadStats(stats)
	}

	// Oblicz i wyświetl ogólne statystyki
	overallStats := calculateOverallStats(allStats, startTime)
	printThreadStats(overallStats)
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

	// Całkowity czas w aplikacji dla wątku
	stats.TotalAppTime = time.Since(startTime)

	// Średni czas w aplikacji na zapytanie
	stats.AvgAppTimePerQuery = float64(stats.TotalAppTime) / float64(totalQueries)

	return stats
}

func calculateOverallStats(allStats []ThreadStats, startTime time.Time) ThreadStats {
	overallStats := ThreadStats{
		TestName:             config.TestName,
		ThreadNum:            -1, // -1 oznacza, że to ogólne statystyki
		TotalStatusDurations: make(map[string]float64),
		AvgStatusDurations:   make(map[string]float64),
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
	overallStats.AvgAppTimePerQuery = float64(overallStats.TotalAppTime) / float64(totalThreads)

	return overallStats
}

func printThreadStats(stats ThreadStats) {
	// Konwertuj strukturę na JSON
	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		fmt.Println("Błąd podczas konwersji do JSON:", err)
		return
	}

	// Wyświetl JSON
	fmt.Println(string(jsonData))
}
