package service

import "github.com/hkalban/remote-cron-api/repository"

type BaseService struct {
	ServerService        ServerService
	TaskService          TaskService
	TaskExecutionService TaskExecutionService
	ExecutionService     ExecutionService
}

func NewBaseService(baseRepo repository.BaseRepository) *BaseService {
	return &BaseService{
		ServerService:        NewServerService(baseRepo.ServerRepository),
		TaskService:          NewTaskService(baseRepo.TaskRepository),
		TaskExecutionService: NewTaskExecutionService(baseRepo.TaskExecutionRepository),
		ExecutionService:     NewExecutionService(baseRepo.ExecutionRepository),
	}
}
