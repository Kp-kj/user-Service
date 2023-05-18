USE user_server;

-- DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `user_id` bigint(20) unsigned NOT NULL,
                         `created_at` datetime(3) NOT NULL,
                         `updated_at` datetime(3) NOT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `twitter_id` varchar(256) NOT NULL,
                         PRIMARY KEY (`user_id`),
                         UNIQUE KEY `users_user_id_unique` (`user_id`),
                         KEY `idx_twitter_id` (`twitter_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- DROP TABLE IF EXISTS `user_trees`;
CREATE TABLE `invitations` (
                               `invitationId` bigint(20) NOT NULL ,
                               `created_at` datetime(3) NOT NULL,
                               `updated_at` datetime(3) NOT NULL,
                               `deleted_at` datetime(3) DEFAULT NULL,
                               `inviterId` bigint(20) NOT NULL,
                               `inviteeId` bigint(20) NOT NULL,
                               PRIMARY KEY (`invitationId`),
                               KEY `idx_inviter_id` (`inviterId`),
                               KEY `idx_invitee_id` (`inviteeId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- DROP TABLE IF EXISTS `user_trees`;
CREATE TABLE `invitation_trees` (
                                    `invitationTreeId` bigint(20) NOT NULL ,
                                    `created_at` datetime(3) NOT NULL,
                                    `updated_at` datetime(3) NOT NULL,
                                    `deleted_at` datetime(3) DEFAULT NULL,
                                    `userId` bigint(20) NOT NULL,
                                    `parentId` bigint(20) NOT NULL,
                                    `depth` int NOT NULL,
                                    PRIMARY KEY (`invitationTreeId`),
                                    KEY `idx_user_id` (`userId`),
                                    KEY `idx_parent_id` (`parentId`),
                                    KEY `idx_depth` (`depth`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `profiles` (
                            `profileId` bigint(20) NOT NULL ,
                            `created_at` datetime(3) NOT NULL,
                            `updated_at` datetime(3) NOT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `userId` bigint(20) NOT NULL,
                            `userName` varchar(256) NOT NULL,
                            `avatar` varchar(256) NOT NULL,
                            PRIMARY KEY (`profileId`),
                            KEY `idx_user_id` (`userId`),
                            KEY `idx_user_name` (`userName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
