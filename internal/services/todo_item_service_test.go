package services

import (
	"github.com/stretchr/testify/assert"
	"golang-ture/internal/models"
	"golang-ture/internal/repositories"
	"testing"
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
					Status:      "Backlog",
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
					Status:      "Backlog",
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
				service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
					Done:        true,
				})
			},
			want: []models.TodoItem{
				{Id: 1, Title: "title", Description: "description", Status: "Backlog", Done: true},
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
				service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
					Done:        true,
				})
			},
			input: args{
				itemId: 1,
			},
			want: models.TodoItem{Id: 1, Title: "title", Description: "description", Status: "Backlog", Done: true},
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
				service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
					Done:        true,
				})
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
				service.Create(models.TodoItem{
					Id:          1,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
				})
			},
			input: args{
				itemId: 1,
				input: models.UpdateTodoItemInput{
					Title:       stringPointer("new title"),
					Description: stringPointer("new description"),
					Status:      stringPointer("In_Progress"),
					Done:        boolPointer(true),
				},
			},
			want: models.TodoItem{Id: 1, Title: "new title", Description: "new description", Status: "In_Progress", Done: true},
		},
		{
			name: "OK_WithoutStatus",
			mock: func() {
				service.Create(models.TodoItem{
					Id:          2,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
				})
			},
			input: args{
				itemId: 2,
				input: models.UpdateTodoItemInput{
					Title:       stringPointer("new title"),
					Description: stringPointer("new description"),
					Done:        boolPointer(true),
				},
			},
			want: models.TodoItem{Id: 2, Title: "new title", Description: "new description", Status: "Backlog", Done: true},
		},
		{
			name: "OK_WithoutStatusAndDone",
			mock: func() {
				service.Create(models.TodoItem{
					Id:          3,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
				})
			},
			input: args{
				itemId: 3,
				input: models.UpdateTodoItemInput{
					Title:       stringPointer("new title"),
					Description: stringPointer("new description"),
				},
			},
			want: models.TodoItem{Id: 3, Title: "new title", Description: "new description", Status: "Backlog", Done: false},
		},
		{
			name: "OK_WithoutDoneAndDescription",
			mock: func() {
				service.Create(models.TodoItem{
					Id:          4,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
				})
			},
			input: args{
				itemId: 4,
				input: models.UpdateTodoItemInput{
					Title:  stringPointer("new title"),
					Status: stringPointer("In_Progress"),
				},
			},
			want: models.TodoItem{Id: 4, Title: "new title", Description: "description", Status: "In_Progress", Done: false},
		},
		{
			name: "OK_NoInputFields",
			mock: func() {
				service.Create(models.TodoItem{
					Id:          5,
					Title:       "title",
					Description: "description",
					Status:      "Backlog",
				})
			},
			input: args{
				itemId: 5,
			},
			want: models.TodoItem{Id: 5, Title: "title", Description: "description", Status: "Backlog", Done: false},
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
