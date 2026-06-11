package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	fakedata "test3/generatefakedata"

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

func checkConnectionAndRunTest(id int) {
	// if id > 10 && id < 21 {
	// 	time.Sleep(20 * time.Second)
	// }
	db, err := setConnection()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	} else {
		fmt.Println("Połączenie z bazą danych działa poprawnie. Wątek: ", id)
		for l := 0; l < 9999; l++ {
			testDb(db, id)
		}
	}
	defer db.Close()

}

func multiThread(workersCount int) {

	if workersCount < 1 {
		fmt.Println("Liczba workerów musi być większa niż 0.")
		return
	} else if workersCount == 1 {
		checkConnectionAndRunTest(0)
	} else if workersCount > 1 {
		var wg sync.WaitGroup

		for i := 0; i < 20; i++ {
			wg.Add(1)

			go func(id int) {
				defer wg.Done()
				checkConnectionAndRunTest(id + 1)
			}(i)
		}

		wg.Wait()

	}

}

func executeQuery(db *sql.DB, query string, id int) {
	// fmt.Printf("Wykonuję zapytanie: %s\n", query)
	rows, err := db.Query(query)
	if err != nil {
		//log.Printf("Zapytanie, które spowodowało błąd: %s\n", query)
		log.Printf("Nie udało się wykonać zapytania: \n")
		return

	}
	fmt.Printf("Succes wątek id: %d \n", id)
	defer rows.Close()

}

func testDb(db *sql.DB, id int) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	whatTable := []string{"badges", "comments", "posthistory", "postlinks", "posts", "users", "votes"}
	randomIndex := r.Intn(len(whatTable))
	if randomIndex == 0 {
		data := fakedata.GenerateBadge()
		query := fmt.Sprintf("INSERT INTO badges (id, user_id, badge_name, badge_date, class, tag_based) VALUES (%d, %d, '%s', '%s', %d, '%s');",
			data.ID, data.UserID, data.BadgeName, data.BadgeDate.Format("2006-01-02 15:04:05"), data.Class, data.TagBased)

		executeQuery(db, query, id)
	} else if randomIndex == 1 {
		data := fakedata.GenerateComments()
		query := fmt.Sprintf("INSERT INTO comments (id, post_id, score, comment_text, creation_date, user_id, content_license) VALUES (%d, %d, %d, '%s', '%s', %d, '%s');",
			data.ID, data.PostID, data.Score, data.CommentText, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.ContentLicense)

		executeQuery(db, query, id)

	} else if randomIndex == 2 {
		data := fakedata.GeneratePostHistory()
		query := fmt.Sprintf("INSERT INTO post_history (id, post_history_type_id, post_id, revision_guid, creation_date, user_id, post_text, content_license) VALUES (%d, %d, %d, '%s', '%s', %d, '%s', '%s');",
			data.ID, data.PostHistoryTypeID, data.PostID, data.RevisionGUID, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.PostText, data.ContentLicense)

		executeQuery(db, query, id)

	} else if randomIndex == 3 {
		data := fakedata.GeneratePostsLinks()
		query := fmt.Sprintf("INSERT INTO post_links (id, creation_date, post_id, related_post_id, link_type_id) VALUES (%d, '%s', %d, %d, %d);",
			data.ID, data.CreationDate.Format("2006-01-02 15:04:05"), data.PostID, data.RelatedPostID, data.LinkTypeID)

		executeQuery(db, query, id)

	} else if randomIndex == 4 {
		data := fakedata.GeneratePosts()
		query := fmt.Sprintf("INSERT INTO posts (id, post_type_id, accepted_answer_id, parent_id, creation_date, score, view_count, post_body, owner_user_id, last_editor_user_id, last_edit_date, last_activity_date, post_title, tags, answer_count, comment_count, content_license) VALUES (%d, %d, %d, %d, '%s', %d, %d, '%s', %d, %d, '%s', '%s', '%s', '%s', %d, %d, '%s');",
			data.ID, data.PostTypeID, data.AcceptedAnswerID, data.ParentID, data.CreationDate.Format("2006-01-02 15:04:05"), data.Score, data.ViewCount, data.PostBody, data.OwnerUserID, data.LastEditorUserID, data.LastEditDate.Format("2006-01-02 15:04:05"), data.LastActivityDate.Format("2006-01-02 15:04:05"), data.PostTitle, data.Tags, data.AnswerCount, data.CommentCount, data.ContentLicense)
		executeQuery(db, query, id)

	} else if randomIndex == 5 {
		data := fakedata.GenerateUsers()
		query := fmt.Sprintf("INSERT INTO users (id, reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes, downvotes, account_id) VALUES (%d, %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d);",
			data.ID, data.Reputation, data.CreationDate.Format("2006-01-02 15:04:05"), data.DisplayName, data.LastAccessDate.Format("2006-01-02 15:04:05"), data.WebsiteURL, data.Location, data.AboutMe, data.Views, data.Upvotes, data.Downvotes, data.AccountID)
		executeQuery(db, query, id)

	} else if randomIndex == 6 {
		data := fakedata.GenerateVote()
		query := fmt.Sprintf("INSERT INTO votes (id, post_id, vote_type_id, creation_date) VALUES (%d, %d, %d, '%s');",
			data.ID, data.PostID, data.VoteTypeID, data.CreationDate.Format("2006-01-02 15:04:05"))
		executeQuery(db, query, id)
	}

}
