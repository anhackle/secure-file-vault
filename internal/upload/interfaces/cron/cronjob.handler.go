package cron

import (
	"context"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/upload/application/usecase"
	"go.uber.org/zap"
)

type DeleteExpiredFileCronHanlder struct {
	DeleteExpiredFileService usecase.IDeleteExpiredFileService
}

func NewDeleteExpiredFileCronHanlder(deleteExpiredFileService usecase.IDeleteExpiredFileService) *DeleteExpiredFileCronHanlder {
	return &DeleteExpiredFileCronHanlder{
		DeleteExpiredFileService: deleteExpiredFileService,
	}
}

func (dfch *DeleteExpiredFileCronHanlder) RegisterDeleteCron() {
	// Register the cron job for deleting expired files
	global.Cron.AddFunc("*/5 * * * * *", func() {
		if err := dfch.DeleteExpiredFileService.DeleteExpiredFiles(context.Background()); err != nil {
			global.Logger.Error("Failed to delete expired files", zap.Error(err))
		} else {
			global.Logger.Info("Expired files deleted successfully")
		}
	})
}
