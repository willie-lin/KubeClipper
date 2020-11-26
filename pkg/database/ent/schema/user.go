package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("username"),
		field.String("password").MinLen(12).MaxLen(128).Sensitive(),
		field.String("email").MinLen(8),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
