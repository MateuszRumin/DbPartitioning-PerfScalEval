package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	userHashPartitionCount = 4
	usersPerHashQuery      = 1000
)

func generateUserHashTests() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	userIDs, err := GetIDs(db, "SELECT id FROM users")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów użytkowników: %v", err)
		return
	}

	if len(userIDs) == 0 {
		log.Println("Tabela users nie zawiera identyfikatorów")
		return
	}

	// Podział istniejących identyfikatorów według HASH(id), przy 4 partycjach.
	var partitionPools [userHashPartitionCount][]int

	for _, id := range userIDs {
		partitionNo := id % userHashPartitionCount

		// Zabezpieczenie na wypadek wartości ujemnych.
		if partitionNo < 0 {
			partitionNo += userHashPartitionCount
		}

		partitionPools[partitionNo] = append(
			partitionPools[partitionNo],
			id,
		)
	}

	for partitionNo, ids := range partitionPools {
		if len(ids) < usersPerHashQuery {
			log.Printf(
				"Partycja p%d zawiera tylko %d identyfikatorów; wymagane minimum: %d",
				partitionNo,
				len(ids),
				usersPerHashQuery,
			)
			return
		}
	}

	file, err := os.Create("users_hash_tests.go")
	if err != nil {
		log.Printf("Błąd tworzenia pliku: %v", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if _, err := writer.WriteString("package main\n\n"); err != nil {
		log.Printf("Błąd zapisu nagłówka: %v", err)
		return
	}

	if _, err := writer.WriteString(
		"var UsersHashTests = []string{\n",
	); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	queryNumber := 1

	/*
		Q01-Q06:
		Każde zapytanie wybiera 1000 użytkowników z jednej partycji.

		Partycje są zmieniane, aby wynik nie zależał wyłącznie od p0.
	*/
	onePartitionPlans := []int{0, 1, 2, 3, 0, 2}

	for _, partitionNo := range onePartitionPlans {
		ids, err := sampleUniqueIDs(
			r,
			partitionPools[partitionNo],
			usersPerHashQuery,
		)
		if err != nil {
			log.Printf(
				"Błąd generowania Q%02d: %v",
				queryNumber,
				err,
			)
			return
		}

		query := buildUsersHashAggregateQuery(ids)

		description := fmt.Sprintf(
			"Q%02d: 1000 użytkowników z jednej partycji HASH p%d",
			queryNumber,
			partitionNo,
		)

		if err := writeGeneratedQuery(
			writer,
			description,
			query,
		); err != nil {
			log.Printf("Błąd zapisu Q%02d: %v", queryNumber, err)
			return
		}

		queryNumber++
	}

	/*
		Q07-Q12:
		Każde zapytanie wybiera po 500 użytkowników z dwóch partycji.
	*/
	twoPartitionPlans := [][2]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 0},
		{0, 2},
		{1, 3},
	}

	for _, plan := range twoPartitionPlans {
		firstIDs, err := sampleUniqueIDs(
			r,
			partitionPools[plan[0]],
			usersPerHashQuery/2,
		)
		if err != nil {
			log.Printf(
				"Błąd generowania Q%02d: %v",
				queryNumber,
				err,
			)
			return
		}

		secondIDs, err := sampleUniqueIDs(
			r,
			partitionPools[plan[1]],
			usersPerHashQuery/2,
		)
		if err != nil {
			log.Printf(
				"Błąd generowania Q%02d: %v",
				queryNumber,
				err,
			)
			return
		}

		ids := append(firstIDs, secondIDs...)

		// Kolejność wartości IN nie powinna sugerować podziału na partycje.
		r.Shuffle(len(ids), func(i, j int) {
			ids[i], ids[j] = ids[j], ids[i]
		})

		query := buildUsersHashAggregateQuery(ids)

		description := fmt.Sprintf(
			"Q%02d: 1000 użytkowników z dwóch partycji HASH p%d i p%d",
			queryNumber,
			plan[0],
			plan[1],
		)

		if err := writeGeneratedQuery(
			writer,
			description,
			query,
		); err != nil {
			log.Printf("Błąd zapisu Q%02d: %v", queryNumber, err)
			return
		}

		queryNumber++
	}

	/*
		Q13-Q18:
		Każde zapytanie wybiera po 250 użytkowników
		z każdej z czterech partycji.
	*/
	for testNo := 0; testNo < 6; testNo++ {
		ids := make([]int, 0, usersPerHashQuery)

		for partitionNo := 0; partitionNo < userHashPartitionCount; partitionNo++ {
			selected, err := sampleUniqueIDs(
				r,
				partitionPools[partitionNo],
				usersPerHashQuery/userHashPartitionCount,
			)
			if err != nil {
				log.Printf(
					"Błąd generowania Q%02d dla p%d: %v",
					queryNumber,
					partitionNo,
					err,
				)
				return
			}

			ids = append(ids, selected...)
		}

		r.Shuffle(len(ids), func(i, j int) {
			ids[i], ids[j] = ids[j], ids[i]
		})

		query := buildUsersHashAggregateQuery(ids)

		description := fmt.Sprintf(
			"Q%02d: 1000 użytkowników ze wszystkich czterech partycji HASH",
			queryNumber,
		)

		if err := writeGeneratedQuery(
			writer,
			description,
			query,
		); err != nil {
			log.Printf("Błąd zapisu Q%02d: %v", queryNumber, err)
			return
		}

		queryNumber++
	}

	// Q19: zakres id — brak skutecznego pruning dla szerokiego zakresu HASH.
	rangeQuery := `
SELECT
    COUNT(*) AS user_count,
    COALESCE(SUM(reputation), 0) AS reputation_sum,
    AVG(reputation) AS reputation_avg,
    MIN(creation_date) AS first_created,
    MAX(creation_date) AS last_created,
    COALESCE(SUM(CHAR_LENGTH(display_name)), 0) AS processed_name_length
FROM users
WHERE id >= 5000000
  AND id < 5100000;`

	if err := writeGeneratedQuery(
		writer,
		"Q19: Zakres id obejmujący wszystkie partycje HASH",
		normalizeSQL(rangeQuery),
	); err != nil {
		log.Printf("Błąd zapisu Q19: %v", err)
		return
	}

	// Q20: filtr poza kluczem partycjonowania.
	nonPartitionKeyQuery := `
SELECT
    COUNT(*) AS user_count,
    COALESCE(SUM(reputation), 0) AS reputation_sum,
    AVG(reputation) AS reputation_avg,
    MIN(creation_date) AS first_created,
    MAX(creation_date) AS last_created,
    COALESCE(SUM(CHAR_LENGTH(display_name)), 0) AS processed_name_length
FROM users
WHERE creation_date >= '2009-01-01'
  AND creation_date < '2010-01-01';`

	if err := writeGeneratedQuery(
		writer,
		"Q20: Filtr po creation_date bez warunku po kluczu partycjonowania",
		normalizeSQL(nonPartitionKeyQuery),
	); err != nil {
		log.Printf("Błąd zapisu Q20: %v", err)
		return
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
		"Zapisano 20 zapytań testowych HASH do pliku users_hash_tests.go",
	)
}

