package selectqueries

var UsertRaports = []string{
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

var PostAnalize = []string{
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

var CommentRaport = []string{
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

var BadgeRaport = []string{
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

var VoteRaport = []string{
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

var PostHistoryRaport = []string{
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
