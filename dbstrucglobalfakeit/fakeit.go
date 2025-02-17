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
	PostID        int       `fake:"{number:1,1000}"` // powiązany post
	RelatedPostID int       `fake:"{number:1,1000}"` // post powiązany
	LinkTypeID    int       `fake:"{number:1,5}"`    // typ linku
}

type Post struct {
	ID               int       `gorm:"primaryKey"`
	PostTypeID       int       `fake:"{number:1,10}"`  // Usunięto skip, dodano generator
	AcceptedAnswerID int       `fake:"{number:1,100}"` // Usunięto skip, dodano generator
	ParentID         int       `fake:"{number:1,100}"` // Usunięto skip, dodano generator
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

type User struct {
	ID             int       `gorm:"primaryKey"`
	Reputation     int       `fake:"{number:0,100000}"` // reputacja może być duża
	CreationDate   time.Time `fake:"{date}"`            // data rejestracji
	DisplayName    string    `fake:"{username}"`        // przykładowa nazwa użytkownika
	LastAccessDate time.Time `fake:"{date}"`            // ostatni dostęp
	WebsiteURL     string    `fake:"{url}"`             // adres strony WWW
	Location       string    `fake:"{city}"`            // lokalizacja
	AboutMe        string    `fake:"{paragraph:3}"`     // dłuższy opis
	Views          int       `fake:"{number:0,1000}"`   // liczba wyświetleń profilu
	Upvotes        int       `fake:"{number:0,500}"`    // głosy pozytywne
	Downvotes      int       `fake:"{number:0,200}"`    // głosy negatywne
	AccountID      int       `fake:"{number:1,1000}"`   // powiązane konto
}

type Vote struct {
	ID           int       `gorm:"primaryKey"`
	PostID       int       `fake:"{number:1,1000}"`
	VoteTypeID   int       `fake:"{number:1,10}"`
	CreationDate time.Time `fake:"{date}"`
}
