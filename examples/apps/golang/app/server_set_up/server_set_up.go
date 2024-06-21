package server_set_up

import (
	"basic-go-rest-api/first_controller"
	"basic-go-rest-api/second_controller"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"os"
	"time"
)

func StartServer() {
	PORT := os.Getenv("PORT")

	if PORT == ""{
		PORT = "3000"
	}

	router := ConfigureRoutes()

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Addr:           ":" + PORT,
		Handler:        loggedRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	server.ListenAndServe()
}

func ConfigureRoutes() *mux.Router {
	log.Println("Starting with Route Configuration...")
	router := mux.NewRouter()

	firstController := router.PathPrefix("/api/first-controller").Subrouter()
	secondController := router.PathPrefix("/api/second-controller").Subrouter()

	firstController.HandleFunc("", first_controller.GetHandler).Methods("GET")
	secondController.HandleFunc("", second_controller.GetHandler).Methods("GET")

	firstController.HandleFunc("", first_controller.PostHandler).Methods("POST")
	secondController.HandleFunc("", second_controller.PostHandler).Methods("POST")

	firstController.HandleFunc("/{id}/{gender}", first_controller.PutHandler).Methods("PUT")
	secondController.HandleFunc("/{id}/{gender}", second_controller.PutHandler).Methods("PUT")

	firstController.HandleFunc("/{id}/{gender}", first_controller.DeletHandler).Methods("DELETE")
	secondController.HandleFunc("/{id}/{gender}", second_controller.DeletHandler).Methods("DELETE")

	log.Println("Router configuration has been completed successfuly.")

	return router
}
