package main

var simpleSelect = []string{
	"SELECT * FROM posts WHERE id = 100;",
	"SELECT * FROM posts WHERE id = 3000;",
	"SELECT * FROM posts WHERE id = 34525;",
	"SELECT * FROM posts WHERE id = 8422345;",
	"SELECT * FROM posts WHERE id = 3536;",
	"SELECT * FROM posts WHERE id = 82356;",
	"SELECT * FROM posts WHERE id = 63568;",
	"SELECT * FROM posts WHERE id = 24515625;",
	"SELECT * FROM posts WHERE id = 2341658;",
	"SELECT * FROM posts WHERE id =4554534;",
	"SELECT * FROM posts WHERE id =324566;",
	"SELECT * FROM posts WHERE id =951231;",
	"SELECT * FROM posts WHERE id =45;",
	//idselcectPostLimit
	"SELECT * FROM posts WHERE id = 100 limit 1;",
	"SELECT * FROM posts WHERE id = 3000 limit 1;",
	"SELECT * FROM posts WHERE id = 34525 limit 1;",
	"SELECT * FROM posts WHERE id = 8422345 limit 1;",
	"SELECT * FROM posts WHERE id = 3536 limit 1;",
	"SELECT * FROM posts WHERE id = 82356 limit 1;",
	"SELECT * FROM posts WHERE id = 63568 limit 1;",
	"SELECT * FROM posts WHERE id = 24515625 limit 1;",
	"SELECT * FROM posts WHERE id = 2341658 limit 1;",
	"SELECT * FROM posts WHERE id =4554534 limit 1;",
	"SELECT * FROM posts WHERE id =324566 limit 1;",
	"SELECT * FROM posts WHERE id =951231 limit 1;",
	"SELECT * FROM posts WHERE id =45 limit 1;",
	//postTypeselcectPost
	"SELECT * FROM posts WHERE post_type_id = 15;",
	"SELECT * FROM posts WHERE post_type_id = 3;",
	"SELECT * FROM posts WHERE post_type_id = 8;",
	"SELECT * FROM posts WHERE post_type_id = 7;",
	"SELECT * FROM posts WHERE post_type_id = 20;",
	"SELECT * FROM posts WHERE post_type_id = 1;",
	//postTypeselcectPostLimit
	"SELECT * FROM posts WHERE post_type_id = 15 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 3 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 8 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 7 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 20 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 1 limit 10;",
	//accepdedansweridcectPost
	"SELECT * FROM posts WHERE accepted_answer_id IS NOT NULL;",
	"SELECT * FROM posts WHERE accepted_answer_id IS NULL;",
	//accepdedansweridcectPostLimit
	"SELECT * FROM posts WHERE accepted_answer_id IS NOT NULL limit 10;",
	"SELECT * FROM posts WHERE accepted_answer_id IS NULL limit 10;",
	//parentIDselectPost
	"SELECT * FROM posts WHERE parent_id = 39 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 50 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 23425 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 2463 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 53 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 83 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 324234 limit 2;",
	//creationDateSelectPost
	"SELECT * FROM posts WHERE creation_date = '2023-01-01';",
	"SELECT * FROM posts WHERE creation_date = '2022-05-15';",
	"SELECT * FROM posts WHERE creation_date = '2021-12-31';",
	"SELECT * FROM posts WHERE creation_date = '2023-07-20';",
	"SELECT * FROM posts WHERE creation_date = '2020-03-10';",
	"SELECT * FROM posts WHERE creation_date = '2023-09-01';",
	"SELECT * FROM posts WHERE creation_date = '2022-11-25';",
	"SELECT * FROM posts WHERE creation_date = '2021-08-14';",
	"SELECT * FROM posts WHERE creation_date = '2023-04-05';",
	"SELECT * FROM posts WHERE creation_date = '2020-10-30';",
	//scoreSelectPost
	"SELECT * FROM posts WHERE score = 0;",
	"SELECT * FROM posts WHERE score = 10;",
	"SELECT * FROM posts WHERE score = 50;",
	"SELECT * FROM posts WHERE score = 100;",
	"SELECT * FROM posts WHERE score = 200;",
	"SELECT * FROM posts WHERE score = 500;",
	"SELECT * FROM posts WHERE score = 1000;",
	"SELECT * FROM posts WHERE score = 1500;",
	"SELECT * FROM posts WHERE score = 2000;",
	"SELECT * FROM posts WHERE score = 2500;",
	//viewCountSelectPost
	"SELECT * FROM posts WHERE view_count = 100;",
	"SELECT * FROM posts WHERE view_count = 500;",
	"SELECT * FROM posts WHERE view_count = 1000;",
	"SELECT * FROM posts WHERE view_count = 5000;",
	"SELECT * FROM posts WHERE view_count = 10000;",
	"SELECT * FROM posts WHERE view_count = 20000;",
	"SELECT * FROM posts WHERE view_count = 50000;",
	"SELECT * FROM posts WHERE view_count = 100000;",
	"SELECT * FROM posts WHERE view_count = 200000;",
	"SELECT * FROM posts WHERE view_count = 500000;",
	//postBodySelectPost
	"SELECT * FROM posts WHERE post_body LIKE '%error%';",
	"SELECT * FROM posts WHERE post_body LIKE '%bug%';",
	"SELECT * FROM posts WHERE post_body LIKE '%fix%';",
	"SELECT * FROM posts WHERE post_body LIKE '%solution%';",
	"SELECT * FROM posts WHERE post_body LIKE '%help%';",
	"SELECT * FROM posts WHERE post_body LIKE '%question%';",
	"SELECT * FROM posts WHERE post_body LIKE '%example%';",
	"SELECT * FROM posts WHERE post_body LIKE '%code%';",
	"SELECT * FROM posts WHERE post_body LIKE '%test%';",
	"SELECT * FROM posts WHERE post_body LIKE '%issue%';",
	//ownerUserIDSelectPost
	"SELECT * FROM posts WHERE owner_user_id = 100;",
	"SELECT * FROM posts WHERE owner_user_id = 500;",
	"SELECT * FROM posts WHERE owner_user_id = 1000;",
	"SELECT * FROM posts WHERE owner_user_id = 1500;",
	"SELECT * FROM posts WHERE owner_user_id = 2000;",
	"SELECT * FROM posts WHERE owner_user_id = 2500;",
	"SELECT * FROM posts WHERE owner_user_id = 3000;",
	"SELECT * FROM posts WHERE owner_user_id = 3500;",
	"SELECT * FROM posts WHERE owner_user_id = 4000;",
	"SELECT * FROM posts WHERE owner_user_id = 4500;",
	//lastEditorUserIDSelectPost
	"SELECT * FROM posts WHERE last_editor_user_id = 100;",
	"SELECT * FROM posts WHERE last_editor_user_id = 500;",
	"SELECT * FROM posts WHERE last_editor_user_id = 1000;",
	"SELECT * FROM posts WHERE last_editor_user_id = 1500;",
	"SELECT * FROM posts WHERE last_editor_user_id = 2000;",
	"SELECT * FROM posts WHERE last_editor_user_id = 2500;",
	"SELECT * FROM posts WHERE last_editor_user_id = 3000;",
	"SELECT * FROM posts WHERE last_editor_user_id = 3500;",
	"SELECT * FROM posts WHERE last_editor_user_id = 4000;",
	"SELECT * FROM posts WHERE last_editor_user_id = 4500;",
	//lastEditDateSelectPost
	"SELECT * FROM posts WHERE last_edit_date = '2023-01-01';",
	"SELECT * FROM posts WHERE last_edit_date = '2022-05-15';",
	"SELECT * FROM posts WHERE last_edit_date = '2021-12-31';",
	"SELECT * FROM posts WHERE last_edit_date = '2023-07-20';",
	"SELECT * FROM posts WHERE last_edit_date = '2020-03-10';",
	"SELECT * FROM posts WHERE last_edit_date = '2023-09-01';",
	"SELECT * FROM posts WHERE last_edit_date = '2022-11-25';",
	"SELECT * FROM posts WHERE last_edit_date = '2021-08-14';",
	"SELECT * FROM posts WHERE last_edit_date = '2023-04-05';",
	"SELECT * FROM posts WHERE last_edit_date = '2020-10-30';",
	//lastActivityDateSelectPost
	"SELECT * FROM posts WHERE last_activity_date = '2023-01-01';",
	"SELECT * FROM posts WHERE last_activity_date = '2022-05-15';",
	"SELECT * FROM posts WHERE last_activity_date = '2021-12-31';",
	"SELECT * FROM posts WHERE last_activity_date = '2023-07-20';",
	"SELECT * FROM posts WHERE last_activity_date = '2020-03-10';",
	"SELECT * FROM posts WHERE last_activity_date = '2023-09-01';",
	"SELECT * FROM posts WHERE last_activity_date = '2022-11-25';",
	"SELECT * FROM posts WHERE last_activity_date = '2021-08-14';",
	"SELECT * FROM posts WHERE last_activity_date = '2023-04-05';",
	"SELECT * FROM posts WHERE last_activity_date = '2020-10-30';",
	//postTitleSelectPost
	"SELECT * FROM posts WHERE post_title LIKE '%error%';",
	"SELECT * FROM posts WHERE post_title LIKE '%bug%';",
	"SELECT * FROM posts WHERE post_title LIKE '%fix%';",
	"SELECT * FROM posts WHERE post_title LIKE '%solution%';",
	"SELECT * FROM posts WHERE post_title LIKE '%help%';",
	"SELECT * FROM posts WHERE post_title LIKE '%question%';",
	"SELECT * FROM posts WHERE post_title LIKE '%example%';",
	"SELECT * FROM posts WHERE post_title LIKE '%code%';",
	"SELECT * FROM posts WHERE post_title LIKE '%test%';",
	"SELECT * FROM posts WHERE post_title LIKE '%issue%';",
	//tagsSelectPost
	"SELECT * FROM posts WHERE tags LIKE '%java%';",
	"SELECT * FROM posts WHERE tags LIKE '%python%';",
	"SELECT * FROM posts WHERE tags LIKE '%javascript%';",
	"SELECT * FROM posts WHERE tags LIKE '%sql%';",
	"SELECT * FROM posts WHERE tags LIKE '%c%23%';", // C#
	"SELECT * FROM posts WHERE tags LIKE '%php%';",
	"SELECT * FROM posts WHERE tags LIKE '%html%';",
	"SELECT * FROM posts WHERE tags LIKE '%css%';",
	"SELECT * FROM posts WHERE tags LIKE '%react%';",
	"SELECT * FROM posts WHERE tags LIKE '%nodejs%';",
	//answerCountSelectPost
	"SELECT * FROM posts WHERE answer_count = 0;",
	"SELECT * FROM posts WHERE answer_count = 1;",
	"SELECT * FROM posts WHERE answer_count = 5;",
	"SELECT * FROM posts WHERE answer_count = 10;",
	"SELECT * FROM posts WHERE answer_count = 20;",
	"SELECT * FROM posts WHERE answer_count = 50;",
	"SELECT * FROM posts WHERE answer_count = 100;",
	"SELECT * FROM posts WHERE answer_count = 200;",
	"SELECT * FROM posts WHERE answer_count = 500;",
	"SELECT * FROM posts WHERE answer_count = 1000;",
	//commentCountSelectPost
	"SELECT * FROM posts WHERE comment_count = 0;",
	"SELECT * FROM posts WHERE comment_count = 1;",
	"SELECT * FROM posts WHERE comment_count = 5;",
	"SELECT * FROM posts WHERE comment_count = 10;",
	"SELECT * FROM posts WHERE comment_count = 20;",
	"SELECT * FROM posts WHERE comment_count = 50;",
	"SELECT * FROM posts WHERE comment_count = 100;",
	"SELECT * FROM posts WHERE comment_count = 200;",
	"SELECT * FROM posts WHERE comment_count = 500;",
	"SELECT * FROM posts WHERE comment_count = 1000;",
	//contentLicenseSelectPost
	"SELECT * FROM posts WHERE content_license = 'CC BY-SA 4.0';",
	"SELECT * FROM posts WHERE content_license = 'MIT';",
	"SELECT * FROM posts WHERE content_license = 'Apache 2.0';",
	"SELECT * FROM posts WHERE content_license = 'GPLv3';",
	"SELECT * FROM posts WHERE content_license = 'BSD';",
	"SELECT * FROM posts WHERE content_license = 'CC BY 3.0';",
	"SELECT * FROM posts WHERE content_license = 'CC BY 2.0';",
	"SELECT * FROM posts WHERE content_license = 'CC0';",
	"SELECT * FROM posts WHERE content_license = 'AGPL';",
	"SELECT * FROM posts WHERE content_license = 'LGPL';",
	//multiConditionSelectPost
	"SELECT * FROM posts WHERE post_type_id = 1 AND score > 100;",
	"SELECT * FROM posts WHERE tags LIKE '%java%' AND answer_count > 5;",
	"SELECT * FROM posts WHERE creation_date > '2023-01-01' AND view_count < 1000;",
	"SELECT * FROM posts WHERE owner_user_id = 100 AND comment_count > 10;",
	"SELECT * FROM posts WHERE post_body LIKE '%error%' AND score > 50;",
	//groupBySelectPost
	"SELECT post_type_id, COUNT(*) FROM posts GROUP BY post_type_id;",
	"SELECT owner_user_id, COUNT(*) FROM posts GROUP BY owner_user_id;",
	"SELECT tags, COUNT(*) FROM posts GROUP BY tags;",
	"SELECT content_license, COUNT(*) FROM posts GROUP BY content_license;",
	"SELECT YEAR(creation_date), COUNT(*) FROM posts GROUP BY YEAR(creation_date);",
	//orderBySelectPost
	"SELECT * FROM posts ORDER BY creation_date DESC LIMIT 20;",
	"SELECT * FROM posts ORDER BY score DESC LIMIT 10;",
	"SELECT * FROM posts ORDER BY view_count DESC LIMIT 15;",
	"SELECT * FROM posts ORDER BY last_activity_date DESC LIMIT 30;",
	"SELECT * FROM posts ORDER BY answer_count DESC LIMIT 25;",
	//betweenSelectPost
	"SELECT * FROM posts WHERE creation_date BETWEEN '2022-01-01' AND '2022-12-31';",
	"SELECT * FROM posts WHERE score BETWEEN 50 AND 100;",
	"SELECT * FROM posts WHERE view_count BETWEEN 1000 AND 5000;",
	"SELECT * FROM posts WHERE last_edit_date BETWEEN '2023-01-01' AND '2023-12-31';",
	"SELECT * FROM posts WHERE answer_count BETWEEN 5 AND 20;",
	//fullScanSelectPost
	"SELECT * FROM posts LIMIT 100;",
	"SELECT * FROM posts LIMIT 500;",
	"SELECT * FROM posts LIMIT 1000;",
	"SELECT * FROM posts LIMIT 5000;",
	"SELECT * FROM posts LIMIT 10000;",
	//countSelectPost
	"SELECT COUNT(*) FROM posts;",
	"SELECT COUNT(*) FROM posts WHERE post_type_id = 1;",
	"SELECT COUNT(*) FROM posts WHERE score > 100;",
	"SELECT COUNT(*) FROM posts WHERE tags LIKE '%java%';",
	"SELECT COUNT(*) FROM posts WHERE creation_date > '2023-01-01';",
	//analizeUserActivity
	"SELECT owner_user_id, COUNT(*) AS post_count FROM posts WHERE creation_date > '2023-01-01' GROUP BY owner_user_id HAVING post_count > 5;",
	"SELECT owner_user_id, AVG(score) AS avg_score FROM posts GROUP BY owner_user_id ORDER BY avg_score DESC LIMIT 10;",
	"SELECT owner_user_id, SUM(view_count) AS total_views FROM posts GROUP BY owner_user_id ORDER BY total_views DESC LIMIT 10;",
	"SELECT owner_user_id, MAX(score) AS max_score FROM posts GROUP BY owner_user_id ORDER BY max_score DESC LIMIT 10;",
	"SELECT owner_user_id, COUNT(*) AS post_count FROM posts WHERE post_type_id = 1 GROUP BY owner_user_id HAVING post_count > 10;",
	//analizePostTags
	"SELECT tags, COUNT(*) AS post_count FROM posts WHERE creation_date > '2022-01-01' GROUP BY tags ORDER BY post_count DESC LIMIT 10;",
	"SELECT tags, AVG(score) AS avg_score FROM posts GROUP BY tags HAVING COUNT(*) > 50 ORDER BY avg_score DESC LIMIT 10;",
	"SELECT tags, SUM(answer_count) AS total_answers FROM posts GROUP BY tags ORDER BY total_answers DESC LIMIT 10;",
	"SELECT tags, SUM(comment_count) AS total_comments FROM posts GROUP BY tags ORDER BY total_comments DESC LIMIT 10;",
	"SELECT tags, AVG(view_count) AS avg_views FROM posts GROUP BY tags HAVING COUNT(*) > 100 ORDER BY avg_views DESC LIMIT 10;",
	//analizePostDate
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS post_count FROM posts GROUP BY YEAR(creation_date) ORDER BY year DESC;",
	"SELECT MONTH(creation_date) AS month, COUNT(*) AS post_count FROM posts WHERE YEAR(creation_date) = 2023 GROUP BY MONTH(creation_date) ORDER BY month;",
	"SELECT DATE(creation_date) AS date, COUNT(*) AS post_count FROM posts GROUP BY DATE(creation_date) ORDER BY date DESC LIMIT 10;",
	"SELECT YEAR(creation_date) AS year, AVG(score) AS avg_score FROM posts GROUP BY YEAR(creation_date) ORDER BY year DESC;",
	"SELECT WEEK(creation_date) AS week, COUNT(*) AS post_count FROM posts WHERE YEAR(creation_date) = 2023 GROUP BY WEEK(creation_date) ORDER BY week;",
	//viewMarkAnalize
	"SELECT score, COUNT(*) AS post_count FROM posts GROUP BY score ORDER BY score DESC LIMIT 10;",
	"SELECT view_count, COUNT(*) AS post_count FROM posts GROUP BY view_count ORDER BY view_count DESC LIMIT 10;",
	"SELECT score, AVG(view_count) AS avg_views FROM posts GROUP BY score ORDER BY score DESC LIMIT 10;",
	"SELECT view_count, AVG(score) AS avg_score FROM posts GROUP BY view_count ORDER BY view_count DESC LIMIT 10;",
	"SELECT score, SUM(answer_count) AS total_answers FROM posts GROUP BY score ORDER BY score DESC LIMIT 10;",
	//answersComsAnalize
	"SELECT answer_count, COUNT(*) AS post_count FROM posts GROUP BY answer_count ORDER BY answer_count DESC LIMIT 10;",
	"SELECT comment_count, COUNT(*) AS post_count FROM posts GROUP BY comment_count ORDER BY comment_count DESC LIMIT 10;",
	"SELECT answer_count, AVG(score) AS avg_score FROM posts GROUP BY answer_count ORDER BY answer_count DESC LIMIT 10;",
	"SELECT comment_count, AVG(view_count) AS avg_views FROM posts GROUP BY comment_count ORDER BY comment_count DESC LIMIT 10;",
	"SELECT answer_count, SUM(comment_count) AS total_comments FROM posts GROUP BY answer_count ORDER BY answer_count DESC LIMIT 10;",
	//liceneAnalize
	"SELECT content_license, COUNT(*) AS post_count FROM posts GROUP BY content_license ORDER BY post_count DESC;",
	"SELECT content_license, AVG(score) AS avg_score FROM posts GROUP BY content_license ORDER BY avg_score DESC;",
	"SELECT content_license, SUM(view_count) AS total_views FROM posts GROUP BY content_license ORDER BY total_views DESC;",
	"SELECT content_license, MAX(answer_count) AS max_answers FROM posts GROUP BY content_license ORDER BY max_answers DESC;",
	"SELECT content_license, AVG(comment_count) AS avg_comments FROM posts GROUP BY content_license ORDER BY avg_comments DESC;",
	//conditionsAnalize
	"SELECT * FROM posts WHERE tags LIKE '%python%' AND post_type_id = 1 AND creation_date > '2023-01-01' ORDER BY score DESC LIMIT 20;",
	"SELECT * FROM posts WHERE score > 100 AND view_count > 1000 AND answer_count > 5 ORDER BY creation_date DESC LIMIT 10;",
	"SELECT * FROM posts WHERE post_body LIKE '%error%' AND tags LIKE '%java%' AND comment_count > 10 ORDER BY view_count DESC LIMIT 15;",
	"SELECT * FROM posts WHERE creation_date BETWEEN '2022-01-01' AND '2022-12-31' AND score > 50 AND answer_count > 2 ORDER BY view_count DESC LIMIT 20;",
	"SELECT * FROM posts WHERE owner_user_id = 100 AND post_type_id = 1 AND creation_date > '2023-01-01' ORDER BY score DESC LIMIT 10;",
	//agregatePostAnalize
	"SELECT post_type_id, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(view_count) AS total_views FROM posts GROUP BY post_type_id ORDER BY post_count DESC;",
	"SELECT tags, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(answer_count) AS total_answers FROM posts GROUP BY tags ORDER BY post_count DESC LIMIT 10;",
	"SELECT owner_user_id, COUNT(*) AS post_count, AVG(score) AS avg_score, MAX(view_count) AS max_views FROM posts GROUP BY owner_user_id ORDER BY post_count DESC LIMIT 10;",
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(comment_count) AS total_comments FROM posts GROUP BY YEAR(creation_date) ORDER BY year DESC;",
	"SELECT content_license, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(answer_count) AS total_answers FROM posts GROUP BY content_license ORDER BY post_count DESC;",
	//subQueryAnalize
	"SELECT * FROM posts WHERE view_count > (SELECT AVG(view_count) FROM posts WHERE post_type_id = 1);",
	"SELECT * FROM posts WHERE answer_count > (SELECT AVG(answer_count) FROM posts WHERE tags LIKE '%java%');",
	"SELECT * FROM posts WHERE comment_count > (SELECT AVG(comment_count) FROM posts WHERE creation_date > '2023-01-01');",
	"SELECT * FROM posts WHERE owner_user_id IN (SELECT owner_user_id FROM posts GROUP BY owner_user_id HAVING COUNT(*) > 50);",
	//joinAnalize
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
	//optAnalize
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
	//indexAnalize
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
	//funWinTrendAnalize
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
	//agregateAnalize
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
	//advanceAnalize
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
