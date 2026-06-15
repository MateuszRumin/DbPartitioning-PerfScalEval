package generatefakedata

import (
	"fmt"

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
	return generated
}

func GeneratePostHistory() PostHistory {
	var generated PostHistory

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GeneratePostsLinks() PostLink {
	var generated PostLink

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GeneratePosts() Post {
	var generated Post

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GenerateUsers() User {
	var generated User

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GenerateVote() Vote {
	var generated Vote

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
