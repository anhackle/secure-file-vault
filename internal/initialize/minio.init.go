package initialize

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

func InitMinio() {
	client, err := minio.New(
		fmt.Sprintf("%s:%d", global.Config.Minio.Host, global.Config.Minio.Port),
		&minio.Options{
			Creds:  credentials.NewStaticV4(global.Config.Minio.Username, global.Config.Minio.Password, ""),
			Secure: false, // false if using http://
		},
	)
	if err != nil {
		global.Logger.Error("Minio initialization error", zap.Error(err))
	}

	global.MinioClient = client
}
