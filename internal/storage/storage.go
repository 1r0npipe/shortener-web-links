package storage

import "github.com/1r0npipe/shortener-web-links/internal/model"

type Storage struct {
}

// The interface to work with external storage solution
// it is like abstruction layer when you need to add, delete or get data from/to
// also will try to impletent cleanup - auto process of cleaning expired data
type Storager interface {
	// New - allocate new storage client
	New(typeSt string) (*Storage, error)
	// Get data by key from storage: return found data or error
	Get(key string) (model.Info, error)
	// Put data to the storage and return error if not possible and why
	Put(key string, m model.Info) error
	// Delete data by key from storage, it returns error if not pssible
	Delete(key string, m model.Info) error
	// CleanUp the expired data from data storage
	CleanUp() error
}
