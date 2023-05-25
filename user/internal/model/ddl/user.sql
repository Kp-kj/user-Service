
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
#                           CONSTRAINT `fk_users_invitations_user_id` FOREIGN KEY (`user_id`) REFERENCES `invitation` (`inviter_id`)
#                          CONSTRAINT `fk_users_authentications_user_id` FOREIGN KEY (`user_id`) REFERENCES `authentications` (`user_id`),
#                          CONSTRAINT `fk_users_profiles_user_id` FOREIGN KEY (`user_id`) REFERENCES `profiles` (`user_id`),
#                          CONSTRAINT `fk_users_blacks_user_id` FOREIGN KEY (`user_id`) REFERENCES `blacks` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- Invitation --
CREATE TABLE `invitation` (
                              `invitation_id` BIGINT(20) NOT NULL,
                              `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` DATETIME DEFAULT NULL,
                              `deleted_at` datetime DEFAULT NULL,
                              `inviter_id` BIGINT(20) NOT NULL,
                              `invitee_id` BIGINT(20) NOT NULL,
                              PRIMARY KEY (`invitation_id`),
                              INDEX `idx_inviter_id` (`inviter_id`),
                              INDEX `idx_invitee_id` (`invitee_id`),
                              UNIQUE (`invitation_id`)
);

-- InvitationTree --
CREATE TABLE `invitation_tree` (
                                   `invitation_tree_id` BIGINT(20) NOT NULL,
                                   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                                   `updated_at` DATETIME DEFAULT NULL,
                                   `deleted_at` datetime DEFAULT NULL,
                                   `user_id` BIGINT(20) NOT NULL,
                                   `parent_id` BIGINT(20) NOT NULL,
                                   `depth` INT NOT NULL,
                                   PRIMARY KEY (`invitation_tree_id`),
                                   INDEX `idx_user_id` (`user_id`),
                                   INDEX `idx_parent_id` (`parent_id`),
                                   INDEX `idx_depth` (`depth`),
                                   UNIQUE (`invitation_tree_id`)
);

-- profile --
CREATE TABLE `profile` (
                           `profile_id` BIGINT(20) NOT NULL ,
                           `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` DATETIME DEFAULT NULL,
                           `deleted_at` DATETIME DEFAULT NULL,
                           `user_id` BIGINT(20) NOT NULL,
                           `twitter_name` VARCHAR(256) NOT NULL,   # twitter 名字 带@的 不会变
                           `user_name` VARCHAR(256) NOT NULL,      # user name 可能会变 用户在twitter修改的
                           `avatar` VARCHAR(256) NOT NULL,         # avatar url  头像
                            `is_new` TINYINT(1) NOT NULL DEFAULT 1, # 1 是否是新用户 1 是 0 不是
                           PRIMARY KEY (`profile_id`),
                           INDEX `idx_user_id` (`user_id`),
                           INDEX `idx_user_name` (`user_name`),
                           InDex `idx_twitter_name` (`twitter_name`),
                           INDEX  `idx_is_new` (`is_new`),
                            UNIQUE (`profile_id`)
);
