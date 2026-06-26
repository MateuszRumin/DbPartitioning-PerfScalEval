package main

import (
	"database/sql"
	"fmt"
	fakedata "insertSPs/generatefakedata"
	"math/rand"
	"time"
)

func testDb(db *sql.DB, id int) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	whatTable := []string{"badges", "comments", "posthistory", "postlinks", "posts", "users", "votes"}
	randomIndex := r.Intn(len(whatTable))
	switch whatTable[randomIndex] {
	case "badges":
		data := fakedata.GenerateBadge()
		query := fmt.Sprintf("INSERT INTO badges (user_id, badge_name, badge_date, class, tag_based) VALUES ( %d, '%s', '%s', %d, '%s');",
			data.UserID, data.BadgeName, data.BadgeDate.Format("2006-01-02 15:04:05"), data.Class, data.TagBased)

		if err := executeQuery(db, query, id); err != nil {
			return
		}
	case "comments":
		data := fakedata.GenerateComments()
		query := fmt.Sprintf("INSERT INTO comments (post_id, score, comment_text, creation_date, user_id, content_license) VALUES ( %d, %d, '%s', '%s', %d, '%s');",
			data.PostID, data.Score, data.CommentText, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.ContentLicense)

		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "posthistory":
		data := fakedata.GeneratePostHistory()
		query := fmt.Sprintf("INSERT INTO post_history (post_history_type_id, post_id, revision_guid, creation_date, user_id, post_text, content_license) VALUES ( %d, %d, '%s', '%s', %d, '%s', '%s');",
			data.PostHistoryTypeID, data.PostID, data.RevisionGUID, data.CreationDate.Format("2006-01-02 15:04:05"), data.UserID, data.PostText, data.ContentLicense)

		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "postlinks":
		data := fakedata.GeneratePostsLinks()
		query := fmt.Sprintf("INSERT INTO post_links ( creation_date, post_id, related_post_id, link_type_id) VALUES ('%s', %d, %d, %d);",
			data.CreationDate.Format("2006-01-02 15:04:05"), data.PostID, data.RelatedPostID, data.LinkTypeID)

		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "posts":
		data := fakedata.GeneratePosts()
		query := fmt.Sprintf("INSERT INTO posts (post_type_id, accepted_answer_id, parent_id, creation_date, score, view_count, post_body, owner_user_id, last_editor_user_id, last_edit_date, last_activity_date, post_title, tags, answer_count, comment_count, content_license) VALUES ( %d, %d, %d, '%s', %d, %d, '%s', %d, %d, '%s', '%s', '%s', '%s', %d, %d, '%s');",
			data.PostTypeID, data.AcceptedAnswerID, data.ParentID, data.CreationDate.Format("2006-01-02 15:04:05"), data.Score, data.ViewCount, data.PostBody, data.OwnerUserID, data.LastEditorUserID, data.LastEditDate.Format("2006-01-02 15:04:05"), data.LastActivityDate.Format("2006-01-02 15:04:05"), data.PostTitle, data.Tags, data.AnswerCount, data.CommentCount, data.ContentLicense)
		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "users":
		data := fakedata.GenerateUsers()
		query := fmt.Sprintf("INSERT INTO users ( reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes, downvotes, account_id) VALUES ( %d, '%s', '%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d);",
			data.Reputation, data.CreationDate.Format("2006-01-02 15:04:05"), data.DisplayName, data.LastAccessDate.Format("2006-01-02 15:04:05"), data.WebsiteURL, data.Location, data.AboutMe, data.Views, data.Upvotes, data.Downvotes, data.AccountID)
		if err := executeQuery(db, query, id); err != nil {
			return
		}

	case "votes":
		data := fakedata.GenerateVote()
		query := fmt.Sprintf("INSERT INTO votes (post_id, vote_type_id, creation_date) VALUES (%d, %d, '%s');",
			data.PostID, data.VoteTypeID, data.CreationDate.Format("2006-01-02 15:04:05"))
		if err := executeQuery(db, query, id); err != nil {
			return
		}
	}

}
