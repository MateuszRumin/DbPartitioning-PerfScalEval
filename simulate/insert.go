package main

import (
	"fmt"
	"log"
	"math/rand"
	fakedata "simulate/generatefakedata"
	"time"
)

func ChooseTableInsert() string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	whatTable := []string{"badges", "comments", "posthistory", "postlinks", "posts", "users", "votes"}
	randomIndex := r.Intn(len(whatTable))
	switch whatTable[randomIndex] {
	case "badges":
		data := fakedata.GenerateBadge()
		return fmt.Sprintf("INSERT INTO badges (user_id, badge_name, badge_date, class, tag_based) VALUES ( %d, '%s', '%s', %d, '%s');",
			data.UserID, data.BadgeName, data.BadgeDate.Format("2006-01-02 15:04:05"), data.Class, data.TagBased)

	case "comments":
		data := fakedata.GenerateComments()
		return fmt.Sprintf("INSERT INTO comments (post_id, score, comment_text, creation_date, user_id, content_license) VALUES ( %d, %d, '%s', '%s', %d, '%s');",
			data.PostID, data.Score, data.CommentText, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.ContentLicense)

	case "posthistory":
		data := fakedata.GeneratePostHistory()
		return fmt.Sprintf("INSERT INTO post_history (post_history_type_id, post_id, revision_guid, creation_date, user_id, post_text, content_license) VALUES ( %d, %d, '%s', '%s', %d, '%s', '%s');",
			data.PostHistoryTypeID, data.PostID, data.RevisionGUID, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.PostText, data.ContentLicense)

	case "postlinks":
		data := fakedata.GeneratePostsLinks()
		return fmt.Sprintf("INSERT INTO post_links ( creation_date, post_id, related_post_id, link_type_id) VALUES ('%s', %d, %d, %d);",
			data.CreationDate.Format("2006-01-02 15:04:05"), data.PostID, data.RelatedPostID, data.LinkTypeID)

	case "posts":
		data := fakedata.GeneratePosts()
		return fmt.Sprintf("INSERT INTO posts (post_type_id, accepted_answer_id, parent_id, creation_date, score, view_count, post_body, owner_user_id, last_editor_user_id, last_edit_date, last_activity_date, post_title, tags, answer_count, comment_count, content_license) VALUES ( %d, %d, %d, '%s', %d, %d, '%s', %d, %d, '%s', '%s', '%s', '%s', %d, %d, '%s');",
			data.PostTypeID, data.AcceptedAnswerID, data.ParentID, data.CreationDate.Format("2006-01-02 15:04:05"), data.Score, data.ViewCount, data.PostBody, data.OwnerUserID, data.LastEditorUserID, data.LastEditDate.Format("2006-01-02 15:04:05"), data.LastActivityDate.Format("2006-01-02 15:04:05"), data.PostTitle, data.Tags, data.AnswerCount, data.CommentCount, data.ContentLicense)

	case "users":
		data := fakedata.GenerateUsers()
		return fmt.Sprintf("INSERT INTO users ( reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes, downvotes, account_id) VALUES ( %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d);",
			data.Reputation, data.CreationDate.Format("2006-01-02 15:04:05"), data.DisplayName, data.LastAccessDate.Format("2006-01-02 15:04:05"), data.WebsiteURL, data.Location, data.AboutMe, data.Views, data.Upvotes, data.Downvotes, data.AccountID)

	case "votes":
		data := fakedata.GenerateVote()
		return fmt.Sprintf("INSERT INTO votes (post_id, vote_type_id, creation_date) VALUES (%d, %d, '%s');",
			data.PostID, data.VoteTypeID, data.CreationDate.Format("2006-01-02 15:04:05"))

	}
	return ""

}

func insertTest(deadline time.Time, id int) {
	// if id > 10 && id < 21 {
	// 	time.Sleep(20 * time.Second)
	// }
	db, err := setConnection()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	}
	defer db.Close()

	var qr []QueryResults

	for time.Now().Before(deadline) {

		query := ChooseTableInsert()

		start := time.Now()

		if err := executeQuery(db, query, id); err != nil {
			return
		} else {
			stop := time.Now()
			duration := time.Since(start)

			qr = append(qr, QueryResults{
				qtype:    "Insert",
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

		db2.Query(fmt.Sprintf("INSERT INTO QueryResults (query_type,timeEnded,duration_ms) VALUES ('%s','%s','%d')",
			d.qtype, d.end.Format("2006-01-02 15:04:05"), d.duration.Milliseconds()))

	}

}
