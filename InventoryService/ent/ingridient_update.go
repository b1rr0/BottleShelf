// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"inventoryService/m/v2/ent/ingridient"
	"inventoryService/m/v2/ent/predicate"
	"inventoryService/m/v2/ent/stockpile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IngridientUpdate is the builder for updating Ingridient entities.
type IngridientUpdate struct {
	config
	hooks    []Hook
	mutation *IngridientMutation
}

// Where appends a list predicates to the IngridientUpdate builder.
func (iu *IngridientUpdate) Where(ps ...predicate.Ingridient) *IngridientUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *IngridientUpdate) SetName(s string) *IngridientUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetAlcohol sets the "alcohol" field.
func (iu *IngridientUpdate) SetAlcohol(f float64) *IngridientUpdate {
	iu.mutation.ResetAlcohol()
	iu.mutation.SetAlcohol(f)
	return iu
}

// SetNillableAlcohol sets the "alcohol" field if the given value is not nil.
func (iu *IngridientUpdate) SetNillableAlcohol(f *float64) *IngridientUpdate {
	if f != nil {
		iu.SetAlcohol(*f)
	}
	return iu
}

// AddAlcohol adds f to the "alcohol" field.
func (iu *IngridientUpdate) AddAlcohol(f float64) *IngridientUpdate {
	iu.mutation.AddAlcohol(f)
	return iu
}

// SetIsDry sets the "isDry" field.
func (iu *IngridientUpdate) SetIsDry(b bool) *IngridientUpdate {
	iu.mutation.SetIsDry(b)
	return iu
}

// SetNillableIsDry sets the "isDry" field if the given value is not nil.
func (iu *IngridientUpdate) SetNillableIsDry(b *bool) *IngridientUpdate {
	if b != nil {
		iu.SetIsDry(*b)
	}
	return iu
}

// SetMeasurmentUnit sets the "measurmentUnit" field.
func (iu *IngridientUpdate) SetMeasurmentUnit(value ingridient.MeasurmentUnit) *IngridientUpdate {
	iu.mutation.SetMeasurmentUnit(value)
	return iu
}

// SetNillableMeasurmentUnit sets the "measurmentUnit" field if the given value is not nil.
func (iu *IngridientUpdate) SetNillableMeasurmentUnit(value *ingridient.MeasurmentUnit) *IngridientUpdate {
	if value != nil {
		iu.SetMeasurmentUnit(*value)
	}
	return iu
}

// AddStockIDs adds the "stock" edge to the Stockpile entity by IDs.
func (iu *IngridientUpdate) AddStockIDs(ids ...int) *IngridientUpdate {
	iu.mutation.AddStockIDs(ids...)
	return iu
}

// AddStock adds the "stock" edges to the Stockpile entity.
func (iu *IngridientUpdate) AddStock(s ...*Stockpile) *IngridientUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iu.AddStockIDs(ids...)
}

// Mutation returns the IngridientMutation object of the builder.
func (iu *IngridientUpdate) Mutation() *IngridientMutation {
	return iu.mutation
}

// ClearStock clears all "stock" edges to the Stockpile entity.
func (iu *IngridientUpdate) ClearStock() *IngridientUpdate {
	iu.mutation.ClearStock()
	return iu
}

// RemoveStockIDs removes the "stock" edge to Stockpile entities by IDs.
func (iu *IngridientUpdate) RemoveStockIDs(ids ...int) *IngridientUpdate {
	iu.mutation.RemoveStockIDs(ids...)
	return iu
}

// RemoveStock removes "stock" edges to Stockpile entities.
func (iu *IngridientUpdate) RemoveStock(s ...*Stockpile) *IngridientUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iu.RemoveStockIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IngridientUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IngridientUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IngridientUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IngridientUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IngridientUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := ingridient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Ingridient.name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.MeasurmentUnit(); ok {
		if err := ingridient.MeasurmentUnitValidator(v); err != nil {
			return &ValidationError{Name: "measurmentUnit", err: fmt.Errorf(`ent: validator failed for field "Ingridient.measurmentUnit": %w`, err)}
		}
	}
	return nil
}

