package repository

import (
	"errors"

	"github.com/hkalban/remote-cron-api/model"
	"gorm.io/gorm"
)

type serverRepository struct {
	DB *gorm.DB
}

type ServerRepository interface {
	Save(server *model.Server) (*model.Server, error)
	FindAll() ([]model.Server, error)
	Find(id int) (model.Server, error)
	Delete(server *model.Server) error
	Migrate() error
}

func NewServerRepository(db *gorm.DB) ServerRepository {
	return &serverRepository{
		DB: db,
	}
}

func (u *serverRepository) Save(server *model.Server) (*model.Server, error) {
	return server, u.DB.Create(server).Error
}

func (u *serverRepository) FindAll() ([]model.Server, error) {
	var servers []model.Server
	err := u.DB.Find(&servers).Error
	return servers, err
}

func (u *serverRepository) Find(id int) (model.Server, error) {
	var server model.Server
	result := u.DB.Find(&server, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return server, result.Error
	}

	return server, nil
}

func (u *serverRepository) Delete(server *model.Server) error {
	return u.DB.Delete(&server).Error
}

func (u *serverRepository) Migrate() error {
	return u.DB.AutoMigrate(&model.Server{})
}
