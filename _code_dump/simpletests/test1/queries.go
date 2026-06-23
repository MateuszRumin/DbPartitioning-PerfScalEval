package main

var simpleSelect = []string{
	"SELECT * FROM posts WHERE answer_count = 5;",
	"SELECT * FROM posts WHERE tags LIKE '%sql%';",
	"SELECT * FROM posts WHERE last_edit_date = '2023-07-20';",
	"SELECT * FROM posts WHERE tags LIKE '%c%23%';", // C#
	"SELECT * FROM posts ORDER BY score DESC LIMIT 10;",
	"SELECT * FROM posts LIMIT 500;",
}
