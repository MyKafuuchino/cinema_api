CREATE TABLE `movies`
(
    `id`           integer PRIMARY KEY,
    `title`        varchar(100),
    `description`  text,
    `genre`        varchar(50),
    `duration`     integer,
    `release_date` date,
    `created_at`   timestamp,
    `updated_at`   timestamp
);