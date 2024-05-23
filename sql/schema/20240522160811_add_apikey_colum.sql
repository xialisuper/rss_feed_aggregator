-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- Step 1: Add api_key column without NOT NULL constraint
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) UNIQUE;

-- Step 2: Generate unique api_key for existing users
UPDATE users
SET api_key = encode(sha256(random()::text::bytea), 'hex')
WHERE api_key IS NULL;

-- Step 3: Set api_key column to NOT NULL and create trigger function and trigger
ALTER TABLE users
ALTER COLUMN api_key SET NOT NULL;

CREATE OR REPLACE FUNCTION generate_api_key()
RETURNS TRIGGER AS $$
BEGIN
  NEW.api_key := encode(sha256(random()::text::bytea), 'hex');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_api_key
BEFORE INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION generate_api_key();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
