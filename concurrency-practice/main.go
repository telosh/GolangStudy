package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type TaskRequest struct {
	TaskCount int `json:"taskCount"`
}

type TaskResult struct {
	TaskID    int     `json:"taskId"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	Duration  float64 `json:"duration"`
}

// 重い処理をシミュレートする関数
func heavyTask(taskID int, wg *sync.WaitGroup, results chan<- TaskResult) {
	defer wg.Done()

	startTime := time.Now()
	// 重い処理をシミュレート（1秒待機）
	time.Sleep(1 * time.Second)
	endTime := time.Now()

	results <- TaskResult{
		TaskID:    taskID,
		StartTime: startTime.Format(time.RFC3339),
		EndTime:   endTime.Format(time.RFC3339),
		Duration:  endTime.Sub(startTime).Seconds(),
	}
}

// 並行処理なしでタスクを実行
func sequentialTasks(count int) []TaskResult {
	results := make([]TaskResult, count)
	for i := 0; i < count; i++ {
		startTime := time.Now()
		time.Sleep(1 * time.Second)
		endTime := time.Now()

		results[i] = TaskResult{
			TaskID:    i,
			StartTime: startTime.Format(time.RFC3339),
			EndTime:   endTime.Format(time.RFC3339),
			Duration:  endTime.Sub(startTime).Seconds(),
		}
	}
	return results
}

// goroutineを使用して並行処理でタスクを実行
func concurrentTasks(count int) []TaskResult {
	var wg sync.WaitGroup
	results := make(chan TaskResult, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go heavyTask(i, &wg, results)
	}

	// すべてのgoroutineが完了するのを待つ
	go func() {
		wg.Wait()
		close(results)
	}()

	// 結果をスライスに変換
	var taskResults []TaskResult
	for result := range results {
		taskResults = append(taskResults, result)
	}
	return taskResults
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.TaskCount <= 0 {
		http.Error(w, "Task count must be positive", http.StatusBadRequest)
		return
	}

	// 並行処理なしの実行時間を計測
	seqStart := time.Now()
	sequentialResults := sequentialTasks(req.TaskCount)
	seqDuration := time.Since(seqStart).Seconds()

	// 並行処理ありの実行時間を計測
	conStart := time.Now()
	concurrentResults := concurrentTasks(req.TaskCount)
	conDuration := time.Since(conStart).Seconds()

	response := Response{
		Message: "Task execution completed",
		Data: map[string]interface{}{
			"sequential": map[string]interface{}{
				"duration": seqDuration,
				"results":  sequentialResults,
			},
			"concurrent": map[string]interface{}{
				"duration": conDuration,
				"results":  concurrentResults,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/tasks", taskHandler)

	log.Println("Concurrency practice server starting on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
} 