Drop Database if EXISTS testdbp;

CREATE DATABASE testdb;

USE testdb;

DROP TABLE IF EXISTS badges;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS post_history;
DROP TABLE IF EXISTS post_links;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS votes;


CREATE TABLE users (
	id INT NOT NULL PRIMARY KEY,
    reputation INT,
    creation_date DATETIME,
    display_name VARCHAR(200),
    last_access_date DATETIME,
    website_url VARCHAR(1000),
    location VARCHAR(200),
    about_me VARCHAR(10000),
    views INT,
    upvotes INT,
    downvotes INT,
    account_id INT
)
PARTITION BY HASH (id)
PARTITIONS 4;






CREATE TABLE posts (
    id INT NOT NULL,
    post_type_id INT,
    accepted_answer_id INT,
    parent_id INT,
    creation_date DATETIME NOT NULL,
    score INT,
    view_count INT,
    post_body VARCHAR(10000),
    owner_user_id INT,
    last_editor_user_id INT,
    last_edit_date DATETIME,
    last_activity_date DATETIME,
    post_title VARCHAR(500),
    tags VARCHAR(500),
    answer_count INT,
    comment_count INT,
    content_license VARCHAR(100),

    PRIMARY KEY (id, creation_date)
)
PARTITION BY RANGE COLUMNS (creation_date) (
    PARTITION before_2008_q3 VALUES LESS THAN ('2008-07-01'),
    PARTITION p2008_q3 VALUES LESS THAN ('2008-10-01'),
    PARTITION p2008_q4 VALUES LESS THAN ('2009-01-01'),
    PARTITION p2009_q1 VALUES LESS THAN ('2009-04-01'),
    PARTITION p2009_q2 VALUES LESS THAN ('2009-07-01'),
    PARTITION p2009_q3 VALUES LESS THAN ('2009-10-01'),
    PARTITION p2009_q4 VALUES LESS THAN ('2010-01-01'),
    PARTITION p2010_q1 VALUES LESS THAN ('2010-04-01'),
    PARTITION p2010_q2 VALUES LESS THAN ('2010-07-01'),
    PARTITION p2010_q3 VALUES LESS THAN ('2010-10-01'),
    PARTITION p2010_q4 VALUES LESS THAN ('2011-01-01'),
    PARTITION p_future VALUES LESS THAN (MAXVALUE)
);



SELECT COUNT(*)
FROM posts p
LEFT JOIN users u ON u.id = p.owner_user_id
WHERE p.owner_user_id IS NOT NULL
  AND u.id IS NULL;

UPDATE posts p
LEFT JOIN users u ON u.id = p.owner_user_id
SET p.owner_user_id = NULL
WHERE p.owner_user_id IS NOT NULL
  AND u.id IS NULL;



CREATE TABLE comments (
	id INT NOT NULL,
    post_id INT NOT NULL,
    score INT,
    comment_text VARCHAR(4000),
    creation_date DATETIME,
    user_id INT NULL,
    content_license VARCHAR(100),
    PRIMARY KEY (id,post_id)
    
)
PARTITION BY KEY (id,post_id)
PARTITIONS 4;

ALTER TABLE comments MODIFY id INT NOT NULL;
ALTER TABLE comments ADD PRIMARY KEY (id);







CREATE TABLE post_history (
  id INT NOT NULL,
  post_history_type_id INT,
  post_id INT NOT NULL,
  revision_guid VARCHAR(100),
  creation_date DATETIME,
  user_id INT NULL,
  post_text VARCHAR(10000),
  content_license VARCHAR(100),
  PRIMARY KEY (id,post_id)
)
PARTITION BY KEY (id,post_id)
PARTITIONS 8;

ALTER TABLE post_history MODIFY id INT NOT NULL;
ALTER TABLE post_history ADD PRIMARY KEY (id);






CREATE TABLE post_links (
	id INT,
    creation_date DATETIME,
    post_id INT NULL,
    related_post_id INT NULL,
    link_type_id INT
);

ALTER TABLE post_links MODIFY id INT NOT NULL;
ALTER TABLE post_links ADD PRIMARY KEY (id);




CREATE TABLE votes (
	id INT NOT NULL,
    post_id INT NOT NULL,
    vote_type_id INT,
    creation_date DATETIME,
    PRIMARY KEY (id,post_id)
)
PARTITION BY KEY (id,post_id)
PARTITIONS 4;



ALTER TABLE votes MODIFY id INT NOT NULL;
ALTER TABLE votes ADD PRIMARY KEY (id);



CREATE TABLE badges (
  id INT NOT NULL,
  user_id INT NULL,
  badge_name VARCHAR(500),
  badge_date DATETIME,
  class INT,
  tag_based VARCHAR(10),
  PRIMARY KEY (id)
);

ALTER TABLE badges MODIFY id INT NOT NULL;
ALTER TABLE badges ADD PRIMARY KEY (id);




CREATE TABLE tags (
	id INT,
    tag_name VARCHAR(100),
    tag_count INT,
    except_post_id INT,
    wiki_post_id INT
);

ALTER TABLE tags MODIFY id INT NOT NULL;
ALTER TABLE tags ADD PRIMARY KEY (id);




ALTER TABLE users MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE posts MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE comments MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE post_history MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE post_links MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE votes MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE badges MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE tags MODIFY id INT NOT NULL AUTO_INCREMENT;
