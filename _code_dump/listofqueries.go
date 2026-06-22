package none

// ===== PK LOOKUPS =====

{80, PostByIDQuery{}},
`SELECT * FROM posts WHERE id = %d;`,
{40, UserByIDQuery{}},
`SELECT * FROM users WHERE id = %d;`
{25, PostHiById{}},
`SELECT * FROM post_history WHERE id = %d;`

// ===== USER-CENTRIC =====

{60, PostByOvnId{}},
`SELECT * FROM posts WHERE owner_user_id = %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC;`
{20, PostByLastEditorId{}},
`SELECT * FROM posts WHERE last_editor_user_id = %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC;`
{40, UserPosts{}},
`SELECT p.* FROM users u JOIN posts p ON u.id = p.owner_user_id WHERE u.id = %d AND p.creation_date >= '%s' AND p.creation_date < '%s' ORDER BY creation_date DESC;`
{25, ComByUsrID{}},
`SELECT * FROM comments WHERE user_id = %d;`

// ===== POST DETAILS =====

{50, ComByPosId{}},
{40, PostComments{}},
`SELECT c.* FROM posts p JOIN comments c ON p.id = c.post_id WHERE p.id = %d AND p.creation_date >= '%s'AND p.creation_date < '%s' ORDER BY p.creation_date DESC;`
{35, PostVotes{}},
`SELECT v.* FROM posts p JOIN votes v ON p.id = v.post_id WHERE p.id = %d AND p.creation_date >= '%s' AND p.creation_date < '%s' ORDER BY p.creation_date DESC;`
{25, VotCountByPosId{}},
`SELECT COUNT(*) FROM votes WHERE post_id = %d;`

// ===== PARTITION PRUNING TESTS =====

{80, SingleDayLookUp{}},
`SELECT * FROM posts WHERE creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC;`
{70, SingleMonthLookUp{}},
`SELECT * FROM posts WHERE creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC;`
{50, DateRandomArea{}},
`SELECT * FROM posts WHERE creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC LIMIT %d;`

// ===== PARTITION + SECONDARY FILTER =====

{45, DateRandomAreaScore{}},
`SELECT * FROM posts WHERE creation_date >= '%s' AND creation_date < '%s' AND score > %d ORDER BY creation_date DESC LIMIT %d;`
{45, DateRandomAreaViewCount{}},
`SELECT * FROM posts WHERE creation_date >= '%s' AND creation_date < '%s' AND view_count > %d ORDER BY creation_date DESC LIMIT %d;`
{30, DateRandomAreaOrderScore{}},
`SELECT id, score, creation_date FROM posts WHERE creation_date >= '%s' AND creation_date < '%s'  ORDER BY score DESC LIMIT %d;`
// ===== RANGE SCANS =====

{35, PostsByScoreRange{}},
`SELECT * FROM posts WHERE score BETWEEN %d AND %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC	LIMIT %d, %d;`
{35, PostsByViewRange{}},
`SELECT * FROM posts WHERE view_count BETWEEN %d AND %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC LIMIT %d, %d;`

// ===== COMMON FILTERS =====

{20, PostByScoreHigh{}},
`SELECT * FROM posts WHERE score > %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC LIMIT %d, %d;`
{15, PostByScoreLow{}},
`SELECT * FROM posts WHERE score < %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC LIMIT %d, %d;`
{20, PostByViewCoHigh{}},
`SELECT * FROM posts WHERE view_count > %d AND creation_date >= '%s' AND creation_date < '%s'  ORDER BY creation_date DESC LIMIT %d, %d;`
{15, PostByViewCoLow{}},
`SELECT * FROM posts WHERE view_count < %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY creation_date DESC LIMIT %d, %d;`
// ===== JOINS =====

{20, UserPostVotes{}},
`SELECT c.* FROM users u JOIN posts p ON u.id = p.owner_user_id JOIN comments c ON p.id = c.post_id WHERE u.id = %d AND p.creation_date >= '%s' AND p.creation_date < '%s' ORDER BY p.creation_date DESC;`

// ===== HISTORY =====

