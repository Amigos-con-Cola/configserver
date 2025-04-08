package config

import "testing"

func TestStores_StoreValuesAtKeys(t *testing.T) {
	inMemory := NewDefaultKeyValueStore()

	testKeyValueStore_StoresValuesAtKey(inMemory, t)
}

func testKeyValueStore_StoresValuesAtKey(kvs KeyValueStore, t *testing.T) {
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for key, value := range data {
		kvs.Set(key, value)
	}

	for key := range data {
		_, ok := kvs.Get(key)
		if !ok {
			t.Errorf("Expected %s to be in store, but it wasn't!", key)
		}
	}
}
