package dto

import "mime/multipart"

type FileUploaded struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
