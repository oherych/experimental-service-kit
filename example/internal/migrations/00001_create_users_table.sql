-- +goose Up
CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username text NOT NULL,
       email text NOT NULL
);

INSERT INTO users(username, email) VALUES
    ('root', 'root@example.com'),
    ('manager', 'manager@example.com');

-- +goose Down
DROP TABLE users;