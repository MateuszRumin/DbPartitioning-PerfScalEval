package main

var Posts = []string{
	// Q01 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2008-07-01' AND creation_date < '2008-10-01';",
	// Q02 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2008-10-01' AND creation_date < '2009-01-01';",
	// Q03 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-01-01' AND creation_date < '2009-04-01';",
	// Q04 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-04-01' AND creation_date < '2009-07-01';",
	// Q05 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-07-01' AND creation_date < '2009-10-01';",
	// Q06 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-10-01' AND creation_date < '2010-01-01';",
	// Q07 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2010-01-01' AND creation_date < '2010-04-01';",
	// Q08 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2010-04-01' AND creation_date < '2010-07-01';",
	// Q09 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2010-07-01' AND creation_date < '2010-10-01';",
	// Q10 — 1 partycja
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2010-10-01' AND creation_date < '2011-01-01';",
	// Q11 — 2 partycje
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-01-01' AND creation_date < '2009-07-01';",
	// Q12 — 2 partycje
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-07-01' AND creation_date < '2010-01-01';",
	// Q13 — 2 partycje
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2010-01-01' AND creation_date < '2010-07-01';",
	// Q14 — 2 partycje
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2010-07-01' AND creation_date < '2011-01-01';",
	// Q15 — 3 partycje
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-01-01' AND creation_date < '2009-10-01';",
	// Q16 — 4 partycje
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-01-01' AND creation_date < '2010-01-01';",
	// Q17 — 4 partycje
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2010-01-01' AND creation_date < '2011-01-01';",
	// Q18 — 6 partycji
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-01-01' AND creation_date < '2010-07-01';",
	// Q19 — 8 partycji
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2009-01-01' AND creation_date < '2011-01-01';",
	// Q20 — 10 partycji
	"SELECT id, post_type_id, creation_date, score, owner_user_id FROM posts WHERE creation_date >= '2008-07-01' AND creation_date < '2011-01-01';",
}

