package controller

func (h *BaseController) RegisterRoutes() *BaseController {
	h.registerServerRoutes()
	h.registerTaskRoutes()

	return h
}

func (h *BaseController) registerServerRoutes() {
	h.router.HandleFunc(h.basePath+"/servers", h.ServerController.GetServers).Methods("GET")
	h.router.HandleFunc(h.basePath+"/servers/{id:[0-9]+}", h.ServerController.GetServer).Methods("GET")
	h.router.HandleFunc(h.basePath+"/servers", h.ServerController.AddServer).Methods("POST")
	// TODO: update endpoint
	// TODO: delete endpoint
}

func (h *BaseController) registerTaskRoutes() {
	h.router.HandleFunc(h.basePath+"/tasks", h.TaskController.GetTasks).Methods("GET")
	h.router.HandleFunc(h.basePath+"/tasks/{id:[0-9]+}", h.TaskController.GetTask).Methods("GET")
	h.router.HandleFunc(h.basePath+"/tasks", h.TaskController.AddTask).Methods("POST")
	h.router.HandleFunc(h.basePath+"/tasks/{id:[0-9]+}/executions", h.TaskController.GetLatestTaskExecutions).Methods("GET")
	// TODO: update endpoint
	// TODO: delete endpoint
}
