CREATE TABLE `post` (
    `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created`       bigint(20) NOT NULL,
    `rating`        varchar(25) CHARACTER SET utf8 DEFAULT NULL,
    `title`         varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '',
    `article`       varchar(8000) CHARACTER SET utf8 NOT NULL DEFAULT '',
    `deleted_at`    datetime NOT NULL,
    `created_at`    datetime NOT NULL,
    `updated_at`    datetime NOT NULL,
    PRIMARY KEY     (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
