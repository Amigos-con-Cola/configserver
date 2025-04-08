package config

var (
	DEV_STORE  KeyValueStore
	PROD_STORE KeyValueStore
)

func init() {
	DEV_STORE = NewDefaultKeyValueStore()
	PROD_STORE = NewDefaultKeyValueStore()
}
