package generatefakedata

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func GenerateBadge() Badge {

	var generated Badge

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}

	return generated
}

func GenerateComments() Comment {
	var generated Comment

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	generated.CreationDate = gofakeit.DateRange(
		time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC),
	)
	return generated
}

func GeneratePostHistory() PostHistory {
	var generated PostHistory

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	generated.CreationDate = gofakeit.DateRange(
		time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC),
	)
	return generated
}

func GeneratePostsLinks() PostLink {
	var generated PostLink

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	generated.CreationDate = gofakeit.DateRange(
		time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC),
	)
	return generated
}

func GeneratePosts() Post {
	var generated Post

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	// generated.CreationDate = gofakeit.DateRange(
	// 	time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC),
	// 	time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC),
	// )
	generated.CreationDate = gofakeit.DateRange(
		time.Date(2008, 7, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2008, 9, 30, 23, 59, 59, 999999999, time.UTC),
	)
	return generated
}

func GenerateUsers() User {
	var generated User

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	generated.CreationDate = gofakeit.DateRange(
		time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC),
	)
	return generated
}

func GenerateVote() Vote {
	var generated Vote

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	generated.CreationDate = gofakeit.DateRange(
		time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC),
	)
	return generated
}
