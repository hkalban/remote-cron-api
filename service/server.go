package service

import (
	"errors"
	"sync"

	"github.com/hkalban/remote-cron-api/model"
	"github.com/hkalban/remote-cron-api/repository"
)

var once sync.Once

// ServerService : represent the server's services
type ServerService interface {
	Validate(server *model.Server) error
	Create(server *model.Server) (*model.Server, error)
	FindAll() ([]model.Server, error)
	Find(id int) (model.Server, error)
}

type serverService struct {
	serverRepository repository.ServerRepository
}

var instance *serverService

//NewServerService: construction function, injected by server repository
func NewServerService(r repository.ServerRepository) *serverService {
	once.Do(func() {
		instance = &serverService{
			serverRepository: r,
		}
	})
	return instance
}

func (*serverService) Validate(server *model.Server) error {
	if server == nil {
		err := errors.New("The server is empty")
		return err
	}
	if server.Hostname == "" {
		err := errors.New("The hostname of server is empty")
		return err
	}
	if server.IP == "" {
		err := errors.New("The IP of server is empty")
		return err
	}
	if server.Username == "" {
		err := errors.New("The Username of server is empty")
		return err
	}
	if server.Password == "" {
		err := errors.New("The Username of server is empty")
		return err
	}
	return nil
}

func (u *serverService) Create(server *model.Server) (*model.Server, error) {
	return u.serverRepository.Save(server)
}

func (u *serverService) FindAll() ([]model.Server, error) {
	return u.serverRepository.FindAll()
}

func (u *serverService) Find(id int) (model.Server, error) {
	return u.serverRepository.Find(id)
}
