package repository

import (
	"context"

	"github.com/anle/codebase/internal/database"
)

type IDeleteExpiredFileRepository interface {
	GetExpiredFiles(ctx context.Context) (expiredFiles []database.GetExpiredMetadataRow, err error)
	DeleteMetadata(ctx context.Context, fileID string) (err error)
	DeleteS3File(ctx context.Context, bucket, S3Key string) (err error)
}
