package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

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

func GenerateUsers() User {
	var generated User
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
