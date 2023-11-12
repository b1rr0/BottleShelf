package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type Org struct {
	ent.Schema
}

func (Org) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StorageKey("id").
			Unique(),
		field.String("name"),
	}
}

func (Org) Edges() []ent.Edge {
	return nil
}
