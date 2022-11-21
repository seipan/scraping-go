CREATE TABLE `articles` (
  `id`       int          COLLATE utf8mb4_bin NOT NULL AUTO_INCREMENT,
  `title`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `img_url`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `content`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;