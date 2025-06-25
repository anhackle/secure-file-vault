package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/upload/domain/entity"
	"github.com/anle/codebase/internal/upload/domain/repository"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type ShareRepository struct {
	queries     *database.Queries
	s3Client    *minio.Client
	redisClient *redis.Client
}

func NewShareRepository(dbConn *sql.DB, s3Client *minio.Client, redisClient *redis.Client) repository.IShareRepository {
	return &ShareRepository{
		queries:     database.New(dbConn),
		s3Client:    s3Client,
		redisClient: redisClient,
	}
}

// GetMetadata implements repository.IShareRepository.
func (s *ShareRepository) GetMetadata(ctx context.Context, fileID string) (metadata *entity.MetadataUploadedFile, err error) {
	result, err := s.queries.GetMetadataByID(ctx, fileID)
	if err != nil {
		return nil, err
	}

	return &entity.MetadataUploadedFile{
		ID:       result.ID,
		S3Key:    result.S3Key,
		MimeType: result.MimeType,
	}, nil
}

func (s *ShareRepository) GetS3File(ctx context.Context, bucket string, S3Key string) (object *minio.Object, err error) {
	object, err = s.s3Client.GetObject(ctx, bucket, S3Key, minio.GetObjectOptions{})
	if err != nil {
		return &minio.Object{}, err
	}

	return object, nil
}

// SaveSharedURLRedis implements repository.IShareRepository.
func (s *ShareRepository) SaveSharedURLRedis(ctx context.Context, key string, value any, mimeType string) (err error) {
	s.redisClient.Watch(ctx, func(tx *redis.Tx) error {
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			err = s.redisClient.HSet(ctx, key, map[string]interface{}{
				"content": value,
				"mime":    mimeType,
			}).Err()
			if err != nil {
				return err
			}

			err = s.redisClient.Expire(ctx, key, 5*time.Minute).Err()
			if err != nil {
				return err
			}

			return nil
		})
		return nil

	}, key)

	return nil
}

// GetFileContent implements repository.IShareRepository.
func (s *ShareRepository) GetFileContent(ctx context.Context, key string, field string) (fileContent any, err error) {
	fileContent, err = s.redisClient.HGet(ctx, key, field).Result()
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// GetMimeType implements repository.IShareRepository.
func (s *ShareRepository) GetMimeType(ctx context.Context, key string, field string) (mimeType string, err error) {
	mimeType, err = s.redisClient.HGet(ctx, key, field).Result()
	if err != nil {
		return "", err
	}

	return mimeType, nil
}
