package dbstrucglobalfakeit

import (
	"time"
)

type Badge struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `fake:"{number:1,100}"`
	BadgeName string    `fake:"{word}"`
	BadgeDate time.Time `fake:"{date}"`
	Class     int       `fake:"{number:1,3}"`
	TagBased  string    `fake:"{randomstring:[true,false]}"`
}

type Comment struct {
	ID             int       `gorm:"primaryKey"`
	PostID         int       `fake:"{number:1,1000}"`
	Score          int       `fake:"{number:0,100}"`
	CommentText    string    `fake:"{sentence:10}"`
	CreationDate   time.Time `fake:"{date}"`
	UserID         int       `fake:"{number:1,100}"`
	ContentLicense string    `fake:"{randomstring:[CC BY-SA, MIT, Apache]}"`
}

type PostHistory struct {
	ID                int       `gorm:"primaryKey"`
	PostHistoryTypeID int       `fake:"{number:1,10}"`
	PostID            int       `fake:"{number:1,1000}"`
	RevisionGUID      string    `fake:"{uuid}"`
	CreationDate      time.Time `fake:"{date}"`
	UserID            int       `fake:"{number:1,100}"`
	PostText          string    `fake:"{paragraph}"`
	ContentLicense    string    `fake:"{randomstring:[CC BY-SA, MIT, Apache]}"`
}

type PostLink struct {
	ID            int       `gorm:"primaryKey"`
	CreationDate  time.Time `fake:"{date}"`
	PostID        int       `fake:"{number:1,1000}"`
	RelatedPostID int       `fake:"{number:1,1000}"`
	LinkTypeID    int       `fake:"{number:1,5}"`
}

type Post struct {
	ID               int       `gorm:"primaryKey"`
	PostTypeID       int       `fake:"{number:1,10}"`
	AcceptedAnswerID int       `fake:"{number:1,100}"`
	ParentID         int       `fake:"{number:1,100}"`
	CreationDate     time.Time `fake:"{date}"`
	Score            int       `fake:"{number:0,1000}"`
	ViewCount        int       `fake:"{number:0,100000}"`
	PostBody         string    `fake:"{sentence:30}"`
	OwnerUserID      int       `fake:"{number:1,100}"`
	LastEditorUserID int       `fake:"{number:1,100}"`
	LastEditDate     time.Time `fake:"{date}"`
	LastActivityDate time.Time `fake:"{date}"`
	PostTitle        string    `fake:"{sentence:5}"`
	Tags             string    `fake:"{word}"`
	AnswerCount      int       `fake:"{number:0,50}"`
	CommentCount     int       `fake:"{number:0,200}"`
	ContentLicense   string    `fake:"{randomstring:[CC BY-SA, MIT, Apache]}"`
}

// id, reputation, creation_date, display_name, last_access_date, website_url, location, about_me, views, upvotes,downvotes, account_id
type User struct {
	ID             int       `gorm:"primaryKey"`
	Reputation     int       `fake:"{number:0,100000}"`
	CreationDate   time.Time `fake:"{date}"`
	DisplayName    string    `fake:"{username}"`
	LastAccessDate time.Time `fake:"{date}"`
	WebsiteURL     string    `fake:"{url}"`
	Location       string    `fake:"{city}"`
	AboutMe        string    `fake:"{paragraph:3}"`
	Views          int       `fake:"{number:0,1000}"`
	Upvotes        int       `fake:"{number:0,500}"`
	Downvotes      int       `fake:"{number:0,200}"`
	AccountID      int       `fake:"{number:1,1000}"`
}

type Vote struct {
	ID           int       `gorm:"primaryKey"`
	PostID       int       `fake:"{number:1,1000}"`
	VoteTypeID   int       `fake:"{number:1,10}"`
	CreationDate time.Time `fake:"{date}"`
}