func (iu *IngridientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ingridient.Table, ingridient.Columns, sqlgraph.NewFieldSpec(ingridient.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(ingridient.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Alcohol(); ok {
		_spec.SetField(ingridient.FieldAlcohol, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.AddedAlcohol(); ok {
		_spec.AddField(ingridient.FieldAlcohol, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.IsDry(); ok {
		_spec.SetField(ingridient.FieldIsDry, field.TypeBool, value)
	}
	if value, ok := iu.mutation.MeasurmentUnit(); ok {
		_spec.SetField(ingridient.FieldMeasurmentUnit, field.TypeEnum, value)
	}
	if iu.mutation.StockCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingridient.StockTable,
			Columns: []string{ingridient.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedStockIDs(); len(nodes) > 0 && !iu.mutation.StockCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingridient.StockTable,
			Columns: []string{ingridient.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.StockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingridient.StockTable,
			Columns: []string{ingridient.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ingridient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IngridientUpdateOne is the builder for updating a single Ingridient entity.
type IngridientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IngridientMutation
}

// SetName sets the "name" field.
func (iuo *IngridientUpdateOne) SetName(s string) *IngridientUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetAlcohol sets the "alcohol" field.
func (iuo *IngridientUpdateOne) SetAlcohol(f float64) *IngridientUpdateOne {
	iuo.mutation.ResetAlcohol()
	iuo.mutation.SetAlcohol(f)
	return iuo
}

// SetNillableAlcohol sets the "alcohol" field if the given value is not nil.
func (iuo *IngridientUpdateOne) SetNillableAlcohol(f *float64) *IngridientUpdateOne {
	if f != nil {
		iuo.SetAlcohol(*f)
	}
	return iuo
}

// AddAlcohol adds f to the "alcohol" field.
func (iuo *IngridientUpdateOne) AddAlcohol(f float64) *IngridientUpdateOne {
	iuo.mutation.AddAlcohol(f)
	return iuo
}

// SetIsDry sets the "isDry" field.
func (iuo *IngridientUpdateOne) SetIsDry(b bool) *IngridientUpdateOne {
	iuo.mutation.SetIsDry(b)
	return iuo
}

// SetNillableIsDry sets the "isDry" field if the given value is not nil.
func (iuo *IngridientUpdateOne) SetNillableIsDry(b *bool) *IngridientUpdateOne {
	if b != nil {
		iuo.SetIsDry(*b)
	}
	return iuo
}

// SetMeasurmentUnit sets the "measurmentUnit" field.
func (iuo *IngridientUpdateOne) SetMeasurmentUnit(iu ingridient.MeasurmentUnit) *IngridientUpdateOne {
	iuo.mutation.SetMeasurmentUnit(iu)
	return iuo
}

// SetNillableMeasurmentUnit sets the "measurmentUnit" field if the given value is not nil.
func (iuo *IngridientUpdateOne) SetNillableMeasurmentUnit(iu *ingridient.MeasurmentUnit) *IngridientUpdateOne {
	if iu != nil {
		iuo.SetMeasurmentUnit(*iu)
	}
	return iuo
}

// AddStockIDs adds the "stock" edge to the Stockpile entity by IDs.
func (iuo *IngridientUpdateOne) AddStockIDs(ids ...int) *IngridientUpdateOne {
	iuo.mutation.AddStockIDs(ids...)
	return iuo
}

// AddStock adds the "stock" edges to the Stockpile entity.
func (iuo *IngridientUpdateOne) AddStock(s ...*Stockpile) *IngridientUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iuo.AddStockIDs(ids...)
}

// Mutation returns the IngridientMutation object of the builder.
func (iuo *IngridientUpdateOne) Mutation() *IngridientMutation {
	return iuo.mutation
}

// ClearStock clears all "stock" edges to the Stockpile entity.
func (iuo *IngridientUpdateOne) ClearStock() *IngridientUpdateOne {
	iuo.mutation.ClearStock()
	return iuo
}

// RemoveStockIDs removes the "stock" edge to Stockpile entities by IDs.
func (iuo *IngridientUpdateOne) RemoveStockIDs(ids ...int) *IngridientUpdateOne {
	iuo.mutation.RemoveStockIDs(ids...)
	return iuo
}

// RemoveStock removes "stock" edges to Stockpile entities.
func (iuo *IngridientUpdateOne) RemoveStock(s ...*Stockpile) *IngridientUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return iuo.RemoveStockIDs(ids...)
}

// Where appends a list predicates to the IngridientUpdate builder.
func (iuo *IngridientUpdateOne) Where(ps ...predicate.Ingridient) *IngridientUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IngridientUpdateOne) Select(field string, fields ...string) *IngridientUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Ingridient entity.
func (iuo *IngridientUpdateOne) Save(ctx context.Context) (*Ingridient, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IngridientUpdateOne) SaveX(ctx context.Context) *Ingridient {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IngridientUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IngridientUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IngridientUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := ingridient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Ingridient.name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.MeasurmentUnit(); ok {
		if err := ingridient.MeasurmentUnitValidator(v); err != nil {
			return &ValidationError{Name: "measurmentUnit", err: fmt.Errorf(`ent: validator failed for field "Ingridient.measurmentUnit": %w`, err)}
		}
	}
	return nil
}

func (iuo *IngridientUpdateOne) sqlSave(ctx context.Context) (_node *Ingridient, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ingridient.Table, ingridient.Columns, sqlgraph.NewFieldSpec(ingridient.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ingridient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ingridient.FieldID)
		for _, f := range fields {
			if !ingridient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ingridient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(ingridient.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Alcohol(); ok {
		_spec.SetField(ingridient.FieldAlcohol, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.AddedAlcohol(); ok {
		_spec.AddField(ingridient.FieldAlcohol, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.IsDry(); ok {
		_spec.SetField(ingridient.FieldIsDry, field.TypeBool, value)
	}
	if value, ok := iuo.mutation.MeasurmentUnit(); ok {
		_spec.SetField(ingridient.FieldMeasurmentUnit, field.TypeEnum, value)
	}
	if iuo.mutation.StockCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingridient.StockTable,
			Columns: []string{ingridient.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedStockIDs(); len(nodes) > 0 && !iuo.mutation.StockCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingridient.StockTable,
			Columns: []string{ingridient.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.StockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingridient.StockTable,
			Columns: []string{ingridient.StockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ingridient{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ingridient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
