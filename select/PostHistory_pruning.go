package main

var PostHistory = []string{
	// Pary poniżej pochodzą kolejno z partycji p0-p7 tabeli post_history.

	// Q01. Pełny klucz z partycji p0 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=1894 AND ph.post_id=1679;",

	// Q02. Pełny klucz z partycji p1 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=1418 AND ph.post_id=1240;",

	// Q03. Pełny klucz z partycji p2 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=2990 AND ph.post_id=2605;",

	// Q04. Pełny klucz z partycji p3 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=2671 AND ph.post_id=2332;",

	// Q05. Pełny klucz z partycji p4 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=2620 AND ph.post_id=2298;",

	// Q06. Pełny klucz z partycji p5 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=1427 AND ph.post_id=1265;",

	// Q07. Pełny klucz z partycji p6 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=2497 AND ph.post_id=2183;",

	// Q08. Pełny klucz z partycji p7 — oczekiwana jedna partycja.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=1661 AND ph.post_id=1474;",

	// Q09. COUNT(*) dla jednego pełnego klucza — jedna partycja.
	"SELECT COUNT(*) AS history_count FROM post_history AS ph WHERE ph.id=1894 AND ph.post_id=1679;",

	// Q10. COUNT(*) dla dwóch pełnych kluczy z dwóch partycji.
	"SELECT COUNT(*) AS history_count FROM post_history AS ph WHERE (ph.id=1894 AND ph.post_id=1679) OR (ph.id=1418 AND ph.post_id=1240);",

	// Q11. COUNT(*) dla czterech pełnych kluczy z czterech partycji.
	"SELECT COUNT(*) AS history_count FROM post_history AS ph WHERE (ph.id=1894 AND ph.post_id=1679) OR (ph.id=1418 AND ph.post_id=1240) OR (ph.id=2990 AND ph.post_id=2605) OR (ph.id=2671 AND ph.post_id=2332);",

	// Q12. COUNT(*) dla pełnych kluczy ze wszystkich ośmiu partycji.
	"SELECT COUNT(*) AS history_count FROM post_history AS ph WHERE (ph.id=1894 AND ph.post_id=1679) OR (ph.id=1418 AND ph.post_id=1240) OR (ph.id=2990 AND ph.post_id=2605) OR (ph.id=2671 AND ph.post_id=2332) OR (ph.id=2620 AND ph.post_id=2298) OR (ph.id=1427 AND ph.post_id=1265) OR (ph.id=2497 AND ph.post_id=2183) OR (ph.id=1661 AND ph.post_id=1474);",

	// Q13. Pobranie i sortowanie rekordu z jednej partycji.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE ph.id=1894 AND ph.post_id=1679 ORDER BY ph.creation_date DESC;",

	// Q14. Pobranie i sortowanie rekordów z dwóch partycji.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE (ph.id=1894 AND ph.post_id=1679) OR (ph.id=1418 AND ph.post_id=1240) ORDER BY ph.creation_date DESC;",

	// Q15. Pobranie i sortowanie rekordów z czterech partycji.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE (ph.id=1894 AND ph.post_id=1679) OR (ph.id=1418 AND ph.post_id=1240) OR (ph.id=2990 AND ph.post_id=2605) OR (ph.id=2671 AND ph.post_id=2332) ORDER BY ph.creation_date DESC;",

	// Q16. Pobranie i sortowanie rekordów ze wszystkich ośmiu partycji.
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date,ph.user_id FROM post_history AS ph WHERE (ph.id=1894 AND ph.post_id=1679) OR (ph.id=1418 AND ph.post_id=1240) OR (ph.id=2990 AND ph.post_id=2605) OR (ph.id=2671 AND ph.post_id=2332) OR (ph.id=2620 AND ph.post_id=2298) OR (ph.id=1427 AND ph.post_id=1265) OR (ph.id=2497 AND ph.post_id=2183) OR (ph.id=1661 AND ph.post_id=1474) ORDER BY ph.creation_date DESC;",

	// Q17. Przypadek kontrolny: tylko id, bez kompletnego klucza KEY(id,post_id).
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date FROM post_history AS ph WHERE ph.id=1894;",

	// Q18. Przypadek kontrolny: tylko post_id, bez kompletnego klucza KEY(id,post_id).
	"SELECT ph.id,ph.post_id,ph.post_history_type_id,ph.creation_date FROM post_history AS ph WHERE ph.post_id=1679 ORDER BY ph.creation_date DESC;",

	// Q19. Przypadek kontrolny: filtr po kolumnach spoza klucza partycjonowania.
	"SELECT COUNT(*) AS history_count,COUNT(DISTINCT ph.post_id) AS edited_posts FROM post_history AS ph WHERE ph.creation_date>='2009-04-01' AND ph.creation_date<'2009-07-01' AND ph.post_history_type_id IN (2,5,8);",

	// Q20. Przypadek kontrolny bez predykatu po kluczu partycjonowania.
	"SELECT COUNT(*) AS history_count,COUNT(DISTINCT ph.post_id) AS edited_posts,COUNT(DISTINCT ph.user_id) AS active_editors FROM post_history AS ph;",
}
