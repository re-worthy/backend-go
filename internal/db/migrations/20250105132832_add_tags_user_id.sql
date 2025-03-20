-- +goose Up
-- +goose StatementBegin
ALTER TABLE tags ADD COLUMN user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tags DROP COLUMN user_id;
-- +goose StatementEnd
