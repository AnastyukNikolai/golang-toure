package services

import (
	"golang-ture/internal/models"
	"golang-ture/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	storage := repositories.NewStorage()
	r := repositories.NewRepositoryStorage(storage)
	service := NewService(r)

	type args struct {
		item models.TodoItem
	}

	tests := []struct {
		name    string
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				item: models.TodoItem{
					Id:          1,
					Title:       "test title",
					Description: "test description",
					Status:      models.TodoItemStatus(0).String(),
					Done:        false,
				},
			},
			want: 1,
		},
		{
			name: "Already exist",
			input: args{
				item: models.TodoItem{
					Id:          1,
					Title:       "test title",
					Description: "test description",
					Status:      models.TodoItemStatus(0).String(),
					Done:        false,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := service.Create(tt.input.item)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetAll(t *testing.T) {

	storage := repositories.NewStorage()
	r := repositories.NewRepositoryStorage(storage)
	service := NewService(r)

	tests := []struct {
		name    string
		mock    func()
		want    []models.TodoItem
		wantErr bool
	}{
		{
			name: "Empty",
			mock: func() {},
			want: nil,
		},
		{
			name: "Ok",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
					Done:        true,
				})
				if err != nil {
					panic(err)
				}
			},
			want: []models.TodoItem{
				{Id: 1, Title: "title", Description: "description", Status: models.TodoItemStatus(0).String(), Done: true},
			},
		},

		//{
		//	name: "Error",
		//	mock: func() {},
		//	wantErr: true,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := service.GetAll()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	storage := repositories.NewStorage()
	r := repositories.NewRepositoryStorage(storage)
	service := NewService(r)

	type args struct {
		itemId int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    models.TodoItem
		wantErr bool
	}{
		{
			name: "Not Found",
			mock: func() {},
			input: args{
				itemId: 404,
			},
			wantErr: true,
		},
		{
			name: "Ok",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
					Done:        true,
				})
				if err != nil {
					panic(err)
				}
			},
			input: args{
				itemId: 1,
			},
			want: models.TodoItem{Id: 1, Title: "title", Description: "description", Status: models.TodoItemStatus(0).String(), Done: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := service.GetById(tt.input.itemId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	storage := repositories.NewStorage()
	r := repositories.NewRepositoryStorage(storage)
	service := NewService(r)

	type args struct {
		itemId int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
					Done:        true,
				})
				if err != nil {
					panic(err)
				}
			},
			input: args{
				itemId: 1,
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			mock: func() {},
			input: args{
				itemId: 404,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := service.Delete(tt.input.itemId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	storage := repositories.NewStorage()
	r := repositories.NewRepositoryStorage(storage)
	service := NewService(r)

	type args struct {
		itemId int
		input  models.UpdateTodoItemInput
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    models.TodoItem
		wantErr bool
	}{
		{
			name: "OK_AllFields",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
				})
				if err != nil {
					panic(err)
				}
			},
			input: args{
				itemId: 1,
				input: models.UpdateTodoItemInput{
					Title:       stringPointer("new title"),
					Description: stringPointer("new description"),
					Status:      intPointer(1),
					Done:        boolPointer(true),
				},
			},
			want: models.TodoItem{Id: 1, Title: "new title", Description: "new description", Status: models.TodoItemStatus(1).String(), Done: true},
		},
		{
			name: "OK_WithoutStatus",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          2,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
				})
				if err != nil {
					panic(err)
				}
			},
			input: args{
				itemId: 2,
				input: models.UpdateTodoItemInput{
					Title:       stringPointer("new title"),
					Description: stringPointer("new description"),
					Done:        boolPointer(true),
				},
			},
			want: models.TodoItem{Id: 2, Title: "new title", Description: "new description", Status: models.TodoItemStatus(0).String(), Done: true},
		},
		{
			name: "OK_WithoutStatusAndDone",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          3,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
				})
				if err != nil {
					panic(err)
				}
			},
			input: args{
				itemId: 3,
				input: models.UpdateTodoItemInput{
					Title:       stringPointer("new title"),
					Description: stringPointer("new description"),
				},
			},
			want: models.TodoItem{Id: 3, Title: "new title", Description: "new description", Status: models.TodoItemStatus(0).String(), Done: false},
		},
		{
			name: "OK_WithoutDoneAndDescription",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          4,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
				})
				if err != nil {
					panic(err)
				}
			},
			input: args{
				itemId: 4,
				input: models.UpdateTodoItemInput{
					Title:  stringPointer("new title"),
					Status: intPointer(1),
				},
			},
			want: models.TodoItem{Id: 4, Title: "new title", Description: "description", Status: models.TodoItemStatus(1).String(), Done: false},
		},
		{
			name: "OK_NoInputFields",
			mock: func() {
				_, err := service.Create(models.TodoItem{
					Id:          5,
					Title:       "title",
					Description: "description",
					Status:      models.TodoItemStatus(0).String(),
				})
				if err != nil {
					panic(err)
				}
			},
			input: args{
				itemId: 5,
			},
			want: models.TodoItem{Id: 5, Title: "title", Description: "description", Status: models.TodoItemStatus(0).String(), Done: false},
		},
		{
			name: "Not Found",
			mock: func() {},
			input: args{
				itemId: 404,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := service.Update(tt.input.itemId, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func stringPointer(s string) *string {
	return &s
}

func boolPointer(b bool) *bool {
	return &b
}

func intPointer(b int) *int {
	return &b
}
