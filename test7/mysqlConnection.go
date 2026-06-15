package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	fakedata "test7/generatefakedata"

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

func testDbSelect(db *sql.DB, id int, simpleSelect []string) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(simpleSelect))
	fmt.Println("Worker: ", id, " wykonuje zapytanie: ", simpleSelect[randomIndex])
	err := executeQuery(db, simpleSelect[randomIndex], id)
	if err != nil {
		log.Printf("[worker %d] błąd: %v", id, err)
		return err
	}
	return nil

}

func testDbInsert(db *sql.DB, id int) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	whatTable := []string{"badges", "comments", "posthistory", "postlinks", "posts", "users", "votes"}
	randomIndex := r.Intn(len(whatTable))
	switch whatTable[randomIndex] {
	case "badges":
		data := fakedata.GenerateBadge()
		query := fmt.Sprintf("INSERT INTO badges (id, user_id, badge_name, badge_date, class, tag_based) VALUES (%d, %d, '%s', '%s', %d, '%s');",
			data.ID, data.UserID, data.BadgeName, data.BadgeDate.Format("2006-01-02 15:04:05"), data.Class, data.TagBased)

		if err := executeQuery(db, query, id); err != nil {
			return
		}
	case "comments":
		data := fakedata.GenerateComments()
		query := fmt.Sprintf("INSERT INTO comments (id, post_id, score, comment_text, creation_date, user_id, content_license) VALUES (%d, %d, %d, '%s', '%s', %d, '%s');",
			data.ID, data.PostID, data.Score, data.CommentText, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.ContentLicense)

		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "posthistory":
		data := fakedata.GeneratePostHistory()
		query := fmt.Sprintf("INSERT INTO post_history (id, post_history_type_id, post_id, revision_guid, creation_date, user_id, post_text, content_license) VALUES (%d, %d, %d, '%s', '%s', %d, '%s', '%s');",
			data.ID, data.PostHistoryTypeID, data.PostID, data.RevisionGUID, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.PostText, data.ContentLicense)

		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "postlinks":
		data := fakedata.GeneratePostsLinks()
		query := fmt.Sprintf("INSERT INTO post_links (id, creation_date, post_id, related_post_id, link_type_id) VALUES (%d, '%s', %d, %d, %d);",
			data.ID, data.CreationDate.Format("2006-01-02 15:04:05"), data.PostID, data.RelatedPostID, data.LinkTypeID)

		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "posts":
		data := fakedata.GeneratePosts()
		query := fmt.Sprintf("INSERT INTO posts (id, post_type_id, accepted_answer_id, parent_id, creation_date, score, view_count, post_body, owner_user_id, last_editor_user_id, last_edit_date, last_activity_date, post_title, tags, answer_count, comment_count, content_license) VALUES (%d, %d, %d, %d, '%s', %d, %d, '%s', %d, %d, '%s', '%s', '%s', '%s', %d, %d, '%s');",
			data.ID, data.PostTypeID, data.AcceptedAnswerID, data.ParentID, data.CreationDate.Format("2006-01-02 15:04:05"), data.Score, data.ViewCount, data.PostBody, data.OwnerUserID, data.LastEditorUserID, data.LastEditDate.Format("2006-01-02 15:04:05"), data.LastActivityDate.Format("2006-01-02 15:04:05"), data.PostTitle, data.Tags, data.AnswerCount, data.CommentCount, data.ContentLicense)
		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "users":
		data := fakedata.GenerateUsers()
		query := fmt.Sprintf("INSERT INTO users (id, reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes, downvotes, account_id) VALUES (%d, %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d);",
			data.ID, data.Reputation, data.CreationDate.Format("2006-01-02 15:04:05"), data.DisplayName, data.LastAccessDate.Format("2006-01-02 15:04:05"), data.WebsiteURL, data.Location, data.AboutMe, data.Views, data.Upvotes, data.Downvotes, data.AccountID)
		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "votes":
		data := fakedata.GenerateVote()
		query := fmt.Sprintf("INSERT INTO votes (id, post_id, vote_type_id, creation_date) VALUES (%d, %d, %d, '%s');",
			data.ID, data.PostID, data.VoteTypeID, data.CreationDate.Format("2006-01-02 15:04:05"))
		if err := executeQuery(db, query, id); err != nil {
			return
		}
	}

}

func testDbUpdate(db *sql.DB, id int) {

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

func checkConnectionAndRunTestSelect(id int, deadline time.Time, simpleSelect []string) {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	for time.Now().Before(deadline) {

		err := testDbSelect(db, id, simpleSelect)
		if err != nil {
			continue
		}

	}
}

func checkConnectionAndRunTestInsert(id int, deadline time.Time) {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	for time.Now().Before(deadline) {
		testDbInsert(db, id)

	}
}

func checkConnectionAndRunTestUpdate(id int, deadline time.Time) {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	for time.Now().Before(deadline) {
		testDbUpdate(db, id)

	}
}

func multiThreadSelect(workersCount int) {
	var wg sync.WaitGroup

	deadline := time.Now().Add(5 * time.Minute)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			checkConnectionAndRunTestSelect(id, deadline, simpleSelect)

		}(i)
	}

	wg.Wait()
}

func multiThreadInsert(workersCount int) {
	var wg sync.WaitGroup

	deadline := time.Now().Add(5 * time.Minute)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			checkConnectionAndRunTestInsert(id, deadline)
		}(i)
	}

	wg.Wait()
}

func multiThreadUpdate(workersCount int) {
	var wg sync.WaitGroup

	deadline := time.Now().Add(5 * time.Minute)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			checkConnectionAndRunTestUpdate(id, deadline)
		}(i)
	}

	wg.Wait()
}
