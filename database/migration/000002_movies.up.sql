CREATE TABLE `movies`
(
    `id`           integer PRIMARY KEY AUTO_INCREMENT,
    `title`        varchar(100),
    `slug`         varchar(100) UNIQUE,
    `description`  text,
    `genre`        varchar(50),
    `duration`     integer,
    `release_date` date,
    `created_at`   timestamp,
    `updated_at`   timestamp
);