// WAŻNE:
// comments jest partycjonowane przez KEY(id,post_id), dlatego do deterministycznego
// pruningu potrzebne są pełne, rzeczywiste pary (id,post_id).
// Wartości p0-p3 poniżej należy zastąpić wynikiem Comments_partition_keys.sql.
// Pierwsza para (9,47239) pochodzi z wcześniejszego testu; pozostałe są znacznikami technicznymi.

var Comments = []string{
	// Q01. Pełny klucz z partycji p0 — oczekiwana jedna partycja.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE c.id=9 AND c.post_id=47239;",

	// Q02. Pełny klucz z partycji p1 — po podmianie oczekiwana jedna partycja.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE c.id=10000124 AND c.post_id=456790;",

	// Q03. Pełny klucz z partycji p2 — po podmianie oczekiwana jedna partycja.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE c.id=10000125 AND c.post_id=456791;",

	// Q04. Pełny klucz z partycji p3 — po podmianie oczekiwana jedna partycja.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE c.id=10000126 AND c.post_id=456792;",

	// Q05. COUNT(*) dla jednego pełnego klucza — jedna partycja.
	"SELECT COUNT(*) AS comment_count FROM comments AS c WHERE c.id=9 AND c.post_id=47239;",

	// Q06. COUNT(*) dla dwóch pełnych kluczy z dwóch partycji.
	"SELECT COUNT(*) AS comment_count FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790);",

	// Q07. COUNT(*) dla trzech pełnych kluczy z trzech partycji.
	"SELECT COUNT(*) AS comment_count FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790) OR (c.id=10000125 AND c.post_id=456791);",

	// Q08. COUNT(*) dla pełnych kluczy ze wszystkich czterech partycji.
	"SELECT COUNT(*) AS comment_count FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790) OR (c.id=10000125 AND c.post_id=456791) OR (c.id=10000126 AND c.post_id=456792);",

	// Q09. Pobranie i sortowanie rekordu z jednej partycji.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE c.id=9 AND c.post_id=47239 ORDER BY c.creation_date DESC;",

	// Q10. Pobranie i sortowanie rekordów z dwóch partycji.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790) ORDER BY c.creation_date DESC;",

	// Q11. Pobranie i sortowanie rekordów z trzech partycji.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790) OR (c.id=10000125 AND c.post_id=456791) ORDER BY c.creation_date DESC;",

	// Q12. Pobranie i sortowanie rekordów ze wszystkich czterech partycji.
	"SELECT c.id,c.post_id,c.score,c.creation_date,c.user_id FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790) OR (c.id=10000125 AND c.post_id=456791) OR (c.id=10000126 AND c.post_id=456792) ORDER BY c.creation_date DESC;",

	// Q13. Agregaty dla jednego pełnego klucza — jedna partycja.
	"SELECT COUNT(*) AS comment_count,SUM(c.score) AS total_score,MIN(c.creation_date) AS first_comment,MAX(c.creation_date) AS last_comment FROM comments AS c WHERE c.id=9 AND c.post_id=47239;",

	// Q14. Te same agregaty dla dwóch pełnych kluczy z dwóch partycji.
	"SELECT COUNT(*) AS comment_count,SUM(c.score) AS total_score,MIN(c.creation_date) AS first_comment,MAX(c.creation_date) AS last_comment FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790);",

	// Q15. Te same agregaty dla pełnych kluczy ze wszystkich czterech partycji.
	"SELECT COUNT(*) AS comment_count,SUM(c.score) AS total_score,MIN(c.creation_date) AS first_comment,MAX(c.creation_date) AS last_comment FROM comments AS c WHERE (c.id=9 AND c.post_id=47239) OR (c.id=10000124 AND c.post_id=456790) OR (c.id=10000125 AND c.post_id=456791) OR (c.id=10000126 AND c.post_id=456792);",

	// Q16. Przypadek kontrolny: tylko id, bez kompletnego klucza KEY(id,post_id).
	"SELECT c.id,c.post_id,c.score,c.creation_date FROM comments AS c WHERE c.id=9;",

	// Q17. Przypadek kontrolny: tylko post_id, bez kompletnego klucza KEY(id,post_id).
	"SELECT c.id,c.post_id,c.score,c.creation_date FROM comments AS c WHERE c.post_id=47239 ORDER BY c.creation_date;",

	// Q18. Przypadek kontrolny: filtr po dacie, która nie jest kluczem partycjonowania.
	"SELECT COUNT(*) AS comment_count,COUNT(DISTINCT c.post_id) AS commented_posts FROM comments AS c WHERE c.creation_date>='2009-04-01' AND c.creation_date<'2009-07-01';",

	// Q19. Przypadek kontrolny: filtr po user_id i score, bez pełnego klucza partycjonowania.
	"SELECT COUNT(*) AS comment_count,SUM(c.score) AS total_score FROM comments AS c WHERE c.user_id=1 AND c.score>=0;",

	// Q20. Przypadek kontrolny bez predykatu po kluczu partycjonowania.
	"SELECT COUNT(*) AS comment_count,COUNT(DISTINCT c.post_id) AS commented_posts,COUNT(DISTINCT c.user_id) AS active_users FROM comments AS c;",
}
