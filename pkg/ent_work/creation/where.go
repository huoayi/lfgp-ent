// Code generated by ent, DO NOT EDIT.

package creation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/predicate"
	"github.com/huoayi/lfgp-ent/pkg/enum"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldDeletedAt, v))
}

// Parameter applies equality check predicate on the "parameter" field. It's identical to ParameterEQ.
func Parameter(v string) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldParameter, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldURL, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldDeletedAt, v))
}

// CreationTypeEQ applies the EQ predicate on the "creation_type" field.
func CreationTypeEQ(v enum.CreationType) predicate.Creation {
	vc := v
	return predicate.Creation(sql.FieldEQ(FieldCreationType, vc))
}

// CreationTypeNEQ applies the NEQ predicate on the "creation_type" field.
func CreationTypeNEQ(v enum.CreationType) predicate.Creation {
	vc := v
	return predicate.Creation(sql.FieldNEQ(FieldCreationType, vc))
}

// CreationTypeIn applies the In predicate on the "creation_type" field.
func CreationTypeIn(vs ...enum.CreationType) predicate.Creation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Creation(sql.FieldIn(FieldCreationType, v...))
}

// CreationTypeNotIn applies the NotIn predicate on the "creation_type" field.
func CreationTypeNotIn(vs ...enum.CreationType) predicate.Creation {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Creation(sql.FieldNotIn(FieldCreationType, v...))
}

// ParameterEQ applies the EQ predicate on the "parameter" field.
func ParameterEQ(v string) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldParameter, v))
}

// ParameterNEQ applies the NEQ predicate on the "parameter" field.
func ParameterNEQ(v string) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldParameter, v))
}

// ParameterIn applies the In predicate on the "parameter" field.
func ParameterIn(vs ...string) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldParameter, vs...))
}

// ParameterNotIn applies the NotIn predicate on the "parameter" field.
func ParameterNotIn(vs ...string) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldParameter, vs...))
}

// ParameterGT applies the GT predicate on the "parameter" field.
func ParameterGT(v string) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldParameter, v))
}

// ParameterGTE applies the GTE predicate on the "parameter" field.
func ParameterGTE(v string) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldParameter, v))
}

// ParameterLT applies the LT predicate on the "parameter" field.
func ParameterLT(v string) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldParameter, v))
}

// ParameterLTE applies the LTE predicate on the "parameter" field.
func ParameterLTE(v string) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldParameter, v))
}

// ParameterContains applies the Contains predicate on the "parameter" field.
func ParameterContains(v string) predicate.Creation {
	return predicate.Creation(sql.FieldContains(FieldParameter, v))
}

// ParameterHasPrefix applies the HasPrefix predicate on the "parameter" field.
func ParameterHasPrefix(v string) predicate.Creation {
	return predicate.Creation(sql.FieldHasPrefix(FieldParameter, v))
}

// ParameterHasSuffix applies the HasSuffix predicate on the "parameter" field.
func ParameterHasSuffix(v string) predicate.Creation {
	return predicate.Creation(sql.FieldHasSuffix(FieldParameter, v))
}

// ParameterEqualFold applies the EqualFold predicate on the "parameter" field.
func ParameterEqualFold(v string) predicate.Creation {
	return predicate.Creation(sql.FieldEqualFold(FieldParameter, v))
}

// ParameterContainsFold applies the ContainsFold predicate on the "parameter" field.
func ParameterContainsFold(v string) predicate.Creation {
	return predicate.Creation(sql.FieldContainsFold(FieldParameter, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Creation {
	return predicate.Creation(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Creation {
	return predicate.Creation(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Creation {
	return predicate.Creation(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Creation {
	return predicate.Creation(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Creation {
	return predicate.Creation(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Creation {
	return predicate.Creation(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Creation {
	return predicate.Creation(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Creation {
	return predicate.Creation(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Creation {
	return predicate.Creation(sql.FieldContainsFold(FieldURL, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Creation {
	return predicate.Creation(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Creation {
	return predicate.Creation(sql.FieldNotIn(FieldUserID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Creation {
	return predicate.Creation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Creation {
	return predicate.Creation(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Creation) predicate.Creation {
	return predicate.Creation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Creation) predicate.Creation {
	return predicate.Creation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Creation) predicate.Creation {
	return predicate.Creation(sql.NotPredicates(p))
}
