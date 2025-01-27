CREATE TABLE `users`
(
    `id`         integer PRIMARY KEY AUTO_INCREMENT,
    `full_name`  varchar(100),
    `email`      varchar(100) UNIQUE,
    `password`   varchar(255),
    `role`       enum ('USER','ADMIN') DEFAULT 'USER',
    `created_at` timestamp,
    `updated_at` timestamp
);