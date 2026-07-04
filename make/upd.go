package main

import (
	"bufio"
	"fmt"
	"log"
	fakedata "make/generatefakedata"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generatePostUpdates() {
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

	if len(idp) == 0 {
		log.Println("Tabela posts nie zawiera identyfikatorów")
		return
	}

	file, err := os.Create("post_update.go")
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

	if _, err := writer.WriteString("var PostUpdate = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	const updateCount = 5000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < updateCount; i++ {
		data := fakedata.GeneratePosts()

		query := fmt.Sprintf(
			"UPDATE posts SET "+
				"score = %d, "+
				"view_count = %d, "+
				"post_body = '%s', "+
				"last_edit_date = '%s', "+
				"last_activity_date = '%s', "+
				"post_title = '%s', "+
				"tags = '%s', "+
				"answer_count = %d, "+
				"comment_count = %d, "+
				"content_license = '%s' "+
				"WHERE id = %d;",
			data.Score,
			data.ViewCount,
			escapeSQL(data.PostBody),
			data.LastEditDate.Format("2006-01-02 15:04:05"),
			data.LastActivityDate.Format("2006-01-02 15:04:05"),
			escapeSQL(data.PostTitle),
			escapeSQL(data.Tags),
			data.AnswerCount,
			data.CommentCount,
			escapeSQL(data.ContentLicense),
			idp[r.Intn(len(idp))],
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
		"Zapisano %d zapytań do pliku post_update.go",
		updateCount,
	)
}

func generateCommentUpdates() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	idc, err := GetIDs(db, "SELECT id FROM comments")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów komentarzy: %v", err)
		return
	}

	idu, err := GetIDs(db, "SELECT id FROM users")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów użytkowników: %v", err)
		return
	}

	if len(idc) == 0 {
		log.Println("Tabela comments nie zawiera identyfikatorów")
		return
	}

	if len(idu) == 0 {
		log.Println("Tabela users nie zawiera identyfikatorów")
		return
	}

	file, err := os.Create("comment_update.go")
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

	if _, err := writer.WriteString("var CommentUpdate = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	const updateCount = 5000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < updateCount; i++ {
		data := fakedata.GenerateComments()

		query := fmt.Sprintf(
			"UPDATE comments SET "+
				"score = %d, "+
				"comment_text = '%s', "+
				"creation_date = '%s', "+
				"user_id = %d, "+
				"content_license = '%s' "+
				"WHERE id = %d;",
			data.Score,
			escapeSQL(data.CommentText),
			data.CreationDate.Format("2006-01-02 15:04:05"),
			idu[r.Intn(len(idu))],
			escapeSQL(data.ContentLicense),
			idc[r.Intn(len(idc))],
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
		"Zapisano %d zapytań do pliku comment_update.go",
		updateCount,
	)
}

func generatePostHistoryUpdates() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	idph, err := GetIDs(db, "SELECT id FROM post_history")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów historii postów: %v", err)
		return
	}

	if len(idph) == 0 {
		log.Println("Tabela post_history nie zawiera identyfikatorów")
		return
	}

	file, err := os.Create("post_history_update.go")
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
		"var PostHistoryUpdate = []string{\n",
	); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	const updateCount = 5000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < updateCount; i++ {
		data := fakedata.GeneratePostHistory()

		query := fmt.Sprintf(
			"UPDATE post_history SET "+
				"post_history_type_id = %d, "+
				"revision_guid = '%s', "+
				"creation_date = '%s', "+
				"post_text = '%s', "+
				"content_license = '%s' "+
				"WHERE id = %d;",
			data.PostHistoryTypeID,
			escapeSQL(data.RevisionGUID),
			data.CreationDate.Format("2006-01-02 15:04:05"),
			escapeSQL(data.PostText),
			escapeSQL(data.ContentLicense),
			idph[r.Intn(len(idph))],
		)

		if _, err := fmt.Fprintf(
			writer,
			"\t%s,\n",
			strconv.Quote(query),
		); err != nil {
			log.Printf(
				"Błąd zapisu zapytania nr %d: %v",
				i,
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
		"Zapisano %d zapytań do pliku post_history_update.go",
		updateCount,
	)

}

func generateUserUpdates() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	idu, err := GetIDs(db, "SELECT id FROM users")
	if err != nil {
		log.Printf("Błąd pobierania identyfikatorów użytkowników: %v", err)
		return
	}

	if len(idu) == 0 {
		log.Println("Tabela users nie zawiera identyfikatorów")
		return
	}

	file, err := os.Create("user_update.go")
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

	if _, err := writer.WriteString("var UserUpdate = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	const updateCount = 5000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < updateCount; i++ {
		data := fakedata.GenerateUsers()

		query := fmt.Sprintf(
			"UPDATE users SET "+
				"reputation = %d, "+
				"display_name = '%s', "+
				"website_url = '%s', "+
				"location = '%s', "+
				"about_me = '%s', "+
				"views = %d, "+
				"upvotes = %d, "+
				"downvotes = %d "+
				"WHERE id = %d;",
			data.Reputation,
			escapeSQL(data.DisplayName),
			escapeSQL(data.WebsiteURL),
			escapeSQL(data.Location),
			escapeSQL(data.AboutMe),
			data.Views,
			data.Upvotes,
			data.Downvotes,
			idu[r.Intn(len(idu))],
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
		"Zapisano %d zapytań do pliku user_update.go",
		updateCount,
	)
}
