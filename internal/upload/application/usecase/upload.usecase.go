package usecase

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/upload/application/service"
	"github.com/anle/codebase/internal/upload/domain/repository"
	"github.com/anle/codebase/response"
)

type IUploadService interface {
	Upload(ctx context.Context, fileHeader *multipart.FileHeader) (result int, err error)
}

type UploadService struct {
	UploadRepository repository.IUploadRepository
}

func NewUploadService(UploadRepository repository.IUploadRepository) IUploadService {
	return &UploadService{
		UploadRepository: UploadRepository,
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

	// Read file content to []byte
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	// Encrypt the file content before saving
	encryptedFileName, err := service.EncryptGCM(fileContent, []byte(global.Config.Upload.Key))
	if err != nil {
		return response.ErrCodeInternal, err
	}

	// Write the encrypted file content to a new file
	dstPath := filepath.Join("/tmp", service.GenerateUUID()+filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(dstPath)
	if err != nil {
		return response.ErrCodeInternal, err
	}
	defer dst.Close()

	_, err = io.Copy(dst, strings.NewReader(encryptedFileName))
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}
