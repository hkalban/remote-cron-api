package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/hkalban/remote-cron-api/service"
)

// BaseScheduler will hold everything that controller needs
type BaseScheduler struct {
	scheduler *gocron.Scheduler
	services  service.BaseService
}

// NewBaseScheduler returns a new BaseHandler
func NewBaseScheduler(services service.BaseService) *BaseScheduler {
	return &BaseScheduler{
		scheduler: gocron.NewScheduler(time.UTC),
		services:  services,
	}
}
