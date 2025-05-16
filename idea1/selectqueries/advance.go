package selectqueries

var advanceAnalize = []string{
	//Rekurencyjne CTE i hierarchie
	//1. Drzewo dyskusji (rekurencyjne CTE dla postów i komentarzy)
	"WITH RECURSIVE PostTree AS ( SELECT id, parent_id, post_title, owner_user_id FROM posts WHERE parent_id IS NULL UNION ALL SELECT p.id, p.parent_id, p.post_title, p.owner_user_id FROM posts p INNER JOIN PostTree pt ON p.parent_id = pt.id)SELECT * FROM PostTree;",
	//2. Hierarchia użytkowników na podstawie odznak (rekurencyjne CTE)
	"WITH RECURSIVE UserBadges AS (SELECT user_id, badge_name, badge_date FROM badges WHERE user_id = 1 -- Start od konkretnego użytkownika UNION ALL SELECT b.user_id, b.badge_name, b.badge_date FROM badges b INNER JOIN UserBadges ub ON b.user_id = ub.user_id) SELECT * FROM UserBadges;",
	//Analiza anomalii
	//3. Detekcja anomalii w głosach (posty z nieprawidłową liczbą głosów)
	"SELECT post_id,COUNT(*) AS vote_count,AVG(COUNT(*)) OVER () AS avg_votes, STDDEV(COUNT(*)) OVER () AS std_dev FROM votes GROUP BY post_id HAVING COUNT(*) > (SELECT AVG(vote_count) + 3*STDDEV(vote_count) FROM votes);",
	//4. Anomalie w komentarzach (komentarze z nieprawidłowym wynikiem)
	"SELECT id, post_id, score FROM comments WHERE score > (SELECT AVG(score) + 3*STDDEV(score) FROM comments);",
	//Złożone agregacje i raporty
	//5. Ranking użytkowników według liczby odznak
	"SELECT u.display_name,COUNT(b.id) AS badge_count, RANK() OVER (ORDER BY COUNT(b.id) DESC) AS rank FROM users u LEFT JOIN badges b ON u.id = b.user_id GROUP BY u.id;",
	//6. Średnia liczba głosów na post wg typu postu
	"SELECT post_type_id,AVG(vote_count) AS avg_votes FROM (SELECT p.post_type_id, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id) AS subquery GROUP BY post_type_id;",
	//7. Liczba postów i komentarzy na użytkownika
	"SELECT u.display_name,COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id;",
	//Analiza trendów czasowych
	//8. Liczba postów wg miesięcy
	"SELECT YEAR(creation_date) AS year, MONTH(creation_date) AS month, COUNT(*) AS post_count FROM posts GROUP BY YEAR(creation_date), MONTH(creation_date);",
	//9. Aktywność użytkowników wg tygodni
	"SELECT YEAR(creation_date) AS year, WEEK(creation_date) AS week, COUNT(DISTINCT owner_user_id) AS active_users FROM posts GROUP BY YEAR(creation_date), WEEK(creation_date);",
	//10. Trendy w głosach wg dni
	"SELECT DATE(creation_date) AS vote_date, COUNT(*) AS vote_count FROM votes GROUP BY DATE(creation_date);",
	//Analiza relacji między tabelami
	//11. Posty z największą liczbą powiązanych postów
	"SELECT p.post_title,COUNT(pl.related_post_id) AS related_count FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.id ORDER BY related_count DESC;",
	//12. Użytkownicy z największą liczbą edycji postów
	"SELECT u.display_name, COUNT(ph.id) AS edit_count FROM users u LEFT JOIN post_history ph ON u.id = ph.user_id GROUP BY u.id ORDER BY edit_count DESC;",
	//13. Posty z największą liczbą komentarzy i głosów
	"SELECT p.post_title, COUNT(DISTINCT c.id) AS comment_count, COUNT(DISTINCT v.id) AS vote_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id ORDER BY comment_count DESC, vote_count DESC;",
	//Analiza odznak
	//14. Najczęściej przyznawane odznaki
	"SELECT badge_name,COUNT(*) AS badge_count FROM badges GROUP BY badge_name ORDER BY badge_count DESC;",
	//15. Użytkownicy z największą liczbą odznak w 2023 roku
	"SELECT u.display_name, COUNT(b.id) AS badge_count FROM users u LEFT JOIN badges b ON u.id = b.user_id WHERE YEAR(b.badge_date) = 2023 GROUP BY u.id ORDER BY badge_count DESC;",
	//16. Odznaki oparte na tagach (tag_based)
	"SELECT badge_name,COUNT(*) AS badge_count FROM badges WHERE tag_based = 'true' GROUP BY badge_name;",
	//Analiza historii postów
	//17. Posty z największą liczbą edycji
	"SELECT p.post_title,COUNT(ph.id) AS edit_count FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id ORDER BY edit_count DESC;",
	//18. Ostatnia edycja dla każdego postu
	"SELECT p.post_title, MAX(ph.creation_date) AS last_edit FROM posts p LEFT JOIN post_history ph ON p.id = ph.post_id GROUP BY p.id;",
	//19. Użytkownicy z największą liczbą edycji
	"SELECT u.display_name, COUNT(ph.id) AS edit_count FROM users u LEFT JOIN post_history ph ON u.id = ph.user_id GROUP BY u.id ORDER BY edit_count DESC;",
	//Analiza głosów
	//20. Posty z największą liczbą głosów
	"SELECT p.post_title, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id ORDER BY vote_count DESC;",
	//21. Typy głosów wg postów
	"SELECT p.post_title, v.vote_type_id, COUNT(v.id) AS vote_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id GROUP BY p.id, v.vote_type_id;",
	//22. Użytkownicy z największą liczbą głosów
	"SELECT u.display_name, COUNT(v.id) AS vote_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN votes v ON p.id = v.post_id GROUP BY u.id ORDER BY vote_count DESC;",
	//Analiza komentarzy
	//23. Najdłuższe komentarze
	"SELECT c.comment_text, LENGTH(c.comment_text) AS comment_length FROM comments c ORDER BY comment_length DESC LIMIT 10;",
	//24. Komentarze z najwyższym wynikiem
	"SELECT c.comment_text, c.score FROM comments c ORDER BY c.score DESC LIMIT 10;",
	//25. Użytkownicy z największą liczbą komentarzy
	"SELECT u.display_name, COUNT(c.id) AS comment_count FROM users u LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ORDER BY comment_count DESC;",
	//Analiza postów
	//26. Posty z największą liczbą wyświetleń
	"SELECT post_title, view_count FROM posts ORDER BY view_count DESC LIMIT 10;",
	//27. Posty z największą liczbą odpowiedzi
	"SELECT post_title,answer_count FROM posts ORDER BY answer_count DESC LIMIT 10;",
	//28. Posty z największą liczbą komentarzy
	"SELECT post_title, comment_count FROM posts ORDER BY comment_count DESC LIMIT 10;",
	//29. Użytkownicy z największą reputacją
	"SELECT display_name, reputation FROM users ORDER BY reputation DESC LIMIT 10;",
	//30. Użytkownicy z największą liczbą wyświetleń profilu
	"SELECT display_name, views FROM users ORDER BY views DESC LIMIT 10;",
	//31. Użytkownicy z największą liczbą głosów (upvotes/downvotes)
	"SELECT display_name,upvotes,downvotes FROM users ORDER BY upvotes DESC, downvotes DESC LIMIT 10",
	//Analiza linków między postami
	//32. Posty z największą liczbą linków do innych postów
	"SELECT p.post_title,COUNT(pl.related_post_id) AS link_count FROM posts p LEFT JOIN post_links pl ON p.id = pl.post_id GROUP BY p.id ORDER BY link_count DESC;",
	//33. Najczęściej linkowane posty
	"SELECT p.post_title, COUNT(pl.post_id) AS linked_count FROM posts p LEFT JOIN post_links pl ON p.id = pl.related_post_id GROUP BY p.id ORDER BY linked_count DESC;",
	//34. Typy linków między postami
	"SELECT link_type_id, COUNT(*) AS link_count FROM post_links GROUP BY link_type_id;",
	//Analiza licencji treści
	//35. Liczba postów wg licencji treści
	"SELECT content_license,COUNT(*) AS post_count FROM posts GROUP BY content_license;",
	//36. Liczba komentarzy wg licencji treści
	"SELECT content_license, COUNT(*) AS comment_count FROM comments GROUP BY content_license;",
	//37. Liczba edycji wg licencji treści
	"SELECT content_license, COUNT(*) AS edit_count FROM post_history GROUP BY content_license;",
	//38. Użytkownicy z największą liczbą postów wg lokalizacji
	"SELECT location, COUNT(p.id) AS post_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY location ORDER BY post_count DESC;",
	//39. Użytkownicy z największą reputacją wg lokalizacji
	"SELECT location, AVG(reputation) AS avg_reputation FROM users GROUP BY location ORDER BY avg_reputation DESC;",
	//40. Użytkownicy z największą liczbą odznak wg lokalizacji
	"SELECT location,COUNT(b.id) AS badge_count FROM users u LEFT JOIN badges b ON u.id = b.user_id GROUP BY location ORDER BY badge_count DESC;",
}
