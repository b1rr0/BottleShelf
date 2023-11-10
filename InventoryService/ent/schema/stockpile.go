package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Stockpile holds the schema definition for the Stockpile entity.
type Stockpile struct {
	ent.Schema
}

// Fields of the Stockpile.
func (Stockpile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("ingridientId", uuid.UUID{}),
		field.UUID("ownerId", uuid.UUID{}),
		field.Int64("quantity"),
	}
}

// Edges of the Stockpile.
func (Stockpile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ingridient", Ingridient.Type).Ref("stock").Unique(),
	}
}
