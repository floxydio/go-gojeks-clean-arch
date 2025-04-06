package schema

import (
	"entgo.io/ent/schema/edge"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("phone").Unique().NotEmpty(),
		field.String("password").NotEmpty(),
		field.Enum("role").Values("user", "driver", "admin"),
		field.Bool("is_verified").Default(false),
		field.Time("created_at").Default(time.Now()).Annotations(
			entsql.Default("CURRENT_TIMESTAMP"),
		),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_trips", Trip.Type),
		edge.To("payments", Payment.Type),
		edge.To("user_balance", Wallet.Type),
		edge.To("given_ratings", TripRating.Type),
		edge.To("received_ratings", TripRating.Type),
		edge.To("user_driver", DriverProfile.Type),
	}
}
