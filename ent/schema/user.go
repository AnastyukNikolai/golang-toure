package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the migrations definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255),
		field.String("username").MaxLen(500),
		field.String("password").Sensitive(),
	}
}

// Edges of the TodoItem.
func (User) Edges() []ent.Edge {
	return nil
}
