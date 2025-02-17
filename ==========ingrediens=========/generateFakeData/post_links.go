package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type PostLink struct {
	ID            int       `gorm:"primaryKey"`
	CreationDate  time.Time `fake:"{date}"`
	PostID        int       `fake:"{number:1,1000}"` // powiązany post
	RelatedPostID int       `fake:"{number:1,1000}"` // post powiązany
	LinkTypeID    int       `fake:"{number:1,5}"`    // typ linku
}

func GeneratePostsLinks() PostLink {
	var generated PostLink
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
