package models

import "errors"

type TodoItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UpdateTodoItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateTodoItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

func (i *TodoItem) ModifyID(newID string) {
	i.Id = newID
}

func (i *TodoItem) ModifyTitle(newTitle string) {
	i.Title = newTitle
}

func (i *TodoItem) ModifyDescription(newDescription string) {
	i.Description = newDescription
}

func (i *TodoItem) ModifyDone(newDone bool) {
	i.Done = newDone
}
