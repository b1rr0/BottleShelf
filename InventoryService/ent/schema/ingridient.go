package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Ingridient holds the schema definition for the Ingridient entity.
type Ingridient struct {
	ent.Schema
}

// Fields of the Ingridient.
func (Ingridient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id"),
		field.String("name").Unique().NotEmpty(),
		field.Float("alcohol").Default(0),
		field.Bool("isDry").Default(false),
		field.Enum("measurmentUnit").Values("ml", "g", "pcs").Default("ml"),
	}
}

// Edges of the Ingridient.
func (Ingridient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stock", Stockpile.Type),
	}
}
