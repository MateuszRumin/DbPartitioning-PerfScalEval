package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Badge struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `fake:"{number:1,100}"`
	BadgeName string    `fake:"{word}"`
	BadgeDate time.Time `fake:"{date}"`
	Class     int       `fake:"{number:1,3}"`
	TagBased  string    `fake:"{randomstring:[true,false]}"`
}

func generateBadge() Badge {

	var generated Badge

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}

	return generated
}
