package sqlgenerate

import (
	"fmt"
	"math/rand/v2"
)

type QuestionPage struct{}
type QuestionAnswers struct{}
type AcceptedAnswer struct{}
type CommentsForPost struct{}
type NewestQuestions struct{}
type ActiveQuestions struct{}
type HotQuestions struct{}
type UnansweredQuestions struct{}
type TagNewestQuestions struct{}
type UserProfile struct{}
type UserRecentPosts struct{}
type UserRecentComments struct{}
type UserBadges struct{}
type TextSearchPosts struct{}
type RelatedPosts struct{}
type PostTimeline struct{}
type VoteBreakdown struct{}
type PopularTags struct{}

func (QuestionPage) Generate(r *rand.Rand, idp []int, _ []int) string {
	postID, ok := randomID(r, idp)
	if !ok {
		return noOpQuery
	}

	// idp może zawierać zarówno pytania, jak i odpowiedzi. Najpierw wyznaczany
	// jest korzeń wątku, dzięki czemu generator nie wymaga osobnej puli questionID.
	return fmt.Sprintf(`SELECT
		q.id,
		q.post_title,
		q.post_body,
		q.tags,
		q.score,
		q.view_count,
		q.answer_count,
		q.comment_count,
		q.accepted_answer_id,
		q.creation_date,
		q.last_activity_date,
		u.id AS owner_id,
		u.display_name,
		u.reputation
	FROM posts seed
	JOIN posts q
	  ON q.id = CASE
		WHEN seed.post_type_id = 1 THEN seed.id
		ELSE seed.parent_id
	  END
	LEFT JOIN users u ON u.id = q.owner_user_id
	WHERE seed.id = %d
	  AND q.post_type_id = 1
	LIMIT 1;`, postID)
}

func (QuestionAnswers) Generate(r *rand.Rand, idp []int, _ []int) string {
	postID, ok := randomID(r, idp)
	if !ok {
		return noOpQuery
	}

	return fmt.Sprintf(`SELECT
		a.id,
		a.post_body,
		a.score,
		a.comment_count,
		a.creation_date,
		a.last_edit_date,
		u.id AS owner_id,
		u.display_name,
		u.reputation,
		(a.id = q.accepted_answer_id) AS is_accepted
	FROM posts seed
	JOIN posts q
	  ON q.id = CASE
		WHEN seed.post_type_id = 1 THEN seed.id
		ELSE seed.parent_id
	  END
	JOIN posts a ON a.parent_id = q.id AND a.post_type_id = 2
	LEFT JOIN users u ON u.id = a.owner_user_id
	WHERE seed.id = %d
	ORDER BY is_accepted DESC, a.score DESC, a.creation_date ASC;`, postID)
}

func (AcceptedAnswer) Generate(r *rand.Rand, idp []int, _ []int) string {
	postID, ok := randomID(r, idp)
	if !ok {
		return noOpQuery
	}

	return fmt.Sprintf(`SELECT
		a.id,
		a.post_body,
		a.score,
		a.creation_date,
		u.display_name,
		u.reputation
	FROM posts seed
	JOIN posts q
	  ON q.id = CASE
		WHEN seed.post_type_id = 1 THEN seed.id
		ELSE seed.parent_id
	  END
	JOIN posts a ON a.id = q.accepted_answer_id
	LEFT JOIN users u ON u.id = a.owner_user_id
	WHERE seed.id = %d
	LIMIT 1;`, postID)
}

func (CommentsForPost) Generate(r *rand.Rand, idp []int, _ []int) string {
	postID, ok := randomID(r, idp)
	if !ok {
		return noOpQuery
	}

	return fmt.Sprintf(`SELECT
		c.id,
		c.comment_text,
		c.score,
		c.creation_date,
		u.id AS user_id,
		u.display_name,
		u.reputation
	FROM comments c
	LEFT JOIN users u ON u.id = c.user_id
	WHERE c.post_id = %d
	ORDER BY c.creation_date ASC;`, postID)
}

