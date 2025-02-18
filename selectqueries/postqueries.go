package selectqueries

var idselcectPost = []string{
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
}

var idselcectPostLimit = []string{
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
}

var postTypeselcectPost = []string{
	"SELECT * FROM posts WHERE post_type_id = 15;",
	"SELECT * FROM posts WHERE post_type_id = 3;",
	"SELECT * FROM posts WHERE post_type_id = 8;",
	"SELECT * FROM posts WHERE post_type_id = 7;",
	"SELECT * FROM posts WHERE post_type_id = 20;",
	"SELECT * FROM posts WHERE post_type_id = 1;",
}

var postTypeselcectPostLimit = []string{
	"SELECT * FROM posts WHERE post_type_id = 15 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 3 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 8 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 7 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 20 limit 10;",
	"SELECT * FROM posts WHERE post_type_id = 1 limit 10;",
}

var accepdedansweridcectPost = []string{
	"SELECT * FROM posts WHERE accepted_answer_id IS NOT NULL;",
	"SELECT * FROM posts WHERE accepted_answer_id IS NULL;",
}
var accepdedansweridcectPostLimit = []string{
	"SELECT * FROM posts WHERE accepted_answer_id IS NOT NULL limit 10;",
	"SELECT * FROM posts WHERE accepted_answer_id IS NULL limit 10;",
}

var parentIDselectPost = []string{
	"SELECT * FROM posts WHERE parent_id = 39 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 50 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 23425 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 2463 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 53 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 83 limit 2;",
	"SELECT * FROM posts WHERE parent_id = 324234 limit 2;",
}

var creationDateSelectPost = []string{
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
}

var scoreSelectPost = []string{
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
}
var viewCountSelectPost = []string{
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
}

var postBodySelectPost = []string{
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
}

var ownerUserIDSelectPost = []string{
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
}

var lastEditorUserIDSelectPost = []string{
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
}

var lastEditDateSelectPost = []string{
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
}

var lastActivityDateSelectPost = []string{
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
}

var postTitleSelectPost = []string{
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
}

var tagsSelectPost = []string{
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
}

var answerCountSelectPost = []string{
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
}

var commentCountSelectPost = []string{
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
}

var contentLicenseSelectPost = []string{
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
}

var multiConditionSelectPost = []string{
	"SELECT * FROM posts WHERE post_type_id = 1 AND score > 100;",
	"SELECT * FROM posts WHERE tags LIKE '%java%' AND answer_count > 5;",
	"SELECT * FROM posts WHERE creation_date > '2023-01-01' AND view_count < 1000;",
	"SELECT * FROM posts WHERE owner_user_id = 100 AND comment_count > 10;",
	"SELECT * FROM posts WHERE post_body LIKE '%error%' AND score > 50;",
}

var groupBySelectPost = []string{
	"SELECT post_type_id, COUNT(*) FROM posts GROUP BY post_type_id;",
	"SELECT owner_user_id, COUNT(*) FROM posts GROUP BY owner_user_id;",
	"SELECT tags, COUNT(*) FROM posts GROUP BY tags;",
	"SELECT content_license, COUNT(*) FROM posts GROUP BY content_license;",
	"SELECT YEAR(creation_date), COUNT(*) FROM posts GROUP BY YEAR(creation_date);",
}

var orderBySelectPost = []string{
	"SELECT * FROM posts ORDER BY creation_date DESC LIMIT 20;",
	"SELECT * FROM posts ORDER BY score DESC LIMIT 10;",
	"SELECT * FROM posts ORDER BY view_count DESC LIMIT 15;",
	"SELECT * FROM posts ORDER BY last_activity_date DESC LIMIT 30;",
	"SELECT * FROM posts ORDER BY answer_count DESC LIMIT 25;",
}

var betweenSelectPost = []string{
	"SELECT * FROM posts WHERE creation_date BETWEEN '2022-01-01' AND '2022-12-31';",
	"SELECT * FROM posts WHERE score BETWEEN 50 AND 100;",
	"SELECT * FROM posts WHERE view_count BETWEEN 1000 AND 5000;",
	"SELECT * FROM posts WHERE last_edit_date BETWEEN '2023-01-01' AND '2023-12-31';",
	"SELECT * FROM posts WHERE answer_count BETWEEN 5 AND 20;",
}

