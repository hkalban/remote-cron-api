package repository

import (
	"errors"

	"github.com/hkalban/remote-cron-api/model"
	"gorm.io/gorm"
)

type taskRepository struct {
	DB *gorm.DB
}

type TaskRepository interface {
	Save(task *model.Task) (*model.Task, error)
	FindAll() ([]model.Task, error)
	Find(id int) (model.Task, error)
	FindServer(id int) (model.Server, error)
	FindLatestTaskExecutions(task_id int) ([]model.Execution, error)
	Delete(task *model.Task) error
	Migrate() error
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		DB: db,
	}
}

func (u *taskRepository) Save(task *model.Task) (*model.Task, error) {
	return task, u.DB.Create(task).Error
}

func (u *taskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := u.DB.Preload("Servers").Find(&tasks).Error
	return tasks, err
}

func (u *taskRepository) Find(id int) (model.Task, error) {
	var task model.Task
	result := u.DB.Preload("Servers").Preload("TaskExecutions").Find(&task, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return task, result.Error
	}

	return task, nil
}

func (u *taskRepository) FindLatestTaskExecutions(task_id int) ([]model.Execution, error) {
	var taskExecution model.TaskExecution
	result := u.DB.Order("created_at desc").Preload("Execution").Offset(1).First(&taskExecution, "task_id = ?", task_id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return taskExecution.Execution, nil
}

func (u *taskRepository) FindServer(id int) (model.Server, error) {
	var server model.Server
	result := u.DB.Find(&server, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return server, result.Error
	}

	return server, nil
}

func (u *taskRepository) Delete(task *model.Task) error {
	return u.DB.Delete(&task).Error
}

func (u *taskRepository) Migrate() error {
	return u.DB.AutoMigrate(&model.Task{})
}