func (NewestQuestions) Generate(r *rand.Rand, _ []int, _ []int) string {
	start, end := randomWindow(r)
	limit := randomPageSize(r)
	offset := randomOffset(r, limit)

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_title,
		p.tags,
		p.score,
		p.view_count,
		p.answer_count,
		p.creation_date,
		u.display_name
	FROM posts p
	LEFT JOIN users u ON u.id = p.owner_user_id
	WHERE p.post_type_id = 1
	  AND p.creation_date >= '%s'
	  AND p.creation_date < '%s'
	ORDER BY p.creation_date DESC
	LIMIT %d OFFSET %d;`,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
		limit,
		offset,
	)
}

func (ActiveQuestions) Generate(r *rand.Rand, _ []int, _ []int) string {
	start, end := randomWindow(r)
	limit := randomPageSize(r)

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_title,
		p.tags,
		p.score,
		p.answer_count,
		p.last_activity_date
	FROM posts p
	WHERE p.post_type_id = 1
	  AND p.creation_date < '%s'
	  AND p.last_activity_date >= '%s'
	  AND p.last_activity_date < '%s'
	ORDER BY p.last_activity_date DESC
	LIMIT %d;`,
		end.Format("2006-01-02"),
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
		limit,
	)
}

func (HotQuestions) Generate(r *rand.Rand, _ []int, _ []int) string {
	start, end := randomWindow(r)
	limit := randomFrom(r, []int{10, 15, 20, 30})

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_title,
		p.tags,
		p.score,
		p.view_count,
		p.answer_count,
		p.comment_count,
		p.creation_date
	FROM posts p
	WHERE p.post_type_id = 1
	  AND p.creation_date >= '%s'
	  AND p.creation_date < '%s'
	ORDER BY (
		LOG10(GREATEST(p.view_count, 1)) * 2.0 +
		p.answer_count * 1.5 +
		p.comment_count * 0.5 +
		p.score * 0.2
	) DESC
	LIMIT %d;`,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
		limit,
	)
}

func (UnansweredQuestions) Generate(r *rand.Rand, _ []int, _ []int) string {
	start, end := randomWindow(r)
	limit := randomPageSize(r)

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_title,
		p.tags,
		p.score,
		p.view_count,
		p.creation_date
	FROM posts p
	WHERE p.post_type_id = 1
	  AND p.answer_count = 0
	  AND p.creation_date >= '%s'
	  AND p.creation_date < '%s'
	ORDER BY p.creation_date DESC
	LIMIT %d;`,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
		limit,
	)
}

func (TagNewestQuestions) Generate(r *rand.Rand, _ []int, _ []int) string {
	start, end := randomWindow(r)
	limit := randomPageSize(r)
	offset := randomOffset(r, limit)
	tag := randomTag(r)

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_title,
		p.tags,
		p.score,
		p.view_count,
		p.answer_count,
		p.creation_date
	FROM posts p
	WHERE p.post_type_id = 1
	  AND p.creation_date >= '%s'
	  AND p.creation_date < '%s'
	  AND p.tags LIKE %s
	ORDER BY p.creation_date DESC
	LIMIT %d OFFSET %d;`,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
		sqlString("%<"+tag+">%"),
		limit,
		offset,
	)
}

func (UserProfile) Generate(r *rand.Rand, _ []int, idu []int) string {
	userID, ok := randomID(r, idu)
	if !ok {
		return noOpQuery
	}

	return fmt.Sprintf(`SELECT
		u.id,
		u.display_name,
		u.reputation,
		u.creation_date,
		u.last_access_date,
		u.location,
		u.website_url,
		u.about_me,
		u.views,
		u.upvotes,
		u.downvotes
	FROM users u
	WHERE u.id = %d
	LIMIT 1;`, userID)
}

func (UserRecentPosts) Generate(r *rand.Rand, _ []int, idu []int) string {
	userID, ok := randomID(r, idu)
	if !ok {
		return noOpQuery
	}
	limit := randomFrom(r, []int{10, 15, 20, 30, 50})

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_type_id,
		p.parent_id,
		p.post_title,
		p.score,
		p.view_count,
		p.creation_date
	FROM posts p
	WHERE p.owner_user_id = %d
	ORDER BY p.creation_date DESC
	LIMIT %d;`, userID, limit)
}

