// Code generated by ent, DO NOT EDIT.

package request

import (
	"fume/internal/storage/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Request {
	return predicate.Request(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Request {
	return predicate.Request(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Request {
	return predicate.Request(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Request {
	return predicate.Request(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Request {
	return predicate.Request(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Request {
	return predicate.Request(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Request {
	return predicate.Request(sql.FieldLTE(FieldID, id))
}

// Endpoint applies equality check predicate on the "endpoint" field. It's identical to EndpointEQ.
func Endpoint(v string) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldEndpoint, v))
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldMethod, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldUpdatedAt, v))
}

// EndpointEQ applies the EQ predicate on the "endpoint" field.
func EndpointEQ(v string) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldEndpoint, v))
}

// EndpointNEQ applies the NEQ predicate on the "endpoint" field.
func EndpointNEQ(v string) predicate.Request {
	return predicate.Request(sql.FieldNEQ(FieldEndpoint, v))
}

// EndpointIn applies the In predicate on the "endpoint" field.
func EndpointIn(vs ...string) predicate.Request {
	return predicate.Request(sql.FieldIn(FieldEndpoint, vs...))
}

// EndpointNotIn applies the NotIn predicate on the "endpoint" field.
func EndpointNotIn(vs ...string) predicate.Request {
	return predicate.Request(sql.FieldNotIn(FieldEndpoint, vs...))
}

// EndpointGT applies the GT predicate on the "endpoint" field.
func EndpointGT(v string) predicate.Request {
	return predicate.Request(sql.FieldGT(FieldEndpoint, v))
}

// EndpointGTE applies the GTE predicate on the "endpoint" field.
func EndpointGTE(v string) predicate.Request {
	return predicate.Request(sql.FieldGTE(FieldEndpoint, v))
}

// EndpointLT applies the LT predicate on the "endpoint" field.
func EndpointLT(v string) predicate.Request {
	return predicate.Request(sql.FieldLT(FieldEndpoint, v))
}

// EndpointLTE applies the LTE predicate on the "endpoint" field.
func EndpointLTE(v string) predicate.Request {
	return predicate.Request(sql.FieldLTE(FieldEndpoint, v))
}

// EndpointContains applies the Contains predicate on the "endpoint" field.
func EndpointContains(v string) predicate.Request {
	return predicate.Request(sql.FieldContains(FieldEndpoint, v))
}

// EndpointHasPrefix applies the HasPrefix predicate on the "endpoint" field.
func EndpointHasPrefix(v string) predicate.Request {
	return predicate.Request(sql.FieldHasPrefix(FieldEndpoint, v))
}

// EndpointHasSuffix applies the HasSuffix predicate on the "endpoint" field.
func EndpointHasSuffix(v string) predicate.Request {
	return predicate.Request(sql.FieldHasSuffix(FieldEndpoint, v))
}

// EndpointEqualFold applies the EqualFold predicate on the "endpoint" field.
func EndpointEqualFold(v string) predicate.Request {
	return predicate.Request(sql.FieldEqualFold(FieldEndpoint, v))
}

// EndpointContainsFold applies the ContainsFold predicate on the "endpoint" field.
func EndpointContainsFold(v string) predicate.Request {
	return predicate.Request(sql.FieldContainsFold(FieldEndpoint, v))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.Request {
	return predicate.Request(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.Request {
	return predicate.Request(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.Request {
	return predicate.Request(sql.FieldNotIn(FieldMethod, vs...))
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.Request {
	return predicate.Request(sql.FieldGT(FieldMethod, v))
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.Request {
	return predicate.Request(sql.FieldGTE(FieldMethod, v))
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.Request {
	return predicate.Request(sql.FieldLT(FieldMethod, v))
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.Request {
	return predicate.Request(sql.FieldLTE(FieldMethod, v))
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.Request {
	return predicate.Request(sql.FieldContains(FieldMethod, v))
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.Request {
	return predicate.Request(sql.FieldHasPrefix(FieldMethod, v))
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.Request {
	return predicate.Request(sql.FieldHasSuffix(FieldMethod, v))
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.Request {
	return predicate.Request(sql.FieldEqualFold(FieldMethod, v))
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.Request {
	return predicate.Request(sql.FieldContainsFold(FieldMethod, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Request {
	return predicate.Request(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Request {
	return predicate.Request(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Request {
	return predicate.Request(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Request {
	return predicate.Request(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Request {
	return predicate.Request(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Request) predicate.Request {
	return predicate.Request(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Request) predicate.Request {
	return predicate.Request(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Request) predicate.Request {
	return predicate.Request(sql.NotPredicates(p))
}
