-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS uploaded_files (
    id CHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    original_name TEXT NOT NULL,
    s3_key TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    file_size BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    download_count INT NOT NULL DEFAULT 5,
    is_deleted INT NOT NULL DEFAULT FALSE,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS uploaded_files;
-- +goose StatementEnd
