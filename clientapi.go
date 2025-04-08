package config

import "errors"

type Config map[string]string

// Set the given value to the specified key in the given environment.
func Set(env Env, key string, value string) error {
	store, ok := getStoreFor(env)
	if !ok {
		return errors.New("Failed to find store for env: " + string(env))
	}

	if !checkValidConfigurationKey(key) {
		return errors.New("Invalid configuration key: " + key)
	}

	return store.Set(key, value)
}

// Get the value associated with the specified key and environment.
func Get(env Env, key string) (string, bool) {
	store, ok := getStoreFor(env)
	if !ok {
		return "", false
	}

	if !checkValidConfigurationKey(key) {
		return "", false
	}

	return store.Get(key)
}

// Get all the configuration values available for the given environment.
func GetAll(env Env) (Config, bool) {
	store, ok := getStoreFor(env)
	if !ok {
		return nil, false
	}
	configKeys := allConfigurationKeys()
	config := make(map[string]string, len(configKeys))
	for _, key := range configKeys {
		value, ok := store.Get(key)
		if !ok {
			return nil, false
		}
		config[key] = value
	}
	return config, true
}
