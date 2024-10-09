-- Obsess user definition

CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `email` varchar(255) NOT NULL DEFAULT '',
    `password_hash` varchar(255) NOT NULL DEFAULT '',
    `name` varchar(255) NOT NULL DEFAULT '',
    `profile_image` varchar(255) DEFAULT '',
    `email_verified` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'メール認証が済んでいれば1,そうでなければ0',
    `plan` varchar(255) utf8mb4_general_ci DEFAULT '' COMMENT 'ユーザーがどの有料プランに入っているかを表す文字列',
    `created_at` datetime NOT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_email` (`email`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT 'user情報(id = 0はメールアドレスのない仮アカウント)';