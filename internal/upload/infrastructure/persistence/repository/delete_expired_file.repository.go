package repository

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/upload/domain/repository"
	"github.com/minio/minio-go/v7"
)

type DeleteExpiredFileRepository struct {
	queries  *database.Queries
	s3Client *minio.Client
}

func NewDeleteExpiredFileRepository(dbConn *sql.DB, s3Client *minio.Client) repository.IDeleteExpiredFileRepository {
	return &DeleteExpiredFileRepository{
		queries:  database.New(dbConn),
		s3Client: s3Client,
	}
}

// GetExpiredFiles implements repository.IDeleteExpiredFileRepository.
func (dfr *DeleteExpiredFileRepository) GetExpiredFiles(ctx context.Context) (expiredFiles []database.GetExpiredMetadataRow, err error) {
	expiredFiles, err = dfr.queries.GetExpiredMetadata(ctx)
	if err != nil {
		return []database.GetExpiredMetadataRow{}, err
	}

	return expiredFiles, nil
}

// DeleteMetadata implements repository.IDeleteExpiredFileRepository.
func (dfr *DeleteExpiredFileRepository) DeleteMetadata(ctx context.Context, fileID string) (err error) {
	_, err = dfr.queries.DeleteMetadata(ctx, fileID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteS3File implements repository.IDeleteExpiredFileRepository.
func (dfr *DeleteExpiredFileRepository) DeleteS3File(ctx context.Context, bucket string, S3Key string) (err error) {
	err = dfr.s3Client.RemoveObject(ctx, bucket, S3Key, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
