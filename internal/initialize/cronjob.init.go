package initialize

import "github.com/robfig/cron/v3"

func InitCronJob() *cron.Cron {
	c := cron.New(cron.WithSeconds())
	return c
}
