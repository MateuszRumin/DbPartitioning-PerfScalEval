package selectqueries

//rankingi, średnie kroczące, kumulowane sumy i porównania
var optAnalize = []string{
	//fulscans
	//Tabela badges
	"SELECT * FROM badges WHERE badge_name LIKE '%gold%';",
	//Tabela comments
	"SELECT * FROM comments WHERE comment_text LIKE '%thanks%';",
	//Tabela post_history
	"SELECT * FROM post_history WHERE post_text LIKE '%update%';",
	//Tabela post_links
	"SELECT * FROM post_links WHERE link_type_id = 1;",
	//Tabela posts
	"SELECT * FROM posts WHERE post_body LIKE '%error%';",
	//Tabela users
	"SELECT * FROM users WHERE display_name LIKE '%john%';'",
	//Tabela votes
	"SELECT * FROM votes WHERE vote_type_id = 2;",
	//Użytkownicy z odznakami za 2023 rok
	"EXPLAIN SELECT u.display_name, b.badge_name FROM users u INNER JOIN badges b ON u.id = b.user_id WHERE YEAR(b.badge_date) = 2023;",
	//Posty z komentarzami użytkownika
	"EXPLAIN SELECT p.post_title, c.comment_text  FROM posts p INNER JOIN comments c ON p.id = c.post_id WHERE p.owner_user_id = 100;",
	//Posty z historią edycji
	"EXPLAIN SELECT p.post_title, ph.post_text  FROM posts p INNER JOIN post_history ph ON p.id = ph.post_id WHERE ph.post_history_type_id = 5;",
	//Najnowsze odznaki
	"SELECT * FROM badges ORDER BY badge_date DESC LIMIT 10;",
	//Najczęściej komentowane posty
	"SELECT p.post_title, COUNT(c.id) AS comment_count FROM posts p INNER JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY comment_count DESC LIMIT 10;",
	//Użytkownicy z najwyższą reputacją
	"SELECT * FROM users ORDER BY reputation DESC LIMIT 10;",
	//Liczba odznak na użytkownika
	"SELECT user_id, COUNT(*) AS badge_count FROM badges GROUP BY user_id HAVING badge_count > 10;",
	//Średnia liczba komentarzy na post
	"SELECT AVG(comment_count) AS avg_comments FROM posts;",
	//Liczba głosów na typ głosu
	"SELECT vote_type_id, COUNT(*) AS vote_count FROM votes GROUP BY vote_type_id;",
	//Użytkownicy z >100 odznakami
	"SELECT * FROM users WHERE id IN (SELECT user_id FROM badges GROUP BY user_id HAVING COUNT(*) > 100);",
	//Posty z >1000 wyświetleniami i >50 komentarzami
	"SELECT * FROM posts WHERE view_count > 1000 AND id IN (SELECT post_id FROM comments GROUP BY post_id HAVING COUNT(*) > 50);",
	//Ranking użytkowników według liczby odznak
	"SELECT u.display_name, COUNT(b.id) OVER (PARTITION BY b.user_id) AS badge_count, RANK() OVER (ORDER BY COUNT(b.id) DESC) AS rank FROM users u LEFT JOIN badges b ON u.id = b.user_id;",
	//Średnia krocząca głosów na post
	"SELECT post_id, creation_date, AVG(vote_type_id) OVER (ORDER BY creation_date ROWS BETWEEN 7 PRECEDING AND CURRENT ROW) AS avg_votes_7d FROM votes;",
	//Pełny skan tabeli posts
	"SELECT * FROM posts WHERE score > 0 LIMIT 10000;",
	//Pełny skan tabeli comments
	"SELECT * FROM comments WHERE score > 0 LIMIT 10000;",
	//Pełny skan tabeli votes
	"SELECT * FROM votes WHERE vote_type_id = 2 LIMIT 10000;",
}

var indexAnalize = []string{
	//Sprawdzenie wykorzystania indeksów (EXPLAIN)
	//Tabela badges
	"EXPLAIN SELECT * FROM badges WHERE user_id = 100;",
	//Tabela comments
	"EXPLAIN SELECT * FROM comments WHERE post_id = 500 AND creation_date > '2023-01-01';",
	//Tabela post_history
	"EXPLAIN SELECT * FROM post_history WHERE post_id = 1000 AND user_id = 200;",
	//Tabela post_links
	"EXPLAIN SELECT * FROM post_links WHERE post_id = 300 AND related_post_id = 400;",
	//Tabela posts
	"EXPLAIN SELECT * FROM posts WHERE owner_user_id = 100 AND creation_date BETWEEN '2023-01-01' AND '2023-12-31';",
	//Tabela users
	"EXPLAIN SELECT * FROM users WHERE reputation > 1000 AND creation_date > '2022-01-01';",
	//Tabela votes
	"EXPLAIN SELECT * FROM votes WHERE post_id = 200 AND vote_type_id = 1;",
	//tworzenie indexów
	//Tabela badges
	"CREATE INDEX idx_badges_user_id ON badges(user_id);",
	"CREATE INDEX idx_badges_badge_date ON badges(badge_date);",
	//Tabela comments
	"CREATE INDEX idx_comments_post_id ON comments(post_id);",
	"CREATE INDEX idx_comments_creation_date ON comments(creation_date);",
	//Tabela post_history
	"CREATE INDEX idx_post_history_post_id ON post_history(post_id);",
	"CREATE INDEX idx_post_history_user_id ON post_history(user_id);",
	//Tabela post_links
	"CREATE INDEX idx_post_links_post_id ON post_links(post_id);",
	"CREATE INDEX idx_post_links_related_post_id ON post_links(related_post_id);",
	//Tabela posts
	"CREATE INDEX idx_posts_owner_user_id ON posts(owner_user_id);",
	"CREATE INDEX idx_posts_creation_date ON posts(creation_date);",
	//Tabela users
	"CREATE INDEX idx_users_reputation ON users(reputation);",
	"CREATE INDEX idx_users_creation_date ON users(creation_date);",
	//Tabela votes
	"CREATE INDEX idx_votes_post_id ON votes(post_id);",
	"CREATE INDEX idx_votes_vote_type_id ON votes(vote_type_id);",
	//Tabela badges
	"SELECT * FROM badges WHERE user_id = 100 AND badge_date > '2023-01-01';",
	//Tabela comments
	"SELECT * FROM comments WHERE post_id = 500 AND creation_date > '2023-01-01';",
	//Tabela post_history
	"SELECT * FROM post_history WHERE post_id = 1000 AND user_id = 200;",
	//Tabela post_links
	"SELECT * FROM post_links WHERE post_id = 300 AND related_post_id = 400;",
	//Tabela posts
	"SELECT * FROM posts WHERE owner_user_id = 100 AND creation_date BETWEEN '2023-01-01' AND '2023-12-31';",
	//Tabela users
	"SELECT * FROM users WHERE reputation > 1000 AND creation_date > '2022-01-01';",
	//Tabela votes
	"SELECT * FROM votes WHERE post_id = 200 AND vote_type_id = 1;",
}
