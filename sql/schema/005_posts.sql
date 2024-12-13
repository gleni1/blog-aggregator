-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title VARCHAR NOT NULL,
    url VARCHAR UNIQUE,
    description VARCHAR NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
