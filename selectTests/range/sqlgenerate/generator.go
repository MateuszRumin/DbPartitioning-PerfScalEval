package sqlgenerate

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Faked struct {
	PostTitle string `fake:"{sentence:3}"`
	Tags      string `fake:"{word}"`
}

func randomArea(r *rand.Rand) (time.Time, time.Time) {
	start := time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))
	randomDate2 := randomDate.AddDate(0, 1+r.IntN(3), 0)

	return randomDate, randomDate2
}

func (q NewQuestions) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT p.id, p.post_title, p.score, p.view_count, p.answer_count, p.creation_date, u.display_name FROM posts p
	LEFT JOIN users u ON u.id = p.owner_user_id
	WHERE p.post_type_id = 1
	ORDER BY p.creation_date DESC LIMIT %d`,
		r.IntN(100))
}

func (q UserActivity) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT id, post_title, score, creation_date, post_type_id FROM posts
	WHERE owner_user_id = %d AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC
	LIMIT %d;
	`, r.IntN(idu), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(50))
}

func (q TopControversialPosts) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT id, post_title, score FROM posts 
	WHERE post_type_id = 1 AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY ABS(score) ASC, view_count DESC LIMIT %d;
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(100))
}

func (q TopCommentedPosts) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT id, post_title, comment_count, score FROM posts
	WHERE post_type_id = 1 AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY comment_count DESC LIMIT %d;
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(100))
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

func (q UserRankingByPostPopularity30d) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 0, 30)

	return fmt.Sprintf(`
	SELECT u.id, u.display_name, COUNT(p.id) AS post_count, AVG(p.score) AS avg_score FROM users u
	JOIN posts p ON p.owner_user_id = u.id
	WHERE p.creation_date >= '%s' AND p.creation_date < '%s'
	GROUP BY u.id, u.display_name ORDER BY post_count DESC, avg_score DESC LIMIT %d;
	;
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(100))
}

func (q TextSearchPosts) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	var fake Faked
	var condition string

	err := gofakeit.Struct(&fake)
	if err != nil {
		condition = "zabieg"
	} else {
		condition = fake.PostTitle
	}

	return fmt.Sprintf(`SELECT id, post_title, score, creation_date FROM posts
	WHERE post_type_id = 1 AND creation_date >= '%s' AND creation_date < '%s'  AND (post_title LIKE CONCAT('%%', '%s', '%%') OR post_body LIKE CONCAT('%%', '%s', '%%'))
	ORDER BY score DESC LIMIT %d;
	
	`, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), condition, condition, 1+r.IntN(50))
}

func (q TagSearchPosts) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	var fake Faked
	var condition string

	err := gofakeit.Struct(&fake)
	if err != nil {
		condition = "mysql"
	} else {
		condition = fake.Tags
	}

	return fmt.Sprintf(`SELECT id, post_title, tags, score FROM posts
	WHERE tags LIKE CONCAT('%%<', '%s', '>%%') AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY score DESC LIMIT %d;
	`, condition, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(50))
}

func (q QuestionsForTag) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	var fake Faked
	var condition string

	err := gofakeit.Struct(&fake)
	if err != nil {
		condition = "mysql"
	} else {
		condition = fake.Tags
	}

	return fmt.Sprintf(`SELECT id, post_title, score, creation_date FROM posts
	WHERE post_type_id = 1 AND tags LIKE CONCAT('%%<', '%s', '>%%') AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC
	LIMIT %d;
	`, condition, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(50))
}

func (q PostByScoreHigh) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
	WHERE score > %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC LIMIT %d;`,
		r.IntN(100), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100))
}

func (q PostByScoreLow) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE score < %d AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC LIMIT %d;`,
		r.IntN(100), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 1+r.IntN(100))
}

func (q PostByViewCoHigh) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE view_count > %d AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC LIMIT %d,;`,
		r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100))
}

func (q PostByViewCoLow) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE view_count < %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC LIMIT %d;`,
		r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100))
}

func (q PostsByScoreRange) Generate(r *rand.Rand, idp int, idu int) string {
	a := r.IntN(1000)

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE score BETWEEN %d AND %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC	LIMIT %d;`,
		a, a+r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100))
}

func (q PostsByViewRange) Generate(r *rand.Rand, idp int, idu int) string {
	a := r.IntN(1000)

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE view_count BETWEEN %d AND %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC LIMIT %d;`,
		a, a+r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100))
}

func (q PostByAnswerCount) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE comment_count > %d AND creation_date >= '%s' AND creation_date < '%s'
	LIMIT %d;`,
		r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100))
}

func (q PostByOwnerAndScore) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE owner_user_id = %d AND score > %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY score DESC LIMIT %d;`,
		r.IntN(idu), r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100))
}

func (q UserPosts) Generate(r *rand.Rand, idp int, idu int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT p.* FROM users u JOIN posts p
	ON u.id = p.owner_user_id
    WHERE u.id = %d AND p.creation_date >= '%s' AND p.creation_date < '%s'
	ORDER BY creation_date DESC;`,
		r.IntN(idu), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q AveragePostViews) Generate(r *rand.Rand, idp int, idu int) string {
	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT AVG(view_count) FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s';`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q PostPerUser) Generate(r *rand.Rand, idp int, idu int) string {
	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT owner_user_id, COUNT(*) FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s'
	GROUP BY owner_user_id limit %d;	`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100)+1)
}

func (q SingleDayLookUp) Generate(r *rand.Rand, idp int, idu int) string {
	randomDate, _ := randomArea(r)
	rd2 := randomDate.AddDate(0, 0, 1)

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC limit %d;`,
		randomDate.Format("2006-01-02"), rd2.Format("2006-01-02"), r.IntN(100)+1)
}

func (q DateRandomAreaScore) Generate(r *rand.Rand, idp int, idu int) string {
	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE creation_date >= '%s' AND creation_date < '%s' AND score > %d 
	ORDER BY creation_date DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 500+r.IntN(500), 5+r.IntN(100))
}

func (q DateRandomAreaViewCount) Generate(r *rand.Rand, idp int, idu int) string {
	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT * FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s' AND view_count > %d
	ORDER BY creation_date DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 10+r.IntN(10000), 5+r.IntN(100))
}

func (q DateRandomAreaOrderScore) Generate(r *rand.Rand, idp int, idu int) string {
	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT id, score, creation_date FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY score DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 5+r.IntN(100))
}
