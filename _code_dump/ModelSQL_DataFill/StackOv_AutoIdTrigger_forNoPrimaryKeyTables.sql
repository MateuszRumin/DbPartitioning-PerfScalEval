-- Tworzenie triggerów dla tabel, aby automatycznie nadawać unikalne ID

DELIMITER $$


CREATE TRIGGER before_insert_badges
BEFORE INSERT ON badges
FOR EACH ROW
BEGIN
    SET NEW.id = (SELECT COALESCE(MAX(id), 0) + 1 FROM badges);
END$$


CREATE TRIGGER before_insert_comments
BEFORE INSERT ON comments
FOR EACH ROW
BEGIN
    SET NEW.id = (SELECT COALESCE(MAX(id), 0) + 1 FROM comments);
END$$


CREATE TRIGGER before_insert_post_history
BEFORE INSERT ON post_history
FOR EACH ROW
BEGIN
    SET NEW.id = (SELECT COALESCE(MAX(id), 0) + 1 FROM post_history);
END$$


CREATE TRIGGER before_insert_post_links
BEFORE INSERT ON post_links
FOR EACH ROW
BEGIN
    SET NEW.id = (SELECT COALESCE(MAX(id), 0) + 1 FROM post_links);
END$$


CREATE TRIGGER before_insert_posts
BEFORE INSERT ON posts
FOR EACH ROW
BEGIN
    SET NEW.id = (SELECT COALESCE(MAX(id), 0) + 1 FROM posts);
END$$


CREATE TRIGGER before_insert_users
BEFORE INSERT ON users
FOR EACH ROW
BEGIN
    SET NEW.id = (SELECT COALESCE(MAX(id), 0) + 1 FROM users);
END$$


CREATE TRIGGER before_insert_votes
BEFORE INSERT ON votes
FOR EACH ROW
BEGIN
    SET NEW.id = (SELECT COALESCE(MAX(id), 0) + 1 FROM votes);
END$$

DELIMITER ;
