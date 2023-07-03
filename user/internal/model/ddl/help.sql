-- helpCategory --
CREATE TABLE `helpCategory` (
    `helpCategory_id` bigint(20) NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `category_status` tinyint NOT NULL DEFAULT 1,  # 0:下架 1:上架
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

-- helpDocument --
CREATE TABLE `helpdocument` (
       `helpDocument_id` bigint(20) NOT NULL,
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
       `updated_at` datetime DEFAULT NULL,
       `deleted_at` datetime DEFAULT NULL,
       `helpCategory_id` bigint NOT NULL,
       `document_status` tinyint NOT NULL DEFAULT '1',  # 0:下架 1:上架
       PRIMARY KEY (`helpDocument_id`)
);

-- helpDocumentTranslation --
CREATE TABLE `helpdocumenttranslation` (
      `helpDocumentTranslation_id` bigint(20) NOT NULL,  # 文章翻譯id
      `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
      `updated_at` datetime DEFAULT NULL,
      `deleted_at` datetime DEFAULT NULL,
      `helpDocument_id` bigint NOT NULL,                 # 文章id
      `language_code` varchar(256) NOT NULL,             # 語系
      `document_title` varchar(256) NOT NULL,            # 文章標題
      `document_content` text NOT NULL,                  # 文章內容
      PRIMARY KEY (`helpDocumentTranslation_id`),
      UNIQUE KEY `idx_helpDocument_id_language_id` (`helpDocument_id`,`language_code`)
);

