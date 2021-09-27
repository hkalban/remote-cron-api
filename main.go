package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hkalban/remote-cron-api/controller"
	"github.com/hkalban/remote-cron-api/repository"
	"github.com/hkalban/remote-cron-api/scheduler"
	"github.com/hkalban/remote-cron-api/service"
)

func main() {
	api_base := "/v1/api"
	api_port := 3000

	// Initialize Repositories
	db := repository.Connect()
	br, err := repository.NewBaseRepository(db).CreateSchema()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Services
	bs := service.NewBaseService(*br)

	// Initialize scheduler
	scheduler := scheduler.NewBaseScheduler(*bs)
	scheduler.Start()

	// Initialize Controllers
	router := mux.NewRouter()
	controller.NewBaseController(api_base, router, *bs, *scheduler).RegisterRoutes()

	// Start http server
	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", api_port),
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf(fmt.Sprintf("remote-cron-api server is listening on port %d", api_port))
	if err := apiServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
