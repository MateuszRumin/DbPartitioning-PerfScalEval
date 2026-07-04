package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

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

func GeneratePostHistory() PostHistory {
	var generated PostHistory
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
