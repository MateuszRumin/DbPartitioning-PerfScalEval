package queries

import (
	"fmt"
	"perfscaleval/config"
	"perfscaleval/dbstrucglobalfakeit"
)

// var (
//
//	createBadge       string = "INSERT INTO badges (id, user_id, badge_name, badge_date, class, tag_based) VALUES (?, ?, ?, ?, ?, ?);"
//	createComment     string = "INSERT INTO comments (id, post_id, score, comment_text, creation_date, user_id, content_license) VALUES (?, ?, ?, ?, ?, ?, ?);"
//	createPostHistory string = "INSERT INTO post_history (id, post_history_type_id, post_id, revision_guid, creation_date, user_id, post_text, content_license) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
//	createPostLink    string = "INSERT INTO post_links (id, creation_date, post_id, related_post_id, link_type_id) VALUES (?, ?, ?, ?, ?);"
//	createPost        string = "INSERT INTO posts ( id, post_type_id, accepted_answer_id, parent_id, creation_date, score, view_count, post_body, owner_user_id, last_editor_user_id, last_edit_date, last_activity_date, post_title, tags, answer_count, comment_count, content_license ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
//	createUser        string = "INSERT INTO users (id, reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes,downvotes, account_id ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
//	createVote        string = "INSERT INTO votes (id, post_id, vote_type_id, creation_date) VALUES (?, ?, ?, ?);"
//
// )
var (
	createBadge       string = "INSERT INTO badges (user_id, badge_name, badge_date, class, tag_based) VALUES (?, ?, ?, ?, ?);"
	createComment     string = "INSERT INTO comments (post_id, score, comment_text, creation_date, user_id, content_license) VALUES (?, ?, ?, ?, ?, ?);"
	createPostHistory string = "INSERT INTO post_history (post_history_type_id, post_id, revision_guid, creation_date, user_id, post_text, content_license) VALUES (?, ?, ?, ?, ?, ?, ?);"
	createPostLink    string = "INSERT INTO post_links (creation_date, post_id, related_post_id, link_type_id) VALUES (?, ?, ?, ?);"
	createPost        string = "INSERT INTO posts (post_type_id, accepted_answer_id, parent_id, creation_date, score, view_count, post_body, owner_user_id, last_editor_user_id, last_edit_date, last_activity_date, post_title, tags, answer_count, comment_count, content_license ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	createUser        string = "INSERT INTO users (reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes,downvotes, account_id ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	createVote        string = "INSERT INTO votes (post_id, vote_type_id, creation_date) VALUES (?, ?, ?);"
)

// id, user_id, badge_name, badge_date, class, tag_based
func CreateBadge(data dbstrucglobalfakeit.Badge, conNumer int) {
	_, err := config.Connections[conNumer].Query(createBadge,
		data.UserID,
		data.BadgeName,
		data.BadgeDate,
		data.Class,
		data.TagBased,
	)
	if err != nil {
		fmt.Printf("Błąd dodawania Badge: %v \n", err)
	}
}

// id, post_id, score, comment_text, creation_date, user_id, content_license
func CreateComment(data dbstrucglobalfakeit.Comment, conNumer int) {
	_, err := config.Connections[conNumer].Query(createComment,
		data.PostID,
		data.Score,
		data.CommentText,
		data.CreationDate,
		data.UserID,
		data.ContentLicense,
	)
	if err != nil {
		fmt.Printf("Błąd dodawania Comment %v \n", err)
	}
}

// id, post_history_type_id, post_id, revision_guid, creation_date, user_id, post_text, content_license
func CreatePostHistory(data dbstrucglobalfakeit.PostHistory, conNumer int) {
	_, err := config.Connections[conNumer].Query(createPostHistory,
		data.PostHistoryTypeID,
		data.PostID,
		data.RevisionGUID,
		data.CreationDate,
		data.UserID,
		data.PostText,
		data.ContentLicense,
	)
	if err != nil {
		fmt.Printf("Błąd dodawaniaPostHistory%v \n", err)
	}
}

// id, creation_date, post_id, related_post_id, link_type_id
func CreatePostLink(data dbstrucglobalfakeit.PostLink, conNumer int) {
	_, err := config.Connections[conNumer].Query(createPostLink,
		data.CreationDate,
		data.PostID,
		data.RelatedPostID,
		data.LinkTypeID,
	)
	if err != nil {
		fmt.Printf("Błąd dodawania PostLink %v \n", err)
	}
}

// id, post_type_id, accepted_answer_id, parent_id, creation_date, score, view_count, post_body, owner_user_id, last_editor_user_id, last_edit_date, last_activity_date, post_title, tags, answer_count, comment_count, content_license
func CreatePost(data dbstrucglobalfakeit.Post, conNumer int) {
	_, err := config.Connections[conNumer].Query(createPost,
		data.PostTypeID,
		data.AcceptedAnswerID,
		data.ParentID,
		data.CreationDate,
		data.Score,
		data.ViewCount,
		data.PostBody,
		data.OwnerUserID,
		data.LastEditorUserID,
		data.LastEditDate,
		data.LastActivityDate,
		data.PostTitle,
		data.Tags,
		data.AnswerCount,
		data.CommentCount,
		data.ContentLicense,
	)
	if err != nil {
		fmt.Printf("Błąd dodawania Post %v \n", err)
	}

}

// id, reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes,downvotes, account_id
func CreateUser(data dbstrucglobalfakeit.User, conNumer int) {
	_, err := config.Connections[conNumer].Query(createUser,
		data.Reputation,
		data.CreationDate,
		data.DisplayName,
		data.LastAccessDate,
		data.WebsiteURL,
		data.Location,
		data.AboutMe,
		data.Views,
		data.Upvotes,
		data.Downvotes,
		data.AccountID,
	)
	if err != nil {
		fmt.Printf("Błąd dodawaniaUser %v \n", err)
	}
}

// id, post_id, vote_type_id, creation_date
func CreateVote(data dbstrucglobalfakeit.Vote, conNumer int) {
	_, err := config.Connections[conNumer].Query(createVote,
		data.PostID,
		data.VoteTypeID,
		data.CreationDate,
	)
	if err != nil {
		fmt.Printf("Błąd dodawaniaVote %v \n", err)
	}
}
