-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS users
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    name         VARCHAR(255) NOT NULL,
    mail_address VARCHAR(255) NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NULL,
    deleted_at   TIMESTAMP    NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;