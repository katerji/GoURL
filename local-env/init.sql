DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;
CREATE TABLE `user`
(
    `id`       int(11) NOT NULL AUTO_INCREMENT,
    `email`    varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `url`
(
    `id`           int          NOT NULL AUTO_INCREMENT,
    `user_id`      int          NOT NULL,
    `original_url` text         NOT NULL,
    `short_url`    varchar(255) NOT NULL,
    `created_on`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on`   datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `original_url` (`original_url`(600),`user_id`),
    KEY            `short_url` (`short_url`),
    KEY            `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;