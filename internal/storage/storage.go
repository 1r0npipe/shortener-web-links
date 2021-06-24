package storage

import "github.com/1r0npipe/shortener-web-links/internal/model"

type Storage struct {
}

// The interface to work with external storage solution
// it is like abstruction layer when you need to add, delete or get data from/to
// also will try to implement the cleanup - auto process of cleaning expired data
type StorageManager interface {
	// Get data by key from storage: return found data or error
	Get(key string) (model.Item, error)
	// Put data to the storage and return error if not possible and why
	Put(key string, m model.Item) error
	// Delete data by key from storage, it returns error if not pssible
	Delete(key string, m model.Item) error
}

type Codec interface {
	// Marshal encodes a Go value to a slice of bytes.
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal decodes a slice of bytes into a Go value.
	Unmarshal(data []byte, v interface{}) error
}
