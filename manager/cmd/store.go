package main

import (
	"github.com/google/uuid"
	"log"
	"ris/manager/model"
	"sync"
	"time"
)

type Task struct {
	RequestID      string
	Hash           string
	MaxLength      int
	Status         string
	Data           []string
	CompletedParts int
	StartTime      time.Time
	Timeout        time.Duration
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
		StartTime: time.Now(),
		Timeout:   1 * time.Minute,
	}

	log.Printf("Создана новая задача: RequestID=%s", requestId)
	return requestId
}

func getHashStatusById(requestId string) (string, []string, time.Time, time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	task, exists := taskStore[requestId]
	if !exists {
		return "NOT_FOUND", nil, time.Time{}, 0
	}

	return task.Status, task.Data, task.StartTime, task.Timeout
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