var fullScanSelectPost = []string{
	"SELECT * FROM posts LIMIT 100;",
	"SELECT * FROM posts LIMIT 500;",
	"SELECT * FROM posts LIMIT 1000;",
	"SELECT * FROM posts LIMIT 5000;",
	"SELECT * FROM posts LIMIT 10000;",
}

var countSelectPost = []string{
	"SELECT COUNT(*) FROM posts;",
	"SELECT COUNT(*) FROM posts WHERE post_type_id = 1;",
	"SELECT COUNT(*) FROM posts WHERE score > 100;",
	"SELECT COUNT(*) FROM posts WHERE tags LIKE '%java%';",
	"SELECT COUNT(*) FROM posts WHERE creation_date > '2023-01-01';",
}

var analizeUserActivity = []string{
	"SELECT owner_user_id, COUNT(*) AS post_count FROM posts WHERE creation_date > '2023-01-01' GROUP BY owner_user_id HAVING post_count > 5;",
	"SELECT owner_user_id, AVG(score) AS avg_score FROM posts GROUP BY owner_user_id ORDER BY avg_score DESC LIMIT 10;",
	"SELECT owner_user_id, SUM(view_count) AS total_views FROM posts GROUP BY owner_user_id ORDER BY total_views DESC LIMIT 10;",
	"SELECT owner_user_id, MAX(score) AS max_score FROM posts GROUP BY owner_user_id ORDER BY max_score DESC LIMIT 10;",
	"SELECT owner_user_id, COUNT(*) AS post_count FROM posts WHERE post_type_id = 1 GROUP BY owner_user_id HAVING post_count > 10;",
}

var analizePostTags = []string{
	"SELECT tags, COUNT(*) AS post_count FROM posts WHERE creation_date > '2022-01-01' GROUP BY tags ORDER BY post_count DESC LIMIT 10;",
	"SELECT tags, AVG(score) AS avg_score FROM posts GROUP BY tags HAVING COUNT(*) > 50 ORDER BY avg_score DESC LIMIT 10;",
	"SELECT tags, SUM(answer_count) AS total_answers FROM posts GROUP BY tags ORDER BY total_answers DESC LIMIT 10;",
	"SELECT tags, SUM(comment_count) AS total_comments FROM posts GROUP BY tags ORDER BY total_comments DESC LIMIT 10;",
	"SELECT tags, AVG(view_count) AS avg_views FROM posts GROUP BY tags HAVING COUNT(*) > 100 ORDER BY avg_views DESC LIMIT 10;",
}

var analizePostDate = []string{
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS post_count FROM posts GROUP BY YEAR(creation_date) ORDER BY year DESC;",
	"SELECT MONTH(creation_date) AS month, COUNT(*) AS post_count FROM posts WHERE YEAR(creation_date) = 2023 GROUP BY MONTH(creation_date) ORDER BY month;",
	"SELECT DATE(creation_date) AS date, COUNT(*) AS post_count FROM posts GROUP BY DATE(creation_date) ORDER BY date DESC LIMIT 10;",
	"SELECT YEAR(creation_date) AS year, AVG(score) AS avg_score FROM posts GROUP BY YEAR(creation_date) ORDER BY year DESC;",
	"SELECT WEEK(creation_date) AS week, COUNT(*) AS post_count FROM posts WHERE YEAR(creation_date) = 2023 GROUP BY WEEK(creation_date) ORDER BY week;",
}
var viewMarkAnalize = []string{
	"SELECT score, COUNT(*) AS post_count FROM posts GROUP BY score ORDER BY score DESC LIMIT 10;",
	"SELECT view_count, COUNT(*) AS post_count FROM posts GROUP BY view_count ORDER BY view_count DESC LIMIT 10;",
	"SELECT score, AVG(view_count) AS avg_views FROM posts GROUP BY score ORDER BY score DESC LIMIT 10;",
	"SELECT view_count, AVG(score) AS avg_score FROM posts GROUP BY view_count ORDER BY view_count DESC LIMIT 10;",
	"SELECT score, SUM(answer_count) AS total_answers FROM posts GROUP BY score ORDER BY score DESC LIMIT 10;",
}

