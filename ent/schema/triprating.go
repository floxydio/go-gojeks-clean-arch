package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// TripRating holds the schema definition for the TripRating entity.
type TripRating struct {
	ent.Schema
}

// Fields of the TripRating.
func (TripRating) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("trip_id", uuid.UUID{}),
		field.UUID("from_user_id", uuid.UUID{}),
		field.UUID("to_user_id", uuid.UUID{}),
		field.Int("rating"),
		field.Text("comment").Default("-"),
		field.Time("created_at").Default(time.Now()).Annotations(
			entsql.Default("CURRENT_TIMESTAMP"),
		),
	}
}

// Edges of the TripRating.
func (TripRating) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trip", Trip.Type).Ref("ratings").Field("trip_id").Unique().Required(),
		edge.From("from_user", User.Type).
			Ref("given_ratings").
			Field("from_user_id").
			Unique().Required(),
		edge.From("to_user", User.Type).
			Ref("received_ratings").
			Field("to_user_id").
			Unique().Required(),
	}
}
