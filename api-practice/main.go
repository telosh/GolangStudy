package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CalculationRequest struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
	Op   string  `json:"op"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
}

type StringOperationRequest struct {
	Text string `json:"text"`
	Op   string `json:"op"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var result float64
	switch req.Op {
	case "add":
		result = req.Num1 + req.Num2
	case "subtract":
		result = req.Num1 - req.Num2
	case "multiply":
		result = req.Num1 * req.Num2
	case "divide":
		if req.Num2 == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		}
		result = req.Num1 / req.Num2
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Calculation successful",
		Data: CalculationResponse{
			Result: result,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func stringOperationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StringOperationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var result string
	switch req.Op {
	case "uppercase":
		result = strings.ToUpper(req.Text)
	case "lowercase":
		result = strings.ToLower(req.Text)
	case "reverse":
		runes := []rune(req.Text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result = string(runes)
	case "length":
		result = strconv.Itoa(len(req.Text))
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "String operation successful",
		Data:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/string", stringOperationHandler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}