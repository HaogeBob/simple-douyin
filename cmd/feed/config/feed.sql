CREATE TABLE IF NOT EXISTS `vedios`(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'vedio_id',
    `author_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'author_id',
    `play_url`     varchar(128) NOT NULL DEFAULT '' COMMENT 'play_url',
    `cover_url`    varchar(128) NOT NULL DEFAULT '' COMMENT 'cover_url',
    `vedio_title`  varchar(128) NOT NULL DEFAULT '' COMMENT 'cover_url',
    `release_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'release_time',

    PRIMARY KEY ( `id` ) COMMENT 'vedio_id'
)ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;