ALTER TABLE `user` ADD `admin` BOOLEAN NOT NULL AFTER `password`, ADD `moderator` BOOLEAN NOT NULL AFTER `admin`;
