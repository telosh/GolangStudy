package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ArrayOperationRequest struct {
	Numbers []int  `json:"numbers"`
	Operation string `json:"operation"`
}

// 配列の操作を行う関数
func processArray(numbers []int, operation string) interface{} {
	switch operation {
	case "sort":
		sorted := make([]int, len(numbers))
		copy(sorted, numbers)
		sort.Ints(sorted)
		return sorted
	case "sum":
		sum := 0
		for _, num := range numbers {
			sum += num
		}
		return sum
	case "average":
		if len(numbers) == 0 {
			return 0
		}
		sum := 0
		for _, num := range numbers {
			sum += num
		}
		return float64(sum) / float64(len(numbers))
	case "max":
		if len(numbers) == 0 {
			return nil
		}
		max := numbers[0]
		for _, num := range numbers {
			if num > max {
				max = num
			}
		}
		return max
	case "min":
		if len(numbers) == 0 {
			return nil
		}
		min := numbers[0]
		for _, num := range numbers {
			if num < min {
				min = num
			}
		}
		return min
	case "unique":
		seen := make(map[int]bool)
		unique := []int{}
		for _, num := range numbers {
			if !seen[num] {
				seen[num] = true
				unique = append(unique, num)
			}
		}
		return unique
	default:
		return nil
	}
}

func arrayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ArrayOperationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := processArray(req.Numbers, req.Operation)
	if result == nil {
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Array operation completed",
		Data:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/array", arrayHandler)

	log.Println("Learning-1 server starting on :8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
} 