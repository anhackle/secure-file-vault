package handler

import (
	"errors"

	"github.com/anle/codebase/internal/upload/application/dto"
	"github.com/anle/codebase/internal/upload/application/usecase"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type ShareHandler struct {
	ShareService usecase.IShareService
}

func NewShareHandler(ShareService usecase.IShareService) *ShareHandler {
	return &ShareHandler{
		ShareService: ShareService,
	}
}

func (uh *ShareHandler) Share(ctx *gin.Context) {
	var fileID dto.FileID
	if err := ctx.ShouldBindJSON(&fileID); err != nil {
		response.ErrorResponseExternal(ctx, response.ErrCodeExternal, err)
		return
	}

	url, result, _ := uh.ShareService.Share(ctx, fileID.FileID)

	response.HandleResult(ctx, result, url)
}

func (uh *ShareHandler) ProcessSharedURL(ctx *gin.Context) {
	var fileID dto.ShareUri
	if err := ctx.ShouldBindUri(&fileID); err != nil {
		response.ErrorResponseExternal(ctx, response.ErrCodeExternal, err)
		return
	}
	mimeType, fileContent, _, err := uh.ShareService.ProcessSharedURL(ctx, fileID.ID)
	if err != nil {
		response.ErrorResponseInternal(ctx, response.ErrCodeInternal, err)
		return
	}

	str, ok := fileContent.(string)
	if !ok {
		response.ErrorResponseInternal(ctx, response.ErrCodeInternal, errors.New("invalid file data"))
		return
	}

	data := []byte(str)
	ctx.Data(200, mimeType, data)
}
