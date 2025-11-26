package jobs

import "github.com/robfig/cron"

func Init() *cron.Cron {
	c:= cron.New()
	return c
}