{15, PostHiByPosId{}},
`SELECT * FROM post_history WHERE post_id = %d;`
{10, PostHiByUsrId{}},
`SELECT * FROM post_history WHERE post_id = %d;`
{5, PostHiByTypeId{}},
`SELECT * FROM post_history WHERE post_history_type_id = %d;`

// ===== LINKS =====

{10, PostLiByPosId{}},
`SELECT * FROM post_links WHERE post_id = %d;`
{10, PostLiByRelPosId{}},
`SELECT * FROM post_links WHERE related_post_id = %d;`
{5, PostLiByLiTyId{}},
`SELECT * FROM comments WHERE post_id = %d;`

// ===== ANALYTICAL =====

{8, AveragePostViews{}},
`SELECT AVG(view_count) FROM posts WHERE creation_date >= '%s' AND creation_date < '%s';`
{8, AgregatePostScore{}},
`SELECT MAX(score) FROM posts WHERE creation_date >= '%s' AND creation_date < '%s';`
{8, PostPerUser{}},
`SELECT owner_user_id, COUNT(*) FROM posts WHERE creation_date >= '%s' AND creation_date < '%s' GROUP BY owner_user_id;	`
{8, AgregateVotes{}},
`SELECT post_id, COUNT(*) FROM votes WHERE creation_date >= '%s' AND creation_date < '%s' GROUP BY post_id;`

// ===== EXPENSIVE REPORTS =====

{8, PostByOwnerAndScore{}},
`SELECT * FROM posts WHERE owner_user_id = %d AND score > %d AND creation_date >= '%s' AND creation_date < '%s' ORDER BY score DESC LIMIT %d, %d;`
{8, PostByAnswerCount{}},
`SELECT * FROM posts WHERE comment_count > %d AND creation_date >= '%s' AND creation_date < '%s' LIMIT %d, %d;`
{8, PostTopViewed{}},
`SELECT id, score, view_count, creation_date FROM posts WHERE creation_date >= '%s' AND creation_date < '%s'  ORDER BY score DESC LIMIT %d;`

// ===== RECENT DATA =====

{5, RecentPosts{}},
`SELECT * FROM posts ORDER BY creation_date DESC LIMIT %d,%d;`

// ===== COMMENTS =====

{5, ComByScorHigh{}},
`SELECT * FROM comments WHERE score > %d LIMIT %d, %d;`
{5, ComByScorLow{}},
`SELECT * FROM comments WHERE score < %d LIMIT %d, %d;`


Users Records: 3473095  Deleted: 0  Skipped: 0  Warnings: 0
Query OK, 3473095 rows affected (42,65 sec)


Posts Records: 21736594  Deleted: 0  Skipped: 0  Warnings: 0
Query OK, 21736594 rows affected (10 min 49,04 sec)
sanity clining set user_id nul 3 records


Comments: Records: 36585420  Deleted: 0  Skipped: 0  Warnings: 0
Query OK, 36585420 rows affected (5 min 54,66 sec)
post_id null satnity clining 4055601
user_id nnull sanity clining 635216
Afret cleaning Records: 
add fk (10 min 33,00 sec)

post_history: Records: 59019898  Deleted: 0  Skipped: 0  Warnings: 0
Query OK, 59019898 rows affected (21 min 44,79 sec)
post_id null satnity clining delated  6729152
user_id nnull sanity clining set null 1400405
ad fk   52290746 rows affected (27 min 32,86 sec)

postlinks: Records: 2271053  Deleted: 0  Skipped: 0  Warnings: 0
Query OK, 2271053 rows affected (17,17 sec)
post_id1 or post_id2 null satnity clining delated  553350
ad fk 1717703 rows affected (1 min 13,64 sec)

votes: Records: 67258370  Deleted: 0  Skipped: 0  Warnings: 0
Query OK, 67258370 rows affected (7 min 38,07 sec)
post_id null satnity clining delated 8923429
ad fk 58334941 rows affected (13 min 14,78 sec)


badges: Records: 12783309  Deleted: 0  Skipped: 0  Warnings: 0
12783309 rows affected (1 min 33,26 sec)
ad fk 12783309 rows affected (1 min 9,50 sec)

tags: Records: 38205  Deleted: 0  Skipped: 0  Warnings: 0
38205 rows affected (0,43 sec)


