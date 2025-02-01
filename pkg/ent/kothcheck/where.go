// Code generated by ent, DO NOT EDIT.

package kothcheck

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldName, v))
}

// File applies equality check predicate on the "file" field. It's identical to FileEQ.
func File(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldFile, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldWeight, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldContainsFold(FieldName, v))
}

// FileEQ applies the EQ predicate on the "file" field.
func FileEQ(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldFile, v))
}

// FileNEQ applies the NEQ predicate on the "file" field.
func FileNEQ(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNEQ(FieldFile, v))
}

// FileIn applies the In predicate on the "file" field.
func FileIn(vs ...string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldIn(FieldFile, vs...))
}

// FileNotIn applies the NotIn predicate on the "file" field.
func FileNotIn(vs ...string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNotIn(FieldFile, vs...))
}

// FileGT applies the GT predicate on the "file" field.
func FileGT(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGT(FieldFile, v))
}

// FileGTE applies the GTE predicate on the "file" field.
func FileGTE(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGTE(FieldFile, v))
}

// FileLT applies the LT predicate on the "file" field.
func FileLT(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLT(FieldFile, v))
}

// FileLTE applies the LTE predicate on the "file" field.
func FileLTE(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLTE(FieldFile, v))
}

// FileContains applies the Contains predicate on the "file" field.
func FileContains(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldContains(FieldFile, v))
}

// FileHasPrefix applies the HasPrefix predicate on the "file" field.
func FileHasPrefix(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldHasPrefix(FieldFile, v))
}

// FileHasSuffix applies the HasSuffix predicate on the "file" field.
func FileHasSuffix(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldHasSuffix(FieldFile, v))
}

// FileEqualFold applies the EqualFold predicate on the "file" field.
func FileEqualFold(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEqualFold(FieldFile, v))
}

// FileContainsFold applies the ContainsFold predicate on the "file" field.
func FileContainsFold(v string) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldContainsFold(FieldFile, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.KothCheck {
	return predicate.KothCheck(sql.FieldLTE(FieldWeight, v))
}

// HasStatuses applies the HasEdge predicate on the "statuses" edge.
func HasStatuses() predicate.KothCheck {
	return predicate.KothCheck(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusesTable, StatusesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusesWith applies the HasEdge predicate on the "statuses" edge with a given conditions (other predicates).
func HasStatusesWith(preds ...predicate.KothStatus) predicate.KothCheck {
	return predicate.KothCheck(func(s *sql.Selector) {
		step := newStatusesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KothCheck) predicate.KothCheck {
	return predicate.KothCheck(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KothCheck) predicate.KothCheck {
	return predicate.KothCheck(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.KothCheck) predicate.KothCheck {
	return predicate.KothCheck(sql.NotPredicates(p))
}
