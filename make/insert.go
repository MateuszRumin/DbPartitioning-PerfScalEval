package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	fakedata "make/generatefakedata"
)

func generateUserInserts() {
	file, err := os.Create("user_insert.go")
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

	if _, err := writer.WriteString("var UserInsert = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	const insertCount = 35000

	for i := 0; i < insertCount; i++ {
		data := fakedata.GenerateUsers()

		query := fmt.Sprintf(
			"INSERT INTO users "+
				"(reputation, creation_date, display_name, last_access_date, "+
				"website_url, location, about_me, views, upvotes, downvotes, account_id) "+
				"VALUES (%d, '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d);",
			data.Reputation,
			data.CreationDate.Format("2006-01-02 15:04:05"),
			escapeSQL(data.DisplayName),
			data.LastAccessDate.Format("2006-01-02 15:04:05"),
			escapeSQL(data.WebsiteURL),
			escapeSQL(data.Location),
			escapeSQL(data.AboutMe),
			data.Views,
			data.Upvotes,
			data.Downvotes,
			data.AccountID,
		)

		if _, err := fmt.Fprintf(
			writer,
			"\t%s,\n",
			strconv.Quote(query),
		); err != nil {
			log.Printf("Błąd zapisu zapytania nr %d: %v", i, err)
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
		"Zapisano %d zapytań do pliku user_insert.go",
		insertCount,
	)
}

func generateCommentInserts() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	idp, err := GetIDs(db, "SELECT id FROM posts")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów postów: %v", err)
		return
	}

	idu, err := GetIDs(db, "SELECT id FROM users")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów użytkowników: %v", err)
		return
	}

	if len(idp) == 0 {
		log.Println("Tabela posts nie zawiera identyfikatorów")
		return
	}

	if len(idu) == 0 {
		log.Println("Tabela users nie zawiera identyfikatorów")
		return
	}

	file, err := os.Create("comment_insert.go")
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

	if _, err := writer.WriteString("var CommentInsert = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	const insertCount = 35000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < insertCount; i++ {
		data := fakedata.GenerateComments()

		query := fmt.Sprintf(
			"INSERT INTO comments "+
				"(post_id, score, comment_text, creation_date, user_id, content_license) "+
				"VALUES (%d, %d, '%s', '%s', %d, '%s');",
			idp[r.Intn(len(idp))],
			data.Score,
			escapeSQL(data.CommentText),
			data.CreationDate.Format("2006-01-02 15:04:05"),
			idu[r.Intn(len(idu))],
			escapeSQL(data.ContentLicense),
		)

		if _, err := fmt.Fprintf(
			writer,
			"\t%s,\n",
			strconv.Quote(query),
		); err != nil {
			log.Printf("Błąd zapisu zapytania nr %d: %v", i, err)
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
		"Zapisano %d zapytań do pliku comment_insert.go",
		insertCount,
	)
}

func generatePostHistoryInserts() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	idp, err := GetIDs(db, "SELECT id FROM posts")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów postów: %v", err)
		return
	}

	idu, err := GetIDs(db, "SELECT id FROM users")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów użytkowników: %v", err)
		return
	}

	if len(idp) == 0 {
		log.Println("Tabela posts nie zawiera identyfikatorów")
		return
	}

	if len(idu) == 0 {
		log.Println("Tabela users nie zawiera identyfikatorów")
		return
	}

	file, err := os.Create("post_history_insert.go")
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

	if _, err := writer.WriteString("var PostHistoryInsert = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	const insertCount = 35000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < insertCount; i++ {
		data := fakedata.GeneratePostHistory()

		query := fmt.Sprintf(
			"INSERT INTO post_history "+
				"(post_history_type_id, post_id, revision_guid, creation_date, "+
				"user_id, post_text, content_license) "+
				"VALUES (%d, %d, '%s', '%s', %d, '%s', '%s');",
			data.PostHistoryTypeID,
			idp[r.Intn(len(idp))],
			escapeSQL(data.RevisionGUID),
			data.CreationDate.Format("2006-01-02 15:04:05"),
			idu[r.Intn(len(idu))],
			escapeSQL(data.PostText),
			escapeSQL(data.ContentLicense),
		)

		if _, err := fmt.Fprintf(
			writer,
			"\t%s,\n",
			strconv.Quote(query),
		); err != nil {
			log.Printf("Błąd zapisu zapytania nr %d: %v", i, err)
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
		"Zapisano %d zapytań do pliku post_history_insert.go",
		insertCount,
	)
}
