package usecase

import (
	"context"

	"github.com/anle/codebase/internal/upload/domain/repository"
)

type IDeleteExpiredFileService interface {
	DeleteExpiredFiles(ctx context.Context) (err error)
}

type DeleteExpiredFileService struct {
	DeleteExpiredFileRepository repository.IDeleteExpiredFileRepository
}

func NewDeleteExpiredFileService(deleteExpiredFileRepository repository.IDeleteExpiredFileRepository) IDeleteExpiredFileService {
	return &DeleteExpiredFileService{
		DeleteExpiredFileRepository: deleteExpiredFileRepository,
	}
}

// DeleteExpiredFiles implements IDeleteExpiredFileService.
func (dfs *DeleteExpiredFileService) DeleteExpiredFiles(ctx context.Context) (err error) {
	expiredFiles, err := dfs.DeleteExpiredFileRepository.GetExpiredFiles(ctx)
	for _, file := range expiredFiles {
		if err := dfs.DeleteExpiredFileRepository.DeleteMetadata(ctx, file.ID); err != nil {
			return err
		}
		if err := dfs.DeleteExpiredFileRepository.DeleteS3File(ctx, "uploaded-files", file.S3Key); err != nil {
			return err
		}
	}

	return nil
}
