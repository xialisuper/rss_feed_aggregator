-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds
ADD COLUMN last_fetched_at timestamptz;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
