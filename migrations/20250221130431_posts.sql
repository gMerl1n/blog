-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id SERIAL UNIQUE NOT NULL,
    title VARCHAR(255) UNIQUE NOT NULL,
    body VARCHAR(255) UNIQUE NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
