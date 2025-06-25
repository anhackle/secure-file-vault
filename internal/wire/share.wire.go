//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/anle/codebase/internal/upload/application/usecase"
	"github.com/anle/codebase/internal/upload/infrastructure/persistence/repository"
	"github.com/anle/codebase/internal/upload/interfaces/http/handler"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
)

func InitShareRouterHandler(dbc *sql.DB, s3Client *minio.Client) (*handler.ShareHandler, error) {
	wire.Build(
		repository.NewShareRepository,
		usecase.NewShareService,
		handler.NewShareHandler,
	)

	return new(handler.ShareHandler), nil
}
