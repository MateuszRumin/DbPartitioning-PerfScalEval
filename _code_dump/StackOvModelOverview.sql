| badges | CREATE TABLE `badges` (
  `id` int NOT NULL,
  `user_id` int DEFAULT NULL,
  `badge_name` varchar(500) DEFAULT NULL,
  `badge_date` datetime DEFAULT NULL,
  `class` int DEFAULT NULL,
  `tag_based` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_badges_user_id` (`user_id`),
  CONSTRAINT `fk_badges_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |

| comments | CREATE TABLE `comments` (
  `id` int NOT NULL,
  `post_id` int DEFAULT NULL,
  `score` int DEFAULT NULL,
  `comment_text` varchar(4000) DEFAULT NULL,
  `creation_date` datetime DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `content_license` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |

| post_history | CREATE TABLE `post_history` (
  `id` int NOT NULL,
  `post_history_type_id` int DEFAULT NULL,
  `post_id` int DEFAULT NULL,
  `revision_guid` varchar(100) DEFAULT NULL,
  `creation_date` datetime DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `post_text` varchar(10000) DEFAULT NULL,
  `content_license` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |


| post_links | CREATE TABLE `post_links` (
  `id` int NOT NULL,
  `creation_date` datetime DEFAULT NULL,
  `post_id` int DEFAULT NULL,
  `related_post_id` int DEFAULT NULL,
  `link_type_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |

| posts | CREATE TABLE `posts` (
  `id` int NOT NULL,
  `post_type_id` int DEFAULT NULL,
  `accepted_answer_id` int DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `creation_date` datetime DEFAULT NULL,
  `score` int DEFAULT NULL,
  `view_count` int DEFAULT NULL,
  `post_body` varchar(10000) DEFAULT NULL,
  `owner_user_id` int DEFAULT NULL,
  `last_editor_user_id` int DEFAULT NULL,
  `last_edit_date` datetime DEFAULT NULL,
  `last_activity_date` datetime DEFAULT NULL,
  `post_title` varchar(500) DEFAULT NULL,
  `tags` varchar(500) DEFAULT NULL,
  `answer_count` int DEFAULT NULL,
  `comment_count` int DEFAULT NULL,
  `content_license` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |




| users | CREATE TABLE `users` (
  `id` int NOT NULL,
  `reputation` int DEFAULT NULL,
  `creation_date` datetime DEFAULT NULL,
  `display_name` varchar(200) DEFAULT NULL,
  `last_access_date` datetime DEFAULT NULL,
  `website_url` varchar(1000) DEFAULT NULL,
  `location` varchar(200) DEFAULT NULL,
  `about_me` varchar(10000) DEFAULT NULL,
  `views` int DEFAULT NULL,
  `upvotes` int DEFAULT NULL,
  `downvotes` int DEFAULT NULL,
  `account_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |


| votes | CREATE TABLE `votes` (
  `id` int NOT NULL,
  `post_id` int DEFAULT NULL,
  `vote_type_id` int DEFAULT NULL,
  `creation_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci |

