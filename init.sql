CREATE TABLE IF NOT EXISTS `go_short_url_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `role` varchar(5) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'user' COMMENT 'admin, user',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE TABLE IF NOT EXISTS `go_short_url_rule` (
  `id` int NOT NULL AUTO_INCREMENT,
  `suffix` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `target` varchar(2000) COLLATE utf8mb4_general_ci NOT NULL,
  `request` int NOT NULL DEFAULT '0',
  `user_id` int NOT NULL,
  `expires` int NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `suffix` (`suffix`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `go_short_url_rule_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `go_short_url_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;