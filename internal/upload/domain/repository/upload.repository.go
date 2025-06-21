package repository

import "context"

type IUploadRepository interface {
	Register(ctx context.Context) (err error)
}
