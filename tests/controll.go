package tests

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"perfscaleval/confmodel"
	"perfscaleval/connection"
)

func CheckEnabledConnection() {
	var wg sync.WaitGroup
	go func(index int) {
		defer wg.Done()
		if confmodel.Connection1.Enable {
			TestcontrollSelect(confmodel.Connection1)
		} else {
			fmt.Println("connection1 disabled")
		}
	}(1)
	go func(index int) {
		defer wg.Done()
		if confmodel.Connection2.Enable {
			TestcontrollSelect(confmodel.Connection2)
		} else {
			fmt.Println("connection2 disabled")
		}
	}(2)

}

func TestcontrollSelect(connectionSettings confmodel.ConnectionSetting) {
	var wg sync.WaitGroup

	startTime := time.Now()
	if len(confmodel.Plan) == 0 {
		fmt.Println("Brak planów")
	} else {

		for _, plan := range confmodel.Plan {
			if len(plan.Group) == 0 {
				fmt.Println("Brak grup")
			} else {

				for i, group := range plan.Group {
					go func(connIndex int, group confmodel.TestGroup) {
						allStats := make([]ThreadStats, group.ThreadNumber)
						defer wg.Done()
						if len(group.Steps) == 0 {
							fmt.Println("Brak kroków")
						} else {

							for j := 0; j < group.ThreadNumber; j++ {
								go func(connIndex int) {
									defer wg.Done()
									connection, err := connection.SetConnection(connectionSettings.User, connectionSettings.Password, connectionSettings.Host, connectionSettings.Port, connectionSettings.DatabaseName)
									if err != nil {
										fmt.Println("cannot connect")
										return
									}

									results := executeQueries(group.Steps, connection, connIndex)

									allStats[connIndex] = calculateThreadStats(results, startTime, connIndex)

								}(j)
							}
						}

					}(i, group)

				}

			}

		}
	}

}

func executeQueries(queries []confmodel.TestStep, conn *sql.DB, threadNum int) []QueryResult {
	// Włącz profilowanie (jeśli baza danych to obsługuje)
	_, err := conn.Exec("SET profiling = 1")
	if err != nil {
		fmt.Printf("Błąd podczas włączania profilowania (wątek %d): %v\n", threadNum, err)
		return nil
	}

	var queryResults []QueryResult

	for _, query := range queries {
		// Wykonaj zapytanie
		_, err := conn.Exec(query.Query)
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
			Query:    query.Query,
			Duration: totalDuration,
			Stages:   stages,
		})
	}

	return queryResults
}

func calculateThreadStats(results []QueryResult, startTime time.Time, threadNum int) ThreadStats {
	stats := ThreadStats{
		TestName:             "Test",
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
