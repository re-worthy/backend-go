-- +goose Up
-- +goose StatementBegin
ALTER TABLE transactions RENAME COLUMN createdAt TO created_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions RENAME COLUMN created_at TO createdAt;
-- +goose StatementEnd
