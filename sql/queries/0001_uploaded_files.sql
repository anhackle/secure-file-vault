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
DELETE FROM `uploaded_files`
WHERE id = ?;