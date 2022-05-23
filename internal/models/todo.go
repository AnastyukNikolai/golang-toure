package models

import (
	"errors"
	"golang-ture/ent"
)

type TodoItemStatus int

//go:generate go run github.com/dmarkham/enumer -type=TodoItemStatus
const (
	Backlog TodoItemStatus = iota
	InProgress
	Review
	Done
)

//go:generate go run $GOPATH/src/golang-ture/cmd/gen/generator_v2.go ./.TodoItem
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Done        bool   `json:"done"`
}

type UpdateTodoItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *int    `json:"status" binding:"required,gte=0,lte=3"`
	Done        *bool   `json:"done"`
}

func (i UpdateTodoItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

func (i *TodoItem) ModifyID(newID int) {
	i.Id = newID
}

func (i *TodoItem) ModifyTitle(newTitle string) {
	i.Title = newTitle
}

func (i *TodoItem) ModifyDescription(newDescription string) {
	i.Description = newDescription
}

func (i *TodoItem) ModifyStatus(newStatus int) {
	i.Status = TodoItemStatus(newStatus).String()
}

func (i *TodoItem) ModifyDone(newDone bool) {
	i.Done = newDone
}

func (i *TodoItem) FormatDBTodoItemToModel(DBTodoItem *ent.TodoItem) {
	i.Id = DBTodoItem.ID
	i.Title = DBTodoItem.Title
	i.Description = DBTodoItem.Description
	i.Status = TodoItemStatus(DBTodoItem.Status).String()
	i.Done = DBTodoItem.Done
}
