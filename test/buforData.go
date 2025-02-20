package test

import (
	"perfscaleval/config"
	"perfscaleval/dbstrucglobalfakeit"
	"perfscaleval/generatefakedata"
)

type BuforData struct {
	Badges      []dbstrucglobalfakeit.Badge
	Comments    []dbstrucglobalfakeit.Comment
	PostHistory []dbstrucglobalfakeit.PostHistory
	PostLinks   []dbstrucglobalfakeit.PostLink
	Posts       []dbstrucglobalfakeit.Post
	Users       []dbstrucglobalfakeit.User
	Votes       []dbstrucglobalfakeit.Vote
}

func makeBuforDataForThread() BuforData {
	bufor := BuforData{
		Badges:      make([]dbstrucglobalfakeit.Badge, 0, config.BuforCreateSize),
		Comments:    make([]dbstrucglobalfakeit.Comment, 0, config.BuforCreateSize),
		PostHistory: make([]dbstrucglobalfakeit.PostHistory, 0, config.BuforCreateSize),
		PostLinks:   make([]dbstrucglobalfakeit.PostLink, 0, config.BuforCreateSize),
		Posts:       make([]dbstrucglobalfakeit.Post, 0, config.BuforCreateSize),
		Users:       make([]dbstrucglobalfakeit.User, 0, config.BuforCreateSize),
		Votes:       make([]dbstrucglobalfakeit.Vote, 0, config.BuforCreateSize),
	}

	// Generowanie danych i dodawanie ich do bufora
	for i := 0; i < config.BuforCreateSize; i++ {
		bufor.Badges = append(bufor.Badges, generatefakedata.GenerateBadge())
		bufor.Comments = append(bufor.Comments, generatefakedata.GenerateComments())
		bufor.PostHistory = append(bufor.PostHistory, generatefakedata.GeneratePostHistory())
		bufor.PostLinks = append(bufor.PostLinks, generatefakedata.GeneratePostsLinks())
		bufor.Posts = append(bufor.Posts, generatefakedata.GeneratePosts())
		bufor.Users = append(bufor.Users, generatefakedata.GenerateUsers())
		bufor.Votes = append(bufor.Votes, generatefakedata.GenerateVote())
	}

	return bufor
}
