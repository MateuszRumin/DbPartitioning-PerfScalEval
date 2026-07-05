Drop Database if EXISTS testdb;

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
	id INT,
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
);
ALTER TABLE users
  MODIFY id INT NOT NULL;

ALTER TABLE users
  ADD PRIMARY KEY (id);

ALTER TABLE users
  MODIFY id INT NOT NULL AUTO_INCREMENT;





CREATE TABLE posts (
	id INT,
    post_type_id INT NOT NULL,
    accepted_answer_id INT,
    parent_id INT,
    creation_date DATETIME,
    score INT,
    view_count INT,
    post_body VARCHAR(10000),
    owner_user_id INT NULL,
    last_editor_user_id INT,
    last_edit_date DATETIME,
    last_activity_date DATETIME,
    post_title VARCHAR(500),
    tags VARCHAR(500),
    answer_count INT,
    comment_count INT,
    content_license VARCHAR(100),
    PRIMARY KEY (id)
);

ALTER TABLE posts MODIFY id INT NOT NULL;
ALTER TABLE posts ADD PRIMARY KEY (id);
ALTER TABLE posts MODIFY id INT NOT NULL AUTO_INCREMENT;



ALTER TABLE posts 
ADD CONSTRAINT fk_posts_user_id FOREIGN KEY (owner_user_id) REFERENCES users(id);




CREATE TABLE comments (
	  id INT,
    post_id INT NULL,
    score INT,
    comment_text VARCHAR(4000),
    creation_date DATETIME,
    user_id INT NULL,
    content_license VARCHAR(100)
    
);

ALTER TABLE comments MODIFY id INT NOT NULL;
ALTER TABLE comments ADD PRIMARY KEY (id);

ALTER TABLE comments MODIFY id INT NOT NULL AUTO_INCREMENT;




ALTER TABLE comments
ADD CONSTRAINT fk_comments_post_id FOREIGN KEY (post_id) REFERENCES posts(id),
ADD CONSTRAINT fk_comments_user_id FOREIGN KEY (user_id) REFERENCES users(id);


CREATE TABLE post_history (
  id INT,
  post_history_type_id INT,
  post_id INT NULL,
  revision_guid VARCHAR(100),
  creation_date DATETIME,
  user_id INT NULL,
  post_text VARCHAR(10000),
  content_license VARCHAR(100)
);

ALTER TABLE post_history MODIFY id INT NOT NULL;
ALTER TABLE post_history ADD PRIMARY KEY (id);

ALTER TABLE post_history MODIFY id INT NOT NULL AUTO_INCREMENT;





ALTER TABLE post_history
ADD CONSTRAINT fk_post_history_post_id FOREIGN KEY (post_id) REFERENCES posts(id),
ADD CONSTRAINT fk_post_history_user_id FOREIGN KEY (user_id) REFERENCES users(id);


CREATE TABLE post_links (
	id INT,
    creation_date DATETIME,
    post_id INT NULL,
    related_post_id INT NULL,
    link_type_id INT
);

ALTER TABLE post_links MODIFY id INT NOT NULL;
ALTER TABLE post_links ADD PRIMARY KEY (id);

ALTER TABLE post_links MODIFY id INT NOT NULL AUTO_INCREMENT;


ALTER TABLE post_links
ADD CONSTRAINT fk_post_links_post_id FOREIGN KEY (post_id) REFERENCES posts(id),
ADD CONSTRAINT fk_related_post_links_post_id FOREIGN KEY (related_post_id) REFERENCES posts(id);



CREATE TABLE votes (
	id INT,
    post_id INT NULL,
    vote_type_id INT,
    creation_date DATETIME
);

ALTER TABLE votes MODIFY id INT NOT NULL;
ALTER TABLE votes ADD PRIMARY KEY (id);

ALTER TABLE votes MODIFY id INT NOT NULL AUTO_INCREMENT;





ALTER TABLE votes ADD CONSTRAINT fk_votes_post_id
FOREIGN KEY (post_id) REFERENCES posts(id);



CREATE TABLE badges (
  id INT,
  user_id INT NULL,
  badge_name VARCHAR(500),
  badge_date DATETIME,
  class INT,
  tag_based VARCHAR(10)
);




ALTER TABLE badges MODIFY id INT NOT NULL;
ALTER TABLE badges ADD PRIMARY KEY (id);




ALTER TABLE badges MODIFY id INT NOT NULL AUTO_INCREMENT;






ALTER TABLE badges ADD CONSTRAINT fk_badges_user_id
FOREIGN KEY (user_id) REFERENCES users(id);



CREATE TABLE tags (
	id INT,
    tag_name VARCHAR(100),
    tag_count INT,
    except_post_id INT,
    wiki_post_id INT
);

ALTER TABLE tags MODIFY id INT NOT NULL;
ALTER TABLE tags ADD PRIMARY KEY (id);
ALTER TABLE tags MODIFY id INT NOT NULL AUTO_INCREMENT;



