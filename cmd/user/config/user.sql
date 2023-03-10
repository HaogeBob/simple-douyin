CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
  `follow_count` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Follow_count',
  `follower_count` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Follower_count',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='User account table';