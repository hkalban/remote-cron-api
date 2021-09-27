package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hkalban/remote-cron-api/model"
	"github.com/hkalban/remote-cron-api/service"
)

type serverController struct {
	serverService service.ServerService
}

type ServerController interface {
	GetServer(w http.ResponseWriter, r *http.Request)
	GetServers(w http.ResponseWriter, r *http.Request)
	AddServer(w http.ResponseWriter, r *http.Request)
}

func NewServerController(s service.ServerService) ServerController {
	return &serverController{
		serverService: s,
	}
}

func (u *serverController) GetServers(w http.ResponseWriter, r *http.Request) {
	servers, err := u.serverService.FindAll()
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json, err := json.Marshal(servers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setSuccessHeaders(w)
	w.Write(json)
}

func (u *serverController) GetServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	server, err := u.serverService.Find(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json, err := json.Marshal(server)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setSuccessHeaders(w)
	w.Write(json)
}

func (u *serverController) AddServer(w http.ResponseWriter, r *http.Request) {
	var server model.Server

	if err := (json.NewDecoder(r.Body).Decode(&server)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err1 := (u.serverService.Validate(&server)); err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	u.serverService.Create(&server)

	setAcceptedHeaders(w)
}
