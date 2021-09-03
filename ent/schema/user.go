package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			NamedValues(
				"Actived", "ACTIVED",
				"InActived", "INACTIVED",
			),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.String("name"),
		field.String("email"),
		field.String("password"),
		field.String("refresh_token"),
		field.String("provider_id"),
		field.String("provider_name"),
		field.String("role").Default("user"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
