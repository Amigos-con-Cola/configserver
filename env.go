package config

type Env string

const (
	ENV_PROD = "prod"
	ENV_DEV  = "dev"
)

func getStoreFor(env Env) (KeyValueStore, bool) {
	switch env {
	case ENV_PROD:
		return PROD_STORE, true
	case ENV_DEV:
		return DEV_STORE, true
	default:
		return nil, false
	}
}

func clearStores() {
	PROD_STORE.Clear()
	DEV_STORE.Clear()
}
