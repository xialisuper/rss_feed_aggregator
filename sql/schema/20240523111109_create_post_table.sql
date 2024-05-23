-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  -- 确保 UUID 生成函数可用

CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- 使用 UUID 并自动生成
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255),
    url VARCHAR(255) UNIQUE,
    description TEXT,
    published_at TIMESTAMP WITH TIME ZONE,
    feed_id UUID,  -- 使用 UUID 类型
    FOREIGN KEY (feed_id) REFERENCES feeds(id)  -- 确保外键也是 UUID 类型
);

CREATE OR REPLACE FUNCTION update_published_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_published_at_trigger
BEFORE UPDATE ON posts
FOR EACH ROW
EXECUTE PROCEDURE update_published_at();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
