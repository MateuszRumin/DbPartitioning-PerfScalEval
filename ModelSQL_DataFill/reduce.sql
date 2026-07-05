
#small database

DELETE FROM votes
WHERE creation_date >= '2010-12-31';

DELETE FROM post_links
WHERE creation_date >= '2010-12-31';

DELETE FROM post_history
WHERE creation_date >= '2010-12-31';

DELETE FROM comments
WHERE creation_date >= '2010-12-31';

DELETE FROM badges
WHERE badge_date >= '2010-12-31';

DELETE c
FROM post_history c
JOIN posts p ON p.id = c.post_id
WHERE p.creation_date >= '2010-12-31';





DELETE FROM posts
WHERE creation_date >= '2010-12-31';


DELETE c
FROM comments c
JOIN users p ON p.id = c.user_id
WHERE p.creation_date >= '2010-12-31';


DELETE c
FROM badges c
JOIN users p ON p.id = c.user_id
WHERE p.creation_date >= '2010-12-31';

DELETE c
FROM post_history c
JOIN users p ON p.id = c.user_id
WHERE p.creation_date >= '2010-12-31';

DELETE ph
FROM comments ph
JOIN posts p ON p.id=ph.post_id
JOIN users u ON u.id = p.owner_user_id
WHERE u.creation_date >= '2010-12-31';

DELETE c
FROM posts c
JOIN users p ON p.id = c.owner_user_id
WHERE p.creation_date >= '2010-12-31';



DELETE FROM users
WHERE creation_date >= '2010-12-31';