var answersComsAnalize = []string{
	"SELECT answer_count, COUNT(*) AS post_count FROM posts GROUP BY answer_count ORDER BY answer_count DESC LIMIT 10;",
	"SELECT comment_count, COUNT(*) AS post_count FROM posts GROUP BY comment_count ORDER BY comment_count DESC LIMIT 10;",
	"SELECT answer_count, AVG(score) AS avg_score FROM posts GROUP BY answer_count ORDER BY answer_count DESC LIMIT 10;",
	"SELECT comment_count, AVG(view_count) AS avg_views FROM posts GROUP BY comment_count ORDER BY comment_count DESC LIMIT 10;",
	"SELECT answer_count, SUM(comment_count) AS total_comments FROM posts GROUP BY answer_count ORDER BY answer_count DESC LIMIT 10;",
}
var liceneAnalize = []string{
	"SELECT content_license, COUNT(*) AS post_count FROM posts GROUP BY content_license ORDER BY post_count DESC;",
	"SELECT content_license, AVG(score) AS avg_score FROM posts GROUP BY content_license ORDER BY avg_score DESC;",
	"SELECT content_license, SUM(view_count) AS total_views FROM posts GROUP BY content_license ORDER BY total_views DESC;",
	"SELECT content_license, MAX(answer_count) AS max_answers FROM posts GROUP BY content_license ORDER BY max_answers DESC;",
	"SELECT content_license, AVG(comment_count) AS avg_comments FROM posts GROUP BY content_license ORDER BY avg_comments DESC;",
}

var conditionsAnalize = []string{
	"SELECT * FROM posts WHERE tags LIKE '%python%' AND post_type_id = 1 AND creation_date > '2023-01-01' ORDER BY score DESC LIMIT 20;",
	"SELECT * FROM posts WHERE score > 100 AND view_count > 1000 AND answer_count > 5 ORDER BY creation_date DESC LIMIT 10;",
	"SELECT * FROM posts WHERE post_body LIKE '%error%' AND tags LIKE '%java%' AND comment_count > 10 ORDER BY view_count DESC LIMIT 15;",
	"SELECT * FROM posts WHERE creation_date BETWEEN '2022-01-01' AND '2022-12-31' AND score > 50 AND answer_count > 2 ORDER BY view_count DESC LIMIT 20;",
	"SELECT * FROM posts WHERE owner_user_id = 100 AND post_type_id = 1 AND creation_date > '2023-01-01' ORDER BY score DESC LIMIT 10;",
}

var agregatePostAnalize = []string{
	"SELECT post_type_id, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(view_count) AS total_views FROM posts GROUP BY post_type_id ORDER BY post_count DESC;",
	"SELECT tags, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(answer_count) AS total_answers FROM posts GROUP BY tags ORDER BY post_count DESC LIMIT 10;",
	"SELECT owner_user_id, COUNT(*) AS post_count, AVG(score) AS avg_score, MAX(view_count) AS max_views FROM posts GROUP BY owner_user_id ORDER BY post_count DESC LIMIT 10;",
	"SELECT YEAR(creation_date) AS year, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(comment_count) AS total_comments FROM posts GROUP BY YEAR(creation_date) ORDER BY year DESC;",
	"SELECT content_license, COUNT(*) AS post_count, AVG(score) AS avg_score, SUM(answer_count) AS total_answers FROM posts GROUP BY content_license ORDER BY post_count DESC;",
}

var subQueryAnalize = []string{
	"SELECT * FROM posts WHERE score > (SELECT AVG(score) FROM posts);",
	"SELECT * FROM posts WHERE view_count > (SELECT AVG(view_count) FROM posts WHERE post_type_id = 1);",
	"SELECT * FROM posts WHERE answer_count > (SELECT AVG(answer_count) FROM posts WHERE tags LIKE '%java%');",
	"SELECT * FROM posts WHERE comment_count > (SELECT AVG(comment_count) FROM posts WHERE creation_date > '2023-01-01');",
	"SELECT * FROM posts WHERE owner_user_id IN (SELECT owner_user_id FROM posts GROUP BY owner_user_id HAVING COUNT(*) > 50);",
}
