package selectqueries

var SubAndCteAnalize = []string{
	//Przykład 1: Użytkownicy z więcej niż 5 odznakami
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u INNER JOIN badges b ON u.id = b.user_id GROUP BY u.id HAVING badge_count > 5;",
	//Przykład 2: Użytkownicy z odznakami w 2023 roku (CTE)
	"WITH Badges2023 AS (SELECT user_id, badge_name FROM badges WHERE YEAR(badge_date) = 2023 ) SELECT u.display_name, b.badge_name FROM users u INNER JOIN Badges2023 b ON u.id = b.user_id;",
	//Przykład 3: Użytkownicy z największą liczbą odznak
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u INNER JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY badge_count DESC LIMIT 10;",
	//Przykład 4: Posty z więcej niż 10 komentarzami
	"SELECT p.post_title, COUNT(c.id) AS comment_count FROM posts p INNER JOIN comments c ON p.id = c.post_id GROUP BY p.id HAVING comment_count > 10;",
	//Przykład 5: Posty bez komentarzy (podzapytanie)
	"SELECT p.post_title FROM posts p WHERE p.id NOT IN (SELECT post_id FROM comments);",
	//Przykład 6: Posty z najnowszym komentarzem (CTE)
	"WITH LatestComment AS ( SELECT post_id, MAX(creation_date) AS last_comment_date FROM comments GROUP BY post_id ) SELECT p.post_title, lc.last_comment_date FROM posts p INNER JOIN LatestComment lc ON p.id = lc.post_id;",
	//Przykład 7: Posty z więcej niż 50 głosami
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p INNER JOIN votes v ON p.id = v.post_id GROUP BY p.id HAVING vote_count > 50;",
	//Przykład 8: Użytkownicy, którzy oddali najwięcej głosów
	"SELECT u.display_name, COUNT(v.id) AS vote_count FROM users u INNER JOIN votes v ON u.id = v.user_id GROUP BY u.id ORDER BY vote_count DESC LIMIT 10;",
	//Przykład 9: Posty z największą liczbą głosów w 2023 roku (CTE)
	"WITH Votes2023 AS ( SELECT post_id, COUNT(id) AS vote_count FROM votes WHERE YEAR(creation_date) = 2023 GROUP BY post_id ) SELECT p.post_title, v.vote_count FROM posts p INNER JOIN Votes2023 v ON p.id = v.post_id ORDER BY v.vote_count DESC LIMIT 10;",
	//Przykład 10: Posty z największą liczbą edycji
	"SELECT p.post_title, COUNT(ph.id) AS edit_count FROM posts p INNER JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id ORDER BY edit_count DESC LIMIT 10;",
	//Przykład 11: Ostatnia edycja każdego postu (CTE)
	"WITH LatestEdit AS ( SELECT post_id, MAX(creation_date) AS last_edit_date FROM post_history GROUP BY post_id)SELECT p.post_title, le.last_edit_date FROM posts p INNER JOIN LatestEdit le ON p.id = le.post_id;",
	//Przykład 12: Użytkownicy, którzy edytowali najwięcej postów
	"SELECT u.display_name, COUNT(ph.id) AS edit_count FROM users u INNER JOIN post_history ph ON u.id = ph.user_id GROUP BY u.id ORDER BY edit_count DESC LIMIT 10;",
	//Przykład 13: Posty z największą liczbą powiązanych postów
	"SELECT p.post_title, COUNT(pl.related_post_id) AS related_count FROM posts p INNER JOIN post_links pl ON p.id = pl.post_id GROUP BY p.id ORDER BY related_count DESC LIMIT 10;",
	//Przykład 14: Posty bez powiązanych postów (podzapytanie)
	"SELECT p.post_title FROM posts p WHERE p.id NOT IN (SELECT post_id FROM post_links);",
	//Przykład 15: Najczęściej powiązane posty (CTE)
	"WITH RelatedPosts AS ( SELECT related_post_id, COUNT(*) AS link_count FROM post_links GROUP BY related_post_id ) SELECT p.post_title, rp.link_count FROM posts p INNER JOIN RelatedPosts rp ON p.id = rp.related_post_id ORDER BY rp.link_count DESC LIMIT 10;",
	//Przykład 16: Użytkownicy z reputacją > 1000 i > 100 postami
	"SELECT u.display_name, u.reputation, COUNT(p.id) AS post_count FROM users u INNER JOIN posts p ON u.id = p.owner_user_id WHERE u.reputation > 1000 GROUP BY u.id HAVING post_count > 100;",
	//Przykład 17: Użytkownicy z najwyższą reputacją (CTE)
	"WITH TopUsers AS (SELECT id, display_name, reputation FROM users ORDER BY reputation DESC LIMIT 10 ) SELECT tu.display_name, tu.reputation FROM TopUsers tu;",
	//Przykład 18: Użytkownicy z największą liczbą głosów i odznak
	"SELECT u.display_name, COUNT(v.id) AS vote_count, COUNT(b.id) AS badge_count FROM users u LEFT JOIN votes v ON u.id = v.user_id LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY vote_count DESC, badge_count DESC LIMIT 10;",
	//Przykład 19: Użytkownicy z największą liczbą postów i komentarzy
	"SELECT u.display_name, COUNT(p.id) AS post_count, COUNT(c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ORDER BY post_count DESC, comment_count DESC LIMIT 10;",
	//Przykład 20: Użytkownicy z największą liczbą aktywności (CTE)
	"WITH UserActivity AS ( SELECT u.id, u.display_name, COUNT(p.id) AS post_count,COUNT(c.id) AS comment_count, COUNT(v.id) AS vote_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id LEFT JOIN votes v ON u.id = v.user_id GROUP BY u.id )SELECT ua.display_name, ua.post_count, ua.comment_count, ua.vote_count FROM UserActivity ua ORDER BY (ua.post_count + ua.comment_count + ua.vote_count) DESC LIMIT 10;",
	//Przykład 21: Liczba postów wg miesięcy (CTE)
	"WITH MonthlyPosts AS (SELECT YEAR(creation_date) AS year, MONTH(creation_date) AS month, COUNT(*) AS post_count FROM posts GROUP BY YEAR(creation_date), MONTH(creation_date)) SELECT mp.year, mp.month, mp.post_count FROM MonthlyPosts mp ORDER BY mp.year DESC, mp.month DESC;",
	//Przykład 22: Liczba komentarzy wg tygodni
	"SELECT YEAR(creation_date) AS year, WEEK(creation_date) AS week, COUNT(*) AS comment_count FROM comments GROUP BY YEAR(creation_date), WEEK(creation_date) ORDER BY year DESC, week DESC;",
	//Przykład 23: Liczba głosów wg dni
	"SELECT DATE(creation_date) AS day, COUNT(*) AS vote_count FROM votes GROUP BY DATE(creation_date) ORDER BY day DESC;",
	//Przykład 24: Najpopularniejsze tagi w 2023 roku
	"SELECT tags, COUNT(*) AS post_count FROM posts WHERE YEAR(creation_date) = 2023 GROUP BY tags ORDER BY post_count DESC LIMIT 10;",
	//Przykład 25: Tagi z największą liczbą odpowiedzi
	"SELECT tags, SUM(answer_count) AS total_answers FROM posts GROUP BY tags ORDER BY total_answers DESC LIMIT 10;",
	//Przykład 26: Tagi z największą liczbą wyświetleń
	"SELECT tags, SUM(view_count) AS total_views FROM posts GROUP BY tags ORDER BY total_views DESC LIMIT 10;",
	//Przykład 27: Liczba postów wg licencji treści
	"SELECT content_license, COUNT(*) AS post_count FROM posts GROUP BY content_license ORDER BY post_count DESC;",
	//Przykład 28: Liczba komentarzy wg licencji treści
	"SELECT content_license, COUNT(*) AS comment_count FROM comments GROUP BY content_license ORDER BY comment_count DESC;",
	//Przykład 29: Liczba edycji wg licencji treści
	"SELECT content_license, COUNT(*) AS edit_count FROM post_history GROUP BY content_license ORDER BY edit_count DESC;",
	//Przykład 30: Sprawdzenie osieroconych komentarzy
	"SELECT c.* FROM comments c LEFT JOIN posts p ON c.post_id = p.id WHERE p.id IS NULL;",
	//Przykład 31: Sprawdzenie osieroconych głosów
	"SELECT v.* FROM votes v LEFT JOIN posts p ON v.post_id = p.id WHERE p.id IS NULL;",
	//Przykład 32: Sprawdzenie osieroconych odznak
	"SELECT b.* FROM badges b LEFT JOIN users u ON b.user_id = u.id WHERE u.id IS NULL;",
	//Przykład 33: Drzewo dyskusji (rekurencyjne CTE)
	"WITH RECURSIVE PostTree AS (SELECT id, parent_id, post_title FROM posts WHERE parent_id IS NULL UNION ALL SELECT p.id, p.parent_id, p.post_title FROM posts p INNER JOIN PostTree pt ON p.parent_id = pt.id) SELECT * FROM PostTree;",
	//Przykład 34: Detekcja anomalii w głosach
	"SELECT post_id, COUNT(*) AS vote_count FROM votes GROUP BY post_id HAVING COUNT(*) > (SELECT AVG(vote_count) + 3*STDDEV(vote_count) FROM (...) );",
	//Przykład 35: Użytkownicy z największą liczbą aktywności w ciągu ostatnich 30 dni
	"SELECT u.display_name, COUNT(p.id) AS post_count, COUNT(c.id) AS comment_count, COUNT(v.id) AS vote_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id AND p.creation_date >= NOW() - INTERVAL 30 DAY LEFT JOIN comments c ON u.id = c.user_id AND c.creation_date >= NOW() - INTERVAL 30 DAY LEFT JOIN votes v ON u.id = v.user_id AND v.creation_date >= NOW() - INTERVAL 30 DAY GROUP BY u.id ORDER BY (post_count + comment_count + vote_count) DESC LIMIT 10;",
	//Przykład 36: Liczba postów wg miesięcy (CTE)
	"WITH MonthlyPosts AS ( SELECT YEAR(creation_date) AS year, MONTH(creation_date) AS month, COUNT(*) AS post_count FROM posts GROUP BY YEAR(creation_date), MONTH(creation_date) ) SELECT mp.year, mp.month, mp.post_count FROM MonthlyPosts mp ORDER BY mp.year DESC, mp.month DESC;",
	//Przykład 37: Liczba komentarzy wg tygodni
	"SELECT YEAR(creation_date) AS year, WEEK(creation_date) AS week, COUNT(*) AS comment_count FROM comments GROUP BY YEAR(creation_date), WEEK(creation_date) ORDER BY year DESC, week DESC;",
	//Przykład 38: Liczba głosów wg dni
	"SELECT DATE(creation_date) AS day, COUNT(*) AS vote_count FROM votes GROUP BY DATE(creation_date) ORDER BY day DESC;",
	//Przykład 39: Liczba odznak wg miesięcy
	"SELECT YEAR(badge_date) AS year, MONTH(badge_date) AS month, COUNT(*) AS badge_count FROM badges GROUP BY YEAR(badge_date), MONTH(badge_date) ORDER BY year DESC, month DESC;",
	//Przykład 40: Liczba edycji wg tygodni
	"SELECT YEAR(creation_date) AS year, WEEK(creation_date) AS week, COUNT(*) AS edit_count FROM post_history GROUP BY YEAR(creation_date), WEEK(creation_date) ORDER BY year DESC, week DESC;",
	//Przykład 41: Najpopularniejsze tagi w 2023 roku
	"SELECT tags, COUNT(*) AS post_count FROM posts WHERE YEAR(creation_date) = 2023 GROUP BY tags ORDER BY post_count DESC LIMIT 10;",
	//Przykład 42: Tagi z największą liczbą odpowiedzi
	"SELECT tags, SUM(answer_count) AS total_answers FROM posts GROUP BY tags ORDER BY total_answers DESC LIMIT 10;",
	//Przykład 43: Tagi z największą liczbą wyświetleń
	"SELECT tags, SUM(view_count) AS total_views FROM posts GROUP BY tags ORDER BY total_views DESC LIMIT 10;",
	//Przykład 44: Tagi z największą liczbą głosów
	"SELECT tags, SUM(score) AS total_score FROM posts GROUP BY tags ORDER BY total_score DESC LIMIT 10;",
	//Przykład 45: Tagi z największą liczbą komentarzy
	"SELECT tags, SUM(comment_count) AS total_comments FROM posts GROUP BY tags ORDER BY total_comments DESC LIMIT 10;",
	//Przykład 46: Liczba postów wg licencji treści
	"SELECT content_license, COUNT(*) AS post_count FROM posts GROUP BY content_license ORDER BY post_count DESC;",
	//Przykład 47: Liczba komentarzy wg licencji treści
	"SELECT content_license, COUNT(*) AS comment_count FROM comments GROUP BY content_license ORDER BY comment_count DESC;",
	//Przykład 48: Liczba edycji wg licencji treści
	"SELECT content_license, COUNT(*) AS edit_count FROM post_history GROUP BY content_license ORDER BY edit_count DESC;",
	//Przykład 49: Liczba głosów wg licencji treści
	"SELECT content_license, COUNT(*) AS vote_count FROM votes GROUP BY content_license ORDER BY vote_count DESC;",
	//Przykład 50: Liczba odznak wg licencji treści
	"SELECT content_license, COUNT(*) AS badge_count",
}
