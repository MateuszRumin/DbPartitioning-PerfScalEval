package main

var joinAnalize = []string{
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
	"SELECT p1.post_title, p2.post_title AS duplicate_post FROM post_links pl INNER JOIN posts p1 ON pl.post_id = p1.id INNER JOIN posts p2 ON pl.related_post_id = p2.id WHERE pl.link_type_id = 3; ", // przykładowy typ linku (duplikat)
	//Przykład 13: Użytkownicy z liczbą postów i komentarzy
	"SELECT u.display_name, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id;",
	//Przykład 14: Użytkownicy z postami, które mają >10 komentarzy
	"SELECT u.display_name, p.post_title, COUNT(c.id) AS comment_count FROM users u INNER JOIN posts p ON u.id = p.owner_user_id INNER JOIN comments c ON p.id = c.post_id GROUP BY u.id, p.id HAVING comment_count > 10;",
	//Przykład 15: Użytkownicy z postami i odznakami
	"SELECT u.display_name, p.post_title, b.badge_name FROM users u INNER JOIN posts p ON u.id = p.owner_user_id LEFT JOIN badges b ON u.id = b.user_id;",
	//Przykład 16: Posty z liczbą głosów
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id;",
	//Przykład 17: Posty z głosami typu "upvote"
	"SELECT p.post_title, COUNT(v.id) AS upvote_count FROM posts p INNER JOIN votes v ON p.id = v.post_id WHERE v.vote_type_id = 2  GROUP BY p.id;", //-- przykładowy typ głosu (upvote)
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
	//subAndCteAnalize
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

var usertRaports = []string{
	//1. Liczba postów na użytkownika
	"SELECT u.display_name, COUNT(p.id) AS post_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.id ORDER BY post_count DESC;",
	//2. Aktywni użytkownicy w 2023 roku
	"SELECT u.display_name, COUNT(p.id) AS post_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id WHERE YEAR(p.creation_date) = 2023 GROUP BY u.id HAVING post_count > 10;",
	//3. Użytkownicy z najwyższą reputacją
	"SELECT display_name, reputation FROM users ORDER BY reputation DESC LIMIT 10;",
	//4. Użytkownicy z największą liczbą odznak
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY badge_count DESC LIMIT 10;",
	//5. Użytkownicy z największą liczbą komentarzy
	"SELECT u.display_name, COUNT(c.id) AS comment_count FROM users u LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ORDER BY comment_count DESC LIMIT 10;",
	//6. Użytkownicy z największą liczbą głosów
	"SELECT u.display_name, COUNT(v.id) AS vote_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN votes v ON p.id = v.post_id GROUP BY u.id ORDER BY vote_count DESC LIMIT 10;",
	//7. Użytkownicy z największą liczbą wyświetleń postów
	"SELECT u.display_name, SUM(p.view_count) AS total_views FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.id ORDER BY total_views DESC LIMIT 10;",
	//8. Użytkownicy z największą liczbą edycji postów
	"SELECT u.display_name, COUNT(ph.id) AS edit_count FROM users u LEFT JOIN post_history ph ON u.id = ph.user_id GROUP BY u.id ORDER BY edit_count DESC LIMIT 10;",
	//9. Użytkownicy z największą liczbą odpowiedzi na posty
	"SELECT u.display_name, SUM(p.answer_count) AS total_answers FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.id ORDER BY total_answers DESC LIMIT 10;",
	//10. Użytkownicy z największą liczbą komentarzy do swoich postów
	"SELECT u.display_name, SUM(p.comment_count) AS total_comments FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.id ORDER BY total_comments DESC LIMIT 10;",
}

var postRaport = []string{
	//1. Najpopularniejsze posty (wg wyświetleń)
	"SELECT post_title, view_count FROM posts ORDER BY view_count DESC LIMIT 10;",
	//2. Posty z największą liczbą głosów
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id ORDER BY vote_count DESC LIMIT 10;",
	//3. Posty z największą liczbą komentarzy
	"SELECT p.post_title, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY comment_count DESC LIMIT 10;",
	//4. Posty z największą liczbą odpowiedzi
	"SELECT post_title, answer_count FROM posts ORDER BY answer_count DESC LIMIT 10;",
	//5. Posty z największą liczbą edycji
	"SELECT p.post_title, COUNT(ph.id) AS edit_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id ORDER BY edit_count DESC LIMIT 10;",
	//6. Posty z największą liczbą linków do innych postów
	"SELECT p.post_title, COUNT(pl.id) AS link_count FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.id ORDER BY link_count DESC LIMIT 10;",
	//7. Posty z największą liczbą odznak
	"SELECT p.post_title, COUNT(b.id) AS badge_count FROM posts p LEFT JOIN badges b ON p.owner_user_id = b.user_id GROUP BY p.id ORDER BY badge_count DESC LIMIT 10;",
	//8. Posty z największą liczbą wyświetleń w 2023 roku
	"SELECT post_title, view_count FROM posts WHERE YEAR(creation_date) = 2023 ORDER BY view_count DESC LIMIT 10;",
	//9. Posty z największą liczbą głosów w 2023 roku
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id WHERE YEAR(p.creation_date) = 2023 GROUP BY p.id ORDER BY vote_count DESC LIMIT 10;",
	//10. Posty z największą liczbą komentarzy w 2023 roku
	"SELECT p.post_title, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id WHERE YEAR(p.creation_date) = 2023 GROUP BY p.id ORDER BY comment_count DESC LIMIT 10;",
}

var commentRaport = []string{
	//1. Najczęściej komentowane posty
	"SELECT p.post_title, COUNT(c.id) AS comment_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY comment_count DESC LIMIT 10;",
	//2. Użytkownicy z największą liczbą komentarzy
	"SELECT u.display_name, COUNT(c.id) AS comment_count FROM users u LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ORDER BY comment_count DESC LIMIT 10;",
	//3. Komentarze z największą liczbą głosów
	"SELECT c.comment_text, c.score FROM comments c ORDER BY c.score DESC LIMIT 10;",
	//4. Średnia liczba komentarzy na post
	"SELECT AVG(comment_count) AS avg_comments_per_post FROM posts;",
	//5. Komentarze z największą liczbą słów
	"SELECT c.comment_text, LENGTH(c.comment_text) - LENGTH(REPLACE(c.comment_text, ' ', '')) + 1 AS word_count FROM comments c ORDER BY word_count DESC LIMIT 10;",
}

var badgeRaport = []string{
	//1. Najczęściej przyznawane odznaki
	"SELECT badge_name, COUNT(*) AS badge_count FROM badges GROUP BY badge_name ORDER BY badge_count DESC LIMIT 10;",
	//2. Użytkownicy z największą liczbą odznak
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id ORDER BY badge_count DESC LIMIT 10;",
	//3. Odznaki przyznane w 2023 roku
	"SELECT badge_name, COUNT(*) AS badge_count FROM badges WHERE YEAR(badge_date) = 2023 GROUP BY badge_name ORDER BY badge_count DESC LIMIT 10;",
	//4. Odznaki przyznane za tagi
	"SELECT badge_name, COUNT(*) AS badge_count FROM badges WHERE tag_based = 'true' GROUP BY badge_name ORDER BY badge_count DESC LIMIT 10;",
	//5. Odznaki przyznane użytkownikom z wysoką reputacją
	"SELECT b.badge_name, COUNT(*) AS badge_count FROM badges b LEFT JOIN users u ON b.user_id = u.id WHERE u.reputation > 1000 GROUP BY b.badge_name ORDER BY badge_count DESC LIMIT 10;",
}

var voteRaport = []string{
	//1. Posty z największą liczbą głosów
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id ORDER BY vote_count DESC LIMIT 10;",
	//2. Użytkownicy z największą liczbą głosów
	"SELECT u.display_name, COUNT(v.id) AS vote_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN votes v ON p.id = v.post_id GROUP BY u.id ORDER BY vote_count DESC LIMIT 10;",
	//3. Głosy w 2023 roku
	"SELECT COUNT(*) AS vote_count FROM votes WHERE YEAR(creation_date) = 2023;",
	//4. Średnia liczba głosów na post
	"SELECT AVG(vote_count) AS avg_votes_per_post FROM ( SELECT COUNT(*) AS vote_count FROM votes GROUP BY post_id) AS vote_counts;",
	//5. Głosy wg typu
	"SELECT vote_type_id, COUNT(*) AS vote_count FROM votes GROUP BY vote_type_id ORDER BY vote_count DESC;",
}

var postHistoryRaport = []string{
	//1. Liczba edycji na post
	"SELECT p.post_title, COUNT(ph.id) AS edit_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id ORDER BY edit_count DESC LIMIT 10;",
	//2. Użytkownicy z największą liczbą edycji
	"SELECT u.display_name, COUNT(ph.id) AS edit_count FROM users u LEFT JOIN post_history ph ON u.id = ph.user_id GROUP BY u.id ORDER BY edit_count DESC LIMIT 10;",
	//3. Etykiety najczęściej edytowanych postów
	"SELECT p.tags, COUNT(ph.id) AS edit_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.tags ORDER BY edit_count DESC LIMIT 10;",
	//4. Etykiety najczęściej edytowanych postów w 2023 roku
	"SELECT p.tags, COUNT(ph.id) AS edit_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id WHERE YEAR(ph.creation_date) = 2023 GROUP BY p.tags ORDER BY edit_count DESC LIMIT 10;",
	//5. Średnia liczba edycji na post
	"SELECT AVG(edit_count) AS avg_edits_per_post FROM ( SELECT COUNT(*) AS edit_count FROM post_history GROUP BY post_id) AS edit_counts;",
}

var performacneAnalize = []string{
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
