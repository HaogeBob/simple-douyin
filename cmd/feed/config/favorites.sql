CREATE TABLE IF NOT EXISTS `favorites`(
    `id`        bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'favorites_id',
    `user_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT  'user_id',
    `video_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'vedio_id',

    PRIMARY KEY ( `id` ) COMMENT 'favorites_id'

)ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;