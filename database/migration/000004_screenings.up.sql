CREATE TABLE `screenings`
(
    `id`             integer PRIMARY KEY AUTO_INCREMENT,
    `movie_id`       integer,
    `cinema_id`      integer,
    `screening_time` timestamp,
    `price`          decimal(10, 2),
    `created_at`     timestamp,
    `updated_at`     timestamp
);

ALTER TABLE `screenings`
    ADD FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `screenings`
    ADD FOREIGN KEY (`cinema_id`) REFERENCES `cinemas` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
