package main

import (
	"context"
	"controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/zhashkevych/scheduler"
	"helpers"
	"log"
	"net/http"
	"schedulers"
	"sqlconnector"
	"strconv"
	"time"
)

func main() {

	log.Println("[INFO] Starting Web Server...")
	log.Println("[INFO] Loading configuration from config.json...")
	configuration := helpers.GetConfiguration()
	log.Println("[INFO] Configuration loaded!")

	log.Println("[INFO] Creating database if needed...")
	sqlconnector.CreateDatabase()
	log.Println("[INFO] Database ready to use!")

	if configuration.Monitoring.EnableScheduler {
		log.Println("[INFO] monitoring callback enabled, registering scheduler...")
		ctx := context.Background()
		worker := scheduler.NewScheduler()
		worker.Add(ctx, schedulers.Scheduler, time.Minute)
		log.Println("[INFO] Registered!")
	}

	log.Println("[INFO] Starting Web Server...")

	//Start Web Server
	r := mux.NewRouter()
	api := r.PathPrefix("").Subrouter()

	api.HandleFunc("/updateService", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateController(w, r)
	}).Methods(http.MethodPost, http.MethodOptions)

	log.Println("[INFO] Web server started. Now listening on *:"+strconv.Itoa(configuration.ServerPort)+".")
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(configuration.ServerPort), r))
}
