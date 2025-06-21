package usecase

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/anle/codebase/internal/upload/application/service"
	"github.com/anle/codebase/internal/upload/domain/repository"
	"github.com/anle/codebase/response"
)

type IUploadService interface {
	Upload(ctx context.Context, fileHeader *multipart.FileHeader) (result int, err error)
}

type UploadService struct {
	AuthRepository repository.IUploadRepository
}

func NewUploadService(AuthRepository repository.IUploadRepository) IUploadService {
	return &UploadService{
		AuthRepository: AuthRepository,
	}
}

func (us *UploadService) Upload(ctx context.Context, fileHeader *multipart.FileHeader) (result int, err error) {
	if fileExtension := filepath.Ext(fileHeader.Filename); !service.CheckFileExtension(fileExtension) {
		return response.ErrCodeExtensionNotAllowed, errors.New("file extension not allowed")
	}

	if fileContentType := fileHeader.Header.Get("Content-Type"); !service.CheckContentType(fileContentType) {
		return response.ErrCodeContentTypeNotAllowd, errors.New("content type not allowed")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return response.ErrCodeInternal, err
	}
	defer file.Close()

	dstPath := filepath.Join("/tmp", service.GenerateUUID()+filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(dstPath)
	if err != nil {
		return response.ErrCodeInternal, err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}
