package repository

import (
	"errors"

	"github.com/hkalban/remote-cron-api/model"
	"gorm.io/gorm"
)

type taskExecutionRepository struct {
	DB *gorm.DB
}

type TaskExecutionRepository interface {
	Save(taskExecution *model.TaskExecution) (*model.TaskExecution, error)
	FindAll() ([]model.TaskExecution, error)
	Find(id int) (model.TaskExecution, error)
	Delete(taskExecution *model.TaskExecution) error
	Migrate() error
}

func NewTaskExecutionRepository(db *gorm.DB) TaskExecutionRepository {
	return &taskExecutionRepository{
		DB: db,
	}
}

func (u *taskExecutionRepository) Save(taskExecution *model.TaskExecution) (*model.TaskExecution, error) {
	return taskExecution, u.DB.Create(taskExecution).Error
}

func (u *taskExecutionRepository) FindAll() ([]model.TaskExecution, error) {
	var tasks []model.TaskExecution
	err := u.DB.Preload("Executions").Find(&tasks).Error
	return tasks, err
}

func (u *taskExecutionRepository) Find(id int) (model.TaskExecution, error) {
	var taskExecution model.TaskExecution
	result := u.DB.Preload("Executions").Find(&taskExecution, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return taskExecution, result.Error
	}

	return taskExecution, nil
}

func (u *taskExecutionRepository) Delete(taskExecution *model.TaskExecution) error {
	return u.DB.Delete(&taskExecution).Error
}

func (u *taskExecutionRepository) Migrate() error {
	return u.DB.AutoMigrate(&model.TaskExecution{})
}
