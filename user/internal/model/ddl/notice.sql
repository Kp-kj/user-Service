-- system notice --
CREATE TABLE `systemNotice` (
                                `systemNotice_id` bigint(20) NOT NULL,
                                `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` datetime DEFAULT NULL,
                                `deleted_at` datetime DEFAULT NULL,
                                `notice_title` varchar(255) NOT NULL,   -- 通知标题
                                `notice_content` text NOT NULL,         -- 通知内容
                                `notice_category` tinyint(1) NOT NULL DEFAULT 0, -- 通知类别   0：手动通知 1：自动通知
                                `notice_status` tinyint(1) NOT NULL DEFAULT 0,   -- 通知状态   0:待发布 1：已发布 2：已下架
                                `notice_start_time` datetime DEFAULT NULL,    -- 自动通知时间
                                PRIMARY KEY (`systemNotice_id`)
);

-- user notice --
CREATE TABLE `userNotice`(
                             `userNotice_id` bigint(20) NOT NULL,
                             `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` datetime DEFAULT NULL,
                             `deleted_at` datetime DEFAULT NULL,
                             `user_id` bigint(20) NOT NULL,    -- 用户id
                             `notice_content` text NOT NULL,         -- 通知内容
                             `notice_status` tinyint(1) NOT NULL DEFAULT 1,   -- 通知状态   1：已发布 2：已下架
                             PRIMARY KEY (`userNotice_id`)
);

-- record notice --
CREATE TABLE `recordNotice` (
                                `recordNotice_id` bigint(20) NOT NULL,
                                `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` datetime DEFAULT NULL,
                                `deleted_at` datetime DEFAULT NULL,
                                `user_id` bigint(20) NOT NULL,    -- 用户id
                                `systemNotice_id` bigint(20), -- 系统通知id
                                `userNotice_id` bigint(20), -- 用户通知id
                                `recordNotice_category` tinyint(1) NOT NULL DEFAULT 0, -- 通知类别   0：系统通知 1：用户通知
                                `recordNotice_status` tinyint(1) NOT NULL DEFAULT 0,   -- 通知状态   0:未读 1：已读
                                PRIMARY KEY (`recordNotice_id`)
);
