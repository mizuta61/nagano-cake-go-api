CREATE TABLE `items` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `introduction` varchar(1000) NOT NULL,
  `price` int NOT NULL,
  `is_active` boolean NOT NULL DEFAULT true,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
  -- FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`) //Genre作ったら適用
) DEFAULT CHARSET=utf8mb4;
