package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/paytax", taxHandler)
	http.ListenAndServe(":8080", nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis-service.learnk8s.svc.cluster.local:6379", // Redis server address
	Password: "",               // No password set
	DB:       0,                // Use default DB
})

func taxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		ClientId string  `json:"client_id"`
		Wage     float64 `json:"wage"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.Wage <= 0 {
		http.Error(w, "Wage cannot be non-positive", http.StatusBadRequest)
		return
	}

    if input.ClientId == "" {
        http.Error(w, "Client ID cannot be empty", http.StatusBadRequest)
        return
    }

	tax := input.Wage * 0.30

	// Store the calculated tax in Redis
	err = rdb.HIncrByFloat(context.Background(), "unpaidtaxes", input.ClientId, tax).Err()
	if err != nil {
		http.Error(w, "Failed to store in Redis", http.StatusInternalServerError)
		return
	}

	// Since we're not returning anything, we just send a status code.
	// Here we're using http.StatusNoContent to indicate success but no content to return.
	w.WriteHeader(http.StatusNoContent)
}
