
type QueryGenerator interface {
	Generate(r *rand.Rand, idp int, idu int) string
}

func randomArea(r *rand.Rand) (time.Time, time.Time) {
	start := time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))
	randomDate2 := randomDate.AddDate(0, 1+r.IntN(3), 0)

	return randomDate, randomDate2
}

func (q PopularPosts7d) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 0, 1)

	return fmt.Sprintf(`SELECT p.id, p.post_title, p.score, p.view_count, p.answer_count FROM posts p
	WHERE p.post_type_id = 1 AND p.creation_date >= '%s' AND p.creation_date < '%s'
	ORDER BY (p.score * 2 + p.view_count * 0.01 + p.answer_count * 5) DESC
	LIMIT %d;
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(100))
}