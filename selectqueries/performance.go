package selectqueries

var Performacne = []string{
	// Pełne skany tabel
	// Przykład 1: Pełny skan tabeli posts
	"SELECT * FROM posts WHERE score > 0 LIMIT 10000;",
	// Przykład 2: Pełny skan tabeli users
	"SELECT * FROM users WHERE reputation > 100 LIMIT 10000;",
	// Przykład 3: Pełny skan tabeli comments
	"SELECT * FROM comments WHERE score > 0 LIMIT 10000;",
	//Przykład 4: Pełny skan tabeli votes
	"SELECT * FROM votes WHERE vote_type_id = 2 LIMIT 10000;",
	// Przykład 5: Pełny skan tabeli badges
	"SELECT * FROM badges WHERE badge_date > '2023-01-01' LIMIT 10000;",
	//Złożone zapytania z JOIN
	// Przykład 6: Posty z komentarzami i użytkownikami
	"SELECT p.post_title, c.comment_text, u.display_name FROM posts p INNER JOIN comments c ON p.id = c.post_id INNER JOIN users u ON c.user_id = u.id LIMIT 1000;",
	// Przykład 7: Użytkownicy z odznakami i ich postami
	"SELECT u.display_name, b.badge_name, p.post_title FROM users u INNER JOIN badges b ON u.id = b.user_id LEFT JOIN posts p ON u.id = p.owner_user_id LIMIT 1000;",
	//Przykład 8: Posty z historią edycji i użytkownikami
	"SELECT p.post_title, ph.post_text, u.display_name FROM posts p INNER JOIN post_history ph ON p.id = ph.post_id INNER JOIN users u ON ph.user_id = u.id LIMIT 1000;",
	//Przykład 9: Posty z linkami i powiązanymi postami
	"SELECT p1.post_title AS post, p2.post_title AS related_post FROM post_links pl INNER JOIN posts p1 ON pl.post_id = p1.id INNER JOIN posts p2 ON pl.related_post_id = p2.id LIMIT 1000;",
	//Przykład 10: Posty z głosami i użytkownikami
	"SELECT p.post_title, v.vote_type_id, u.display_name FROM posts p INNER JOIN votes v ON p.id = v.post_id INNER JOIN users u ON v.user_id = u.id LIMIT 1000;",
	//Agregacje i GROUP BY
	//Przykład 11: Liczba postów na użytkownika
	"SELECT owner_user_id, COUNT(*) AS post_count FROM posts GROUP BY owner_user_id HAVING post_count > 10 LIMIT 1000;",
	//Przykład 12: Liczba komentarzy na post
	"SELECT post_id, COUNT(*) AS comment_count FROM comments GROUP BY post_id HAVING comment_count > 5 LIMIT 1000;",
	//Przykład 13: Liczba odznak na użytkownika
	"SELECT user_id, COUNT(*) AS badge_count FROM badges GROUP BY user_id HAVING badge_count > 3 LIMIT 1000;",
	//Przykład 14: Liczba głosów na post
	"SELECT post_id, COUNT(*) AS vote_count FROM votes GROUP BY post_id HAVING vote_count > 10 LIMIT 1000;",
	//Przykład 15: Liczba edycji na post
	"SELECT post_id, COUNT(*) AS edit_count FROM post_history GROUP BY post_id HAVING edit_count > 2 LIMIT 1000;",
	//Sortowanie i LIMIT
	//Przykład 16: Najlepiej oceniane posty
	"SELECT * FROM posts ORDER BY score DESC LIMIT 1000;",
	//Przykład 17: Najczęściej wyświetlane posty
	"SELECT * FROM posts ORDER BY view_count DESC LIMIT 1000;",
	//Przykład 18: Najnowsze posty
	"SELECT * FROM posts ORDER BY creation_date DESC LIMIT 1000;",
	//Przykład 19: Najnowsze komentarze
	"SELECT * FROM comments ORDER BY creation_date DESC LIMIT 1000;",
	//Przykład 20: Najnowsze głosy
	"SELECT * FROM votes ORDER BY creation_date DESC LIMIT 1000;",
	//Zapytania z podzapytaniami
	//Przykład 21: Użytkownicy z >100 postami
	"SELECT * FROM users WHERE id IN ( SELECT owner_user_id FROM posts GROUP BY owner_user_id HAVING COUNT(*) > 100 );",
	//Przykład 22: Posty z >10 komentarzami
	"SELECT * FROM posts WHERE id IN ( SELECT post_id FROM comments GROUP BY post_id HAVING COUNT(*) > 10 );",
	//Przykład 23: Użytkownicy z odznakami za 2023 rok
	"SELECT * FROM users WHERE id IN ( SELECT user_id FROM badges WHERE YEAR(badge_date) = 2023);",
	//Przykład 24: Posty z >50 głosami
	"SELECT * FROM posts WHERE id IN (SELECT post_id FROM votes GROUP BY post_id HAVING COUNT(*) > 50);",
	//Przykład 25: Użytkownicy z >1000 punktami reputacji
	"SELECT * FROM users WHERE reputation > 1000;",
	//Zapytania z funkcjami okiennymi
	//Przykład 26: Ranking użytkowników według liczby postów
	"SELECT display_name,COUNT(*) OVER (PARTITION BY owner_user_id) AS post_count, RANK() OVER (ORDER BY COUNT(*) DESC) AS rank FROM posts INNER JOIN users ON posts.owner_user_id = users.id;",
	//Przykład 27: Średnia krocząca głosów na post
	"SELECT post_id, creation_date, AVG(score) OVER (ORDER BY creation_date ROWS BETWEEN 7 PRECEDING AND CURRENT ROW) AS avg_score_7d FROM posts;",
	//Przykład 28: Liczba komentarzy na post w czasie
	"SELECT post_id, creation_date, COUNT(*) OVER (PARTITION BY post_id ORDER BY creation_date) AS comment_count FROM comments;",
	//Przykład 29: Suma głosów na użytkownika
	"SELECT user_id, SUM(vote_type_id) OVER (PARTITION BY user_id) AS total_votes FROM votes;",
	//Przykład 30: Liczba odznak na użytkownika w czasie
	"SELECT user_id,badge_date,COUNT(*) OVER (PARTITION BY user_id ORDER BY badge_date) AS badge_count FROM badges;",
	// Przykład 31: Posty z najnowszą historią edycji
	"WITH LatestEdits AS (SELECT post_id, MAX(creation_date) AS last_edit FROM post_history GROUP BY post_id ) SELECT p.post_title, le.last_edit FROM posts p INNER JOIN LatestEdits le ON p.id = le.post_id;",
	//Przykład 32: Użytkownicy z największą liczbą odznak
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u INNER JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY badge_count DESC LIMIT 10;",
	//Przykład 33: Posty z największą liczbą głosów i komentarzy
	"SELECT p.post_title, COUNT(v.id) AS vote_count, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY vote_count DESC, comment_count DESC LIMIT 10;",
	//Przykład 34: Użytkownicy z największą liczbą postów i komentarzy
	"SELECT u.display_name, COUNT(p.id) AS post_count, COUNT(c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ORDER BY post_count DESC, comment_count DESC LIMIT 10;",
	//Przykład 35: Posty z największą liczbą edycji
	"SELECT p.post_title, COUNT(ph.id) AS edit_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id ORDER BY edit_count DESC LIMIT 10;",
	//Testy wydajnościowe z dużymi zbiorami danych
	//Przykład 36: Pełny skan tabeli posts z sortowaniem
	"SELECT * FROM posts ORDER BY creation_date DESC LIMIT 10000;",
	//Przykład 37: Pełny skan tabeli comments z filtrem
	"SELECT * FROM comments WHERE score > 0 LIMIT 10000;",
	//Przykład 38: Pełny skan tabeli votes z JOIN
	"SELECT v.*, p.post_title FROM votes v INNER JOIN posts p ON v.post_id = p.id LIMIT 10000;",
	//Przykład 39: Pełny skan tabeli badges z GROUP BY
	"SELECT user_id, COUNT(*) AS badge_count FROM badges GROUP BY user_id HAVING badge_count > 5 LIMIT 10000;",
	//Przykład 40: Pełny skan tabeli post_history z JOIN
	"SELECT ph.*, p.post_title FROM post_history ph INNER JOIN posts p ON ph.post_id = p.id LIMIT 10000;",
	//Testy złożonych JOIN i agregacji
	//Przykład 41: Posty z liczbą komentarzy i głosów
	"SELECT p.post_title, COUNT(c.id) AS comment_count, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id ORDER BY comment_count DESC, vote_count DESC LIMIT 1000;",
	//Przykład 42: Użytkownicy z liczbą postów, komentarzy i odznak
	"SELECT u.display_name, COUNT(p.id) AS post_count, COUNT(c.id) AS comment_count, COUNT(b.id) AS badge_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY post_count DESC, comment_count DESC, badge_count DESC LIMIT 1000;",
	//Przykład 43: Posty z największą liczbą edycji i komentarzy
	"SELECT p.post_title, COUNT(ph.id) AS edit_count, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY edit_count DESC, comment_count DESC LIMIT 1000;",
	//Przykład 44: Użytkownicy z największą liczbą głosów i odznak
	"SELECT u.display_name, COUNT(v.id) AS vote_count, COUNT(b.id) AS badge_count FROM users u LEFT JOIN votes v ON u.id = v.user_id LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY vote_count DESC, badge_count DESC LIMIT 1000;",
	//Przykład 45: Posty z największą liczbą linków i komentarzy
	"SELECT p.post_title, COUNT(pl.id) AS link_count, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY link_count DESC, comment_count DESC LIMIT 1000;",
	//Testy złożonych zapytań z podzapytaniami
	//Przykład 46: Użytkownicy z >100 postami i >1000 punktami reputacji
	"SELECT * FROM users WHERE id IN ( SELECT owner_user_id FROM posts GROUP BY owner_user_id HAVING COUNT(*) > 100) AND reputation > 1000;",
	//Przykład 47: Posty z >10 komentarzami i >50 głosami
	"SELECT * FROM posts WHERE id IN ( SELECT post_id FROM comments GROUP BY post_id HAVING COUNT(*) > 10 ) AND id IN ( SELECT post_id FROM votes GROUP BY post_id HAVING COUNT(*) > 50 );",
	// Przykład 48: Użytkownicy z odznakami za 2023 rok i >100 postami
	"SELECT * FROM users WHERE id IN ( SELECT user_id FROM badges WHERE YEAR(badge_date) = 2023 ) AND id IN ( SELECT owner_user_idFROM posts GROUP BY owner_user_id HAVING COUNT(*) > 100 );",
	//Przykład 49: Posty z >5 edycjami i >10 komentarzami
	"SELECT * FROM posts WHERE id IN ( SELECT post_id FROM post_history GROUP BY post_id HAVING COUNT(*) > 5 ) AND id IN ( SELECT post_id FROM comments GROUP BY post_id HAVING COUNT(*) > 10 );",
	//Przykład 50: Użytkownicy z >50 głosami i >10 odznakami
	"SELECT * FROM users WHERE id IN (SELECT user_id FROM votes GROUP BY user_id HAVING COUNT(*) > 50 ) AND id IN (SELECT user_id FROM badges GROUP BY user_id HAVING COUNT(*) > 10 );",
}
