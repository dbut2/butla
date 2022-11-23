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

INSERT INTO `links` (`code`, `url`, `expiry`, `ip`) VALUES ('default', 'shorten', NULL, NULL);