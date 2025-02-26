-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens (
    id SERIAL UNIQUE NOT NULL,
    token VARCHAR(2048) NOT NULL,
    refresh_token VARCHAR(2048) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
