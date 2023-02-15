CREATE DATABASE `shortener`;

USE `shortener`;

CREATE TABLE `links` (
    `id` int(11) NOT NULL,
    `code` varchar(36) NOT NULL,
    `url` text NOT NULL,
    `expiry` datetime DEFAULT CURRENT_TIMESTAMP,
    `ip` varchar(45) DEFAULT '0.0.0.0'
);

ALTER TABLE `links`
    ADD PRIMARY KEY (`id`),
    ADD UNIQUE KEY `code` (`code`);

ALTER TABLE `links`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

CREATE TABLE `users` (
    `id`       int(11) NOT NULL,
    `username` varchar(36) NOT NULL,
    `email`    varchar(72) NOT NULL,
    `name`     varchar(36) NOT NULL,
    `password` varchar(36) NOT NULL
);

ALTER TABLE `users`
    ADD PRIMARY KEY (`id`),
    ADD UNIQUE KEY `username` (`username`);

ALTER TABLE `users`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

INSERT INTO `links` (`code`, `url`, `expiry`, `ip`) VALUES ('default', 'https://localhost:8088/shorten', NULL, NULL);
INSERT INTO `users` (`username`, `email`, `name`, `password`) VALUES ('admin', 'admin@example.com', 'Admin', 'password');
