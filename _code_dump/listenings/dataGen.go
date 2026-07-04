
type Post struct {
	ID               int32     `gorm:"primaryKey;"`
	PostTypeID       int       `fake:"{number:1,10}"`
	AcceptedAnswerID int       `fake:"{number:1,100}"`
	ParentID         int       `fake:"{number:1,100}"`
	CreationDate     time.Time `fake:"{date}"`
	Score            int       `fake:"{number:0,1000}"`
	ViewCount        int       `fake:"{number:0,100000}"`
	PostBody         string    `fake:"{sentence:30}"`
	OwnerUserID      int32     `fake:"{number:1,100}"`
	LastEditorUserID int32     `fake:"{number:1,100}"`
	LastEditDate     time.Time `fake:"{date}"`
	LastActivityDate time.Time `fake:"{date}"`
	PostTitle        string    `fake:"{sentence:5}"`
	Tags             string    `fake:"{word}"`
	AnswerCount      int       `fake:"{number:0,50}"`
	CommentCount     int       `fake:"{number:0,200}"`
	ContentLicense   string    `fake:"{randomstring:[CC BY-SA, MIT, Apache]}"`
}

func GeneratePosts() Post {
	var generated Post
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