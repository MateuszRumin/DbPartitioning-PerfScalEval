package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Vote struct {
	ID           int       `gorm:"primaryKey"`
	PostID       int       `fake:"{number:1,1000}"`
	VoteTypeID   int       `fake:"{number:1,10}"`
	CreationDate time.Time `fake:"{date}"`
}

func GenerateVote() Vote {
	var generated Vote
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
