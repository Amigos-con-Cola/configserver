package config

import (
	"strings"
	"testing"
)

func TestSettingValuesInEnvDoesntAffectAnotherEnv(t *testing.T) {
	clearStores()

	key := CONFIG_AUTH_CALLBACK_URL

	err := Set(ENV_DEV, key, "value")
	if err != nil {
		t.Errorf("Failed to set key in store: '%s'", key)
	}

	_, ok := Get(ENV_DEV, key)
	if !ok {
		t.Errorf("Expected '%s' to be present in store, but it wasn't!", key)
	}

	_, ok = Get(ENV_PROD, key)
	if ok {
		t.Errorf("Expected '%s' to be absent from store, but it was present!", key)
	}
}

func TestSettingValueToInvalidKeyReturnsError(t *testing.T) {
	err := Set(ENV_DEV, "key", "value")
	if err == nil {
		t.Errorf("Expected setting invalid key to fail, but it didn't!")
	}

	if !strings.Contains(strings.ToLower(err.Error()), "invalid configuration key") {
		t.Errorf("Invalid key error is not descriptive enough")
	}
}

func TestGettingValueFromInvalidEnvReturnsFalse(t *testing.T) {
	_, ok := Get("INVALID_ENV", CONFIG_AUTH_CALLBACK_URL)
	if ok {
		t.Errorf("Expected getting value from invalid env to return false, but it did not!")
	}
}

func TestGettingNonSetValueReturnsFalse(t *testing.T) {
	clearStores()

	_, ok := Get(ENV_DEV, "key")
	if ok {
		t.Errorf("Expected getting a non-set value to return false, but it returned true")
	}
}

func TestSettingValueInInvalidEnvReturnsError(t *testing.T) {
	clearStores()

	err := Set("KKCK", "key", "value")
	if err == nil {
		t.Errorf("Expected setting value for invalid env to fail, but it didn't!")
	}

	if !strings.Contains(err.Error(), "Failed to find store for env") {
		t.Errorf("Error message for invalid env is not descriptive enough")
	}
}

func TestGettingAllConfigurationsReturnsConfigValuesForAllKnownKeys(t *testing.T) {
	clearStores()

	for _, key := range allConfigurationKeys() {
		Set(ENV_DEV, key, "some-value")
	}

	config, ok := GetAll(ENV_DEV)
	if !ok {
		t.Errorf("Failed to get all configuration values")
		return
	}

	for _, key := range allConfigurationKeys() {
		_, ok := config[key]
		if !ok {
			t.Errorf("Key '%s' is not present in returned configuration", key)
		}
	}
}
