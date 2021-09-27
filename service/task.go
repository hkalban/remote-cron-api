package service

import (
	"errors"
	"sync"

	"github.com/hkalban/remote-cron-api/model"
	"github.com/hkalban/remote-cron-api/repository"
)

var tsOnce sync.Once

// TaskService : represent the task's services
type TaskService interface {
	Validate(task *model.Task) error
	Create(task *model.Task) (*model.Task, error)
	FindAll() ([]model.Task, error)
	Find(id int) (model.Task, error)
	FindServer(id int) (model.Server, error)
	FindLatestTaskExecutions(task_id int) ([]model.Execution, error)
}

type taskService struct {
	taskRepository repository.TaskRepository
}

var taskServiceInstance *taskService

//NewTaskService: construction function, injected by task repository
func NewTaskService(r repository.TaskRepository) *taskService {
	tsOnce.Do(func() {
		taskServiceInstance = &taskService{
			taskRepository: r,
		}
	})
	return taskServiceInstance
}

func (*taskService) Validate(task *model.Task) error {
	if task == nil {
		err := errors.New("The task is empty")
		return err
	}
	if task.Name == "" {
		err := errors.New("The name of task is empty")
		return err
	}
	if task.Command == "" {
		err := errors.New("The Command of task is empty")
		return err
	}
	if task.IntervalSeconds == 0 {
		err := errors.New("The Interval in Seconds of task is empty")
		return err
	}
	// if task.TimeoutSeconds == 0 {
	// 	err := errors.New("The Timeout in Seconds of task is empty")
	// 	return err
	// }
	if len(task.Servers) < 1 {
		err := errors.New("The Servers of task is empty")
		return err
	}
	return nil
}

func (u *taskService) Create(task *model.Task) (*model.Task, error) {
	return u.taskRepository.Save(task)
}

func (u *taskService) FindAll() ([]model.Task, error) {
	return u.taskRepository.FindAll()
}

func (u *taskService) Find(id int) (model.Task, error) {
	return u.taskRepository.Find(id)
}

func (u *taskService) FindServer(id int) (model.Server, error) {
	return u.taskRepository.FindServer(id)
}

func (u *taskService) FindLatestTaskExecutions(task_id int) ([]model.Execution, error) {
	return u.taskRepository.FindLatestTaskExecutions(task_id)
}
