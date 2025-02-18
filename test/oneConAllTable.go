package test

import (
	"perfscaleval/generatefakedata"
	"perfscaleval/queries"
)

func oneConAllTable() {

	badge := generatefakedata.GenerateBadge()
	comment := generatefakedata.GenerateComments()
	postHistory := generatefakedata.GeneratePostHistory()
	postLink := generatefakedata.GeneratePostsLinks()
	post := generatefakedata.GeneratePosts()
	user := generatefakedata.GenerateUsers()
	vote := generatefakedata.GenerateVote()

	queries.CreateBadge(badge, 0)
	queries.CreateComment(comment, 0)
	queries.CreatePostHistory(postHistory, 0)
	queries.CreatePostLink(postLink, 0)
	queries.CreatePost(post, 0)
	queries.CreateUser(user, 0)
	queries.CreateVote(vote, 0)

	// fmt.Println("badge")
	// fmt.Printf("%+v\n", badge)
	// fmt.Println("comment")
	// fmt.Printf("%+v\n", comment)
	// fmt.Println("postHistory")
	// fmt.Printf("%+v\n", postHistory)
	// fmt.Println("postLink")
	// fmt.Printf("%+v\n", postLink)
	// fmt.Println("post")
	// fmt.Printf("%+v\n", post)
	// fmt.Println("user")
	// fmt.Printf("%+v\n", user)
	// fmt.Println("vote")
	// fmt.Printf("%+v\n", vote)

	// fmt.Println("first test")

}
