package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type Membership struct {
	ent.Schema
}

func (Membership) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StorageKey("id").
			Unique().
			Default(uuid.New),
		field.String("role"),
	}
}

func (Membership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("member", User.Type).
			Unique(),
		edge.To("org", Org.Type).
			Unique(),
	}
}
