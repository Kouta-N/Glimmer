-- Obsess chapter definition

CREATE TABLE `chapter` (
    `id` int NOT NULL AUTO_INCREMENT,
    `novel_id` int NOT NULL AUTO_INCREMENT,
    `chapter_number` int NOT NULL AUTO_INCREMENT,
    `content` varchar(255) NOT NULL AUTO_INCREMENT,
    `created_at` datetime NOT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `fk_novel` (`novel_id`),
    CONSTRAINT `fk_novel` FOREIGN KEY (`novel_id`) REFERENCES `novel` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;