func (UserRecentComments) Generate(r *rand.Rand, _ []int, idu []int) string {
	userID, ok := randomID(r, idu)
	if !ok {
		return noOpQuery
	}
	limit := randomFrom(r, []int{10, 15, 20, 30})

	return fmt.Sprintf(`SELECT
		c.id,
		c.post_id,
		c.comment_text,
		c.score,
		c.creation_date,
		p.post_title
	FROM comments c
	JOIN posts p ON p.id = c.post_id
	WHERE c.user_id = %d
	ORDER BY c.creation_date DESC
	LIMIT %d;`, userID, limit)
}

func (UserBadges) Generate(r *rand.Rand, _ []int, idu []int) string {
	userID, ok := randomID(r, idu)
	if !ok {
		return noOpQuery
	}
	limit := randomFrom(r, []int{20, 50, 100})

	return fmt.Sprintf(`SELECT
		b.id,
		b.badge_name,
		b.badge_date,
		b.class,
		b.tag_based
	FROM badges b
	WHERE b.user_id = %d
	ORDER BY b.badge_date DESC
	LIMIT %d;`, userID, limit)
}

func (TextSearchPosts) Generate(r *rand.Rand, _ []int, _ []int) string {
	term := randomSearchTerm(r)
	start, end := randomWindow(r)
	limit := randomFrom(r, []int{10, 15, 20, 30})

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_title,
		p.tags,
		p.score,
		p.view_count,
		p.creation_date
	FROM posts p
	WHERE p.post_type_id = 1
	  AND p.creation_date >= '%s'
	  AND p.creation_date < '%s'
	  AND (
		p.post_title LIKE %s
		OR p.post_body LIKE %s
	  )
	ORDER BY p.score DESC, p.creation_date DESC
	LIMIT %d;`,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
		sqlString("%"+term+"%"),
		sqlString("%"+term+"%"),
		limit,
	)
}

func (RelatedPosts) Generate(r *rand.Rand, idp []int, _ []int) string {
	postID, ok := randomID(r, idp)
	if !ok {
		return noOpQuery
	}
	limit := randomFrom(r, []int{5, 10, 15, 20})

	return fmt.Sprintf(`SELECT
		p.id,
		p.post_title,
		p.tags,
		p.score,
		pl.link_type_id
	FROM post_links pl
	JOIN posts p ON p.id = pl.related_post_id
	WHERE pl.post_id = %d
	ORDER BY p.score DESC
	LIMIT %d;`, postID, limit)
}

func (PostTimeline) Generate(r *rand.Rand, idp []int, _ []int) string {
	postID, ok := randomID(r, idp)
	if !ok {
		return noOpQuery
	}
	limit := randomFrom(r, []int{20, 50, 100})

	return fmt.Sprintf(`SELECT
		ph.id,
		ph.post_history_type_id,
		ph.revision_guid,
		ph.creation_date,
		ph.user_id,
		u.display_name,
		ph.post_text,
		ph.content_license
	FROM post_history ph
	LEFT JOIN users u ON u.id = ph.user_id
	WHERE ph.post_id = %d
	ORDER BY ph.creation_date DESC
	LIMIT %d;`, postID, limit)
}

func (VoteBreakdown) Generate(r *rand.Rand, idp []int, _ []int) string {
	postID, ok := randomID(r, idp)
	if !ok {
		return noOpQuery
	}

	return fmt.Sprintf(`SELECT
		v.vote_type_id,
		COUNT(*) AS vote_count
	FROM votes v
	WHERE v.post_id = %d
	GROUP BY v.vote_type_id
	ORDER BY v.vote_type_id;`, postID)
}

func (PopularTags) Generate(r *rand.Rand, _ []int, _ []int) string {
	limit := randomFrom(r, []int{20, 30, 50, 100})

	return fmt.Sprintf(`SELECT
		t.id,
		t.tag_name,
		t.tag_count,
		t.except_post_id,
		t.wiki_post_id
	FROM tags t
	ORDER BY t.tag_count DESC
	LIMIT %d;`, limit)
}
