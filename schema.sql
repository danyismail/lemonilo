CREATE DATABASE `lemonilo` -- lemonilo.users definition
CREATE TABLE `users` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `email` varchar(255) DEFAULT NULL,
    `address` varchar(255) DEFAULT NULL,
    `password` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE = InnoDB AUTO_INCREMENT = 9 DEFAULT CHARSET = utf8;