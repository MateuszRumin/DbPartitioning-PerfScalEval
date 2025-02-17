package test

import (
	"fmt"
	"perfscaleval/generatefakedata"
)

func TestGenerateDataFromOneSourceToAlltable() {

	badge := generatefakedata.GenerateBadge()
	comment := generatefakedata.GenerateComments()
	postHistory := generatefakedata.GeneratePostHistory()
	postLink := generatefakedata.GeneratePostsLinks()
	post := generatefakedata.GeneratePosts()
	user := generatefakedata.GenerateUsers()
	vote := generatefakedata.GenerateVote()

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

	fmt.Println("first test")

}
