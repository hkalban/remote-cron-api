package repository

import (
	"errors"

	"github.com/hkalban/remote-cron-api/model"
	"gorm.io/gorm"
)

type executionRepository struct {
	DB *gorm.DB
}

type ExecutionRepository interface {
	Save(execution *model.Execution) (*model.Execution, error)
	Update(execution *model.Execution) (*model.Execution, error)
	FindAll() ([]model.Execution, error)
	Find(id int) (model.Execution, error)
	FindServer(id int) (model.Server, error)
	Delete(execution *model.Execution) error
	Migrate() error
}

func NewExecutionRepository(db *gorm.DB) ExecutionRepository {
	return &executionRepository{
		DB: db,
	}
}

func (u *executionRepository) Save(execution *model.Execution) (*model.Execution, error) {
	return execution, u.DB.Create(execution).Error
}

func (u *executionRepository) Update(execution *model.Execution) (*model.Execution, error) {
	return execution, u.DB.Save(execution).Error
}

func (u *executionRepository) FindAll() ([]model.Execution, error) {
	var executions []model.Execution
	err := u.DB.Find(&executions).Error
	return executions, err
}

func (u *executionRepository) Find(id int) (model.Execution, error) {
	var execution model.Execution
	result := u.DB.Find(&execution, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return execution, result.Error
	}

	return execution, nil
}

func (u *executionRepository) FindServer(id int) (model.Server, error) {
	var server model.Server
	result := u.DB.Find(&server, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return server, result.Error
	}

	return server, nil
}

func (u *executionRepository) Delete(execution *model.Execution) error {
	return u.DB.Delete(&execution).Error
}

func (u *executionRepository) Migrate() error {
	return u.DB.AutoMigrate(&model.Execution{})
}
