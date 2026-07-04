func escapeSQL(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
func ChooseTable(idb []int, idc []int, idph []int, idp []int, idu []int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	whatTable := []string{"badges", "comments", "posthistory", "postlinks", "posts", "users", "votes"}
	randomIndex := r.Intn(len(whatTable))
	switch whatTable[randomIndex] {
	case "badges":
		data := fakedata.GenerateBadge()
		return fmt.Sprintf("INSERT INTO badges (user_id, badge_name, badge_date, class, tag_based) VALUES ( %d, '%s', '%s', %d, '%s');",
			idu[r.Intn(len(idu))], escapeSQL(data.BadgeName), data.BadgeDate.Format("2006-01-02 15:04:05"), data.Class, escapeSQL(data.TagBased))
	case "comments":
		data := fakedata.GenerateComments()
		return fmt.Sprintf("INSERT INTO comments (post_id, score, comment_text, creation_date, user_id, content_license) VALUES ( %d, %d, '%s', '%s', %d, '%s');",
			idp[r.Intn(len(idp))], data.Score, escapeSQL(data.CommentText), data.CreationDate.Format("2006-01-02 15:04:05"), idu[r.Intn(len(idu))], escapeSQL(data.ContentLicense))
...
	}	