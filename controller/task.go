package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hkalban/remote-cron-api/model"
	"github.com/hkalban/remote-cron-api/scheduler"
	"github.com/hkalban/remote-cron-api/service"
)

type taskController struct {
	taskService service.TaskService
	scheduler   scheduler.BaseScheduler
}

type TaskController interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	GetTasks(w http.ResponseWriter, r *http.Request)
	AddTask(w http.ResponseWriter, r *http.Request)
	GetLatestTaskExecutions(w http.ResponseWriter, r *http.Request)
}

func NewTaskController(s service.TaskService, scheduler scheduler.BaseScheduler) TaskController {
	return &taskController{
		taskService: s,
		scheduler:   scheduler,
	}
}

type TaskRO struct {
	Name            string `json:"name"`
	Command         string `json:"command"`
	IntervalSeconds int    `json:"interval_seconds"`
	TimeoutSeconds  int    `json:"timeout_seconds"`
	Servers         []int  `json:"servers"`
}

func (u *taskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := u.taskService.FindAll()
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setSuccessHeaders(w)
	w.Write(json)
}

func (u *taskController) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	task, err := u.taskService.Find(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setSuccessHeaders(w)
	w.Write(json)
}

func (u *taskController) AddTask(w http.ResponseWriter, r *http.Request) {
	var taskRO TaskRO

	// Parse Request Body
	if err := json.NewDecoder(r.Body).Decode(&taskRO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(taskRO.Servers) < 1 {
		http.Error(w, "One or more Server IDs required", http.StatusBadRequest)
	}

	// Get all associated servers
	var servers []*model.Server
	for _, server_id := range taskRO.Servers {
		server, err1 := u.taskService.FindServer(server_id)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		servers = append(servers, &server)
	}

	// Map to DB Entity
	taskEntity := model.Task{
		Name:            taskRO.Name,
		Command:         taskRO.Command,
		IntervalSeconds: taskRO.IntervalSeconds,
		TimeoutSeconds:  taskRO.TimeoutSeconds,
		Servers:         servers,
	}

	if err2 := (u.taskService.Validate(&taskEntity)); err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	u.taskService.Create(&taskEntity)

	u.scheduler.AddNewTask(taskEntity)

	setAcceptedHeaders(w)
}

func (u *taskController) GetLatestTaskExecutions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	latestExecutions, err := u.taskService.FindLatestTaskExecutions(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json, err := json.Marshal(latestExecutions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setSuccessHeaders(w)
	w.Write(json)
}
