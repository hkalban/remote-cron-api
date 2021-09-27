package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BaseRepository struct {
	ServerRepository        ServerRepository
	TaskRepository          TaskRepository
	TaskExecutionRepository TaskExecutionRepository
	ExecutionRepository     ExecutionRepository
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		ServerRepository:        NewServerRepository(db),
		TaskRepository:          NewTaskRepository(db),
		TaskExecutionRepository: NewTaskExecutionRepository(db),
		ExecutionRepository:     NewExecutionRepository(db),
	}
}

func Connect() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=remote_cron port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func (u *BaseRepository) CreateSchema() (*BaseRepository, error) {
	if srErr := u.ServerRepository.Migrate(); srErr != nil {
		return u, srErr
	}
	if trErr := u.TaskRepository.Migrate(); trErr != nil {
		return u, trErr
	}
	if terErr := u.TaskExecutionRepository.Migrate(); terErr != nil {
		return u, terErr
	}
	if erErr := u.ExecutionRepository.Migrate(); erErr != nil {
		return u, erErr
	}

	return u, nil
}
