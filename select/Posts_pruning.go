package main

var Posts = []string{
	// Q01. COUNT(*) dla jednego kwartału — oczekiwany pruning do jednej partycji.
	"SELECT COUNT(*) AS post_count FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-04-01';",

	// Q02. COUNT(*) dla dwóch kolejnych kwartałów — oczekiwane dwie partycje.
	"SELECT COUNT(*) AS post_count FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-07-01';",

	// Q03. COUNT(*) dla trzech kolejnych kwartałów — oczekiwane trzy partycje.
	"SELECT COUNT(*) AS post_count FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-10-01';",

	// Q04. COUNT(*) dla całego roku — oczekiwane cztery partycje kwartalne.
	"SELECT COUNT(*) AS post_count FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2010-01-01';",

	// Q05. Agregaty dla pytań z jednego kwartału — jedna partycja.
	"SELECT COUNT(*) AS question_count, SUM(p.score) AS total_score, AVG(p.view_count) AS avg_views, SUM(p.answer_count) AS total_answers FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-04-01' AND p.post_type_id = 1;",

	// Q06. Te same agregaty dla dwóch kwartałów — dwie partycje.
	"SELECT COUNT(*) AS question_count, SUM(p.score) AS total_score, AVG(p.view_count) AS avg_views, SUM(p.answer_count) AS total_answers FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-07-01' AND p.post_type_id = 1;",

	// Q07. Te same agregaty dla trzech kwartałów — trzy partycje.
	"SELECT COUNT(*) AS question_count, SUM(p.score) AS total_score, AVG(p.view_count) AS avg_views, SUM(p.answer_count) AS total_answers FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-10-01' AND p.post_type_id = 1;",

	// Q08. Te same agregaty dla całego roku — cztery partycje kwartalne.
	"SELECT COUNT(*) AS question_count, SUM(p.score) AS total_score, AVG(p.view_count) AS avg_views, SUM(p.answer_count) AS total_answers FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2010-01-01' AND p.post_type_id = 1;",

	// Q09. Sortowanie i LIMIT dla jednego kwartału — jedna partycja.
	"SELECT p.id, p.post_title, p.post_type_id, p.score, p.view_count, p.answer_count, p.creation_date FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-04-01' ORDER BY p.score DESC, p.view_count DESC LIMIT 100;",

	// Q10. To samo sortowanie dla dwóch kwartałów — dwie partycje.
	"SELECT p.id, p.post_title, p.post_type_id, p.score, p.view_count, p.answer_count, p.creation_date FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-07-01' ORDER BY p.score DESC, p.view_count DESC LIMIT 100;",

	// Q11. To samo sortowanie dla trzech kwartałów — trzy partycje.
	"SELECT p.id, p.post_title, p.post_type_id, p.score, p.view_count, p.answer_count, p.creation_date FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-10-01' ORDER BY p.score DESC, p.view_count DESC LIMIT 100;",

	// Q12. To samo sortowanie dla całego roku — cztery partycje kwartalne.
	"SELECT p.id, p.post_title, p.post_type_id, p.score, p.view_count, p.answer_count, p.creation_date FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2010-01-01' ORDER BY p.score DESC, p.view_count DESC LIMIT 100;",

	// Q13. Grupowanie według typu postu dla jednego kwartału — jedna partycja.
	"SELECT p.post_type_id, COUNT(*) AS post_count, SUM(p.score) AS total_score, AVG(p.score) AS avg_score FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-04-01' GROUP BY p.post_type_id ORDER BY p.post_type_id;",

	// Q14. To samo grupowanie dla dwóch kwartałów — dwie partycje.
	"SELECT p.post_type_id, COUNT(*) AS post_count, SUM(p.score) AS total_score, AVG(p.score) AS avg_score FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-07-01' GROUP BY p.post_type_id ORDER BY p.post_type_id;",

	// Q15. To samo grupowanie dla trzech kwartałów — trzy partycje.
	"SELECT p.post_type_id, COUNT(*) AS post_count, SUM(p.score) AS total_score, AVG(p.score) AS avg_score FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2009-10-01' GROUP BY p.post_type_id ORDER BY p.post_type_id;",

	// Q16. To samo grupowanie dla całego roku — cztery partycje kwartalne.
	"SELECT p.post_type_id, COUNT(*) AS post_count, SUM(p.score) AS total_score, AVG(p.score) AS avg_score FROM posts AS p WHERE p.creation_date >= '2009-01-01' AND p.creation_date < '2010-01-01' GROUP BY p.post_type_id ORDER BY p.post_type_id;",

	// Q17. Wąski zakres znajdujący się całkowicie w jednej partycji.
	"SELECT COUNT(*) AS post_count, MIN(p.creation_date) AS first_post, MAX(p.creation_date) AS last_post FROM posts AS p WHERE p.creation_date >= '2009-05-01' AND p.creation_date < '2009-06-01';",

	// Q18. Zakres przecinający granicę dwóch partycji kwartalnych.
	"SELECT COUNT(*) AS post_count, MIN(p.creation_date) AS first_post, MAX(p.creation_date) AS last_post FROM posts AS p WHERE p.creation_date >= '2009-06-15' AND p.creation_date < '2009-07-15';",

	// Q19. Zakres obejmujący cały badany zbiór — pruning nie eliminuje partycji,
	// ale predykat nadal jest zapisany bezpośrednio po kluczu partycjonowania.
	"SELECT COUNT(*) AS post_count, SUM(p.score) AS total_score FROM posts AS p WHERE p.creation_date >= '2008-01-01' AND p.creation_date < '2013-01-01';",

	// Q20. Przypadek kontrolny bez warunku po creation_date — brak skutecznego pruning posts.
	"SELECT COUNT(*) AS post_count, SUM(p.score) AS total_score FROM posts AS p WHERE p.post_type_id = 1 AND p.score >= 10;",
}
