package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TodoItem holds the migrations definition for the TodoItem entity.
type TodoItem struct {
	ent.Schema
}

// Fields of the TodoItem.
func (TodoItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255),
		field.String("description").MaxLen(500),
		field.String("status").MaxLen(25),
		field.Bool("done").Default(false),
	}
}

// Edges of the TodoItem.
func (TodoItem) Edges() []ent.Edge {
	return nil
}
