// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"inventoryService/m/v2/ent/ingridient"
	"inventoryService/m/v2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IngridientDelete is the builder for deleting a Ingridient entity.
type IngridientDelete struct {
	config
	hooks    []Hook
	mutation *IngridientMutation
}

// Where appends a list predicates to the IngridientDelete builder.
func (id *IngridientDelete) Where(ps ...predicate.Ingridient) *IngridientDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *IngridientDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *IngridientDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *IngridientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ingridient.Table, sqlgraph.NewFieldSpec(ingridient.FieldID, field.TypeUUID))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// IngridientDeleteOne is the builder for deleting a single Ingridient entity.
type IngridientDeleteOne struct {
	id *IngridientDelete
}

// Where appends a list predicates to the IngridientDelete builder.
func (ido *IngridientDeleteOne) Where(ps ...predicate.Ingridient) *IngridientDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *IngridientDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ingridient.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *IngridientDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}