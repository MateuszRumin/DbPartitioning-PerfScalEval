package main

var Users = []string{
	// Q01. COUNT(*) dla identyfikatorów należących wyłącznie do jednej partycji HASH (id % 4 = 0).
	"SELECT COUNT(*) AS user_count FROM users AS u WHERE u.id IN (232136,232756,234152,238100);",

	// Q02. COUNT(*) dla identyfikatorów z dwóch partycji HASH (reszty 0 i 1).
	"SELECT COUNT(*) AS user_count FROM users AS u WHERE u.id IN (232136,232756,100001,650001);",

	// Q03. COUNT(*) dla identyfikatorów z trzech partycji HASH (reszty 0, 1 i 2).
	"SELECT COUNT(*) AS user_count FROM users AS u WHERE u.id IN (232136,100001,650001,1500002,650002);",

	// Q04. COUNT(*) dla identyfikatorów ze wszystkich czterech partycji HASH.
	"SELECT COUNT(*) AS user_count FROM users AS u WHERE u.id IN (232136,100001,1500002,5100003);",

	// Q05. Agregaty dla użytkowników z jednej partycji HASH.
	"SELECT COUNT(*) AS user_count,SUM(u.reputation) AS total_reputation,AVG(u.views) AS avg_views,SUM(u.upvotes) AS total_upvotes FROM users AS u WHERE u.id IN (232136,232756,234152,238100);",

	// Q06. Te same agregaty dla dwóch partycji HASH.
	"SELECT COUNT(*) AS user_count,SUM(u.reputation) AS total_reputation,AVG(u.views) AS avg_views,SUM(u.upvotes) AS total_upvotes FROM users AS u WHERE u.id IN (232136,232756,100001,650001);",

	// Q07. Te same agregaty dla trzech partycji HASH.
	"SELECT COUNT(*) AS user_count,SUM(u.reputation) AS total_reputation,AVG(u.views) AS avg_views,SUM(u.upvotes) AS total_upvotes FROM users AS u WHERE u.id IN (232136,100001,650001,1500002,650002);",

	// Q08. Te same agregaty dla wszystkich czterech partycji HASH.
	"SELECT COUNT(*) AS user_count,SUM(u.reputation) AS total_reputation,AVG(u.views) AS avg_views,SUM(u.upvotes) AS total_upvotes FROM users AS u WHERE u.id IN (232136,100001,1500002,5100003);",

	// Q09. Pobranie i sortowanie użytkowników z jednej partycji HASH.
	"SELECT u.id,u.display_name,u.reputation,u.views,u.upvotes,u.downvotes FROM users AS u WHERE u.id IN (232136,232756,234152,238100) ORDER BY u.reputation DESC,u.id;",

	// Q10. To samo dla dwóch partycji HASH.
	"SELECT u.id,u.display_name,u.reputation,u.views,u.upvotes,u.downvotes FROM users AS u WHERE u.id IN (232136,232756,100001,650001) ORDER BY u.reputation DESC,u.id;",

	// Q11. To samo dla trzech partycji HASH.
	"SELECT u.id,u.display_name,u.reputation,u.views,u.upvotes,u.downvotes FROM users AS u WHERE u.id IN (232136,100001,650001,1500002,650002) ORDER BY u.reputation DESC,u.id;",

	// Q12. To samo dla wszystkich czterech partycji HASH.
	"SELECT u.id,u.display_name,u.reputation,u.views,u.upvotes,u.downvotes FROM users AS u WHERE u.id IN (232136,100001,1500002,5100003) ORDER BY u.reputation DESC,u.id;",

	// Q13. Grupowanie dla użytkowników z jednej partycji HASH.
	"SELECT CASE WHEN u.reputation<100 THEN 'low' WHEN u.reputation<1000 THEN 'medium' ELSE 'high' END AS reputation_group,COUNT(*) AS user_count,AVG(u.views) AS avg_views FROM users AS u WHERE u.id IN (232136,232756,234152,238100) GROUP BY reputation_group ORDER BY reputation_group;",

	// Q14. To samo grupowanie dla dwóch partycji HASH.
	"SELECT CASE WHEN u.reputation<100 THEN 'low' WHEN u.reputation<1000 THEN 'medium' ELSE 'high' END AS reputation_group,COUNT(*) AS user_count,AVG(u.views) AS avg_views FROM users AS u WHERE u.id IN (232136,232756,100001,650001) GROUP BY reputation_group ORDER BY reputation_group;",

	// Q15. To samo grupowanie dla trzech partycji HASH.
	"SELECT CASE WHEN u.reputation<100 THEN 'low' WHEN u.reputation<1000 THEN 'medium' ELSE 'high' END AS reputation_group,COUNT(*) AS user_count,AVG(u.views) AS avg_views FROM users AS u WHERE u.id IN (232136,100001,650001,1500002,650002) GROUP BY reputation_group ORDER BY reputation_group;",

	// Q16. To samo grupowanie dla wszystkich czterech partycji HASH.
	"SELECT CASE WHEN u.reputation<100 THEN 'low' WHEN u.reputation<1000 THEN 'medium' ELSE 'high' END AS reputation_group,COUNT(*) AS user_count,AVG(u.views) AS avg_views FROM users AS u WHERE u.id IN (232136,100001,1500002,5100003) GROUP BY reputation_group ORDER BY reputation_group;",

	// Q17. Krótki zakres trzech kolejnych wartości id; optymalizator może przekształcić go do IN i ograniczyć liczbę partycji.
	"SELECT u.id,u.display_name,u.reputation FROM users AS u WHERE u.id>=5000000 AND u.id<5000003 ORDER BY u.id;",

	// Q18. Szeroki zakres id; przy HASH nie daje skutecznego ograniczenia do części partycji.
	"SELECT COUNT(*) AS user_count,SUM(u.reputation) AS total_reputation FROM users AS u WHERE u.id>=5000000 AND u.id<5100000;",

	// Q19. Przypadek kontrolny z filtrem po kolumnie innej niż klucz partycjonowania.
	"SELECT COUNT(*) AS user_count,AVG(u.views) AS avg_views FROM users AS u WHERE u.reputation>=10000;",

	// Q20. Przypadek kontrolny bez predykatu po id; wszystkie partycje users.
	"SELECT COUNT(*) AS user_count,SUM(u.reputation) AS total_reputation,AVG(u.views) AS avg_views FROM users AS u;",
}
