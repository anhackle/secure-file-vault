package repository

import (
	"context"

	"github.com/anle/codebase/internal/upload/domain/entity"
	"github.com/minio/minio-go/v7"
)

type IShareRepository interface {
	GetMetadata(ctx context.Context, fileID string) (metadata *entity.MetadataUploadedFile, err error)
	GetS3File(ctx context.Context, bucket string, S3Key string) (object *minio.Object, err error)
	SaveSharedURLRedis(ctx context.Context, key string, value any, mimeType string) (err error)
	GetMimeType(ctx context.Context, key string, field string) (mimeType string, err error)
	GetFileContent(ctx context.Context, key string, field string) (fileContent any, err error)
}
