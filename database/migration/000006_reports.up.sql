CREATE TABLE `reports`
(
    `id`          integer PRIMARY KEY,
    `report_name` varchar(100),
    `content`     json,
    `created_at`  timestamp
);