package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/internal/api/worker/hash/crack/task", handleTask).Methods("POST")
	r.HandleFunc("/health", healthCheck).Methods("GET")
	log.Println("Starting worker on port " + os.Getenv("WORKER_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("WORKER_PORT"), r))
}
