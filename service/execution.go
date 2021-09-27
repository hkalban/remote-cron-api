package service

import (
	"sync"

	"github.com/hkalban/remote-cron-api/model"
	"github.com/hkalban/remote-cron-api/repository"
)

var eOnce sync.Once

type ExecutionService interface {
	Create(execution *model.Execution) (*model.Execution, error)
	Update(execution *model.Execution) (*model.Execution, error)
	FindAll() ([]model.Execution, error)
	Find(id int) (model.Execution, error)
	FindServer(id int) (model.Server, error)
}

type executionService struct {
	executionRepository repository.ExecutionRepository
}

var executionServiceExecutionInstance *executionService

func NewExecutionService(r repository.ExecutionRepository) *executionService {
	eOnce.Do(func() {
		executionServiceExecutionInstance = &executionService{
			executionRepository: r,
		}
	})
	return executionServiceExecutionInstance
}

func (u *executionService) Create(execution *model.Execution) (*model.Execution, error) {
	return u.executionRepository.Save(execution)
}

func (u *executionService) Update(execution *model.Execution) (*model.Execution, error) {
	return u.executionRepository.Update(execution)
}

func (u *executionService) FindAll() ([]model.Execution, error) {
	return u.executionRepository.FindAll()
}

func (u *executionService) Find(id int) (model.Execution, error) {
	return u.executionRepository.Find(id)
}

func (u *executionService) FindServer(id int) (model.Server, error) {
	return u.executionRepository.FindServer(id)
}
