-- +goose Up
ALTER TABLE feeds
ADD COLUMN created_at TIMESTAMP NOT NULL;

ALTER TABLE feeds
ADD COLUMN updated_at TIMESTAMP NOT NULL;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_fetched_at;
