-- Obsess user definition

CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `email` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ユーザーのメールアドレス',
    `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ユーザーのパスワード',
    `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'ユーザーの名前',
    `profile_image` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'ユーザーのプロフィール画像',
    `email_verified` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'メール認証が済んでいるかどうか',
    `created_at` datetime NOT NULL 'ユーザーの登録日時',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 'ユーザー情報の更新日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_email` (`email`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT 'user情報(id = 0はメールアドレスのない仮アカウント)';