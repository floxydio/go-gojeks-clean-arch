package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("trip_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("payment_method").NotEmpty(),
		field.Float("amount"),
		field.Enum("status").Values("pending", "success", "failed"),
		field.Time("paid_at"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trip", Trip.Type).
			Ref("payment").
			Field("trip_id").
			Unique().Required(),
		edge.From("user", User.Type).
			Ref("payments").
			Field("user_id").Unique().Required(),
	}
}
