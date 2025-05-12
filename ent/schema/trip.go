package schema

import (
	"time"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Trip holds the schema definition for the Trip entity.
type Trip struct {
	ent.Schema
}

// Fields of the Trip.
func (Trip) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("driver_id", uuid.UUID{}).Optional().Nillable(),
		field.Float("pickup_lat"),
		field.Float("pickup_long"),
		field.Float("drop_lat"),
		field.Float("drop_long"),
		field.Enum("status").Values("requested", "accepted", "ongoing", "completed", "cancelled"),
		field.Float("distance_km").Optional(),
		field.String("numeric").Optional(),
		field.Bool("is_paid").Default(false),
		field.Time("created_at").Default(time.Now()).Annotations(
			entsql.Default("CURRENT_TIMESTAMP"),
		),
		field.Time("started_at").Optional(),
		field.Time("completed_at").Optional(),
	}
}

// Edges of the Trip.
func (Trip) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_trips").Field("user_id").Unique().Required(),
		edge.From("driver", User.Type).Ref("driver_trips").Field("driver_id").Unique(),
		edge.To("payment", Payment.Type),
		edge.To("ratings", TripRating.Type),
	}
}
