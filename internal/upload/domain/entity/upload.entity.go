package entity

import "time"

type MetadataUploadedFile struct {
	ID            string    `json:"id"`
	OriginalName  string    `json:"original_name"`
	S3Key         string    `json:"s3_key"`
	MimeType      string    `json:"mime_type"`
	FileSize      int64     `json:"file_size"`
	CreatedAt     time.Time `json:"created_at"`
	ExpiredAt     time.Time `json:"expired_at"`
	DownloadCount int       `json:"download_count"`
	IsDeleted     bool      `json:"is_deleted"`
}
