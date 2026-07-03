package main

import (
	"fmt"
	"math/rand/v2"
)

func (q UserByIDQuery) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["users"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT *
        FROM users
        WHERE id = %d;
    `, id)
}

func (q PostByIDQuery) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["posts"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE id = %d;
    `, id)
}

func (q PostsByScoreEq) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE score = %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q PostByScoreHigh) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE score > %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q PostByScoreLow) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE score < %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q PostByViewCoEq) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE view_count = %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q PostByViewCoHigh) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE view_count > %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q PostByViewCoLow) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE view_count < %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q PostByOvnId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["users"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE owner_user_id = %d;
    `, id)
}

func (q PostByLastEditorId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["users"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM posts
        WHERE last_editor_user_id = %d;
    `, id)
}

func (q PostHiById) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["posts_history"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM post_history
        WHERE id = %d;
    `, id)
}

func (q PostHiByTypeId) Generate(r *rand.Rand, ids map[string][]int) string {

	id := r.IntN(30)

	return fmt.Sprintf(`
        SELECT *
        FROM post_history
        WHERE post_history_type_id = %d;
    `, id)
}

func (q PostHiByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["posts"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM post_history
        WHERE post_id = %d;
    `, id)
}

func (q PostHiByUsrId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["users"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT *
        FROM post_history
        WHERE user_id = %d;
    `, id)
}

func (q PostLiByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["posts"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM post_links
        WHERE post_id = %d;
    `, id)
}

func (q PostLiByRelPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["posts"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM post_links
        WHERE related_post_id = %d;
    `, id)
}

func (q PostLiByLiTyId) Generate(r *rand.Rand, ids map[string][]int) string {

	id := r.IntN(30)

	return fmt.Sprintf(`
        SELECT *
        FROM post_links
        WHERE link_type_id = %d;
    `, id)
}

func (q ComByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["posts"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT *
        FROM comments
        WHERE post_id = %d;
    `, id)
}

func (q ComByScorEq) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM comments
        WHERE score = %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q ComByScorHigh) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM comments
        WHERE score > %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q ComByScorLow) Generate(r *rand.Rand, ids map[string][]int) string {

	s := r.IntN(100)
	l1 := r.IntN(100)
	l2 := r.IntN(100)

	return fmt.Sprintf(`
        SELECT *
        FROM comments
        WHERE score < %d LIMIT %d, %d;
    `, s, l1, l2)
}

func (q ComByUsrID) Generate(r *rand.Rand, ids map[string][]int) string {
	idsu := ids["users"]
	if len(idsu) == 0 {
		return ""
	}
	id := idsu[r.IntN(len(idsu))]

	return fmt.Sprintf(`
        SELECT *
        FROM comments
        WHERE user_id = %d;
    `, id)
}

func (q VotCountByPosId) Generate(r *rand.Rand, ids map[string][]int) string {
	idsp := ids["posts"]
	if len(idsp) == 0 {
		return ""
	}
	id := idsp[r.IntN(len(idsp))]

	return fmt.Sprintf(`
        SELECT COUNT(*)
        FROM votes
        WHERE post_id = %d;
    `, id)
}
