package jobs

import (
	"github.com/go-co-op/gocron"
	"time"
)

type Jobs struct {
	Enabled bool
}

var cron *gocron.Scheduler

var enabled bool

//GetJobService will return the instance of the job service
func GetJobService() (*gocron.Scheduler, bool) {
	if enabled {
		return cron, true
	}
	return cron, false
}

func (j *Jobs) InitCron() {
	cron = gocron.NewScheduler(time.UTC)
	cron.StartAsync()
	j.Enabled = true
	enabled = j.Enabled
}
