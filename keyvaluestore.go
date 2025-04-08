package config

type KeyValueStore interface {
	/// Store the provided value at the specified key.
	Set(key string, value string) error

	/// Get the value associated to the specified key.
	Get(key string) (string, bool)

	/// Clear the store.
	Clear()
}

// / A KeyValueStored that stores configuration in volatile memory.
type InMemoryKeyValueStore struct {
	store map[string]string
}

func (im *InMemoryKeyValueStore) Set(key string, value string) error {
	im.store[key] = value
	return nil
}

func (im *InMemoryKeyValueStore) Get(key string) (string, bool) {
	value, ok := im.store[key]
	return value, ok
}

func (im *InMemoryKeyValueStore) Clear() {
	for k := range im.store {
		delete(im.store, k)
	}
}

func NewDefaultKeyValueStore() KeyValueStore {
	return &InMemoryKeyValueStore{
		store: make(map[string]string),
	}
}
