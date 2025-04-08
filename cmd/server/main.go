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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(cfg)
}

func handleGetOne(w http.ResponseWriter, r *http.Request) {}

func handleSetMany(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("POST /api/v1/{env}", handleSetMany)
	http.HandleFunc("GET /api/v1/{env}", handleGetAll)
	http.HandleFunc("GET /api/v1/{env}/{key}", handleGetOne)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
