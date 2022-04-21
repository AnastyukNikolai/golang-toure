package repositories

import (
	"context"
	"fmt"
	"golang-ture/ent"
	"golang-ture/ent/todoitem"
	"golang-ture/internal/models"
	_ "google.golang.org/genproto/googleapis/rpc/status"
)

type TodoItemRepositoryEnt struct {
	DBClient *ent.Client
}

func NewTodoItemRepositoryEnt(DBClient *ent.Client) *TodoItemRepositoryEnt {
	return &TodoItemRepositoryEnt{DBClient: DBClient}
}

func (r *TodoItemRepositoryEnt) Create(item models.TodoItem) (int, error) {
	fmt.Println(item)
	todoItemDB, err := r.DBClient.TodoItem.
		Create().
		SetTitle(item.Title).
		SetDescription(item.Description).
		SetStatus(0).
		SetDone(false).
		Save(context.Background())
	if err != nil {
		return 0, err
	}
	return todoItemDB.ID, nil
}

func (r *TodoItemRepositoryEnt) GetAll() ([]models.TodoItem, error) {
	var items []models.TodoItem
	todoItemsDB, err := r.DBClient.TodoItem.Query().All(context.Background())
	if err != nil {
		return items, err
	}
	for _, todoItemDB := range todoItemsDB {
		itemStatus := models.TodoItemStatus(todoItemDB.Status).String()
		var todo = models.TodoItem{todoItemDB.ID, todoItemDB.Title, todoItemDB.Description, itemStatus, todoItemDB.Done}
		items = append(items, todo)
	}
	return items, nil
}

func (r *TodoItemRepositoryEnt) GetById(itemId int) (models.TodoItem, error) {
	var todo models.TodoItem
	todoItemDB, err := r.DBClient.TodoItem.Query().Where(todoitem.ID(itemId)).Only(context.Background())
	if err != nil {
		return todo, err
	}
	itemStatus := models.TodoItemStatus(todoItemDB.Status).String()
	todo = models.TodoItem{todoItemDB.ID, todoItemDB.Title, todoItemDB.Description, itemStatus, todoItemDB.Done}
	return todo, nil
}

func (r *TodoItemRepositoryEnt) Delete(itemId int) error {
	err := r.DBClient.TodoItem.DeleteOneID(itemId).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoItemRepositoryEnt) Update(itemId int, input models.UpdateTodoItemInput) (models.TodoItem, error) {
	var todo models.TodoItem
	toDoUpdate := r.DBClient.TodoItem.UpdateOneID(itemId)
	if input.Title != nil {
		toDoUpdate.SetTitle(*input.Title)
	}
	if input.Description != nil {
		toDoUpdate.SetDescription(*input.Description)
	}
	if input.Status != nil {
		toDoUpdate.SetStatus(*input.Status)
	}
	if input.Done != nil {
		toDoUpdate.SetDone(*input.Done)
	}

	todoItemDB, err := toDoUpdate.Save(context.Background())
	if err != nil {
		return todo, err
	}

	itemStatus := models.TodoItemStatus(todoItemDB.Status).String()
	todo = models.TodoItem{todoItemDB.ID, todoItemDB.Title, todoItemDB.Description, itemStatus, todoItemDB.Done}

	return todo, nil
}
