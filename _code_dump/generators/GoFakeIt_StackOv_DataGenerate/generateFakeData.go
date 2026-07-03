package main

import (
	"fmt"
)

func main() {

	badge := generateBadge()
	comment := GenerateComments()
	postHistory := GeneratePostHistory()
	postLink := GeneratePostsLinks()
	post := GeneratePosts()
	user := GenerateUsers()
	vote := GenerateVote()

	fmt.Println("badge")
	fmt.Printf("%+v\n", badge)
	fmt.Println("comment")
	fmt.Printf("%+v\n", comment)
	fmt.Println("postHistory")
	fmt.Printf("%+v\n", postHistory)
	fmt.Println("postLink")
	fmt.Printf("%+v\n", postLink)
	fmt.Println("post")
	fmt.Printf("%+v\n", post)
	fmt.Println("user")
	fmt.Printf("%+v\n", user)
	fmt.Println("vote")
	fmt.Printf("%+v\n", vote)

}