// sampleUniqueIDs wybiera n różnych identyfikatorów z podanej puli.
// Nie kopiuje całej puli i nie wykonuje Perm dla milionów rekordów.
func sampleUniqueIDs(
	r *rand.Rand,
	source []int,
	n int,
) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf(
			"liczba wybieranych identyfikatorów musi być dodatnia",
		)
	}

	if len(source) < n {
		return nil, fmt.Errorf(
			"pula zawiera %d identyfikatorów, wymagane: %d",
			len(source),
			n,
		)
	}

	selected := make([]int, 0, n)
	usedIndexes := make(map[int]struct{}, n)

	for len(selected) < n {
		index := r.Intn(len(source))

		if _, exists := usedIndexes[index]; exists {
			continue
		}

		usedIndexes[index] = struct{}{}
		selected = append(selected, source[index])
	}

	return selected, nil
}

// buildUsersHashAggregateQuery buduje jednakowe zapytanie dla dowolnego
// rozkładu identyfikatorów między partycjami.
func buildUsersHashAggregateQuery(ids []int) string {
	return fmt.Sprintf(
		"SELECT "+
			"COUNT(*) AS user_count,"+
			"COALESCE(SUM(reputation),0) AS reputation_sum,"+
			"AVG(reputation) AS reputation_avg,"+
			"MIN(creation_date) AS first_created,"+
			"MAX(creation_date) AS last_created,"+
			"COALESCE(SUM(CHAR_LENGTH(display_name)),0) AS processed_name_length "+
			"FROM users "+
			"WHERE id IN (%s);",
		idsToSQLList(ids),
	)
}

// idsToSQLList zamienia []int na ciąg: 1,2,3,4.
func idsToSQLList(ids []int) string {
	var builder strings.Builder

	// Przybliżona rezerwacja pamięci.
	builder.Grow(len(ids) * 10)

	for i, id := range ids {
		if i > 0 {
			builder.WriteByte(',')
		}

		builder.WriteString(strconv.Itoa(id))
	}

	return builder.String()
}

// writeGeneratedQuery zapisuje komentarz i poprawnie zacytowany string Go.
func writeGeneratedQuery(
	writer *bufio.Writer,
	description string,
	query string,
) error {
	if _, err := fmt.Fprintf(
		writer,
		"\t// %s\n\t%s,\n",
		description,
		strconv.Quote(query),
	); err != nil {
		return err
	}

	return nil
}

// normalizeSQL usuwa nadmiarowe białe znaki z zapytania wieloliniowego.
func normalizeSQL(query string) string {
	return strings.Join(strings.Fields(query), " ")
}
