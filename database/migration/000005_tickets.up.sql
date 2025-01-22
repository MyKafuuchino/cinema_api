CREATE TABLE `tickets`
(
    `id`           integer PRIMARY KEY,
    `user_id`      integer,
    `screening_id` integer,
    `seat_number`  varchar(10),
    `status`       enum ('booked','paid'),
    `created_at`   timestamp,
    `updated_at`   timestamp
);

ALTER TABLE `tickets`
    ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `tickets`
    ADD FOREIGN KEY (`screening_id`) REFERENCES `screenings` (`id`);
