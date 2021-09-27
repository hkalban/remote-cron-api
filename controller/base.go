package controller

import (
	"github.com/gorilla/mux"
	"github.com/hkalban/remote-cron-api/scheduler"
	"github.com/hkalban/remote-cron-api/service"
)

type BaseController struct {
	basePath         string
	router           *mux.Router
	ServerController ServerController
	TaskController   TaskController
	scheduler        scheduler.BaseScheduler
}

func NewBaseController(bp string, r *mux.Router, br service.BaseService, scheduler scheduler.BaseScheduler) *BaseController {
	return &BaseController{
		basePath:         bp,
		router:           r,
		scheduler:        scheduler,
		ServerController: NewServerController(br.ServerService),
		TaskController:   NewTaskController(br.TaskService, scheduler),
	}
}
