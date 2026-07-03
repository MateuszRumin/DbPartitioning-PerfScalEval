package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

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

func GeneratePosts() Post {
	var generated Post

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
