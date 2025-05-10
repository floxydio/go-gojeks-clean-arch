package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UsersAdmin holds the schema definition for the UsersAdmin entity.
type UsersAdmin struct {
	ent.Schema
}

// Fields of the UsersAdmin.
func (UsersAdmin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.Int("status_admin").Default(0),
	}
}

// Edges of the UsersAdmin.
func (UsersAdmin) Edges() []ent.Edge {
	return nil
}
