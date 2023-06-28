-- helpCategory --
CREATE TABLE `helpCategory` (
    `helpCategory_id` bigint(20) NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `category_status` tinyint NOT NULL DEFAULT 1,  # 0:上架 1:下架
    PRIMARY KEY (`helpCategory_id`)
);


-- helpCategoryTranslation --
CREATE  TABLE   `helpCategoryTranslation` (
    `helpCategoryTranslation_id` bigint(20) NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `helpCategory_id` bigint(20) NOT NULL,
    language_code varchar(256) NOT NULL,
    `category_name` varchar(256) NOT NULL,
    PRIMARY KEY ( `helpCategoryTranslation_id`),
    UNIQUE INDEX `idx_helpCategory_id_language_id` (`helpCategory_id`, `language_code`)
);