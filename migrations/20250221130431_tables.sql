-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email         VARCHAR(255) NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id SERIAL UNIQUE NOT NULL,
    user_id  INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) UNIQUE NOT NULL,
    body VARCHAR(255) UNIQUE NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tokens (
    id SERIAL UNIQUE NOT NULL,
    token VARCHAR(4048) NOT NULL,
    refresh_token VARCHAR(4048) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
