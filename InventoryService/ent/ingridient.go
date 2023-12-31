// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"inventoryService/m/v2/ent/ingridient"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Ingridient is the model entity for the Ingridient schema.
type Ingridient struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Alcohol holds the value of the "alcohol" field.
	Alcohol float64 `json:"alcohol,omitempty"`
	// IsDry holds the value of the "isDry" field.
	IsDry bool `json:"isDry,omitempty"`
	// MeasurmentUnit holds the value of the "measurmentUnit" field.
	MeasurmentUnit ingridient.MeasurmentUnit `json:"measurmentUnit,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IngridientQuery when eager-loading is set.
	Edges        IngridientEdges `json:"edges"`
	selectValues sql.SelectValues
}

// IngridientEdges holds the relations/edges for other nodes in the graph.
type IngridientEdges struct {
	// Stock holds the value of the stock edge.
	Stock []*Stockpile `json:"stock,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StockOrErr returns the Stock value or an error if the edge
// was not loaded in eager-loading.
func (e IngridientEdges) StockOrErr() ([]*Stockpile, error) {
	if e.loadedTypes[0] {
		return e.Stock, nil
	}
	return nil, &NotLoadedError{edge: "stock"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ingridient) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ingridient.FieldIsDry:
			values[i] = new(sql.NullBool)
		case ingridient.FieldAlcohol:
			values[i] = new(sql.NullFloat64)
		case ingridient.FieldName, ingridient.FieldMeasurmentUnit:
			values[i] = new(sql.NullString)
		case ingridient.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ingridient fields.
func (i *Ingridient) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case ingridient.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case ingridient.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case ingridient.FieldAlcohol:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field alcohol", values[j])
			} else if value.Valid {
				i.Alcohol = value.Float64
			}
		case ingridient.FieldIsDry:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isDry", values[j])
			} else if value.Valid {
				i.IsDry = value.Bool
			}
		case ingridient.FieldMeasurmentUnit:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field measurmentUnit", values[j])
			} else if value.Valid {
				i.MeasurmentUnit = ingridient.MeasurmentUnit(value.String)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Ingridient.
// This includes values selected through modifiers, order, etc.
func (i *Ingridient) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryStock queries the "stock" edge of the Ingridient entity.
func (i *Ingridient) QueryStock() *StockpileQuery {
	return NewIngridientClient(i.config).QueryStock(i)
}

// Update returns a builder for updating this Ingridient.
// Note that you need to call Ingridient.Unwrap() before calling this method if this Ingridient
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Ingridient) Update() *IngridientUpdateOne {
	return NewIngridientClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Ingridient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Ingridient) Unwrap() *Ingridient {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ingridient is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Ingridient) String() string {
	var builder strings.Builder
	builder.WriteString("Ingridient(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("alcohol=")
	builder.WriteString(fmt.Sprintf("%v", i.Alcohol))
	builder.WriteString(", ")
	builder.WriteString("isDry=")
	builder.WriteString(fmt.Sprintf("%v", i.IsDry))
	builder.WriteString(", ")
	builder.WriteString("measurmentUnit=")
	builder.WriteString(fmt.Sprintf("%v", i.MeasurmentUnit))
	builder.WriteByte(')')
	return builder.String()
}

// Ingridients is a parsable slice of Ingridient.
type Ingridients []*Ingridient
