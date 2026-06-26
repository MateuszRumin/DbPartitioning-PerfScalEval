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

func (q PostByIDQuery) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE id = %d;`,
		r.IntN(idp))
}

func (q NewQuestions) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT p.id, p.post_title, p.score, p.view_count, p.answer_count, p.creation_date, u.display_name FROM posts p
	LEFT JOIN users u ON u.id = p.owner_user_id
	WHERE p.post_type_id = 1
	ORDER BY p.creation_date DESC LIMIT %d`,
		r.IntN(100))
}

func (q QuestionDetailAndAuthor) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT p.*, u.display_name, u.reputation, u.location FROM posts p
	LEFT JOIN users u ON u.id = p.owner_user_id
	WHERE p.id = %d;`,
		r.IntN(idp))
}

func (q QuestionAnswers) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT p.id, p.post_body, p.score, p.creation_date, u.display_name FROM posts p
	LEFT JOIN users u ON u.id = p.owner_user_id
	WHERE p.parent_id = %d
	ORDER BY p.score DESC, p.creation_date ASC;`, r.IntN(idp))
}

func (q QuestionBestAnswers) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE parent_id = %d
	ORDER BY score DESC
	LIMIT 1;
	`, r.IntN(idp))
}

func (q CommentsForPost) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT c.id, c.comment_text, c.score, c.creation_date, u.display_name FROM comments c
	LEFT JOIN users u ON u.id = c.user_id
	WHERE c.post_id = %d
	ORDER BY c.creation_date ASC;
	`, r.IntN(idp))
}

func (q UserProfile) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT id, display_name, reputation, creation_date, location, website_url, views FROM users
	WHERE id = %d;`, r.IntN(idu))
}

func (q UserActivity) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT id, post_title, score, creation_date, post_type_id FROM posts
	WHERE owner_user_id = %d
	ORDER BY creation_date DESC
	LIMIT 50;
	`, r.IntN(idu))
}

func (q PostHistory) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT ph.id, ph.post_text, ph.creation_date, ph.user_id FROM post_history ph
	WHERE ph.post_id = %d
	ORDER BY ph.creation_date DESC;
	`, r.IntN(idp))
}

func (q ConnectedPosts) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT p2.id, p2.post_title, pl.link_type_id FROM post_links pl
	JOIN posts p2 ON p2.id = pl.related_post_id
	WHERE pl.post_id = %d;
	`, r.IntN(idp))
}

func (q ReputationTrend) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT u.id, u.display_name, u.reputation, u.upvotes, u.downvotes FROM users u
	ORDER BY u.reputation DESC LIMIT %d;
	`, 1+r.IntN(100))
}

func (q TopControversialPosts) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT id, post_title, score FROM posts WHERE post_type_id = 1
	ORDER BY ABS(score) ASC, view_count DESC LIMIT %d;
	`, 1+r.IntN(100))
}

func (q TopCommentedPosts) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT id, post_title, comment_count, score FROM posts
	WHERE post_type_id = 1
	ORDER BY comment_count DESC LIMIT %d;
	`, 1+r.IntN(100))
}

func (q LastGlobalActivity) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT id, post_type_id, creation_date, owner_user_id FROM posts
	ORDER BY creation_date DESC LIMIT %d;
	`, r.IntN(100))
}

func (q ConnectedPostsMultiple) Generate(r *rand.Rand, idp int, idu int) string {

	return fmt.Sprintf(`SELECT p.id, p.post_title, p.answer_count, p.comment_count FROM posts p
	WHERE p.id = %d;
	`, r.IntN(idp))
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
	var fake Faked
	var condition string

	err := gofakeit.Struct(&fake)
	if err != nil {
		condition = "zabieg"
	} else {
		condition = fake.PostTitle
	}

	return fmt.Sprintf(`SELECT id, post_title, score, creation_date FROM posts
	WHERE post_type_id = 1 AND (post_title LIKE CONCAT('%%', '%s', '%%') OR post_body LIKE CONCAT('%%', '%s', '%%'))
	ORDER BY score DESC LIMIT %d;
	
	`, condition, condition, 1+r.IntN(50))
}

func (q TagSearchPosts) Generate(r *rand.Rand, idp int, idu int) string {
	var fake Faked
	var condition string

	err := gofakeit.Struct(&fake)
	if err != nil {
		condition = "mysql"
	} else {
		condition = fake.Tags
	}

	return fmt.Sprintf(`SELECT id, post_title, tags, score FROM posts
	WHERE tags LIKE CONCAT('%%<', '%s', '>%%')
	ORDER BY score DESC LIMIT %d;
	`, condition, 1+r.IntN(50))
}

func (q QuestionsForTag) Generate(r *rand.Rand, idp int, idu int) string {
	var fake Faked
	var condition string

	err := gofakeit.Struct(&fake)
	if err != nil {
		condition = "mysql"
	} else {
		condition = fake.Tags
	}

	return fmt.Sprintf(`SELECT id, post_title, score, creation_date FROM posts
	WHERE post_type_id = 1 AND tags LIKE CONCAT('%%<', '%s', '>%%')
	ORDER BY creation_date DESC
	LIMIT %d;
	`, condition, 1+r.IntN(50))
}
