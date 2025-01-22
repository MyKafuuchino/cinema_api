CREATE TABLE `audits`
(
    `id`        integer PRIMARY KEY,
    `action`    varchar(255),
    `user_id`   integer,
    `timestamp` timestamp,
    `details`   json
);

ALTER TABLE `audits`
    ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
