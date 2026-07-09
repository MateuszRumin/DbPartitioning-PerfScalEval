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
