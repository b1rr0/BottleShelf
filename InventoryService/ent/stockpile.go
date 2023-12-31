// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"inventoryService/m/v2/ent/ingridient"
	"inventoryService/m/v2/ent/stockpile"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Stockpile is the model entity for the Stockpile schema.
type Stockpile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IngridientId holds the value of the "ingridientId" field.
	IngridientId uuid.UUID `json:"ingridientId,omitempty"`
	// OwnerId holds the value of the "ownerId" field.
	OwnerId uuid.UUID `json:"ownerId,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int64 `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StockpileQuery when eager-loading is set.
	Edges            StockpileEdges `json:"edges"`
	ingridient_stock *uuid.UUID
	selectValues     sql.SelectValues
}

// StockpileEdges holds the relations/edges for other nodes in the graph.
type StockpileEdges struct {
	// Ingridient holds the value of the ingridient edge.
	Ingridient *Ingridient `json:"ingridient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// IngridientOrErr returns the Ingridient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StockpileEdges) IngridientOrErr() (*Ingridient, error) {
	if e.loadedTypes[0] {
		if e.Ingridient == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ingridient.Label}
		}
		return e.Ingridient, nil
	}
	return nil, &NotLoadedError{edge: "ingridient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stockpile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case stockpile.FieldID, stockpile.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case stockpile.FieldIngridientId, stockpile.FieldOwnerId:
			values[i] = new(uuid.UUID)
		case stockpile.ForeignKeys[0]: // ingridient_stock
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stockpile fields.
func (s *Stockpile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stockpile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case stockpile.FieldIngridientId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ingridientId", values[i])
			} else if value != nil {
				s.IngridientId = *value
			}
		case stockpile.FieldOwnerId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ownerId", values[i])
			} else if value != nil {
				s.OwnerId = *value
			}
		case stockpile.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				s.Quantity = value.Int64
			}
		case stockpile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field ingridient_stock", values[i])
			} else if value.Valid {
				s.ingridient_stock = new(uuid.UUID)
				*s.ingridient_stock = *value.S.(*uuid.UUID)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Stockpile.
// This includes values selected through modifiers, order, etc.
func (s *Stockpile) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryIngridient queries the "ingridient" edge of the Stockpile entity.
func (s *Stockpile) QueryIngridient() *IngridientQuery {
	return NewStockpileClient(s.config).QueryIngridient(s)
}

// Update returns a builder for updating this Stockpile.
// Note that you need to call Stockpile.Unwrap() before calling this method if this Stockpile
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stockpile) Update() *StockpileUpdateOne {
	return NewStockpileClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Stockpile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stockpile) Unwrap() *Stockpile {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stockpile is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stockpile) String() string {
	var builder strings.Builder
	builder.WriteString("Stockpile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("ingridientId=")
	builder.WriteString(fmt.Sprintf("%v", s.IngridientId))
	builder.WriteString(", ")
	builder.WriteString("ownerId=")
	builder.WriteString(fmt.Sprintf("%v", s.OwnerId))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", s.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// Stockpiles is a parsable slice of Stockpile.
type Stockpiles []*Stockpile
