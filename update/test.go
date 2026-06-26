package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	fakedata "update/generatefakedata"
)

type QueryResults struct {
	qtype    string
	end      time.Time
	duration time.Duration
}

func chooseTable(idb int, idc int, idph int, idp int, idu int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	whatTable := []string{"badges", "comments", "posthistory", "posts", "users"}
	randomIndex := r.Intn(len(whatTable))

	switch whatTable[randomIndex] {
	case "badges":
		data := fakedata.GenerateBadge()
		return fmt.Sprintf("Update badges SET badge_name = '%s', badge_date = '%s', class = %d, tag_based = '%s' WHERE id = %d;",
			data.BadgeName, data.BadgeDate.Format("2006-01-02 15:04:05"), data.Class, data.TagBased, r.Intn(idb)+1)

	case "comments":
		data := fakedata.GenerateComments()
		return fmt.Sprintf("Update comments SET score = %d, comment_text = '%s', creation_date = '%s', user_id = %d, content_license = '%s' WHERE id = %d;",
			data.Score, data.CommentText, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.ContentLicense, r.Intn(idc)+1)

	case "posthistory":
		data := fakedata.GeneratePostHistory()

		return fmt.Sprintf("Update post_history SET post_history_type_id = %d, revision_guid = '%s', creation_date = '%s', post_text = '%s', content_license = '%s' WHERE id = %d;",
			data.PostHistoryTypeID, data.RevisionGUID, data.CreationDate.Format("2006-01-02 15:04:05"), data.PostText, data.ContentLicense, r.Intn(idph)+1)

	case "posts":
		data := fakedata.GeneratePosts()

		return fmt.Sprintf("Update posts SET   score = %d, view_count = %d, post_body = '%s',   last_edit_date = '%s', last_activity_date = '%s', post_title = '%s', tags = '%s', answer_count = %d, comment_count = %d, content_license = '%s' WHERE id = %d;",
			data.Score, data.ViewCount, data.PostBody, data.LastEditDate.Format("2006-01-02 15:04:05"), data.LastActivityDate.Format("2006-01-02 15:04:05"),
			data.PostTitle, data.Tags, data.AnswerCount, data.CommentCount, data.ContentLicense, r.Intn(idp)+1)

	case "users":
		data := fakedata.GenerateUsers()

		return fmt.Sprintf("Update users SET reputation = %d, display_name = '%s',  website_url = '%s', location = '%s', about_me = '%s', views = %d, upvotes = %d, downvotes = %d WHERE id = %d;",
			data.Reputation, data.DisplayName, data.WebsiteURL, data.Location, data.AboutMe, data.Views, data.Upvotes, data.Downvotes, r.Intn(idu)+1)

	}

	return ""

}

func checkConnectionAndRunTest(id int, deadline time.Time, idb int, idc int, idph int, idp int, idu int) {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	var qr []QueryResults

	for time.Now().Before(deadline) {

		query := chooseTable(idb, idc, idph, idp, idu)
		start := time.Now()
		err := executeQuery(db, query, id)
		if err != nil {
			continue
		} else {
			stop := time.Now()
			duration := time.Since(start)

			qr = append(qr, QueryResults{
				qtype:    "Update",
				end:      stop,
				duration: duration,
			})
		}

	}
	db2, err := slc()
	if err != nil {

		return
	}
	defer db.Close()

	for _, d := range qr {

		db2.Query(fmt.Sprintf("INSERT INTO QueryResults (query_type,timeEnded,duration_ms) VALUES ('%s','%s','%d')", d.qtype, d.end.Format("2006-01-02 15:04:05"), d.duration.Milliseconds()))

	}

}
