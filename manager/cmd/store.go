package main

import (
	"github.com/google/uuid"
	"log"
	"ris/manager/model"
	"sync"
)

type Task struct {
	RequestID      string
	Hash           string
	MaxLength      int
	Status         string
	Data           []string
	CompletedParts int
}

var (
	taskStore = make(map[string]*Task)
	mu        sync.Mutex
)

func createTask(hash string, maxLength int) string {
	mu.Lock()
	defer mu.Unlock()

	requestId := uuid.New().String()
	taskStore[requestId] = &Task{
		RequestID: requestId,
		Hash:      hash,
		MaxLength: maxLength,
		Status:    model.IN_PROGRESS,
		Data:      []string{},
	}

	log.Printf("Создана новая задача: RequestID=%s", requestId)
	return requestId
}

func getHashStatusById(requestId string) (string, []string) {
	mu.Lock()
	defer mu.Unlock()

	task, exists := taskStore[requestId]
	if !exists {
		return "NOT_FOUND", nil
	}

	return task.Status, task.Data
}

func appendTaskData(requestId, word string) {
	mu.Lock()
	defer mu.Unlock()

	if task, exists := taskStore[requestId]; exists {
		task.Data = append(task.Data, word)
		task.CompletedParts++
	}
}

func updateTaskStatus(requestId, status string) {
	mu.Lock()
	defer mu.Unlock()

	if task, exists := taskStore[requestId]; exists {
		task.Status = status
	}
}

func countOfCompletedWorkers(requestId string) int {
	mu.Lock()
	defer mu.Unlock()

	if task, exists := taskStore[requestId]; exists {
		return task.CompletedParts
	}

	return 0
}
