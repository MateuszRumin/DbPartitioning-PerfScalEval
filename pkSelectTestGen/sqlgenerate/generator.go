package sqlgenerate

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func randomArea(r *rand.Rand) (time.Time, time.Time) {
	start := time.Date(2008, 07, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2014, 9, 14, 0, 0, 0, 0, time.UTC)

	days := int(end.Sub(start).Hours() / 24)
	randomDate := start.AddDate(0, 0, r.IntN(days+1))
	randomDate2 := randomDate.AddDate(0, 1+r.IntN(3), 0)

	return randomDate, randomDate2
}

func randomId(r *rand.Rand, values []int) (int, bool) {
	if len(values) == 0 {
		return 0, false
	}

	return values[r.IntN(len(values))], true
}

func (q UserByIDQuery) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM users 
	WHERE id = %d;`,
		id)
}

func (q PostByIDQuery) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE id = %d;`,
		id)
}

func (q PostByScoreHigh) Generate(r *rand.Rand, ids map[string][]int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
	WHERE score > %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC LIMIT %d, %d;`,
		r.IntN(100), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q PostByScoreLow) Generate(r *rand.Rand, ids map[string][]int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE score < %d AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC LIMIT %d, %d;`,
		r.IntN(100), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q PostByViewCoHigh) Generate(r *rand.Rand, ids map[string][]int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE view_count > %d AND creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC LIMIT %d, %d;`,
		r.IntN(100), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q PostByViewCoLow) Generate(r *rand.Rand, ids map[string][]int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE view_count < %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC LIMIT %d, %d;`,
		r.IntN(100), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q PostByOvnId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE owner_user_id = %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC;`,
		id, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q PostByLastEditorId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE last_editor_user_id = %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC;`,
		id, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q PostHiById) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts_history"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT *
    FROM post_history WHERE id = %d;`,
		id)
}

func (q PostHiByTypeId) Generate(r *rand.Rand, ids map[string][]int) string {

	id := r.IntN(30)

	return fmt.Sprintf(`SELECT * FROM post_history
    WHERE post_history_type_id = %d;`,
		id)
}

func (q PostHiByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM post_history
    WHERE post_id = %d;`,
		id)
}

func (q PostHiByUsrId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM post_history 
	WHERE user_id = %d;`,
		id)
}

func (q PostLiByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM post_links
    WHERE post_id = %d;`,
		id)
}

func (q PostLiByRelPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM post_links
    WHERE related_post_id = %d;`,
		id)
}

func (q PostLiByLiTyId) Generate(r *rand.Rand, ids map[string][]int) string {

	return fmt.Sprintf(`SELECT * FROM post_links
    WHERE link_type_id = %d;`,
		r.IntN(30))
}

func (q ComByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM comments
    WHERE post_id = %d;`,
		id)
}

func (q ComByScorHigh) Generate(r *rand.Rand, ids map[string][]int) string {

	return fmt.Sprintf(`SELECT * FROM comments WHERE score > %d LIMIT %d, %d;`,
		r.IntN(400), r.IntN(100), r.IntN(100))
}

func (q ComByScorLow) Generate(r *rand.Rand, ids map[string][]int) string {

	return fmt.Sprintf(`SELECT * FROM comments
    WHERE score < %d
	LIMIT %d, %d;`,
		r.IntN(100), r.IntN(100), r.IntN(100))
}

