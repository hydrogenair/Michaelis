CREATE TABLE `user_star`(
    `id` BIGINT UNSIGNED NOT NULL  AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL ,
    `star_id` BIGINT UNSIGNED NOT NULL ,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_star_index` (`user_id`,`star_id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;