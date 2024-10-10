評価テーブル（ratings）
カラム名	型	説明
id	INT (AUTO_INCREMENT)	主キー
novel_id	INT	小説ID（外部キー）
user_id	INT	ユーザーID（外部キー）
rating	INT	評価（1〜5）
created_at	DATETIME	評価日時

CREATE TABLE `rating` (
    `id` int NOT NULL AUTO_INCREMENT,
    `novel_id` int NOT NULL AUTO_INCREMENT,
    `name` int NOT NULL AUTO_INCREMENT,
    `created_at` datetime NOT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `fk_novel` (`novel_id`),
    CONSTRAINT `fk_novel` FOREIGN KEY (`novel_id`) REFERENCES `novel` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;