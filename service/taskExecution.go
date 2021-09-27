package service

import (
	"sync"

	"github.com/hkalban/remote-cron-api/model"
	"github.com/hkalban/remote-cron-api/repository"
)

var tseOnce sync.Once

type TaskExecutionService interface {
	Create(task *model.TaskExecution) (*model.TaskExecution, error)
	FindAll() ([]model.TaskExecution, error)
	Find(id int) (model.TaskExecution, error)
}

type taskExecutionService struct {
	taskExecutionRepository repository.TaskExecutionRepository
}

var taskExecutionServiceExecutionInstance *taskExecutionService

func NewTaskExecutionService(r repository.TaskExecutionRepository) *taskExecutionService {
	tseOnce.Do(func() {
		taskExecutionServiceExecutionInstance = &taskExecutionService{
			taskExecutionRepository: r,
		}
	})
	return taskExecutionServiceExecutionInstance
}

func (u *taskExecutionService) Create(task *model.TaskExecution) (*model.TaskExecution, error) {
	return u.taskExecutionRepository.Save(task)
}

func (u *taskExecutionService) FindAll() ([]model.TaskExecution, error) {
	return u.taskExecutionRepository.FindAll()
}

func (u *taskExecutionService) Find(id int) (model.TaskExecution, error) {
	return u.taskExecutionRepository.Find(id)
}
