package usecase

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/upload/application/service"
	"github.com/anle/codebase/internal/upload/domain/repository"
	"github.com/anle/codebase/response"
)

type IShareService interface {
	Share(ctx context.Context, fileID string) (url string, result int, err error)
	ProcessSharedURL(ctx context.Context, fileID string) (mimeType string, value any, result int, err error)
}

type ShareService struct {
	ShareRepository repository.IShareRepository
}

func NewShareService(ShareRepository repository.IShareRepository) IShareService {
	return &ShareService{
		ShareRepository: ShareRepository,
	}
}

// Share implements IShareService.
func (s *ShareService) Share(ctx context.Context, fileID string) (url string, result int, err error) {
	metadata, err := s.ShareRepository.GetMetadata(ctx, fileID)
	if err != nil && err != sql.ErrNoRows {
		return "", response.ErrCodeInternal, err
	}

	if err != nil && err == sql.ErrNoRows {
		return "", response.ErrCodeFileNotFound, err
	}

	temporaryPath := service.GenerateUUID()
	object, err := s.ShareRepository.GetS3File(ctx, "uploaded-files", metadata.S3Key)
	if err != nil {
		return "", response.ErrCodeInternal, err
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, object); err != nil {
		log.Fatalln(err)
	}

	fileContent, err := service.DecryptGCM(buf.String(), []byte(global.Config.Upload.Key))
	if err != nil {
		return "", response.ErrCodeInternal, err
	}

	err = s.ShareRepository.SaveSharedURLRedis(ctx, temporaryPath, fileContent, metadata.MimeType)
	if err != nil {
		return "", response.ErrCodeExternal, err
	}

	return fmt.Sprintf("/share/%s", temporaryPath), response.ErrCodeSuccess, nil
}

func (s *ShareService) ProcessSharedURL(ctx context.Context, fileID string) (mimeType string, fileContent any, result int, err error) {
	fileContent, err = s.ShareRepository.GetFileContent(ctx, fileID, "content")
	if err != nil {
		return "", nil, response.ErrCodeExternal, err
	}

	mimeType, err = s.ShareRepository.GetMimeType(ctx, fileID, "mime")
	if err != nil {
		return "", nil, response.ErrCodeExternal, err
	}

	return mimeType, fileContent, response.ErrCodeSuccess, nil
}