var comments []string = []string{
	//Q01: Pełny klucz partycjonowania i join do posts oraz users
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.comment_text,u.id AS commenter_id,u.display_name AS commenter_name,p.post_title,p.score AS post_score FROM comments c JOIN posts p ON p.id=c.post_id LEFT JOIN users u ON u.id=c.user_id WHERE c.id=9  AND c.post_id=47239;",
	//Q02: Tylko id, częściowe dopasowanie klucza KEY
	"SELECT c.id,c.post_id,c.score,c.creation_date,u.display_name,p.post_title FROM comments c JOIN posts p ON p.id=c.post_id LEFT JOIN users u ON u.id=c.user_id WHERE c.id=9;",
	//Q03: Tylko post_id, typowe zapytanie aplikacyjne bez pełnego klucza KEY
	"SELECT c.id,c.post_id,c.score,c.creation_date,u.display_name,p.post_title FROM comments c JOIN posts p ON p.id=c.post_id LEFT JOIN users u ON u.id=c.user_id WHERE c.post_id=456789 ORDER BY c.creation_date;",
	//Q04: Cztery pełne pary klucza przez alternatywy OR
	"SELECT c.id,c.post_id,c.score,c.creation_date,u.display_name,p.post_title FROM comments c JOIN posts p ON p.id=c.post_id LEFT JOIN users u ON u.id=c.user_id WHERE (c.id=9 AND c.post_id=47239) OR (c.id=12 AND c.post_id=47428) OR (c.id=25 AND c.post_id=47626) OR (c.id=1 AND c.post_id=35314) ORDER BY c.id;",
	//Q05: Zakres id i join do posts, przypadek kontrolny
	"SELECT c.id,c.post_id,c.score,c.creation_date,p.post_title,u.display_name FROM comments c JOIN posts p ON p.id=c.post_id LEFT JOIN users u ON u.id=c.user_id WHERE c.id>=10000000 AND c.id<10010000 ORDER BY c.id LIMIT 500;",
	//Q06: Filtrowanie po user_id bez warunku po kluczu partycjonowania
	"SELECT u.id,u.display_name,COUNT(c.id) AS comment_count,COUNT(DISTINCT c.post_id) AS commented_posts,COALESCE(SUM(c.score),0) AS total_comment_score FROM comments c JOIN users u ON u.id=c.user_id WHERE c.user_id=1 GROUP BY u.id,u.display_name;",
	//Q07: Komentarze użytkownika z informacjami o postach
	"SELECT c.id,c.post_id,c.score,c.creation_date,p.post_title,p.post_type_id,p.score AS post_score,u.display_name FROM comments c JOIN posts p ON p.id=c.post_id JOIN users u ON u.id=c.user_id WHERE c.user_id=1 ORDER BY c.creation_date DESC LIMIT 100;",
	//Q08: Filtrowanie comments po dacie, brak pruning po KEY(id,post_id)
	"SELECT DATE(c.creation_date) AS day,COUNT(*) AS comment_count,COUNT(DISTINCT c.user_id) AS active_users,COUNT(DISTINCT c.post_id) AS commented_posts,AVG(c.score) AS avg_score FROM comments c WHERE c.creation_date>='2009-04-01' AND c.creation_date<'2009-07-01' GROUP BY DATE(c.creation_date) ORDER BY day;",
	//Q09: Komentarze do pytań z jednej partycji czasowej posts
	"SELECT p.id,p.post_title,p.score AS post_score,COUNT(c.id) AS comment_count,COUNT(DISTINCT c.user_id) AS commenter_count,AVG(c.score) AS avg_comment_score FROM posts p JOIN comments c ON c.post_id=p.id WHERE p.creation_date>='2009-04-01' AND p.creation_date<'2009-07-01' AND p.post_type_id=1 GROUP BY p.id,p.post_title,p.score HAVING COUNT(c.id)>=5 ORDER BY comment_count DESC LIMIT 100;",
	//Q10: Ranking komentujących pytania z wybranego kwartału
	"SELECT u.id,u.display_name,u.reputation,COUNT(c.id) AS comment_count,COUNT(DISTINCT c.post_id) AS commented_posts,COALESCE(SUM(c.score),0) AS total_comment_score FROM posts p JOIN comments c ON c.post_id=p.id JOIN users u ON u.id=c.user_id WHERE p.creation_date>='2009-04-01' AND p.creation_date<'2009-07-01' AND p.post_type_id=1 GROUP BY u.id,u.display_name,u.reputation HAVING COUNT(c.id)>=3 ORDER BY comment_count DESC,total_comment_score DESC LIMIT 100;",
	//Q11: Autorzy postów i komentujący te posty
	"SELECT owner.id AS owner_id,owner.display_name AS owner_name,commenter.id AS commenter_id,commenter.display_name AS commenter_name,COUNT(c.id) AS comment_count,COUNT(DISTINCT p.id) AS post_count FROM posts p JOIN comments c ON c.post_id=p.id LEFT JOIN users owner ON owner.id=p.owner_user_id LEFT JOIN users commenter ON commenter.id=c.user_id WHERE p.creation_date>='2009-04-01' AND p.creation_date<'2009-07-01' GROUP BY owner.id,owner.display_name,commenter.id,commenter.display_name HAVING COUNT(c.id)>=3 ORDER BY comment_count DESC LIMIT 100;",
	//Q12: Komentarze autora pod własnymi postami
	"SELECT u.id,u.display_name,COUNT(c.id) AS self_comment_count,COUNT(DISTINCT p.id) AS self_commented_posts FROM comments c JOIN posts p ON p.id=c.post_id JOIN users u ON u.id=c.user_id AND u.id=p.owner_user_id WHERE p.creation_date>='2009-01-01' AND p.creation_date<'2010-01-01' GROUP BY u.id,u.display_name ORDER BY self_comment_count DESC LIMIT 100;",
	//Q13: Najnowszy komentarz dla każdego postu z wybranego okresu
	"WITH ranked_comments AS (SELECT c.id,c.post_id,c.user_id,c.score,c.creation_date,ROW_NUMBER() OVER (PARTITION BY c.post_id ORDER BY c.creation_date DESC,c.id DESC) AS rn FROM comments c JOIN posts p ON p.id=c.post_id WHERE p.creation_date>='2009-04-01' AND p.creation_date<'2009-07-01') SELECT rc.id,rc.post_id,p.post_title,rc.score,rc.creation_date,u.display_name FROM ranked_comments rc JOIN posts p ON p.id=rc.post_id LEFT JOIN users u ON u.id=rc.user_id WHERE rc.rn=1 ORDER BY rc.creation_date DESC LIMIT 100;",
	//Q14: Ranking postów według liczby komentarzy z funkcją okienkową
	"WITH comment_stats AS (SELECT p.id,p.post_title,p.owner_user_id,COUNT(c.id) AS comment_count,COUNT(DISTINCT c.user_id) AS commenter_count,AVG(c.score) AS avg_comment_score FROM posts p JOIN comments c ON c.post_id=p.id WHERE p.creation_date>='2009-04-01' AND p.creation_date<'2009-07-01' GROUP BY p.id,p.post_title,p.owner_user_id) SELECT cs.id,cs.post_title,u.display_name,cs.comment_count,cs.commenter_count,cs.avg_comment_score,DENSE_RANK() OVER (ORDER BY cs.comment_count DESC) AS comment_rank FROM comment_stats cs LEFT JOIN users u ON u.id=cs.owner_user_id ORDER BY comment_rank LIMIT 100;",
	//Q15: Komentarze do pytań i zaakceptowanych odpowiedziv
	"SELECT q.id AS question_id,q.post_title,a.id AS accepted_answer_id,COUNT(DISTINCT qc.id) AS question_comment_count,COUNT(DISTINCT ac.id) AS answer_comment_count,qu.display_name AS question_author,au.display_name AS answer_author FROM posts q JOIN posts a ON a.id=q.accepted_answer_id LEFT JOIN comments qc ON qc.post_id=q.id LEFT JOIN comments ac ON ac.post_id=a.id LEFT JOIN users qu ON qu.id=q.owner_user_id LEFT JOIN users au ON au.id=a.owner_user_id WHERE q.creation_date>='2009-04-01' AND q.creation_date<'2009-07-01' AND q.post_type_id=1 GROUP BY q.id,q.post_title,a.id,qu.display_name,au.display_name ORDER BY question_comment_count+answer_comment_count DESC LIMIT 100;",
	//Q16: Wybrane pełne klucze w CTE i join do posts oraz users
	"WITH selected_comments AS (SELECT id,post_id,user_id,score,creation_date FROM comments WHERE (id=9  AND post_id=47239) OR (id=12 AND post_id=47428) OR (id=25 AND post_id=47626) OR (id=1 AND post_id=35314)) SELECT sc.id,sc.post_id,sc.score,sc.creation_date,p.post_title,u.display_name FROM selected_comments sc JOIN posts p ON p.id=sc.post_id LEFT JOIN users u ON u.id=sc.user_id ORDER BY sc.id;",
	//Q17: Statystyki komentarzy oraz historii dla wybranych postów
	"WITH selected_posts AS (SELECT DISTINCT c.post_id FROM comments c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=12 AND c.post_id=47428) OR (c.id=25 AND c.post_id=47626) OR (c.id=1 AND c.post_id=35314)),comment_stats AS (SELECT c.post_id,COUNT(*) AS comment_count,COUNT(DISTINCT c.user_id) AS commenter_count FROM comments c JOIN selected_posts sp ON sp.post_id=c.post_id GROUP BY c.post_id),history_stats AS (SELECT ph.post_id,COUNT(*) AS revision_count,COUNT(DISTINCT ph.user_id) AS editor_count FROM post_history ph JOIN selected_posts sp ON sp.post_id=ph.post_id GROUP BY ph.post_id) SELECT p.id,p.post_title,cs.comment_count,cs.commenter_count,hs.revision_count,hs.editor_count FROM selected_posts sp JOIN posts p ON p.id=sp.post_id LEFT JOIN comment_stats cs ON cs.post_id=sp.post_id LEFT JOIN history_stats hs ON hs.post_id=sp.post_id;",
	//Q18: Statystyki komentarzy oraz głosów dla wybranych postów
	"WITH selected_posts AS (SELECT DISTINCT c.post_id FROM comments c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=12 AND c.post_id=47428) OR (c.id=25 AND c.post_id=47626) OR (c.id=1 AND c.post_id=35314)),comment_stats AS (SELECT c.post_id,COUNT(*) AS comment_count,SUM(c.score) AS comment_score FROM comments c JOIN selected_posts sp ON sp.post_id=c.post_id GROUP BY c.post_id),vote_stats AS (SELECT v.post_id,COUNT(*) AS vote_count,COUNT(DISTINCT v.vote_type_id) AS vote_type_count FROM votes v JOIN selected_posts sp ON sp.post_id=v.post_id GROUP BY v.post_id) SELECT p.id,p.post_title,cs.comment_count,cs.comment_score,vs.vote_count,vs.vote_type_count FROM selected_posts sp JOIN posts p ON p.id=sp.post_id LEFT JOIN comment_stats cs ON cs.post_id=sp.post_id LEFT JOIN vote_stats vs ON vs.post_id=sp.post_id;",
	//Q19: Miesięczne statystyki komentarzy według typu postu
	"SELECT YEAR(c.creation_date) AS year_number,MONTH(c.creation_date) AS month_number,p.post_type_id,COUNT(c.id) AS comment_count,COUNT(DISTINCT c.user_id) AS active_commenters,COUNT(DISTINCT c.post_id) AS commented_posts,AVG(c.score) AS avg_comment_score FROM comments c JOIN posts p ON p.id=c.post_id WHERE c.creation_date>='2009-01-01' AND c.creation_date<'2010-01-01' GROUP BY YEAR(c.creation_date),MONTH(c.creation_date),p.post_type_id ORDER BY year_number,month_number,p.post_type_id;",
	//Q20: Pełny przypadek kontrolny bez predykatu po id ani post_id
	"SELECT u.id,u.display_name,COUNT(c.id) AS comment_count,COUNT(DISTINCT c.post_id) AS commented_posts,COALESCE(SUM(c.score),0) AS total_comment_score,AVG(c.score) AS avg_comment_score FROM comments c JOIN users u ON u.id=c.user_id GROUP BY u.id,u.display_name HAVING COUNT(c.id)>=10 ORDER BY comment_count DESC LIMIT 100;",
}
