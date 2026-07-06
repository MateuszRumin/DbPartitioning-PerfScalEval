package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type postHistoryKey struct {
	ID     int64
	PostID int64
}

type generatedPostHistoryQuery struct {
	Description string
	SQL         string
}

func generatePostHistoryQueries() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	// Nazwy partycji są pobierane bezpośrednio z aktualnie używanej bazy.
	partitionNames, err := getPostHistoryPartitionNames(db)
	if err != nil {
		log.Printf("Błąd pobierania partycji tabeli post_history: %v", err)
		return
	}

	// Zestaw został zaprojektowany dla ośmiu partycji.
	if len(partitionNames) != 8 {
		log.Printf(
			"Tabela post_history powinna zawierać 8 partycji, znaleziono: %d",
			len(partitionNames),
		)
		return
	}

	const samplesPerPartition = 18

	// Z każdej partycji pobierane jest 18 istniejących par (id, post_id).
	samples := make([][]postHistoryKey, len(partitionNames))

	for partitionIndex, partitionName := range partitionNames {
		keys, err := getPostHistoryKeysFromPartition(
			db,
			partitionName,
			samplesPerPartition,
		)
		if err != nil {
			log.Printf(
				"Błąd pobierania rekordów z partycji %s: %v",
				partitionName,
				err,
			)
			return
		}

		if len(keys) < samplesPerPartition {
			log.Printf(
				"Partycja %s zawiera za mało rekordów: %d, wymagane: %d",
				partitionName,
				len(keys),
				samplesPerPartition,
			)
			return
		}

		samples[partitionIndex] = keys
	}

	queries := make([]generatedPostHistoryQuery, 0, 20)

	addQuery := func(
		description string,
		keys []postHistoryKey,
		disablePruning bool,
	) {
		queries = append(
			queries,
			generatedPostHistoryQuery{
				Description: description,
				SQL: makePostHistoryScanQuery(
					keys,
					disablePruning,
				),
			},
		)
	}

	// -------------------------------------------------------------------------
	// Q01-Q08
	// Osiem rekordów pochodzących z jednej partycji.
	// Każde zapytanie powinno przeszukać jedną partycję.
	// -------------------------------------------------------------------------

	onePartitionSets := make([][]postHistoryKey, 8)

	for partitionIndex := 0; partitionIndex < 8; partitionIndex++ {
		keys := selectPostHistoryKeys(
			samples,
			[]int{partitionIndex},
			8,
			0,
		)

		onePartitionSets[partitionIndex] = keys

		addQuery(
			fmt.Sprintf(
				"8 rekordów z jednej partycji — %s",
				partitionNames[partitionIndex],
			),
			keys,
			false,
		)
	}

	// -------------------------------------------------------------------------
	// Q09-Q12
	// Po cztery rekordy z dwóch partycji.
	// Każde zapytanie nadal zwraca osiem rekordów.
	// -------------------------------------------------------------------------

	twoPartitionGroups := [][]int{
		{0, 1},
		{2, 3},
		{4, 5},
		{6, 7},
	}

	twoPartitionSets := make([][]postHistoryKey, len(twoPartitionGroups))

	for groupIndex, partitions := range twoPartitionGroups {
		keys := selectPostHistoryKeys(
			samples,
			partitions,
			4,
			8,
		)

		twoPartitionSets[groupIndex] = keys

		addQuery(
			fmt.Sprintf(
				"Po 4 rekordy z dwóch partycji — %s, %s",
				partitionNames[partitions[0]],
				partitionNames[partitions[1]],
			),
			keys,
			false,
		)
	}

	// -------------------------------------------------------------------------
	// Q13-Q16
	// Po dwa rekordy z czterech partycji.
	// Każde zapytanie nadal zwraca osiem rekordów.
	// -------------------------------------------------------------------------

	fourPartitionGroups := []struct {
		Partitions []int
		Offset     int
	}{
		{
			Partitions: []int{0, 1, 2, 3},
			Offset:     12,
		},
		{
			Partitions: []int{4, 5, 6, 7},
			Offset:     12,
		},
		{
			Partitions: []int{0, 2, 4, 6},
			Offset:     14,
		},
		{
			Partitions: []int{1, 3, 5, 7},
			Offset:     14,
		},
	}

	for _, group := range fourPartitionGroups {
		keys := selectPostHistoryKeys(
			samples,
			group.Partitions,
			2,
			group.Offset,
		)

		names := make([]string, 0, len(group.Partitions))

		for _, partitionIndex := range group.Partitions {
			names = append(
				names,
				partitionNames[partitionIndex],
			)
		}

		addQuery(
			fmt.Sprintf(
				"Po 2 rekordy z czterech partycji — %s",
				strings.Join(names, ", "),
			),
			keys,
			false,
		)
	}

	// -------------------------------------------------------------------------
	// Q17-Q18
	// Po jednym rekordzie z każdej z ośmiu partycji.
	// -------------------------------------------------------------------------

	allPartitions := []int{0, 1, 2, 3, 4, 5, 6, 7}

	keysQ17 := selectPostHistoryKeys(
		samples,
		allPartitions,
		1,
		16,
	)

	addQuery(
		"Po 1 rekordzie z każdej z 8 partycji",
		keysQ17,
		false,
	)

	keysQ18 := selectPostHistoryKeys(
		samples,
		allPartitions,
		1,
		17,
	)

	addQuery(
		"Po 1 innym rekordzie z każdej z 8 partycji",
		keysQ18,
		false,
	)

	// -------------------------------------------------------------------------
	// Q19
	// Te same rekordy co Q01, lecz funkcja ABS() uniemożliwia bezpośrednie
	// wykorzystanie predykatów do partition pruning.
	// -------------------------------------------------------------------------

	addQuery(
		"Te same rekordy co Q01 — pruning celowo wyłączony",
		onePartitionSets[0],
		true,
	)

	// -------------------------------------------------------------------------
	// Q20
	// Te same rekordy co Q09, lecz pruning jest celowo wyłączony.
	// -------------------------------------------------------------------------

	addQuery(
		"Te same rekordy co Q09 — pruning celowo wyłączony",
		twoPartitionSets[0],
		true,
	)

	if len(queries) != 20 {
		log.Printf(
			"Nieprawidłowa liczba wygenerowanych zapytań: %d",
			len(queries),
		)
		return
	}

	file, err := os.Create("post_history_queries.go")
	if err != nil {
		log.Printf("Błąd tworzenia pliku: %v", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if _, err := writer.WriteString("package main\n\n"); err != nil {
		log.Printf("Błąd zapisu nagłówka pliku: %v", err)
		return
	}

	if _, err := writer.WriteString("var postHistory = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji tablicy: %v", err)
		return
	}

	for queryIndex, query := range queries {
		if _, err := fmt.Fprintf(
			writer,
			"\t// Q%02d: %s\n",
			queryIndex+1,
			query.Description,
		); err != nil {
			log.Printf(
				"Błąd zapisu komentarza zapytania Q%02d: %v",
				queryIndex+1,
				err,
			)
			return
		}

		if _, err := fmt.Fprintf(
			writer,
			"\t%s,\n\n",
			strconv.Quote(query.SQL),
		); err != nil {
			log.Printf(
				"Błąd zapisu zapytania Q%02d: %v",
				queryIndex+1,
				err,
			)
			return
		}
	}

	if _, err := writer.WriteString("}\n"); err != nil {
		log.Printf("Błąd zapisu zakończenia tablicy: %v", err)
		return
	}

	if err := writer.Flush(); err != nil {
		log.Printf("Błąd opróżniania bufora: %v", err)
		return
	}

	log.Printf(
		"Zapisano %d zapytań do pliku post_history_queries.go",
		len(queries),
	)
}

