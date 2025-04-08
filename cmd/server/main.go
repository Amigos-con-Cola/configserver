package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Amigos-con-Cola/config"
)

func handleGetAll(w http.ResponseWriter, r *http.Request) {
	env := r.PathValue("env")
	log.Printf("Serving request for configuration values for env: %s\n", env)

	cfg, ok := config.GetAll(config.Env(env))
	if !ok {
		log.Printf("Failed to get configuration values for env: %s\n", env)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(cfg)
}

func handleGetOne(w http.ResponseWriter, r *http.Request) {
	env := r.PathValue("env")
	key := r.PathValue("key")
	log.Printf("Serving request for configuration value for env: %s\n", env)

	value, ok := config.Get(config.Env(env), key)
	if !ok {
		log.Printf("Request for key '%s' failed\n", key)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Value string `json:"value"`
	}{
		Value: value,
	})
}

func handleSetMany(w http.ResponseWriter, r *http.Request) {
	env := r.PathValue("env")
	log.Printf("Serving request for setting configuration value for env: %s\n", env)
}

func main() {
	http.HandleFunc("POST /api/v1/{env}", handleSetMany)
	http.HandleFunc("GET /api/v1/{env}", handleGetAll)
	http.HandleFunc("GET /api/v1/{env}/{key}", handleGetOne)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
