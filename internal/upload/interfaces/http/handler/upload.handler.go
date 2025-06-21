package handler

import (
	"github.com/anle/codebase/internal/upload/application/dto"
	"github.com/anle/codebase/internal/upload/application/usecase"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	UploadService usecase.IUploadService
}

func NewUploadHandler(UploadService usecase.IUploadService) *UploadHandler {
	return &UploadHandler{
		UploadService: UploadService,
	}
}

func (uh *UploadHandler) Upload(ctx *gin.Context) {
	var fileUploaded dto.FileUploaded
	if err := ctx.ShouldBind(&fileUploaded); err != nil {
		response.ErrorResponseExternal(ctx, response.ErrCodeExternal, err)
		return
	}

	// Read only 10Mb from request body
	// ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 10<<20)

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		response.ErrorResponseExternal(ctx, response.ErrCodeExternal, err)
		return
	}

	result, _ := uh.UploadService.Upload(ctx, fileHeader)

	response.HandleResult(ctx, result, nil)
}
