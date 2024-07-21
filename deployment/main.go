package main

import (
	"encoding/json"
	"math"
	"net/http"

	"log/slog"

	"github.com/redis/go-redis/v9"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.HandleFunc("/paytax", taxHandler)

	http.ListenAndServe(":8080", nil)
}

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis-service.k8s-in-a-shell.svc.cluster.local:6379",
	Password: "",
	DB:       0,
})

func taxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	var input struct {
		EmployeeId string  `json:"employee_id"`
		Wage       float64 `json:"wage"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to decode input", "error", err)
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	slog.InfoContext(ctx, "Successfully decoded input", "input", input)
	
	if input.Wage <= 0 {
		http.Error(w, "Wage cannot be non-positive", http.StatusBadRequest)
		return
	}

	if input.EmployeeId == "" {
		http.Error(w, "Employee ID cannot be empty", http.StatusBadRequest)
		return
	}

	tax := math.Round(input.Wage * 30) / 100

	// Store the calculated tax in Redis
	err = rdb.HIncrByFloat(ctx, "unpaidtaxes", input.EmployeeId, tax).Err()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to store in Redis", "error", err)
		http.Error(w, "Failed to store in Redis: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Since we're not returning anything, we just send a status code.
	// Here we're using http.StatusNoContent to indicate success but no content to return.
	w.WriteHeader(http.StatusNoContent)
}
