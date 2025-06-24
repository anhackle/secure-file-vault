-- name: GetExpiredMetadata :many
SELECT
    id,
    s3_key,
    expired_at
FROM `uploaded_files`
WHERE is_deleted = 0 AND expired_at < NOW();

-- name: CreateMetadata :execresult
INSERT INTO `uploaded_files` (
    id,
    original_name,
    s3_key,
    mime_type,
    file_size,
    created_at,
    expired_at
)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: DeleteMetadata :execresult
UPDATE `uploaded_files`
SET
    is_deleted = 1
WHERE id = ?;