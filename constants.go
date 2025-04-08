package config

import "slices"

const (
	CONFIG_AUTH_CLIENT_ID     = "AUTH_CLIENT_ID"
	CONFIG_AUTH_CLIENT_SECRET = "AUTH_CLIENT_SECRET"
	CONFIG_AUTH_CALLBACK_URL  = "AUTH_CALLBACK_URL"
	CONFIG_AUTH_ISSUER        = "AUTH_ISSUER"
)

func allConfigurationKeys() []string {
	return []string{
		CONFIG_AUTH_CLIENT_ID,
		CONFIG_AUTH_CLIENT_SECRET,
		CONFIG_AUTH_CALLBACK_URL,
		CONFIG_AUTH_ISSUER,
	}
}

func checkValidConfigurationKey(key string) bool {
	return slices.Contains(allConfigurationKeys(), key)
}
