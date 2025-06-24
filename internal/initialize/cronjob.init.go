package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/robfig/cron/v3"
)

func InitCronJob() *cron.Cron {
	c := cron.New(cron.WithSeconds())
	global.Cron = c
	c.Start()

	return c
}
