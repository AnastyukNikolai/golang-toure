package repositories

import (
	"sync"
)

const (
	TodoItemsKey = "todo_items"
)

type Storage struct {
	sync.Mutex
	Data map[string]map[string]string
}

func NewStorage() *Storage {
	return &Storage{}
}
