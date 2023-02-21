CREATE TABLE IF NOT EXISTS `comments`(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'comment_id',
    `vedio_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'vedio_id',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'user_id',
    `content`  varchar(128) NOT NULL DEFAULT '' COMMENT 'content',
    `creat_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'creat_time',

    PRIMARY KEY ( `id` ) COMMENT 'comment_id'
)ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;