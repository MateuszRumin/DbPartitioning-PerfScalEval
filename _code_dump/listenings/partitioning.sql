CREATE TABLE posts (
    id INT NOT NULL,
    post_type_id INT,
    parent_id INT,
    creation_date DATETIME NOT NULL,
    PRIMARY KEY (id, creation_date)
)PARTITION BY RANGE COLUMNS (creation_date) (
    PARTITION before_2008_q3 VALUES LESS THAN ('2008-07-01'),
    PARTITION p2008_q3 VALUES LESS THAN ('2008-10-01'),
    PARTITION p_future VALUES LESS THAN (MAXVALUE));

CREATE TABLE users (
	id INT NOT NULL PRIMARY KEY,
    reputation INT,
    creation_date DATETIME,
    display_name VARCHAR(200),
)PARTITION BY HASH (id) PARTITIONS 4;

CREATE TABLE comments (
	id INT NOT NULL,
    post_id INT NOT NULL,
    score INT,
    PRIMARY KEY (id,post_id)
)PARTITION BY KEY (id,post_id) PARTITIONS 4;