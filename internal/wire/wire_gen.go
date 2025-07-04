// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"database/sql"
	"github.com/anle/codebase/internal/upload/application/usecase"
	"github.com/anle/codebase/internal/upload/infrastructure/persistence/repository"
	"github.com/anle/codebase/internal/upload/interfaces/cron"
	"github.com/anle/codebase/internal/upload/interfaces/http/handler"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

// Injectors from share.wire.go:

func InitShareRouterHandler(dbc *sql.DB, s3Client *minio.Client, redisClient *redis.Client) (*handler.ShareHandler, error) {
	iShareRepository := repository.NewShareRepository(dbc, s3Client, redisClient)
	iShareService := usecase.NewShareService(iShareRepository)
	shareHandler := handler.NewShareHandler(iShareService)
	return shareHandler, nil
}

// Injectors from upload.wire.go:

func InitUploadRouterHandler(dbc *sql.DB, s3Client *minio.Client) (*handler.UploadHandler, error) {
	iUploadRepository := repository.NewUploadRepository(dbc, s3Client)
	iUploadService := usecase.NewUploadService(iUploadRepository)
	uploadHandler := handler.NewUploadHandler(iUploadService)
	return uploadHandler, nil
}

func InitDeleteExpiredFileService(dbc *sql.DB, s3Client *minio.Client) (*cron.DeleteExpiredFileCronHanlder, error) {
	iDeleteExpiredFileRepository := repository.NewDeleteExpiredFileRepository(dbc, s3Client)
	iDeleteExpiredFileService := usecase.NewDeleteExpiredFileService(iDeleteExpiredFileRepository)
	deleteExpiredFileCronHanlder := cron.NewDeleteExpiredFileCronHanlder(iDeleteExpiredFileService)
	return deleteExpiredFileCronHanlder, nil
}
