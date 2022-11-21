CREATE TABLE `articles` (
  `id`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `title`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `body`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `url`      varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;