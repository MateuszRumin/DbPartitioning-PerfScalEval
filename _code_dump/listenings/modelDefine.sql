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

ALTER TABLE posts MODIFY id INT NOT NULL AUTO_INCREMENT;
ALTER TABLE posts ADD CONSTRAINT fk_posts_user_id FOREIGN KEY (owner_user_id) REFERENCES users(id);