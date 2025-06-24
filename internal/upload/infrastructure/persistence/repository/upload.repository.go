package repository

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/upload/domain/entity"
	"github.com/anle/codebase/internal/upload/domain/repository"
	"github.com/minio/minio-go/v7"
)

type UploadRepository struct {
	queries  *database.Queries
	s3Client *minio.Client
}

func NewUploadRepository(dbConn *sql.DB, s3Client *minio.Client) repository.IUploadRepository {
	return &UploadRepository{
		queries:  database.New(dbConn),
		s3Client: s3Client,
	}
}

// SaveMetadata implements repository.IUploadRepository.
func (ur *UploadRepository) SaveMetadata(ctx context.Context, metadata *entity.MetadataUploadedFile) (err error) {
	_, err = ur.queries.CreateMetadata(ctx, database.CreateMetadataParams{
		ID:           metadata.ID,
		OriginalName: metadata.OriginalName,
		S3Key:        metadata.S3Key,
		MimeType:     metadata.MimeType,
		FileSize:     metadata.FileSize,
		CreatedAt:    sql.NullTime{Time: metadata.CreatedAt, Valid: true},
		ExpiredAt:    sql.NullTime{Time: metadata.ExpiredAt, Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}

// UploadFileToS3 implements repository.IUploadRepository.
func (ur *UploadRepository) UploadFileToS3(ctx context.Context, bucket, s3Key, dstPath string) (err error) {
	_, err = ur.s3Client.FPutObject(ctx, bucket, s3Key, dstPath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// DeleteMetadata implements repository.IUploadRepository.
func (ur *UploadRepository) DeleteMetadata(ctx context.Context, fileID string) (err error) {
	panic("unimplemented")
}
