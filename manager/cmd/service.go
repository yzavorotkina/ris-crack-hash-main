package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

var taskQueue chan struct{}

func initializeTaskQueue() {
	queueSize := os.Getenv("QUEUE_SIZE")
	queueSizeInt, _ := strconv.Atoi(queueSize)
	taskQueue = make(chan struct{}, queueSizeInt)
	for i := 0; i < queueSizeInt; i++ {
		taskQueue <- struct{}{}
	}
}

func Init() {
	initializeTaskQueue()

	r := mux.NewRouter()
	r.HandleFunc("/internal/api/manager/hash/crack", crackHash).Methods("POST")
	r.HandleFunc("/internal/api/manager/hash/status", getHashStatus).Methods("GET")
	r.HandleFunc("/internal/api/manager/hash/crack/request", workerResult).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("MANAGER_PORT"), r))
}

//	порт в константы
