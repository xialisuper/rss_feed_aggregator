-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  -- 确保 UUID 生成函数可用

CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- 使用 UUID 并自动生成
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    feed_id UUID NOT NULL,  -- 使用 UUID 类型
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
DROP TABLE posts;
DROP FUNCTION update_published_at();
DROP EXTENSION "uuid-ossp";
-- +goose StatementEnd
