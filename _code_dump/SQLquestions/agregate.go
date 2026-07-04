package selectqueries

var agregateAnalize = []string{
	// 1. Liczba odznak na użytkownika
	"SELECT user_id, COUNT(*) AS badge_count FROM badges GROUP BY user_id HAVING badge_count > 5;",
	//2. Najczęściej przyznawane odznaki
	"SELECT badge_name, COUNT(*) AS badge_count FROM badges GROUP BY badge_name ORDER BY badge_count DESC LIMIT 10;",
	//3. Średnia liczba odznak na użytkownika
	"SELECT AVG(badge_count) AS avg_badges_per_user FROM (SELECT user_id, COUNT(*) AS badge_count FROM badges GROUP BY user_id) AS user_badges;",
	//4. Liczba odznak wg klasy
	"SELECT class, COUNT(*) AS badge_count FROM badges GROUP BY class ORDER BY badge_count DESC;",
	//5. Użytkownicy z największą liczbą odznak w 2023 roku
	"SELECT user_id, COUNT(*) AS badge_count FROM badges WHERE YEAR(badge_date) = 2023 GROUP BY user_id ORDER BY badge_count DESC LIMIT 10;",
	//6. Liczba komentarzy na post
	"SELECT post_id, COUNT(*) AS comment_count FROM comments GROUP BY post_id HAVING comment_count > 10;",
	//7. Średnia liczba komentarzy na użytkownika
	"SELECT user_id, AVG(comment_count) AS avg_comments FROM (SELECT user_id, COUNT(*) AS comment_count FROM comments GROUP BY user_id) AS user_comments;",
	//8. Najczęściej komentowane posty
	"SELECT post_id, COUNT(*) AS comment_count FROM comments GROUP BY post_id ORDER BY comment_count DESC LIMIT 10;",
	//9. Liczba komentarzy wg roku
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS comment_count FROM comments GROUP BY YEAR(creation_date);",
	//10. Użytkownicy z największą liczbą komentarzy
	"SELECT user_id, COUNT(*) AS comment_count FROM comments GROUP BY user_id ORDER BY comment_count DESC LIMIT 10;",
	//11. Liczba edycji na post
	"SELECT post_id, COUNT(*) AS edit_count FROM post_history GROUP BY post_id HAVING edit_count > 5;",
	//12. Najczęściej edytowane posty
	"SELECT post_id, COUNT(*) AS edit_count FROM post_history GROUP BY post_id ORDER BY edit_count DESC LIMIT 10;",
	//13. Liczba edycji wg typu
	"SELECT post_history_type_id, COUNT(*) AS edit_count FROM post_history GROUP BY post_history_type_id;",
	//14. Użytkownicy z największą liczbą edycji
	"SELECT user_id, COUNT(*) AS edit_count FROM post_history GROUP BY user_id ORDER BY edit_count DESC LIMIT 10;",
	//15. Liczba edycji wg roku
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS edit_count FROM post_history GROUP BY YEAR(creation_date);",
	//16. Liczba linków na post
	"SELECT post_id, COUNT(*) AS link_count FROM post_links GROUP BY post_id HAVING link_count > 5;",
	//17. Najczęściej linkowane posty
	"SELECT post_id, COUNT(*) AS link_count FROM post_links GROUP BY post_id ORDER BY link_count DESC LIMIT 10;",
	//18. Liczba linków wg typu
	"SELECT link_type_id, COUNT(*) AS link_count FROM post_links GROUP BY link_type_id;",
	//19. Posty z największą liczbą linków w 2023 roku
	"SELECT post_id, COUNT(*) AS link_count FROM post_links WHERE YEAR(creation_date) = 2023 GROUP BY post_id ORDER BY link_count DESC LIMIT 10;",
	//20. Liczba linków wg roku
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS link_count FROM post_links GROUP BY YEAR(creation_date);",
	//21. Liczba postów na użytkownika
	"SELECT owner_user_id, COUNT(*) AS post_count FROM posts GROUP BY owner_user_id HAVING post_count > 10;",
	//22. Najczęściej używane tagi
	"SELECT tags, COUNT(*) AS post_count FROM posts GROUP BY tags ORDER BY post_count DESC LIMIT 10;",
	//23. Średnia liczba odpowiedzi na post
	"SELECT AVG(answer_count) AS avg_answers FROM posts;",
	//24. Posty z największą liczbą wyświetleń
	"SELECT id, post_title, view_count FROM posts ORDER BY view_count DESC LIMIT 10;",
	//25. Liczba postów wg typu
	"SELECT post_type_id, COUNT(*) AS post_count FROM posts GROUP BY post_type_id;",
	//26. Użytkownicy z największą reputacją
	"SELECT id, display_name, reputation FROM users ORDER BY reputation DESC LIMIT 10;",
	//27. Średnia reputacja użytkowników
	"SELECT AVG(reputation) AS avg_reputation FROM users;",
	//28. Liczba użytkowników wg lokalizacji
	"SELECT location, COUNT(*) AS user_count FROM users GROUP BY location ORDER BY user_count DESC LIMIT 10;",
	//29. Użytkownicy z największą liczbą wyświetleń profilu
	"SELECT id, display_name, views FROM users ORDER BY views DESC LIMIT 10;",
	//30. Liczba użytkowników wg roku rejestracji
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS user_count FROM users GROUP BY YEAR(creation_date);",
	//31. Liczba głosów na post
	"SELECT post_id, COUNT(*) AS vote_count FROM votes GROUP BY post_id HAVING vote_count > 10;",
	//32. Najczęściej głosowane posty
	"SELECT post_id, COUNT(*) AS vote_count FROM votes GROUP BY post_id ORDER BY vote_count DESC LIMIT 10;",
	//33. Liczba głosów wg typu
	"SELECT vote_type_id, COUNT(*) AS vote_count FROM votes GROUP BY vote_type_id;",
	//34. Użytkownicy z największą liczbą głosów
	"SELECT user_id, COUNT(*) AS vote_count FROM votes GROUP BY user_id ORDER BY vote_count DESC LIMIT 10;",
	//35. Liczba głosów wg roku
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS vote_count FROM votes GROUP BY YEAR(creation_date);",
	//36. Użytkownicy z największą liczbą postów i komentarzy
	"SELECT u.id, u.display_name, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count FROM users u LEFT JOIN posts p ON u.id = p.owner_user_id LEFT JOIN comments c ON u.id = c.user_id GROUP BY u.id ORDER BY post_count DESC, comment_count DESC LIMIT 10;",
	//37. Posty z największą liczbą głosów i komentarzy
	"SELECT p.id, p.post_title, COUNT(DISTINCT v.id) AS vote_count, COUNT(DISTINCT c.id) AS comment_count FROM posts p LEFT JOIN votes v ON p.id = v.post_id LEFT JOIN comments c ON p.id = c.post_id GROUP BY p.id ORDER BY vote_count DESC, comment_count DESC LIMIT 10;",
	//38. Użytkownicy z największą liczbą odznak i postów
	"SELECT u.id, u.display_name, COUNT(DISTINCT b.id) AS badge_count,  COUNT(DISTINCT p.id) AS post_count FROM users u LEFT JOIN badges b ON u.id = b.user_id LEFT JOIN posts p ON u.id = p.owner_user_id GROUP BY u.id ORDER BY badge_count DESC, post_count DESC LIMIT 10;",
	//39. Średnia liczba głosów na post wg typu
	"SELECT p.post_type_id, AVG(vote_count) AS avg_votes FROM posts p LEFT JOIN (SELECT post_id, COUNT(*) AS vote_count FROM votes GROUP BY post_id) AS v ON p.id = v.post_id GROUP BY p.post_type_id;",
	//40. Liczba postów i komentarzy wg roku
	"SELECT YEAR(p.creation_date) AS year, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count FROM posts p LEFT JOIN comments c ON p.id = c.post_id GROUP BY YEAR(p.creation_date);",
	//41. Ranking użytkowników wg liczby postów
	"SELECT id, display_name, COUNT(*) OVER (PARTITION BY owner_user_id) AS post_count, RANK() OVER (ORDER BY COUNT(*) DESC) AS rank FROM posts GROUP BY owner_user_id;",
	//42. Średnia krocząca głosów na post
	"SELECT post_id, creation_date, AVG(vote_count) OVER (ORDER BY creation_date ROWS BETWEEN 7 PRECEDING AND CURRENT ROW) AS avg_votes_7d FROM votes;",
	//43. Skumulowana liczba komentarzy wg czasu
	"SELECT creation_date, SUM(comment_count) OVER (ORDER BY creation_date) AS cumulative_comments FROM (SELECT DATE(creation_date) AS creation_date, COUNT(*) AS comment_count FROM comments GROUP BY DATE(creation_date)) AS daily_comments;",
	// 44. Ranking postów wg liczby wyświetleń
	"SELECT id, post_title, view_count, RANK() OVER (ORDER BY view_count DESC) AS rank FROM posts;",
	//45. Średnia liczba głosów na post wg miesiąca
	"SELECT MONTH(creation_date) AS month, AVG(vote_count) OVER (PARTITION BY MONTH(creation_date)) AS avg_votes  FROM votes;",
	//46. Użytkownicy z >100 postami i >1000 punktami reputacji
	"SELECT u.id, u.display_name FROM users u WHERE u.id IN (SELECT owner_user_id FROM posts GROUP BY owner_user_id HAVING COUNT(*) > 100) AND u.reputation > 1000;",
	// 47. Posty z >10 komentarzami i >50 głosami
	"SELECT p.id, p.post_title FROM posts p WHERE p.id IN (SELECT post_id FROM comments GROUP BY post_id HAVING COUNT(*) > 10) AND p.id IN (SELECT post_id FROM votes GROUP BY post_id HAVING COUNT(*) > 50);",
	//48. Użytkownicy z odznakami i >50 postami
	"SELECT u.id, u.display_name FROM users u WHERE u.id IN (SELECT user_id FROM badges GROUP BY user_id HAVING COUNT(*) > 5) AND u.id IN (SELECT owner_user_id FROM posts GROUP BY owner_user_id HAVING COUNT(*) > 50);",
	//49. Posty z największą liczbą głosów i komentarzy
	"SELECT p.id, p.post_title FROM posts pWHERE p.id IN (SELECT post_id FROM votes GROUP BY post_id ORDER BY COUNT(*) DESC LIMIT 10) AND p.id IN (SELECT post_id FROM comments GROUP BY post_id ORDER BY COUNT(*) DESC LIMIT 10);",
	//50. Użytkownicy z największą liczbą odznak i głosów
	"SELECT u.id, u.display_name FROM users u WHERE u.id IN (SELECT user_id FROM badges GROUP BY user_id ORDER BY COUNT(*) DESC LIMIT 10) AND u.id IN (SELECT user_id FROM votes GROUP BY user_id ORDER BY COUNT(*) DESC LIMIT 10 );",
}
