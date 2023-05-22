-- user --
CREATE TABLE `user` (
                         `user_id` bigint(20) NOT NULL,
                         `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime DEFAULT NULL,
                         `deleted_at` datetime DEFAULT NULL,
                         `twitter_id` varchar(256) NOT NULL,
                         PRIMARY KEY (`user_id`),
                         UNIQUE KEY `idx_twitter_id` (`twitter_id`),
                         KEY `idx_users_deleted_at` (`deleted_at`)
#                          CONSTRAINT `fk_users_authentications_user_id` FOREIGN KEY (`user_id`) REFERENCES `authentications` (`user_id`),
#                          CONSTRAINT `fk_users_profiles_user_id` FOREIGN KEY (`user_id`) REFERENCES `profiles` (`user_id`),
#                          CONSTRAINT `fk_users_blacks_user_id` FOREIGN KEY (`user_id`) REFERENCES `blacks` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
