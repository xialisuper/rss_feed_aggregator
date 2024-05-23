-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- 创建 feed_follows 表并添加唯一约束
CREATE TABLE feed_follows (
  id UUID PRIMARY KEY,
  feed_id UUID NOT NULL,
  user_id UUID NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_feed
    FOREIGN KEY(feed_id) 
      REFERENCES feeds(id)
      ON DELETE CASCADE,
  CONSTRAINT fk_user
    FOREIGN KEY(user_id) 
      REFERENCES users(id)
      ON DELETE CASCADE,
  CONSTRAINT unique_feed_user UNIQUE (feed_id, user_id)
);

-- 创建触发器函数
CREATE OR REPLACE FUNCTION update_feed_follows_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 创建触发器
CREATE TRIGGER update_feed_follows_updated_at_trigger
BEFORE UPDATE ON feed_follows
FOR EACH ROW
EXECUTE FUNCTION update_feed_follows_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
