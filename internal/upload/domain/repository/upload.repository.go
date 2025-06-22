package repository

import (
	"context"

	"github.com/anle/codebase/internal/upload/domain/entity"
)

type IUploadRepository interface {
	SaveMetadata(ctx context.Context, metadata *entity.MetadataUploadedFile) (err error)
	UploadFileToS3(ctx context.Context, bucket, s3Key, dstPath string) (err error)
}
