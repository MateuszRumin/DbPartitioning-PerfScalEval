package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Comment struct {
	ID             int       `gorm:"primaryKey"`
	PostID         int       `fake:"{number:1,1000}"`
	Score          int       `fake:"{number:0,100}"`
	CommentText    string    `fake:"{sentence:10}"`
	CreationDate   time.Time `fake:"{date}"`
	UserID         int       `fake:"{number:1,100}"`
	ContentLicense string    `fake:"{randomstring:[CC BY-SA, MIT, Apache]}"`
}

func GenerateComments() Comment {
	var generated Comment
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
