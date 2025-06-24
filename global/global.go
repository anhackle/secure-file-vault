package global

import (
	"database/sql"

	"github.com/anle/codebase/pkg/logger"
	"github.com/anle/codebase/setting"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

var (
	Config      setting.Config
	Logger      *logger.LoggerZap
	Rdb         *redis.Client
	Mdb         *sql.DB
	MinioClient *minio.Client
	Cron        *cron.Cron
)
