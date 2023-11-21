// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"inventoryService/m/v2/ent/ingridient"
	"inventoryService/m/v2/ent/predicate"
	"inventoryService/m/v2/ent/stockpile"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeIngridient = "Ingridient"
	TypeStockpile  = "Stockpile"
)

// IngridientMutation represents an operation that mutates the Ingridient nodes in the graph.
type IngridientMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	name           *string
	alcohol        *float64
	addalcohol     *float64
	isDry          *bool
	measurmentUnit *ingridient.MeasurmentUnit
	clearedFields  map[string]struct{}
	stock          map[int]struct{}
	removedstock   map[int]struct{}
	clearedstock   bool
	done           bool
	oldValue       func(context.Context) (*Ingridient, error)
	predicates     []predicate.Ingridient
}

var _ ent.Mutation = (*IngridientMutation)(nil)

// ingridientOption allows management of the mutation configuration using functional options.
type ingridientOption func(*IngridientMutation)

// newIngridientMutation creates new mutation for the Ingridient entity.
func newIngridientMutation(c config, op Op, opts ...ingridientOption) *IngridientMutation {
	m := &IngridientMutation{
		config:        c,
		op:            op,
		typ:           TypeIngridient,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withIngridientID sets the ID field of the mutation.
func withIngridientID(id uuid.UUID) ingridientOption {
	return func(m *IngridientMutation) {
		var (
			err   error
			once  sync.Once
			value *Ingridient
		)
		m.oldValue = func(ctx context.Context) (*Ingridient, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Ingridient.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withIngridient sets the old Ingridient of the mutation.
func withIngridient(node *Ingridient) ingridientOption {
	return func(m *IngridientMutation) {
		m.oldValue = func(context.Context) (*Ingridient, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m IngridientMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m IngridientMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Ingridient entities.
func (m *IngridientMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *IngridientMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *IngridientMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Ingridient.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *IngridientMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *IngridientMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Ingridient entity.
// If the Ingridient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IngridientMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *IngridientMutation) ResetName() {
	m.name = nil
}

// SetAlcohol sets the "alcohol" field.
func (m *IngridientMutation) SetAlcohol(f float64) {
	m.alcohol = &f
	m.addalcohol = nil
}

// Alcohol returns the value of the "alcohol" field in the mutation.
func (m *IngridientMutation) Alcohol() (r float64, exists bool) {
	v := m.alcohol
	if v == nil {
		return
	}
	return *v, true
}

// OldAlcohol returns the old "alcohol" field's value of the Ingridient entity.
// If the Ingridient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IngridientMutation) OldAlcohol(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAlcohol is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAlcohol requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAlcohol: %w", err)
	}
	return oldValue.Alcohol, nil
}

// AddAlcohol adds f to the "alcohol" field.
func (m *IngridientMutation) AddAlcohol(f float64) {
	if m.addalcohol != nil {
		*m.addalcohol += f
	} else {
		m.addalcohol = &f
	}
}

// AddedAlcohol returns the value that was added to the "alcohol" field in this mutation.
func (m *IngridientMutation) AddedAlcohol() (r float64, exists bool) {
	v := m.addalcohol
	if v == nil {
		return
	}
	return *v, true
}

// ResetAlcohol resets all changes to the "alcohol" field.
func (m *IngridientMutation) ResetAlcohol() {
	m.alcohol = nil
	m.addalcohol = nil
}

// SetIsDry sets the "isDry" field.
func (m *IngridientMutation) SetIsDry(b bool) {
	m.isDry = &b
}

// IsDry returns the value of the "isDry" field in the mutation.
func (m *IngridientMutation) IsDry() (r bool, exists bool) {
	v := m.isDry
	if v == nil {
		return
	}
	return *v, true
}

// OldIsDry returns the old "isDry" field's value of the Ingridient entity.
// If the Ingridient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IngridientMutation) OldIsDry(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsDry is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsDry requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsDry: %w", err)
	}
	return oldValue.IsDry, nil
}

// ResetIsDry resets all changes to the "isDry" field.
func (m *IngridientMutation) ResetIsDry() {
	m.isDry = nil
}

// SetMeasurmentUnit sets the "measurmentUnit" field.
func (m *IngridientMutation) SetMeasurmentUnit(iu ingridient.MeasurmentUnit) {
	m.measurmentUnit = &iu
}

// MeasurmentUnit returns the value of the "measurmentUnit" field in the mutation.
func (m *IngridientMutation) MeasurmentUnit() (r ingridient.MeasurmentUnit, exists bool) {
	v := m.measurmentUnit
	if v == nil {
		return
	}
	return *v, true
}

// OldMeasurmentUnit returns the old "measurmentUnit" field's value of the Ingridient entity.
// If the Ingridient object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IngridientMutation) OldMeasurmentUnit(ctx context.Context) (v ingridient.MeasurmentUnit, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMeasurmentUnit is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMeasurmentUnit requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMeasurmentUnit: %w", err)
	}
	return oldValue.MeasurmentUnit, nil
}

// ResetMeasurmentUnit resets all changes to the "measurmentUnit" field.
func (m *IngridientMutation) ResetMeasurmentUnit() {
	m.measurmentUnit = nil
}

// AddStockIDs adds the "stock" edge to the Stockpile entity by ids.
func (m *IngridientMutation) AddStockIDs(ids ...int) {
	if m.stock == nil {
		m.stock = make(map[int]struct{})
	}
	for i := range ids {
		m.stock[ids[i]] = struct{}{}
	}
}

// ClearStock clears the "stock" edge to the Stockpile entity.
func (m *IngridientMutation) ClearStock() {
	m.clearedstock = true
}

// StockCleared reports if the "stock" edge to the Stockpile entity was cleared.
func (m *IngridientMutation) StockCleared() bool {
	return m.clearedstock
}

// RemoveStockIDs removes the "stock" edge to the Stockpile entity by IDs.
func (m *IngridientMutation) RemoveStockIDs(ids ...int) {
	if m.removedstock == nil {
		m.removedstock = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.stock, ids[i])
		m.removedstock[ids[i]] = struct{}{}
	}
}

// RemovedStock returns the removed IDs of the "stock" edge to the Stockpile entity.
func (m *IngridientMutation) RemovedStockIDs() (ids []int) {
	for id := range m.removedstock {
		ids = append(ids, id)
	}
	return
}

// StockIDs returns the "stock" edge IDs in the mutation.
func (m *IngridientMutation) StockIDs() (ids []int) {
	for id := range m.stock {
		ids = append(ids, id)
	}
	return
}

// ResetStock resets all changes to the "stock" edge.
func (m *IngridientMutation) ResetStock() {
	m.stock = nil
	m.clearedstock = false
	m.removedstock = nil
}

// Where appends a list predicates to the IngridientMutation builder.
func (m *IngridientMutation) Where(ps ...predicate.Ingridient) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the IngridientMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *IngridientMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Ingridient, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *IngridientMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *IngridientMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Ingridient).
func (m *IngridientMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *IngridientMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, ingridient.FieldName)
	}
	if m.alcohol != nil {
		fields = append(fields, ingridient.FieldAlcohol)
	}
	if m.isDry != nil {
		fields = append(fields, ingridient.FieldIsDry)
	}
	if m.measurmentUnit != nil {
		fields = append(fields, ingridient.FieldMeasurmentUnit)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *IngridientMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case ingridient.FieldName:
		return m.Name()
	case ingridient.FieldAlcohol:
		return m.Alcohol()
	case ingridient.FieldIsDry:
		return m.IsDry()
	case ingridient.FieldMeasurmentUnit:
		return m.MeasurmentUnit()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *IngridientMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case ingridient.FieldName:
		return m.OldName(ctx)
	case ingridient.FieldAlcohol:
		return m.OldAlcohol(ctx)
	case ingridient.FieldIsDry:
		return m.OldIsDry(ctx)
	case ingridient.FieldMeasurmentUnit:
		return m.OldMeasurmentUnit(ctx)
	}
	return nil, fmt.Errorf("unknown Ingridient field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *IngridientMutation) SetField(name string, value ent.Value) error {
	switch name {
	case ingridient.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case ingridient.FieldAlcohol:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAlcohol(v)
		return nil
	case ingridient.FieldIsDry:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsDry(v)
		return nil
	case ingridient.FieldMeasurmentUnit:
		v, ok := value.(ingridient.MeasurmentUnit)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMeasurmentUnit(v)
		return nil
	}
	return fmt.Errorf("unknown Ingridient field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *IngridientMutation) AddedFields() []string {
	var fields []string
	if m.addalcohol != nil {
		fields = append(fields, ingridient.FieldAlcohol)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *IngridientMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case ingridient.FieldAlcohol:
		return m.AddedAlcohol()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *IngridientMutation) AddField(name string, value ent.Value) error {
	switch name {
	case ingridient.FieldAlcohol:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAlcohol(v)
		return nil
	}
	return fmt.Errorf("unknown Ingridient numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *IngridientMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *IngridientMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *IngridientMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Ingridient nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *IngridientMutation) ResetField(name string) error {
	switch name {
	case ingridient.FieldName:
		m.ResetName()
		return nil
	case ingridient.FieldAlcohol:
		m.ResetAlcohol()
		return nil
	case ingridient.FieldIsDry:
		m.ResetIsDry()
		return nil
	case ingridient.FieldMeasurmentUnit:
		m.ResetMeasurmentUnit()
		return nil
	}
	return fmt.Errorf("unknown Ingridient field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *IngridientMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.stock != nil {
		edges = append(edges, ingridient.EdgeStock)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *IngridientMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case ingridient.EdgeStock:
		ids := make([]ent.Value, 0, len(m.stock))
		for id := range m.stock {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *IngridientMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedstock != nil {
		edges = append(edges, ingridient.EdgeStock)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *IngridientMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case ingridient.EdgeStock:
		ids := make([]ent.Value, 0, len(m.removedstock))
		for id := range m.removedstock {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *IngridientMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedstock {
		edges = append(edges, ingridient.EdgeStock)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *IngridientMutation) EdgeCleared(name string) bool {
	switch name {
	case ingridient.EdgeStock:
		return m.clearedstock
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *IngridientMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Ingridient unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *IngridientMutation) ResetEdge(name string) error {
	switch name {
	case ingridient.EdgeStock:
		m.ResetStock()
		return nil
	}
	return fmt.Errorf("unknown Ingridient edge %s", name)
}

// StockpileMutation represents an operation that mutates the Stockpile nodes in the graph.
type StockpileMutation struct {
	config
	op                Op
	typ               string
	id                *int
	ingridientId      *uuid.UUID
	ownerId           *uuid.UUID
	quantity          *int64
	addquantity       *int64
	clearedFields     map[string]struct{}
	ingridient        *uuid.UUID
	clearedingridient bool
	done              bool
	oldValue          func(context.Context) (*Stockpile, error)
	predicates        []predicate.Stockpile
}

var _ ent.Mutation = (*StockpileMutation)(nil)

// stockpileOption allows management of the mutation configuration using functional options.
type stockpileOption func(*StockpileMutation)

// newStockpileMutation creates new mutation for the Stockpile entity.
func newStockpileMutation(c config, op Op, opts ...stockpileOption) *StockpileMutation {
	m := &StockpileMutation{
		config:        c,
		op:            op,
		typ:           TypeStockpile,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStockpileID sets the ID field of the mutation.
func withStockpileID(id int) stockpileOption {
	return func(m *StockpileMutation) {
		var (
			err   error
			once  sync.Once
			value *Stockpile
		)
		m.oldValue = func(ctx context.Context) (*Stockpile, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Stockpile.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStockpile sets the old Stockpile of the mutation.
func withStockpile(node *Stockpile) stockpileOption {
	return func(m *StockpileMutation) {
		m.oldValue = func(context.Context) (*Stockpile, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StockpileMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StockpileMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StockpileMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StockpileMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Stockpile.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetIngridientId sets the "ingridientId" field.
func (m *StockpileMutation) SetIngridientId(u uuid.UUID) {
	m.ingridientId = &u
}

// IngridientId returns the value of the "ingridientId" field in the mutation.
func (m *StockpileMutation) IngridientId() (r uuid.UUID, exists bool) {
	v := m.ingridientId
	if v == nil {
		return
	}
	return *v, true
}

// OldIngridientId returns the old "ingridientId" field's value of the Stockpile entity.
// If the Stockpile object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StockpileMutation) OldIngridientId(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIngridientId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIngridientId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIngridientId: %w", err)
	}
	return oldValue.IngridientId, nil
}

// ResetIngridientId resets all changes to the "ingridientId" field.
func (m *StockpileMutation) ResetIngridientId() {
	m.ingridientId = nil
}

// SetOwnerId sets the "ownerId" field.
func (m *StockpileMutation) SetOwnerId(u uuid.UUID) {
	m.ownerId = &u
}

// OwnerId returns the value of the "ownerId" field in the mutation.
func (m *StockpileMutation) OwnerId() (r uuid.UUID, exists bool) {
	v := m.ownerId
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerId returns the old "ownerId" field's value of the Stockpile entity.
// If the Stockpile object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StockpileMutation) OldOwnerId(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwnerId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwnerId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerId: %w", err)
	}
	return oldValue.OwnerId, nil
}

// ResetOwnerId resets all changes to the "ownerId" field.
func (m *StockpileMutation) ResetOwnerId() {
	m.ownerId = nil
}

// SetQuantity sets the "quantity" field.
func (m *StockpileMutation) SetQuantity(i int64) {
	m.quantity = &i
	m.addquantity = nil
}

// Quantity returns the value of the "quantity" field in the mutation.
func (m *StockpileMutation) Quantity() (r int64, exists bool) {
	v := m.quantity
	if v == nil {
		return
	}
	return *v, true
}

// OldQuantity returns the old "quantity" field's value of the Stockpile entity.
// If the Stockpile object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StockpileMutation) OldQuantity(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldQuantity is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldQuantity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldQuantity: %w", err)
	}
	return oldValue.Quantity, nil
}

// AddQuantity adds i to the "quantity" field.
func (m *StockpileMutation) AddQuantity(i int64) {
	if m.addquantity != nil {
		*m.addquantity += i
	} else {
		m.addquantity = &i
	}
}

// AddedQuantity returns the value that was added to the "quantity" field in this mutation.
func (m *StockpileMutation) AddedQuantity() (r int64, exists bool) {
	v := m.addquantity
	if v == nil {
		return
	}
	return *v, true
}

// ResetQuantity resets all changes to the "quantity" field.
func (m *StockpileMutation) ResetQuantity() {
	m.quantity = nil
	m.addquantity = nil
}

// SetIngridientID sets the "ingridient" edge to the Ingridient entity by id.
func (m *StockpileMutation) SetIngridientID(id uuid.UUID) {
	m.ingridient = &id
}

// ClearIngridient clears the "ingridient" edge to the Ingridient entity.
func (m *StockpileMutation) ClearIngridient() {
	m.clearedingridient = true
}

// IngridientCleared reports if the "ingridient" edge to the Ingridient entity was cleared.
func (m *StockpileMutation) IngridientCleared() bool {
	return m.clearedingridient
}

// IngridientID returns the "ingridient" edge ID in the mutation.
func (m *StockpileMutation) IngridientID() (id uuid.UUID, exists bool) {
	if m.ingridient != nil {
		return *m.ingridient, true
	}
	return
}

// IngridientIDs returns the "ingridient" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// IngridientID instead. It exists only for internal usage by the builders.
func (m *StockpileMutation) IngridientIDs() (ids []uuid.UUID) {
	if id := m.ingridient; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetIngridient resets all changes to the "ingridient" edge.
func (m *StockpileMutation) ResetIngridient() {
	m.ingridient = nil
	m.clearedingridient = false
}

// Where appends a list predicates to the StockpileMutation builder.
func (m *StockpileMutation) Where(ps ...predicate.Stockpile) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StockpileMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StockpileMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Stockpile, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StockpileMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StockpileMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Stockpile).
func (m *StockpileMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StockpileMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.ingridientId != nil {
		fields = append(fields, stockpile.FieldIngridientId)
	}
	if m.ownerId != nil {
		fields = append(fields, stockpile.FieldOwnerId)
	}
	if m.quantity != nil {
		fields = append(fields, stockpile.FieldQuantity)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StockpileMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case stockpile.FieldIngridientId:
		return m.IngridientId()
	case stockpile.FieldOwnerId:
		return m.OwnerId()
	case stockpile.FieldQuantity:
		return m.Quantity()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StockpileMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case stockpile.FieldIngridientId:
		return m.OldIngridientId(ctx)
	case stockpile.FieldOwnerId:
		return m.OldOwnerId(ctx)
	case stockpile.FieldQuantity:
		return m.OldQuantity(ctx)
	}
	return nil, fmt.Errorf("unknown Stockpile field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StockpileMutation) SetField(name string, value ent.Value) error {
	switch name {
	case stockpile.FieldIngridientId:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIngridientId(v)
		return nil
	case stockpile.FieldOwnerId:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerId(v)
		return nil
	case stockpile.FieldQuantity:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetQuantity(v)
		return nil
	}
	return fmt.Errorf("unknown Stockpile field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StockpileMutation) AddedFields() []string {
	var fields []string
	if m.addquantity != nil {
		fields = append(fields, stockpile.FieldQuantity)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StockpileMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case stockpile.FieldQuantity:
		return m.AddedQuantity()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StockpileMutation) AddField(name string, value ent.Value) error {
	switch name {
	case stockpile.FieldQuantity:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddQuantity(v)
		return nil
	}
	return fmt.Errorf("unknown Stockpile numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StockpileMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StockpileMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StockpileMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Stockpile nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StockpileMutation) ResetField(name string) error {
	switch name {
	case stockpile.FieldIngridientId:
		m.ResetIngridientId()
		return nil
	case stockpile.FieldOwnerId:
		m.ResetOwnerId()
		return nil
	case stockpile.FieldQuantity:
		m.ResetQuantity()
		return nil
	}
	return fmt.Errorf("unknown Stockpile field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StockpileMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.ingridient != nil {
		edges = append(edges, stockpile.EdgeIngridient)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StockpileMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case stockpile.EdgeIngridient:
		if id := m.ingridient; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StockpileMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StockpileMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StockpileMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedingridient {
		edges = append(edges, stockpile.EdgeIngridient)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StockpileMutation) EdgeCleared(name string) bool {
	switch name {
	case stockpile.EdgeIngridient:
		return m.clearedingridient
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StockpileMutation) ClearEdge(name string) error {
	switch name {
	case stockpile.EdgeIngridient:
		m.ClearIngridient()
		return nil
	}
	return fmt.Errorf("unknown Stockpile unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StockpileMutation) ResetEdge(name string) error {
	switch name {
	case stockpile.EdgeIngridient:
		m.ResetIngridient()
		return nil
	}
	return fmt.Errorf("unknown Stockpile edge %s", name)
}