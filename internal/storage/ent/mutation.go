// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fume/internal/storage/ent/predicate"
	"fume/internal/storage/ent/request"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRequest = "Request"
)

// RequestMutation represents an operation that mutates the Request nodes in the graph.
type RequestMutation struct {
	config
	op            Op
	typ           string
	id            *int
	endpoint      *string
	method        *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Request, error)
	predicates    []predicate.Request
}

var _ ent.Mutation = (*RequestMutation)(nil)

// requestOption allows management of the mutation configuration using functional options.
type requestOption func(*RequestMutation)

// newRequestMutation creates new mutation for the Request entity.
func newRequestMutation(c config, op Op, opts ...requestOption) *RequestMutation {
	m := &RequestMutation{
		config:        c,
		op:            op,
		typ:           TypeRequest,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRequestID sets the ID field of the mutation.
func withRequestID(id int) requestOption {
	return func(m *RequestMutation) {
		var (
			err   error
			once  sync.Once
			value *Request
		)
		m.oldValue = func(ctx context.Context) (*Request, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Request.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRequest sets the old Request of the mutation.
func withRequest(node *Request) requestOption {
	return func(m *RequestMutation) {
		m.oldValue = func(context.Context) (*Request, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RequestMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RequestMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Request entities.
func (m *RequestMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RequestMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RequestMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Request.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEndpoint sets the "endpoint" field.
func (m *RequestMutation) SetEndpoint(s string) {
	m.endpoint = &s
}

// Endpoint returns the value of the "endpoint" field in the mutation.
func (m *RequestMutation) Endpoint() (r string, exists bool) {
	v := m.endpoint
	if v == nil {
		return
	}
	return *v, true
}

// OldEndpoint returns the old "endpoint" field's value of the Request entity.
// If the Request object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestMutation) OldEndpoint(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEndpoint is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEndpoint requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEndpoint: %w", err)
	}
	return oldValue.Endpoint, nil
}

// ResetEndpoint resets all changes to the "endpoint" field.
func (m *RequestMutation) ResetEndpoint() {
	m.endpoint = nil
}

// SetMethod sets the "method" field.
func (m *RequestMutation) SetMethod(s string) {
	m.method = &s
}

// Method returns the value of the "method" field in the mutation.
func (m *RequestMutation) Method() (r string, exists bool) {
	v := m.method
	if v == nil {
		return
	}
	return *v, true
}

// OldMethod returns the old "method" field's value of the Request entity.
// If the Request object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestMutation) OldMethod(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMethod is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMethod requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMethod: %w", err)
	}
	return oldValue.Method, nil
}

// ResetMethod resets all changes to the "method" field.
func (m *RequestMutation) ResetMethod() {
	m.method = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *RequestMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *RequestMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Request entity.
// If the Request object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *RequestMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *RequestMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *RequestMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Request entity.
// If the Request object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *RequestMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the RequestMutation builder.
func (m *RequestMutation) Where(ps ...predicate.Request) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the RequestMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *RequestMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Request, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *RequestMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *RequestMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Request).
func (m *RequestMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RequestMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.endpoint != nil {
		fields = append(fields, request.FieldEndpoint)
	}
	if m.method != nil {
		fields = append(fields, request.FieldMethod)
	}
	if m.created_at != nil {
		fields = append(fields, request.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, request.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RequestMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case request.FieldEndpoint:
		return m.Endpoint()
	case request.FieldMethod:
		return m.Method()
	case request.FieldCreatedAt:
		return m.CreatedAt()
	case request.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RequestMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case request.FieldEndpoint:
		return m.OldEndpoint(ctx)
	case request.FieldMethod:
		return m.OldMethod(ctx)
	case request.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case request.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Request field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RequestMutation) SetField(name string, value ent.Value) error {
	switch name {
	case request.FieldEndpoint:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEndpoint(v)
		return nil
	case request.FieldMethod:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMethod(v)
		return nil
	case request.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case request.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Request field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RequestMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RequestMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RequestMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Request numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RequestMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RequestMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RequestMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Request nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RequestMutation) ResetField(name string) error {
	switch name {
	case request.FieldEndpoint:
		m.ResetEndpoint()
		return nil
	case request.FieldMethod:
		m.ResetMethod()
		return nil
	case request.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case request.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Request field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RequestMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RequestMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RequestMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RequestMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RequestMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RequestMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RequestMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Request unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RequestMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Request edge %s", name)
}