func (q ComByUsrID) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT * FROM comments 
	WHERE user_id = %d;`,
		id)
}

func (q VotCountByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	return fmt.Sprintf(`SELECT COUNT(*) FROM votes
    WHERE post_id = %d;`,
		id)
}

func (q PostsByScoreRange) Generate(r *rand.Rand, ids map[string][]int) string {
	a := r.IntN(1000)

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE score BETWEEN %d AND %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC	LIMIT %d, %d;`,
		a, a+r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q PostsByViewRange) Generate(r *rand.Rand, ids map[string][]int) string {
	a := r.IntN(1000)

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE view_count BETWEEN %d AND %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC LIMIT %d, %d;`,
		a, a+r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q PostTopViewed) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT id, score, view_count, creation_date FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY score DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 10+r.IntN(100))
}

func (q PostByAnswerCount) Generate(r *rand.Rand, ids map[string][]int) string {

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE comment_count > %d AND creation_date >= '%s' AND creation_date < '%s'
	LIMIT %d, %d;`,
		r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q PostByOwnerAndScore) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT * FROM posts
    WHERE owner_user_id = %d AND score > %d AND creation_date >= '%s' AND creation_date < '%s'
	ORDER BY score DESC LIMIT %d, %d;`,
		id, r.IntN(1000), randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), r.IntN(100), r.IntN(100))
}

func (q UserPosts) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT p.* FROM users u JOIN posts p
	ON u.id = p.owner_user_id
    WHERE u.id = %d AND p.creation_date >= '%s' AND p.creation_date < '%s'
	ORDER BY creation_date DESC;`,
		id, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q PostComments) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT c.* FROM posts p JOIN comments c
	ON p.id = c.post_id
    WHERE p.id = %d AND p.creation_date >= '%s'AND p.creation_date < '%s'
	ORDER BY p.creation_date DESC;`,
		id, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q PostVotes) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["posts"])
	if !ok {
		return ""
	}

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT v.* FROM posts p 
	JOIN votes v ON p.id = v.post_id
    WHERE p.id = %d AND p.creation_date >= '%s' AND p.creation_date < '%s'
	ORDER BY p.creation_date DESC;`,
		id, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q UserPostVotes) Generate(r *rand.Rand, ids map[string][]int) string {
	id, ok := randomId(r, ids["users"])
	if !ok {
		return ""
	}

	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT c.* FROM users u
	JOIN posts p ON u.id = p.owner_user_id
	JOIN comments c ON p.id = c.post_id
	WHERE u.id = %d AND p.creation_date >= '%s' AND p.creation_date < '%s'
	ORDER BY p.creation_date DESC;`,
		id, randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q AgregateVotes) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT post_id, COUNT(*) FROM votes
	WHERE creation_date >= '%s' AND creation_date < '%s' GROUP BY post_id;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q AgregatePostScore) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT MAX(score) FROM posts 
	WHERE creation_date >= '%s' AND creation_date < '%s';`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q AveragePostViews) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT AVG(view_count) FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s';`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q PostPerUser) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, randomDate2 := randomArea(r)

	return fmt.Sprintf(`SELECT owner_user_id, COUNT(*) FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s'
	GROUP BY owner_user_id;	`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"))
}

func (q SingleDayLookUp) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, _ := randomArea(r)
	rd2 := randomDate.AddDate(0, 0, 1)

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE creation_date >= '%s' AND creation_date < '%s'
	ORDER BY creation_date DESC;`,
		randomDate.Format("2006-01-02"), rd2.Format("2006-01-02"))
}

func (q SingleMonthLookUp) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, _ := randomArea(r)
	month1 := time.Date(randomDate.Year(), randomDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	month2 := month1.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC;`,
		month1.Format("2006-01-02"), month2.Format("2006-01-02"))
}

func (q RecentPosts) Generate(r *rand.Rand, ids map[string][]int) string {

	return fmt.Sprintf(`SELECT * FROM posts 
	ORDER BY creation_date DESC LIMIT %d,%d;`,
		r.IntN(100), r.IntN(100))
}

func (q DateRandomArea) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, _ := randomArea(r)

	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY creation_date DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 5+r.IntN(100))
}

func (q DateRandomAreaScore) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT * FROM posts 
	WHERE creation_date >= '%s' AND creation_date < '%s' AND score > %d 
	ORDER BY creation_date DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 500+r.IntN(500), 5+r.IntN(100))
}

func (q DateRandomAreaViewCount) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT * FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s' AND view_count > %d
	ORDER BY creation_date DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 10+r.IntN(10000), 5+r.IntN(100))
}

func (q DateRandomAreaOrderScore) Generate(r *rand.Rand, ids map[string][]int) string {
	randomDate, _ := randomArea(r)
	randomDate2 := randomDate.AddDate(0, 1, 0)

	return fmt.Sprintf(`SELECT id, score, creation_date FROM posts
	WHERE creation_date >= '%s' AND creation_date < '%s' 
	ORDER BY score DESC LIMIT %d;`,
		randomDate.Format("2006-01-02"), randomDate2.Format("2006-01-02"), 5+r.IntN(100))
}
