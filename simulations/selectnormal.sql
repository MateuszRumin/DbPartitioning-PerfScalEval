-- QuestionPage
SELECT q.id, q.post_title, q.post_body, q.tags, q.score, q.view_count, q.answer_count, q.comment_count, q.accepted_answer_id, q.creation_date, q.last_activity_date, u.id AS owner_id, u.display_name, u.reputation FROM posts seed JOIN posts q ON q.id = CASE WHEN seed.post_type_id = 1 THEN seed.id ELSE seed.parent_id END LEFT JOIN users u ON u.id = q.owner_user_id WHERE seed.id = ? AND q.post_type_id = 1 LIMIT 1;

-- QuestionAnswers
SELECT a.id, a.post_body, a.score, a.comment_count, a.creation_date, a.last_edit_date, u.id AS owner_id, u.display_name, u.reputation, (a.id = q.accepted_answer_id) AS is_accepted FROM posts seed JOIN posts q ON q.id = CASE WHEN seed.post_type_id = 1 THEN seed.id ELSE seed.parent_id END JOIN posts a ON a.parent_id = q.id AND a.post_type_id = 2 LEFT JOIN users u ON u.id = a.owner_user_id WHERE seed.id = ? ORDER BY is_accepted DESC, a.score DESC, a.creation_date ASC;

-- AcceptedAnswer
SELECT a.id, a.post_body, a.score, a.creation_date, u.display_name, u.reputation FROM posts seed JOIN posts q ON q.id = CASE WHEN seed.post_type_id = 1 THEN seed.id ELSE seed.parent_id END JOIN posts a ON a.id = q.accepted_answer_id LEFT JOIN users u ON u.id = a.owner_user_id WHERE seed.id = ? LIMIT 1;

-- CommentsForPost
SELECT c.id, c.comment_text, c.score, c.creation_date, u.id AS user_id, u.display_name, u.reputation FROM comments c LEFT JOIN users u ON u.id = c.user_id WHERE c.post_id = ? ORDER BY c.creation_date ASC;

-- NewestQuestions
SELECT p.id, p.post_title, p.tags, p.score, p.view_count, p.answer_count, p.creation_date, u.display_name FROM posts p LEFT JOIN users u ON u.id = p.owner_user_id WHERE p.post_type_id = 1 AND p.creation_date >= ? AND p.creation_date < ? ORDER BY p.creation_date DESC LIMIT ? OFFSET ?;

-- ActiveQuestions
SELECT p.id, p.post_title, p.tags, p.score, p.answer_count, p.last_activity_date FROM posts p WHERE p.post_type_id = 1 AND p.creation_date < ? AND p.last_activity_date >= ? AND p.last_activity_date < ? ORDER BY p.last_activity_date DESC LIMIT ?;

-- HotQuestions
SELECT p.id, p.post_title, p.tags, p.score, p.view_count, p.answer_count, p.comment_count, p.creation_date FROM posts p WHERE p.post_type_id = 1 AND p.creation_date >= ? AND p.creation_date < ? ORDER BY (LOG10(GREATEST(p.view_count, 1)) * 2.0 + p.answer_count * 1.5 + p.comment_count * 0.5 + p.score * 0.2) DESC LIMIT ?;

-- UnansweredQuestions
SELECT p.id, p.post_title, p.tags, p.score, p.view_count, p.creation_date FROM posts p WHERE p.post_type_id = 1 AND p.answer_count = 0 AND p.creation_date >= ? AND p.creation_date < ? ORDER BY p.creation_date DESC LIMIT ?;

-- TagNewestQuestions
SELECT p.id, p.post_title, p.tags, p.score, p.view_count, p.answer_count, p.creation_date FROM posts p WHERE p.post_type_id = 1 AND p.creation_date >= ? AND p.creation_date < ? AND p.tags LIKE CONCAT('%<', ?, '>%') ORDER BY p.creation_date DESC LIMIT ? OFFSET ?;

-- UserProfile
SELECT u.id, u.display_name, u.reputation, u.creation_date, u.last_access_date, u.location, u.website_url, u.about_me, u.views, u.upvotes, u.downvotes FROM users u WHERE u.id = ? LIMIT 1;

-- UserRecentPosts
SELECT p.id, p.post_type_id, p.parent_id, p.post_title, p.score, p.view_count, p.creation_date FROM posts p WHERE p.owner_user_id = ? ORDER BY p.creation_date DESC LIMIT ?;

-- UserRecentComments
SELECT c.id, c.post_id, c.comment_text, c.score, c.creation_date, p.post_title FROM comments c JOIN posts p ON p.id = c.post_id WHERE c.user_id = ? ORDER BY c.creation_date DESC LIMIT ?;

-- UserBadges
SELECT b.id, b.badge_name, b.badge_date, b.class, b.tag_based FROM badges b WHERE b.user_id = ? ORDER BY b.badge_date DESC LIMIT ?;

-- TextSearchPosts
SELECT p.id, p.post_title, p.tags, p.score, p.view_count, p.creation_date FROM posts p WHERE p.post_type_id = 1 AND p.creation_date >= ? AND p.creation_date < ? AND (p.post_title LIKE CONCAT('%', ?, '%') OR p.post_body LIKE CONCAT('%', ?, '%')) ORDER BY p.score DESC, p.creation_date DESC LIMIT ?;

-- RelatedPosts
SELECT p.id, p.post_title, p.tags, p.score, pl.link_type_id FROM post_links pl JOIN posts p ON p.id = pl.related_post_id WHERE pl.post_id = ? ORDER BY p.score DESC LIMIT ?;

-- PostTimeline
SELECT ph.id, ph.post_history_type_id, ph.revision_guid, ph.creation_date, ph.user_id, u.display_name, ph.post_text, ph.content_license FROM post_history ph LEFT JOIN users u ON u.id = ph.user_id WHERE ph.post_id = ? ORDER BY ph.creation_date DESC LIMIT ?;

-- VoteBreakdown
SELECT v.vote_type_id, COUNT(*) AS vote_count FROM votes v WHERE v.post_id = ? GROUP BY v.vote_type_id ORDER BY v.vote_type_id;

-- PopularTags
SELECT t.id, t.tag_name, t.tag_count, t.except_post_id, t.wiki_post_id FROM tags t ORDER BY t.tag_count DESC LIMIT ?;