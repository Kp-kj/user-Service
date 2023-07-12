-- white --
CREATE TABLE `white` (
                         `white_id` BIGINT(20) NOT NULL,
                         `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` DATETIME DEFAULT NULL,
                         `deleted_at` DATETIME DEFAULT NULL,
                         `name` VARCHAR(255) NOT NULL,
                         UNIQUE (`white_id`),
                         PRIMARY KEY (`white_id`)
);
-- white_list --
CREATE TABLE `white_list` (
                         `white_list_id` BIGINT(20) NOT NULL,
                         `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` DATETIME DEFAULT NULL,
                         `deleted_at` datetime DEFAULT NULL,
                         `white_id` BIGINT(20) NOT NULL,
                         `user_id` BIGINT(20) NOT NULL,
                         `name` VARCHAR(255) NOT NULL,
                         UNIQUE (`white_list_id`),
                         PRIMARY KEY (`white_list_id`)

);