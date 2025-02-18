package selectqueries

//rankingi, średnie kroczące, kumulowane sumy i porównania
var FunokieTrendAnalize = []string{
	//Przykład 1: Ranking użytkowników według liczby odznak
	"SELECT u.display_name, COUNT(b.id) OVER (PARTITION BY b.user_id) AS badge_count, RANK() OVER (ORDER BY COUNT(b.id) DESC) AS rank FROM users u LEFT JOIN badges b ON u.id = b.user_id;",
	//Przykład 2: Średnia krocząca liczby odznak na użytkownika w czasie
	"SELECT b.user_id, b.badge_date, AVG(COUNT(b.id)) OVER (ORDER BY b.badge_date ROWS BETWEEN 30 PRECEDING AND CURRENT ROW) AS avg_badges_30d FROM badges b GROUP BY b.user_id, b.badge_date;",
	//Przykład 3: Kumulowana liczba odznak w czasie
	"SELECT b.user_id, b.badge_date, SUM(COUNT(b.id)) OVER (PARTITION BY b.user_id ORDER BY b.badge_date) AS cumulative_badges FROM badges b GROUP BY b.user_id, b.badge_date;",
	//Przykład 4: Użytkownicy z największą liczbą odznak w danym roku
	"SELECT u.display_name, YEAR(b.badge_date) AS year, COUNT(b.id) AS badge_count, RANK() OVER (PARTITION BY YEAR(b.badge_date) ORDER BY COUNT(b.id) DESC) AS rank FROM users u INNER JOIN badges b ON u.id = b.user_id GROUP BY u.display_name, YEAR(b.badge_date);",
	//Przykład 5: Porównanie liczby odznak między użytkownikami
	"SELECT  u1.display_name AS user1, u2.display_name AS user2, COUNT(b1.id) AS badges_user1, COUNT(b2.id) AS badges_user2 FROM users u1 CROSS JOIN users u2 LEFT JOIN badges b1 ON u1.id = b1.user_id LEFT JOIN badges b2 ON u2.id = b2.user_id WHERE u1.id <> u2.id GROUP BY u1.display_name, u2.display_name;",
	//Przykład 6: Ranking postów według liczby komentarzy
	"SELECT p.post_title, COUNT(c.id) AS comment_count, RANK() OVER (ORDER BY COUNT(c.id) DESC) AS rank FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.post_title;",
	//Przykład 7: Średnia krocząca liczby komentarzy na post w czasie
	"SELECT  p.id, p.creation_date, AVG(COUNT(c.id)) OVER (ORDER BY p.creation_date ROWS BETWEEN 7 PRECEDING AND CURRENT ROW) AS avg_comments_7d FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id, p.creation_date;",
	//Przykład 8: Kumulowana liczba komentarzy na post w czasie
	"SELECT p.id, p.creation_date, SUM(COUNT(c.id)) OVER (PARTITION BY p.id ORDER BY p.creation_date) AS cumulative_comments FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id, p.creation_date;",
	//Przykład 9: Posty z największą liczbą komentarzy w danym miesiącu
	"SELECT  p.post_title, MONTH(p.creation_date) AS month, COUNT(c.id) AS comment_count, RANK() OVER (PARTITION BY MONTH(p.creation_date) ORDER BY COUNT(c.id) DESC) AS rank FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.post_title, MONTH(p.creation_date);",
	//Przykład 10: Porównanie liczby komentarzy między postami
	"SELECT p1.post_title AS post1, p2.post_title AS post2, COUNT(c1.id) AS comments_post1, COUNT(c2.id) AS comments_post2 FROM posts p1 CROSS JOIN posts p2 LEFT JOIN comments c1 ON p1.id = c1.post_id LEFT JOIN comments c2 ON p2.id = c2.post_id WHERE p1.id <> p2.id GROUP BY p1.post_title, p2.post_title;",
	//Przykład 11: Ranking postów według liczby głosów
	"SELECT p.post_title,COUNT(v.id) AS vote_count, RANK() OVER (ORDER BY COUNT(v.id) DESC) AS rank FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.post_title;",
	// Przykład 12: Średnia krocząca głosów na post w czasie
	"SELECT p.id, p.creation_date, AVG(COUNT(v.id)) OVER (ORDER BY p.creation_date ROWS BETWEEN 30 PRECEDING AND CURRENT ROW) AS avg_votes_30d FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id, p.creation_date;",
	//Przykład 13: Kumulowana liczba głosów na post w czasie
	"SELECT p.id, p.creation_date, SUM(COUNT(v.id)) OVER (PARTITION BY p.id ORDER BY p.creation_date) AS cumulative_votes FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id, p.creation_date;",
	// Przykład 14: Posty z największą liczbą głosów w danym roku
	"SELECT p.post_title, YEAR(p.creation_date) AS year, COUNT(v.id) AS vote_count, RANK() OVER (PARTITION BY YEAR(p.creation_date) ORDER BY COUNT(v.id) DESC) AS rank FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.post_title, YEAR(p.creation_date);",
	//Przykład 15: Porównanie liczby głosów między postami
	"SELECT p1.post_title AS post1, p2.post_title AS post2, COUNT(v1.id) AS votes_post1, COUNT(v2.id) AS votes_post2 FROM posts p1 CROSS JOIN posts p2 LEFT JOIN votes v1 ON p1.id = v1.post_id LEFT JOIN votes v2 ON p2.id = v2.post_id WHERE p1.id <> p2.id GROUP BY p1.post_title, p2.post_title;",
	//Przykład 16: Ranking użytkowników według liczby edycji
	"SELECT u.display_name, COUNT(ph.id) AS edit_count, RANK() OVER (ORDER BY COUNT(ph.id) DESC) AS rank FROM users u LEFT JOIN post_history ph ON u.id = ph.user_id GROUP BY u.display_name;",
	//Przykład 17: Średnia krocząca liczby edycji na post w czasie
	"SELECT p.id, p.creation_date,AVG(COUNT(ph.id)) OVER (ORDER BY p.creation_date ROWS BETWEEN 7 PRECEDING AND CURRENT ROW) AS avg_edits_7d FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id, p.creation_date;",
	//Przykład 18: Kumulowana liczba edycji na post w czasie
	"SELECT p.id, p.creation_date, SUM(COUNT(ph.id)) OVER (PARTITION BY p.id ORDER BY p.creation_date) AS cumulative_edits FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id, p.creation_date;",
	//Przykład 19: Posty z największą liczbą edycji w danym miesiącu
	"SELECT p.post_title, MONTH(p.creation_date) AS month, COUNT(ph.id) AS edit_count, RANK() OVER (PARTITION BY MONTH(p.creation_date) ORDER BY COUNT(ph.id) DESC) AS rank FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.post_title, MONTH(p.creation_date);",
	//Pzzykład 20: Porównanie liczby edycji między postami
	"SELECT p1.post_title AS post1, p2.post_title AS post2, COUNT(ph1.id) AS edits_post1, COUNT(ph2.id) AS edits_post2 FROM posts p1 CROSS JOIN posts p2 LEFT JOIN post_history ph1 ON p1.id = ph1.post_id LEFT JOIN post_history ph2 ON p2.id = ph2.post_id WHERE p1.id <> p2.id GROUP BY p1.post_title, p2.post_title;",
	//Przykład 21: Ranking postów według liczby linków
	"SELECT p.post_title, COUNT(pl.id) AS link_count, RANK() OVER (ORDER BY COUNT(pl.id) DESC) AS rank FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.post_title;",
	//Przykład 22: Średnia krocząca liczby linków na post w czasie
	"SELECT p.id,p.creation_date, AVG(COUNT(pl.id)) OVER (ORDER BY p.creation_date ROWS BETWEEN 30 PRECEDING AND CURRENT ROW) AS avg_links_30d FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.id, p.creation_date;",
	//Przykład 23: Kumulowana liczba linków na post w czasie
	"SELECT p.id,p.creation_date, SUM(COUNT(pl.id)) OVER (PARTITION BY p.id ORDER BY p.creation_date) AS cumulative_links FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.id, p.creation_date;",
	//Przykład 24: Posty z największą liczbą linków w danym roku
	"SELECT p.post_title, YEAR(p.creation_date) AS year, COUNT(pl.id) AS link_count, RANK() OVER (PARTITION BY YEAR(p.creation_date) ORDER BY COUNT(pl.id) DESC) AS rank FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.post_title, YEAR(p.creation_date);",
	//Przykład 25: Porównanie liczby linków między postami
	"SELECT p1.post_title AS post1,p2.post_title AS post2,COUNT(pl1.id) AS links_post1, COUNT(pl2.id) AS links_post2 FROM posts p1 CROSS JOIN posts p2 LEFT JOIN post_links pl1 ON p1.id = pl1.post_id LEFT JOIN post_links pl2 ON p2.id = pl2.post_id WHERE p1.id <> p2.id GROUP BY p1.post_title, p2.post_title;",
	//Przykład 26: Ranking użytkowników według reputacji
	"SELECT u.display_name,u.reputation, RANK() OVER (ORDER BY u.reputation DESC) AS rank FROM users u;",
	//Pzzykład 27: Średnia krocząca reputacji użytkowników w czasie
	"SELECT u.id,u.creation_date,AVG(u.reputation) OVER (ORDER BY u.creation_date ROWS BETWEEN 365 PRECEDING AND CURRENT ROW) AS avg_reputation_1y FROM users u;",
	//Przykład 28: Kumulowana reputacja użytkowników w czasie
	"SELECT u.id,u.creation_date,SUM(u.reputation) OVER (PARTITION BY u.id ORDER BY u.creation_date) AS cumulative_reputation FROM users u;",
	//Przykład 29: Użytkownicy z największą reputacją w danym roku
	"SELECT u.display_name,YEAR(u.creation_date) AS year,u.reputation,RANK() OVER (PARTITION BY YEAR(u.creation_date) ORDER BY u.reputation DESC) AS rank FROM users u;",
	//Przykład 30: Porównanie reputacji między użytkownikami
	"SELECT u1.display_name AS user1, u2.display_name AS user2, u1.reputation AS reputation_user1, u2.reputation AS reputation_user2 FROM users u1 CROSS JOIN users u2 WHERE u1.id <> u2.id;",
	//Przykład 31: Ranking użytkowników według liczby postów
	"SELECT u.display_name, COUNT(p.id) AS post_count, RANK() OVER (ORDER BY COUNT(p.id) DESC) AS rank FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.display_name;",
	//Przykład 32: Średnia krocząca liczby postów na użytkownika w czasie
	"SELECT u.id,u.creation_date, AVG(COUNT(p.id)) OVER (ORDER BY u.creation_date ROWS BETWEEN 30 PRECEDING AND CURRENT ROW) AS avg_posts_30d FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.id, u.creation_date;",
	//Przykład 33: Kumulowana liczba postów na użytkownika w czasie
	"SELECT u.id, u.creation_date, SUM(COUNT(p.id)) OVER (PARTITION BY u.id ORDER BY u.creation_date) AS cumulative_posts FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.id, u.creation_date;",
	//Przykład 34: Użytkownicy z największą liczbą postów w danym roku
	"SELECT u.display_name, YEAR(p.creation_date) AS year, COUNT(p.id) AS post_count, RANK() OVER (PARTITION BY YEAR(p.creation_date) ORDER BY COUNT(p.id) DESC) AS rank FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.display_name, YEAR(p.creation_date);",
	//Przykład 35: Porównanie liczby postów między użytkownikami
	"SELECT u1.display_name AS user1, u2.display_name AS user2, COUNT(p1.id) AS posts_user1, COUNT(p2.id) AS posts_user2 FROM users u1 CROSS JOIN users u2 LEFT JOIN posts p1 ON u1.id = p1.owner_user_id LEFT JOIN posts p2 ON u2.id = p2.owner_user_id WHERE u1.id <> u2.id GROUP BY u1.display_name, u2.display_name;",
	//Przykład 36: Ranking postów według liczby wyświetleń
	"SELECT p.post_title, p.view_count, RANK() OVER (ORDER BY p.view_count DESC) AS rank FROM posts p;",
	//Przykład 37: Średnia krocząca wyświetleń na post w czasie
	"SELECT p.id, p.creation_date, AVG(p.view_count) OVER (ORDER BY p.creation_date ROWS BETWEEN 7 PRECEDING AND CURRENT ROW) AS avg_views_7d FROM posts p;",
	//Przykład 38: Kumulowana liczba wyświetleń na post w czasie
	"SELECT p.id, p.creation_date, SUM(p.view_count) OVER (PARTITION BY p.id ORDER BY p.creation_date) AS cumulative_views FROM posts p;",
	//Przykład 39: Posty z największą liczbą wyświetleń w danym roku
	"SELECT p.post_title,YEAR(p.creation_date) AS year, p.view_count, RANK() OVER (PARTITION BY YEAR(p.creation_date) ORDER BY p.view_count DESC) AS rank FROM posts p;",
	//Przykład 40: Porównanie liczby wyświetleń między postami
	"SELECT p1.post_title AS post1,p2.post_title AS post2, p1.view_count AS views_post1, p2.view_count AS views_post2 FROM posts p1 CROSS JOIN posts p2 WHERE p1.id <> p2.id;",
	//Przykład 41: Ranking postów według liczby odpowiedzi
	"SELECT p.post_title, p.answer_count, RANK() OVER (ORDER BY p.answer_count DESC) AS rank FROM posts p;",
	//Przykład 42: Średnia krocząca liczby odpowiedzi na post w czasie
	"SELECT p.id, p.creation_date, AVG(p.answer_count) OVER (ORDER BY p.creation_date ROWS BETWEEN 30 PRECEDING AND CURRENT ROW) AS avg_answers_30d FROM posts p;",
	//Przykład 43: Kumulowana liczba odpowiedzi na post w czasie
	"SELECT p.id, p.creation_date, SUM(p.answer_count) OVER (PARTITION BY p.id ORDER BY p.creation_date) AS cumulative_answers FROM posts p;",
	//Przykład 44: Posty z największą liczbą odpowiedzi w danym roku
	"SELECT p.post_title, YEAR(p.creation_date) AS year, p.answer_count, RANK() OVER (PARTITION BY YEAR(p.creation_date) ORDER BY p.answer_count DESC) AS rank FROM posts p;",
	//Przykład 45: Porównanie liczby odpowiedzi między postami
	"SELECT p1.post_title AS post1, p2.post_title AS post2, p1.answer_count AS answers_post1, p2.answer_count AS answers_post2 FROM posts p1 CROSS JOIN posts p2 WHERE p1.id <> p2.id;",
	//Przykład 46: Ranking użytkowników według liczby komentarzy
	"SELECT u.display_name, COUNT(c.id) AS comment_count, RANK() OVER (ORDER BY COUNT(c.id) DESC) AS rank FROM users u LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.display_name;",
	//Przykład 47: Średnia krocząca liczby komentarzy na użytkownika w czasie
	"SELECT u.id, u.creation_date, AVG(COUNT(c.id)) OVER (ORDER BY u.creation_date ROWS BETWEEN 30 PRECEDING AND CURRENT ROW) AS avg_comments_30d FROM users u LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id, u.creation_date;",
	//Przykład 48: Kumulowana liczba komentarzy na użytkownika w czasie
	"SELECT u.id, u.creation_date, SUM(COUNT(c.id)) OVER (PARTITION BY u.id ORDER BY u.creation_date) AS cumulative_comments FROM users u LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id, u.creation_date;",
	//Przykład 49: Użytkownicy z największą liczbą komentarzy w danym roku
	"SELECT u.display_name, YEAR(c.creation_date) AS year, COUNT(c.id) AS comment_count, RANK() OVER (PARTITION BY YEAR(c.creation_date) ORDER BY COUNT(c.id) DESC) AS rank FROM users u LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.display_name, YEAR(c.creation_date);",
	//Przykład 50: Porównanie liczby komentarzy między użytkownikami
	"SELECT u1.display_name AS user1, u2.display_name AS user2, COUNT(c1.id) AS comments_user1, COUNT(c2.id) AS comments_user2 FROM users u1 CROSS JOIN users u2 LEFT JOIN comments c1 ON u1.id = c1.user_id LEFT JOIN comments c2 ON u2.id = c2.user_id WHERE u1.id <> u2.id GROUP BY u1.display_name, u2.display_name;",
}
