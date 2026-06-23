package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	fakedata "test6/generatefakedata"

	_ "github.com/go-sql-driver/mysql"
)

func setConnection() (*sql.DB, error) {

	user := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "testdb"
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

func checkConnectionAndRunTest(id int, deadline time.Time) {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	for time.Now().Before(deadline) {
		testDb(db, id)
	}
}

func multiThread(workersCount int) {
	var wg sync.WaitGroup

	deadline := time.Now().Add(10 * time.Minute)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			checkConnectionAndRunTest(id, deadline)
		}(i)
	}

	wg.Wait()
}

func executeQuery(db *sql.DB, query string, id int) (err error) {
	// fmt.Printf("Wykonuję zapytanie: %s\n", query)
	rows, err := db.Query(query)
	if err != nil {
		//log.Printf("Zapytanie, które spowodowało błąd: %s\n", query)
		log.Printf("Nie udało się wykonać zapytania: \n")
		return err

	}
	fmt.Printf("Succes wątek id: %d \n", id)
	defer rows.Close()
	return nil

}
func executeSimpleQuery(db *sql.DB, query string) (int, error) {
	var result int

	err := db.QueryRow(query).Scan(&result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func random(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func testDb(db *sql.DB, id int) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	whatTable := []string{"badges", "comments", "posthistory", "posts", "users"}
	randomIndex := r.Intn(len(whatTable))
	switch whatTable[randomIndex] {
	case "badges":
		data := fakedata.GenerateBadge()
		dataLow, errLow := executeSimpleQuery(db, "SELECT MIN(id) FROM badges;")
		dataHigh, errHigh := executeSimpleQuery(db, "SELECT MAX(id) FROM badges;")
		if errLow != nil || errHigh != nil {
			return
		}
		randomId := random(dataLow, dataHigh)
		query := fmt.Sprintf("Update badges SET badge_name = '%s', badge_date = '%s', class = %d, tag_based = '%s' WHERE id = %d;", data.BadgeName, data.BadgeDate.Format("2006-01-02 15:04:05"), data.Class, data.TagBased, randomId)
		if err := executeQuery(db, query, id); err != nil {
			return
		}
	case "comments":
		data := fakedata.GenerateComments()
		dataLow, errLow := executeSimpleQuery(db, "SELECT MIN(id) FROM comments;")
		dataHigh, errHigh := executeSimpleQuery(db, "SELECT MAX(id) FROM comments;")
		if errLow != nil || errHigh != nil {
			return
		}
		randomId := random(dataLow, dataHigh)
		query := fmt.Sprintf("Update comments SET score = %d, comment_text = '%s', creation_date = '%s', user_id = %d, content_license = '%s' WHERE id = %d;", data.Score, data.CommentText, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.ContentLicense, randomId)
		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "posthistory":
		data := fakedata.GeneratePostHistory()
		dataLow, errLow := executeSimpleQuery(db, "SELECT MIN(id) FROM post_history;")
		dataHigh, errHigh := executeSimpleQuery(db, "SELECT MAX(id) FROM post_history;")
		if errLow != nil || errHigh != nil {
			return
		}
		randomId := random(dataLow, dataHigh)
		query := fmt.Sprintf("Update post_history SET post_history_type_id = %d, revision_guid = '%s', creation_date = '%s', post_text = '%s', content_license = '%s' WHERE id = %d;", data.PostHistoryTypeID, data.RevisionGUID, data.CreationDate.Format("2006-01-02 15:04:05"), data.PostText, data.ContentLicense, randomId)

		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "posts":
		data := fakedata.GeneratePosts()
		dataLow, errLow := executeSimpleQuery(db, "SELECT MIN(id) FROM posts;")
		dataHigh, errHigh := executeSimpleQuery(db, "SELECT MAX(id) FROM posts;")
		if errLow != nil || errHigh != nil {
			return
		}
		randomId := random(dataLow, dataHigh)
		query := fmt.Sprintf("Update posts SET   score = %d, view_count = %d, post_body = '%s',   last_edit_date = '%s', last_activity_date = '%s', post_title = '%s', tags = '%s', answer_count = %d, comment_count = %d, content_license = '%s' WHERE id = %d;",
			data.Score, data.ViewCount, data.PostBody, data.LastEditDate.Format("2006-01-02 15:04:05"), data.LastActivityDate.Format("2006-01-02 15:04:05"), data.PostTitle, data.Tags, data.AnswerCount, data.CommentCount, data.ContentLicense, randomId)
		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "users":
		data := fakedata.GenerateUsers()
		dataLow, errLow := executeSimpleQuery(db, "SELECT MIN(id) FROM users;")
		dataHigh, errHigh := executeSimpleQuery(db, "SELECT MAX(id) FROM users;")
		if errLow != nil || errHigh != nil {
			return
		}
		randomId := random(dataLow, dataHigh)
		query := fmt.Sprintf("Update users SET reputation = %d, display_name = '%s',  website_url = '%s', location = '%s', about_me = '%s', views = %d, upvotes = %d, downvotes = %d WHERE id = %d;",
			data.Reputation, data.DisplayName, data.WebsiteURL, data.Location, data.AboutMe, data.Views, data.Upvotes, data.Downvotes, randomId)
		if err := executeQuery(db, query, id); err != nil {
			return
		}
	}

}
