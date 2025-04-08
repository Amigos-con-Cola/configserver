package config

var (
	DEV_STORE  KeyValueStore
	PROD_STORE KeyValueStore
)

func init() {
	DEV_STORE = NewDefaultKeyValueStore()
	PROD_STORE = NewDefaultKeyValueStore()

	for _, key := range allConfigurationKeys() {
		DEV_STORE.Set(key, "")
		PROD_STORE.Set(key, "")
	}
}
