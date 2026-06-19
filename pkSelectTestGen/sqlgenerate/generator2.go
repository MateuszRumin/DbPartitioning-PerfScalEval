package sqlgenerate

import (
	"fmt"
	"math/rand/v2"
)

type PostsByScoreRange struct{}

func (q PostsByScoreRange) Generate(r *rand.Rand, ids map[string][]int) string {
	a := r.IntN(1000)
	b := a + r.IntN(1000)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE score BETWEEN %d AND %d LIMIT %d, %d
		
		
		;
    `, a, b, l1, l2)
}

type PostsByViewRange struct{}

func (q PostsByViewRange) Generate(r *rand.Rand, ids map[string][]int) string {
	a := r.IntN(1000)
	b := a + r.IntN(1000)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE view_count BETWEEN %d AND %d LIMIT %d, %d
		
		
		;
    `, a, b, l1, l2)
}

type PostTopViewed struct{}

func (q PostTopViewed) Generate(r *rand.Rand, ids map[string][]int) string {

	l1 := r.IntN(10000)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE score LIMIT %d, %d
		
		
		;
    `, l1, l2)
}

type PostByAnswerCount struct{}

func (q PostByAnswerCount) Generate(r *rand.Rand, ids map[string][]int) string {
	c := r.IntN(1000)
	l1 := r.IntN(10000)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE comment_count > %d LIMIT %d, %d
		
		
		;
    `, c, l1, l2)
}

type PostByOwnerAndScore struct{}

func (q PostByOwnerAndScore) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["users"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]
	c := r.IntN(1000)
	l1 := r.IntN(10000)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE user_id = %d AND score > %d LIMIT %d, %d
		
		
		;
    `, id, c, l1, l2)
}

type UserPosts struct{}

func (q UserPosts) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["users"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT p.*
        FROM users u JOIN posts p
			ON u.id = p.owner_user_id
        WHERE u.id = %d 
		
		
		;
    `, id)
}

type PostComments struct{}

func (q PostComments) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["posts"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT c.*
        FROM posts p JOIN comments c
			ON p.id = c.post_id
        WHERE p.id = %d 
		
		
		;
    `, id)
}

type PostVotes struct{}

func (q PostVotes) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["posts"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT v.*
        FROM posts p JOIN votes v
			ON p.id = v.post_id
        WHERE p.id = %d 
		
		
		;
    `, id)
}

type UserPostVotes struct{}

func (q UserPostVotes) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["users"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT c.*
        FROM users u
		JOIN posts p
			ON u.id = p.owner_user_id
		JOIN comments c
			ON p.id = c.post_id
		WHERE u.id = %d
		
		
		;

    `, id)
}

type AgregateVotes struct{}

func (q AgregateVotes) Generate(r *rand.Rand, ids map[string][]int) string {
	return `
		SELECT post_id, COUNT(*)
		FROM votes
		GROUP BY post_id
		
		
		;
	`
}

type AgregatePostScore struct{}

func (q AgregatePostScore) Generate(r *rand.Rand, ids map[string][]int) string {
	return `
		SELECT MAX(score)
		FROM posts
		
		
		;
	`
}

type AveragePostViews struct{}

func (q AveragePostViews) Generate(r *rand.Rand, ids map[string][]int) string {
	return `
		SELECT AVG(view_count)
		FROM posts
		
		
		;
	`
}

type PostPerUser struct{}

func (q PostPerUser) Generate(r *rand.Rand, ids map[string][]int) string {
	return `
		SELECT owner_user_id, COUNT(*)
		FROM posts
		GROUP BY owner_user_id
		
		
		;
		
	`
}
