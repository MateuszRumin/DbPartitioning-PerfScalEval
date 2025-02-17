package generatefakedata

import (
	"fmt"

	"perfscaleval/dbstrucglobalfakeit"

	"github.com/brianvoe/gofakeit/v7"
)

func GenerateBadge() dbstrucglobalfakeit.Badge {

	var generated dbstrucglobalfakeit.Badge

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}

	return generated
}

func GenerateComments() dbstrucglobalfakeit.Comment {
	var generated dbstrucglobalfakeit.Comment
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GeneratePostHistory() dbstrucglobalfakeit.PostHistory {
	var generated dbstrucglobalfakeit.PostHistory
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GeneratePostsLinks() dbstrucglobalfakeit.PostLink {
	var generated dbstrucglobalfakeit.PostLink
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GeneratePosts() dbstrucglobalfakeit.Post {
	var generated dbstrucglobalfakeit.Post

	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GenerateUsers() dbstrucglobalfakeit.User {
	var generated dbstrucglobalfakeit.User
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}

func GenerateVote() dbstrucglobalfakeit.Vote {
	var generated dbstrucglobalfakeit.Vote
	err := gofakeit.Struct(&generated)
	if err != nil {
		fmt.Println("Błąd generacji")
	}
	return generated
}
