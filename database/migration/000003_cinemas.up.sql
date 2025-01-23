CREATE TABLE `cinemas`
(
    `id`         integer PRIMARY KEY AUTO_INCREMENT,
    `name`       varchar(100),
    `location`   varchar(255),
    `created_at` timestamp,
    `updated_at` timestamp
);