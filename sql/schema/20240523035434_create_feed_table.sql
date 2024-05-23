-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- 创建 feeds 表
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  name VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL UNIQUE,
  user_id UUID NOT NULL,
  CONSTRAINT fk_user
    FOREIGN KEY(user_id) 
      REFERENCES users(id)
      ON DELETE CASCADE
);

-- 创建触发器函数
CREATE OR REPLACE FUNCTION update_feeds_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 创建触发器
CREATE TRIGGER update_feeds_updated_at_trigger
BEFORE UPDATE ON feeds
FOR EACH ROW
EXECUTE FUNCTION update_feeds_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
