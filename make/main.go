package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	fakedata "make/generatefakedata"

	_ "github.com/go-sql-driver/mysql"
)

func escapeSQL(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

func setConnection() (*sql.DB, error) {

	user := "root"
	password := ""
	host := "192.168.50.3"
	port := "3306"
	database := "testdbp"
	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	// Połączenie z bazą danych
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func generateInsertPosts() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	//idb, errb := GetIDs(db, "SELECT id FROM badges")
	//idc, errc := GetIDs(db, "SELECT id FROM comments")
	//idph, errph := GetIDs(db, "SELECT id FROM post_history")
	idp, errp := GetIDs(db, "SELECT id FROM posts")
	idu, erru := GetIDs(db, "SELECT id FROM users")
	// if errb != nil || errc != nil || errph != nil || errp != nil || erru != nil {
	// 	fmt.Printf("Brak danych o indeksach")
	// 	return
	// }
	if errp != nil || erru != nil {
		fmt.Printf("Brak danych o indeksach")
		return
	}
	file, err := os.Create("post_insert.go")
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

	if _, err := writer.WriteString("var PostInsert = []string{\n"); err != nil {
		log.Printf("Błąd zapisu deklaracji zmiennej: %v", err)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 35000; i++ {
		data := fakedata.GeneratePosts()

		query := fmt.Sprintf(
			"INSERT INTO posts "+
				"(post_type_id, accepted_answer_id, parent_id, creation_date, "+
				"score, view_count, post_body, owner_user_id, last_editor_user_id, "+
				"last_edit_date, last_activity_date, post_title, tags, answer_count, "+
				"comment_count, content_license) "+
				"VALUES (%d, %d, %d, '%s', %d, %d, '%s', %d, %d, '%s', "+
				"'%s', '%s', '%s', %d, %d, '%s');",
			data.PostTypeID,
			data.AcceptedAnswerID,
			idp[r.Intn(len(idp))],
			data.CreationDate.Format("2006-01-02 15:04:05"),
			data.Score,
			data.ViewCount,
			escapeSQL(data.PostBody),
			idu[r.Intn(len(idu))],
			idu[r.Intn(len(idu))],
			data.LastEditDate.Format("2006-01-02 15:04:05"),
			data.LastActivityDate.Format("2006-01-02 15:04:05"),
			escapeSQL(data.PostTitle),
			escapeSQL(data.Tags),
			data.AnswerCount,
			data.CommentCount,
			escapeSQL(data.ContentLicense),
		)

		// strconv.Quote tworzy poprawny literał string języka Go.
		if _, err := fmt.Fprintf(writer, "\t%s,\n", strconv.Quote(query)); err != nil {
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

	log.Printf("Zapisano %d zapytań do pliku post_insert.go", 5000)
}

func GetIDs(db *sql.DB, query string) ([]int, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func main() {
	generateInsertPosts()
	// generateUserInserts()
	// generatePostHistoryInserts()
	// generateCommentInserts()

	// generatePostUpdates()
	// generateUserUpdates()
	// generateCommentUpdates()
	// generatePostHistoryUpdates()

}
