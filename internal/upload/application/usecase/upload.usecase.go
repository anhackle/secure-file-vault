package usecase

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/upload/application/service"
	"github.com/anle/codebase/internal/upload/domain/entity"
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
	var (
		fileUUID = service.GenerateUUID()
		S3Key    = fileUUID + service.GetFileExtension(fileHeader.Filename)
	)

	dstPath := filepath.Join("/tmp", fileUUID+service.GetFileExtension(fileHeader.Filename))
	dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY, 0444)
	if err != nil {
		return response.ErrCodeInternal, err
	}
	defer dst.Close()

	_, err = io.Copy(dst, strings.NewReader(encryptedFileName))
	if err != nil {
		return response.ErrCodeInternal, err
	}

	// Save metadata to the database
	err = us.UploadRepository.SaveMetadata(ctx, &entity.MetadataUploadedFile{
		ID:           fileUUID,
		OriginalName: filepath.Base(fileHeader.Filename),
		S3Key:        S3Key,
		MimeType:     fileHeader.Header.Get("Content-Type"),
		FileSize:     fileHeader.Size,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().AddDate(0, 0, 30),
	})
	if err != nil {
		return response.ErrCodeInternal, err
	}

	// Upload the file to S3
	err = us.UploadRepository.UploadFileToS3(ctx, "uploaded-files", S3Key, dstPath)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}
