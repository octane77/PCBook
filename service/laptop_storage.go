package service

import (
	"PCBook/pb/PCBook/proto"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"sync"
)

// ErrAlreadyExists is returned when a record with the same ID already exists
var ErrAlreadyExists = errors.New("already exists")

// LaptopStorage is an interface for laptop storage
type LaptopStorage interface {
	Save(laptop *proto.Laptop) error
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.Mutex
	data  map[string]*proto.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*proto.Laptop),
	}
}

// Save laptop to storage
func (store *InMemoryLaptopStore) Save(laptop *proto.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	//Deep copy
	other := &proto.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("unable to copy laptop data: %v", err)
	}

	store.data[other.Id] = other
	return nil
}
