package selectqueries

var JoinAnalize = []string{
	// Przykład 1: Posty z komentarzami użytkownika
	"SELECT p.post_title, c.comment_text FROM posts p INNER JOIN comments c ON p.id = c.post_id WHERE p.owner_user_id = 100;",
	// Przykład 2: Posty z komentarzami i nazwą użytkownika
	"SELECT p.post_title, c.comment_text, u.display_name FROM posts p INNER JOIN comments c ON p.id = c.post_id INNER JOIN users u ON c.user_id = u.id;",
	// Przykład 3: Posty z liczbą komentarzy
	"SELECT p.post_title, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id;",
	//Przykład 4: Użytkownicy z odznakami za 2023 rok
	"SELECT u.display_name, b.badge_name FROM users u INNER JOIN badges b ON u.id = b.user_id WHERE YEAR(b.badge_date) = 2023;",
	// Przykład 5: Użytkownicy z odznakami typu "tag_based"
	"SELECT u.display_name, b.badge_name FROM users u INNER JOIN badges b ON u.id = b.user_id WHERE b.tag_based = 'true';",
	//Przykład 6: Liczba odznak na użytkownika
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id;",
	//Przykład 7: Posty z najnowszą historią edycji
	"SELECT p.post_title, ph.post_text, ph.creation_date FROM posts p INNER JOIN post_history ph ON p.id = ph.post_id WHERE ph.post_history_type_id = 5; ", // przykładowy typ edycji
	//Przykład 8: Posty z liczbą edycji
	"SELECT p.post_title, COUNT(ph.id) AS edit_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id;",
	//Przykład 9: Użytkownicy, którzy edytowali posty
	"SELECT u.display_name, COUNT(ph.id) AS edit_count FROM users u INNER JOIN post_history ph ON u.id = ph.user_id GROUP BY u.id;",
	//Przykład 10: Posty z linkami do powiązanych postów
	"SELECT p1.post_title, p2.post_title AS related_post FROM post_links pl INNER JOIN posts p1 ON pl.post_id = p1.id INNER JOIN posts p2 ON pl.related_post_id = p2.id;",
	//Przykład 11: Liczba powiązanych postów
	"SELECT p.post_title, COUNT(pl.related_post_id) AS related_count FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.id;",
	//Przykład 12: Posty z linkami typu "duplikat"
	"SELECT p1.post_title, p2.post_title AS duplicate_post FROM post_links pl INNER JOIN posts p1 ON pl.post_id = p1.id INNER JOIN posts p2 ON pl.related_post_id = p2.id WHERE pl.link_type_id = 3; -- przykładowy typ linku (duplikat)",
	//Przykład 13: Użytkownicy z liczbą postów i komentarzy
	"SELECT u.display_name, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id;",
	//Przykład 14: Użytkownicy z postami, które mają >10 komentarzy
	"SELECT u.display_name, p.post_title, COUNT(c.id) AS comment_count FROM users u INNER JOIN posts p ON u.id = p.owner_user_id INNER JOIN comments c ON p.id = c.post_id GROUP BY u.id, p.id HAVING comment_count > 10;",
	//Przykład 15: Użytkownicy z postami i odznakami
	"SELECT u.display_name, p.post_title, b.badge_name FROM users u INNER JOIN posts p ON u.id = p.owner_user_id LEFT JOIN badges b ON u.id = b.user_id;",
	//Przykład 16: Posty z liczbą głosów
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id;",
	//Przykład 17: Posty z głosami typu "upvote"
	"SELECT p.post_title, COUNT(v.id) AS upvote_count FROM posts p INNER JOIN votes v ON p.id = v.post_id WHERE v.vote_type_id = 2 -- przykładowy typ głosu (upvote) GROUP BY p.id;",
	//Przykład 18: Użytkownicy z liczbą głosów na ich posty
	"SELECT u.display_name, COUNT(v.id) AS vote_count FROM users u INNER JOIN posts p ON u.id = p.owner_user_id LEFT JOIN votes v ON p.id = v.post_id GROUP BY u.id;",
	//Przykład 19: Posty z komentarzami, głosami i odznakami
	"SELECT p.post_title, COUNT(DISTINCT c.id) AS comment_count, COUNT(DISTINCT v.id) AS vote_count, COUNT(DISTINCT b.id) AS badge_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id LEFT JOIN votes v ON p.id = v.post_id LEFT JOIN users u ON p.owner_user_id = u.id LEFT JOIN badges b ON u.id = b.user_id GROUP BY p.id;",
	//Przykład 20: Użytkownicy z postami, komentarzami i odznakami
	"SELECT u.display_name, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count, COUNT(DISTINCT b.id) AS badge_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id;",
	//Przykład 21: Posty z komentarzami z ostatniego miesiąca
	"SELECT p.post_title, c.comment_text, c.creation_date FROM posts p  INNER JOIN comments c ON p.id = c.post_id WHERE c.creation_date >= NOW() - INTERVAL 1 MONTH;",
	//Przykład 22: Użytkownicy z odznakami z ostatniego roku
	"SELECT u.display_name, b.badge_name, b.badge_date FROM users u INNER JOIN badges b ON u.id = b.user_id WHERE b.badge_date >= NOW() - INTERVAL 1 YEAR;",
	//Przykład 23: Posty z historią edycji z ostatniego tygodnia
	"SELECT p.post_title, ph.post_text, ph.creation_date FROM posts p INNER JOIN post_history ph ON p.id = ph.post_id WHERE ph.creation_date >= NOW() - INTERVAL 1 WEEK;",
	//Przykład 24: Posty z >1000 wyświetleniami i >50 głosami
	"SELECT p.post_title, p.view_count, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id WHERE p.view_count > 1000 GROUP BY p.id HAVING vote_count > 50;",
	//Przykład 25: Użytkownicy z reputacją >1000 i >10 odznakami
	"SELECT u.display_name, u.reputation, COUNT(b.id) AS badge_count FROM users u LEFT JOIN badges b ON u.id = b.user_id WHERE u.reputation > 1000 GROUP BY u.id HAVING badge_count > 10;",
	//Przykład 26: Najczęściej komentowane posty
	"SELECT p.post_title, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY comment_count DESC LIMIT 10;",
	//Przykład 27: Najczęściej głosowane posty
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id ORDER BY vote_count DESC LIMIT 10;",
	//Przykład 28: Użytkownicy z największą liczbą odznak
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY badge_count DESC LIMIT 10;",
	//Przykład 29: Użytkownicy z najnowszymi postami
	"SELECT u.display_name, p.post_title, p.creation_date FROM users u INNER JOIN posts p ON u.id = p.owner_user_id WHERE p.creation_date = (SELECT MAX(creation_date) FROM posts WHERE owner_user_id = u.id);",
	//Przykład 30: Posty z najnowszymi komentarzami
	"SELECT p.post_title, c.comment_text, c.creation_date FROM posts p INNER JOIN comments c ON p.id = c.post_id WHERE c.creation_date = (SELECT MAX(creation_date) FROM comments WHERE post_id = p.id);",
	//Przykład 31: Średnia liczba komentarzy na post
	"SELECT AVG(comment_count) AS avg_comments FROM ( SELECT p.id, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ) AS comment_stats;",
	//Przykład 32: Średnia liczba głosów na post
	"SELECT AVG(vote_count) AS avg_votes FROM ( SELECT p.id, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id ) AS vote_stats;",
	//Przykład 33: Posty i komentarze użytkownika
	"SELECT p.post_title AS content, 'post' AS type FROM posts p WHERE p.owner_user_id = 100 UNION SELECT c.comment_text AS content, 'comment' AS type FROM comments c WHERE c.user_id = 100;",
	//Przykład 34: Użytkownicy z największą liczbą aktywności
	"WITH UserActivity AS (SELECT u.id, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ) SELECT u.display_name, ua.post_count, ua.comment_count FROM users u INNER JOIN UserActivity ua ON u.id = ua.id ORDER BY (ua.post_count + ua.comment_count) DESC LIMIT 10;",
	//Przykład 35: Ranking postów według liczby komentarzy
	"SELECT p.post_title, COUNT(c.id) AS comment_count, RANK() OVER (ORDER BY COUNT(c.id) DESC) AS rank FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id;",
	//Przykład 36: Posty z tagami i użytkownikami
	"SELECT p.post_title, p.tags, u.display_name FROM posts p INNER JOIN users u ON p.owner_user_id = u.id WHERE p.tags LIKE '%java%';",
	//Przykład 37: Drzewo dyskusji
	"WITH RECURSIVE DiscussionTree AS ( SELECT id, parent_id, post_title FROM posts WHERE parent_id IS NULL UNION ALL SELECT p.id, p.parent_id, p.post_title FROM posts p INNER JOIN DiscussionTree dt ON p.parent_id = dt.id ) SELECT * FROM DiscussionTree;",
	//Przykład 38: Posty bez komentarzy
	"SELECT p.post_title FROM posts p LEFT JOIN comments c ON p.id = c.post_id WHERE c.id IS NULL;",
	//Przykład 39: Użytkownicy bez postów
	"SELECT u.display_name FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id WHERE p.id IS NULL;",
	//Przykład 40: Posty z komentarzami i głosami
	"SELECT p.post_title, COUNT(DISTINCT c.id) AS comment_count,  COUNT(DISTINCT v.id) AS vote_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id;",
	//Przykład 41: Posty z ostatnimi komentarzami
	"SELECT p.post_title, c.comment_text, c.creation_date FROM posts p INNER JOIN comments c ON p.id = c.post_id WHERE c.creation_date = (SELECT MAX(creation_date) FROM comments WHERE post_id = p.id);",
	//Przykład 42: Użytkownicy z tej samej lokalizacji
	"SELECT u1.display_name AS user1, u2.display_name AS user2, u1.location FROM users u1 INNER JOIN users u2 ON u1.location = u2.location AND u1.id <> u2.id;",
	//Przykład 43: Posty z najnowszymi komentarzami i głosami
	"SELECT p.post_title, (SELECT comment_text FROM comments WHERE post_id = p.id ORDER BY creation_date DESC LIMIT 1) AS latest_comment, (SELECT COUNT(*) FROM votes WHERE post_id = p.id) AS vote_count FROM posts p;",
	//Przykład 44: Posty z komentarzami i odpowiedziami
	"SELECT p.post_title, c.comment_text, a.post_title AS answer_title FROM posts p LEFT JOIN comments c ON p.id = c.post_id LEFT JOIN posts a ON p.accepted_answer_id = a.id;",
	//Przykład 45: Posty z komentarzami od użytkowników z reputacją >1000
	"SELECT p.post_title, c.comment_text, u.display_name FROM posts p INNER JOIN comments c ON p.id = c.post_id INNER JOIN users u ON c.user_id = u.id WHERE u.reputation > 1000;",
	//Przykład 46: Posty z komentarzami sortowane po dacie
	"SELECT p.post_title, c.comment_text, c.creation_date FROM posts p INNER JOIN comments c ON p.id = c.post_id ORDER BY c.creation_date DESC;",
	//Przykład 47: Posty z najnowszymi komentarzami (limit 10)
	"SELECT p.post_title, c.comment_text, c.creation_date FROM posts p INNER JOIN comments c ON p.id = c.post_id ORDER BY c.creation_date DESC LIMIT 10;",
	//Przykład 48: Posty z rankingiem komentarzy
	"SELECT p.post_title, c.comment_text, RANK() OVER (PARTITION BY p.id ORDER BY c.creation_date DESC) AS comment_rank FROM posts p INNER JOIN comments c ON p.id = c.post_id;",
	//Przykład 49: Posty z najnowszymi komentarzami i głosami
	"SELECT p.post_title, (SELECT comment_text FROM comments WHERE post_id = p.id ORDER BY creation_date DESC LIMIT 1) AS latest_comment, (SELECT COUNT(*) FROM votes WHERE post_id = p.id) AS vote_count FROM posts p;",
	//Przykład 50: Użytkownicy z największą liczbą aktywności
	"WITH UserActivity AS ( SELECT u.id,  COUNT(DISTINCT p.id) AS post_count,  COUNT(DISTINCT c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ) SELECT u.display_name, ua.post_count, ua.comment_count FROM users u INNER JOIN UserActivity ua ON u.id = ua.id ORDER BY (ua.post_count + ua.comment_count) DESC LIMIT 10;",
}
