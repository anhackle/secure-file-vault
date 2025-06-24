//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/upload/application/usecase"
	"github.com/anle/codebase/internal/upload/infrastructure/persistence/repository"
	"github.com/anle/codebase/internal/upload/interfaces/cron"
	"github.com/anle/codebase/internal/upload/interfaces/http/handler"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
)

func InitUploadRouterHandler(dbc *sql.DB, s3Client *minio.Client) (*handler.UploadHandler, error) {
	wire.Build(
		repository.NewUploadRepository,
		usecase.NewUploadService,
		handler.NewUploadHandler,
	)

	return new(handler.UploadHandler), nil
}

func InitDeleteExpiredFileService(dbc *sql.DB, s3Client *minio.Client) (*cron.DeleteExpiredFileCronHanlder, error) {
	wire.Build(
		repository.NewDeleteExpiredFileRepository,
		usecase.NewDeleteExpiredFileService,
		cron.NewDeleteExpiredFileCronHanlder,
	)

	return new(cron.DeleteExpiredFileCronHanlder), nil
}
