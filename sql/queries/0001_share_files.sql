-- name: GetMetadataByID :one
SELECT
    id,
    s3_key,
    mime_type
FROM `uploaded_files`
WHERE is_deleted = 0 AND id = ?;