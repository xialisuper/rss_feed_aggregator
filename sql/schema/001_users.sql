-- +goose Up
CREATE TABLE users (
  id   UUID PRIMARY KEY,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;
