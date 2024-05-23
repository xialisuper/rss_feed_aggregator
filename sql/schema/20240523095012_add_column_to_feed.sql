-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE feeds
ADD COLUMN last_fetched_at time;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