func getPostHistoryPartitionNames(
	db *sql.DB,
) ([]string, error) {
	const query = `
		SELECT partition_name
		FROM information_schema.partitions
		WHERE table_schema = DATABASE()
		  AND table_name = 'post_history'
		  AND partition_name IS NOT NULL
		ORDER BY partition_ordinal_position
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partitionNames []string

	for rows.Next() {
		var partitionName string

		if err := rows.Scan(&partitionName); err != nil {
			return nil, err
		}

		partitionNames = append(
			partitionNames,
			partitionName,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(partitionNames) == 0 {
		return nil, fmt.Errorf(
			"tabela post_history nie jest partycjonowana albo nie istnieje",
		)
	}

	return partitionNames, nil
}

func getPostHistoryKeysFromPartition(
	db *sql.DB,
	partitionName string,
	limit int,
) ([]postHistoryKey, error) {
	// Nazwa partycji pochodzi z INFORMATION_SCHEMA, ale mimo tego
	// zabezpieczamy identyfikator przed znakiem backtick.
	safePartitionName := strings.ReplaceAll(
		partitionName,
		"`",
		"``",
	)

	query := fmt.Sprintf(
		"SELECT id, post_id "+
			"FROM post_history PARTITION (`%s`) "+
			"ORDER BY id, post_id "+
			"LIMIT %d",
		safePartitionName,
		limit,
	)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	keys := make([]postHistoryKey, 0, limit)

	for rows.Next() {
		var key postHistoryKey

		if err := rows.Scan(
			&key.ID,
			&key.PostID,
		); err != nil {
			return nil, err
		}

		keys = append(keys, key)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return keys, nil
}

func selectPostHistoryKeys(
	samples [][]postHistoryKey,
	partitionIndexes []int,
	keysPerPartition int,
	offset int,
) []postHistoryKey {
	keys := make(
		[]postHistoryKey,
		0,
		len(partitionIndexes)*keysPerPartition,
	)

	for _, partitionIndex := range partitionIndexes {
		for keyIndex := 0; keyIndex < keysPerPartition; keyIndex++ {
			keys = append(
				keys,
				samples[partitionIndex][offset+keyIndex],
			)
		}
	}

	return keys
}

func makePostHistoryScanQuery(
	keys []postHistoryKey,
	disablePruning bool,
) string {
	conditions := make([]string, 0, len(keys))

	for _, key := range keys {
		var condition string

		if disablePruning {
			condition = fmt.Sprintf(
				"(ABS(ph.id)=%d AND ABS(ph.post_id)=%d)",
				key.ID,
				key.PostID,
			)
		} else {
			condition = fmt.Sprintf(
				"(ph.id=%d AND ph.post_id=%d)",
				key.ID,
				key.PostID,
			)
		}

		conditions = append(
			conditions,
			condition,
		)
	}

	return fmt.Sprintf(
		"SELECT ph.id,ph.post_id,ph.post_history_type_id,"+
			"ph.creation_date,ph.user_id "+
			"FROM post_history AS ph USE INDEX () "+
			"WHERE %s;",
		strings.Join(conditions, " OR "),
	)
}
