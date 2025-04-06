package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DriverProfile holds the schema definition for the DriverProfile entity.
type DriverProfile struct {
	ent.Schema
}

// Fields of the DriverProfile.
func (DriverProfile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("license_number"),
		field.String("ktp_number"),
		field.Enum("status").Values("pending", "approved", "reject").Default("pending"),
		field.String("vehicle_info"),
		field.Float("current_lat").Optional().Nillable(),
		field.Float("current_long").Optional().Nillable(),
		field.Bool("is_active").Default(false),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the DriverProfile.
func (DriverProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_driver").Field("user_id").Unique().Required(),
		edge.To("trips_driver", Trip.Type),
	}
}
