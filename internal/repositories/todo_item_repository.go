package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/xid"
	"golang-ture/internal/models"
	_map "golang-ture/pkg/map"
)

type TodoItemRepository struct {
	storage *Storage
}

func NewTodoItemRepository(storage *Storage) *TodoItemRepository {
	return &TodoItemRepository{storage: storage}
}

func (r *TodoItemRepository) Create(item models.TodoItem) (string, error) {
	guid := xid.New().String()
	item.ModifyID(guid)

	jsonItem, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	var newItem map[string]string
	newItem = make(map[string]string)
	newItem[guid] = string(jsonItem)
	r.storage.Lock()
	r.storage.Data[todoItemsKey] = _map.MergeStringMaps(r.storage.Data[todoItemsKey], newItem)
	r.storage.Unlock()

	return guid, nil
}

func (r *TodoItemRepository) GetAll() ([]models.TodoItem, error) {
	var items []models.TodoItem
	for _, value := range r.storage.Data[todoItemsKey] {
		var todo models.TodoItem
		err := json.Unmarshal([]byte(value), &todo)
		if err != nil {
			return nil, err
		}
		items = append(items, todo)
	}
	return items, nil
}

func (r *TodoItemRepository) GetById(itemId string) (models.TodoItem, error) {
	var todo models.TodoItem
	item, ok := r.storage.Data[todoItemsKey][itemId]
	if !ok {
		return todo, errors.New("item not found")
	}
	err := json.Unmarshal([]byte(item), &todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoItemRepository) Delete(itemId string) error {
	_, ok := r.storage.Data[todoItemsKey][itemId]
	if !ok {
		return nil
	}
	r.storage.Lock()
	delete(r.storage.Data[todoItemsKey], itemId)
	r.storage.Unlock()
	return nil
}

func (r *TodoItemRepository) Update(itemId string, input models.UpdateTodoItemInput) error {
	todo, err := r.GetById(itemId)
	if err != nil {
		return err
	}

	if input.Title != nil {
		todo.ModifyTitle(*input.Title)
	}

	if input.Description != nil {
		todo.ModifyDescription(*input.Description)
	}

	if input.Done != nil {
		todo.ModifyDone(*input.Done)

	}

	updatedTodo, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	r.storage.Lock()
	r.storage.Data[todoItemsKey][itemId] = string(updatedTodo)
	r.storage.Unlock()

	fmt.Println(r.storage.Data[todoItemsKey])
	return nil
}
