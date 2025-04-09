package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Amigos-con-Cola/config"
)

var (
	configServerUsername string
	configServerPassword string
)

func authMiddleware(wrapped func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != configServerUsername || password != configServerPassword {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		wrapped(w, r)
	}
}

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
	env := config.Env(r.PathValue("env"))
	log.Printf("Serving request for setting configuration value for env: %s\n", env)

	var payload map[string]string
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("There was an error decoding the request payload: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Consider implementing a transactional API.

	type response struct {
		Success bool     `json:"success"`
		Errors  []string `json:"errors"`
	}

	errors := make([]string, 0)

	for key, value := range payload {
		err := config.Set(env, key, value)
		if err != nil {
			log.Printf("Failed to set value in env (%s): %s = %s\n", env, key, value)
			errors = append(errors, key)
		}
	}

	res := response{
		Success: len(errors) == 0,
		Errors:  errors,
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	configServerUsername = os.Getenv("CONFIG_SERVER_USERNAME")
	if configServerUsername == "" {
		log.Fatalf("CONFIG_SERVER_USERNAME is not set")
	}

	configServerPassword = os.Getenv("CONFIG_SERVER_PASSWORD")
	if configServerPassword == "" {
		log.Fatalf("CONFIG_SERVER_PASSWORD is not set")
	}

	http.HandleFunc("POST /api/v1/{env}", authMiddleware(handleSetMany))
	http.HandleFunc("GET /api/v1/{env}", authMiddleware(handleGetAll))
	http.HandleFunc("GET /api/v1/{env}/{key}", authMiddleware(handleGetOne))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
