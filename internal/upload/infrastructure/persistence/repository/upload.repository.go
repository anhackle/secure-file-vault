package repository

import (
	"context"
	"database/sql"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/upload/domain/repository"
)

type UploadRepository struct {
	queries *database.Queries
}

func NewUploadRepository(dbConn *sql.DB) repository.IUploadRepository {
	return &UploadRepository{
		queries: database.New(dbConn),
	}
}

func (ur *UploadRepository) Register(ctx context.Context) (err error) {
	panic("unimplemented")
}
