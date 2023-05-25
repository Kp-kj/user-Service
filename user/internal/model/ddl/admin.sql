-- admin --
CREATE TABLE `admin_user` (
                              `admin_user_id` BIGINT(20) NOT NULL,
                              `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` DATETIME DEFAULT NULL,
                              `deleted_at` DATETIME DEFAULT NULL,
                              `admin_name` VARCHAR(256) NOT NULL,
                              `admin_passwd` VARCHAR(256) NOT NULL,
                              `admin_status` TINYINT NOT NULL DEFAULT 0,
                              `admin_role` TINYINT NOT NULL DEFAULT 0,
                              PRIMARY KEY (`admin_user_id`),
                              UNIQUE INDEX `idx_admin_name` (`admin_name`)
);