// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scorify/scorify/pkg/ent/audit"
	"github.com/scorify/scorify/pkg/ent/predicate"
	"github.com/scorify/scorify/pkg/structs"
)

// AuditUpdate is the builder for updating Audit entities.
type AuditUpdate struct {
	config
	hooks    []Hook
	mutation *AuditMutation
}

// Where appends a list predicates to the AuditUpdate builder.
func (au *AuditUpdate) Where(ps ...predicate.Audit) *AuditUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAction sets the "action" field.
func (au *AuditUpdate) SetAction(a audit.Action) *AuditUpdate {
	au.mutation.SetAction(a)
	return au
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (au *AuditUpdate) SetNillableAction(a *audit.Action) *AuditUpdate {
	if a != nil {
		au.SetAction(*a)
	}
	return au
}

// SetIP sets the "ip" field.
func (au *AuditUpdate) SetIP(s *structs.Inet) *AuditUpdate {
	au.mutation.SetIP(s)
	return au
}

// SetMessage sets the "message" field.
func (au *AuditUpdate) SetMessage(s string) *AuditUpdate {
	au.mutation.SetMessage(s)
	return au
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (au *AuditUpdate) SetNillableMessage(s *string) *AuditUpdate {
	if s != nil {
		au.SetMessage(*s)
	}
	return au
}

// Mutation returns the AuditMutation object of the builder.
func (au *AuditUpdate) Mutation() *AuditMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuditUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuditUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuditUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuditUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AuditUpdate) check() error {
	if v, ok := au.mutation.Action(); ok {
		if err := audit.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Audit.action": %w`, err)}
		}
	}
	if v, ok := au.mutation.IP(); ok {
		if err := audit.IPValidator(v.String()); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`ent: validator failed for field "Audit.ip": %w`, err)}
		}
	}
	if v, ok := au.mutation.Message(); ok {
		if err := audit.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Audit.message": %w`, err)}
		}
	}
	return nil
}

func (au *AuditUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(audit.Table, audit.Columns, sqlgraph.NewFieldSpec(audit.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Action(); ok {
		_spec.SetField(audit.FieldAction, field.TypeEnum, value)
	}
	if value, ok := au.mutation.IP(); ok {
		_spec.SetField(audit.FieldIP, field.TypeString, value)
	}
	if value, ok := au.mutation.Message(); ok {
		_spec.SetField(audit.FieldMessage, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AuditUpdateOne is the builder for updating a single Audit entity.
type AuditUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuditMutation
}

// SetAction sets the "action" field.
func (auo *AuditUpdateOne) SetAction(a audit.Action) *AuditUpdateOne {
	auo.mutation.SetAction(a)
	return auo
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (auo *AuditUpdateOne) SetNillableAction(a *audit.Action) *AuditUpdateOne {
	if a != nil {
		auo.SetAction(*a)
	}
	return auo
}

// SetIP sets the "ip" field.
func (auo *AuditUpdateOne) SetIP(s *structs.Inet) *AuditUpdateOne {
	auo.mutation.SetIP(s)
	return auo
}

// SetMessage sets the "message" field.
func (auo *AuditUpdateOne) SetMessage(s string) *AuditUpdateOne {
	auo.mutation.SetMessage(s)
	return auo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (auo *AuditUpdateOne) SetNillableMessage(s *string) *AuditUpdateOne {
	if s != nil {
		auo.SetMessage(*s)
	}
	return auo
}

// Mutation returns the AuditMutation object of the builder.
func (auo *AuditUpdateOne) Mutation() *AuditMutation {
	return auo.mutation
}

// Where appends a list predicates to the AuditUpdate builder.
func (auo *AuditUpdateOne) Where(ps ...predicate.Audit) *AuditUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuditUpdateOne) Select(field string, fields ...string) *AuditUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Audit entity.
func (auo *AuditUpdateOne) Save(ctx context.Context) (*Audit, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuditUpdateOne) SaveX(ctx context.Context) *Audit {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuditUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuditUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AuditUpdateOne) check() error {
	if v, ok := auo.mutation.Action(); ok {
		if err := audit.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Audit.action": %w`, err)}
		}
	}
	if v, ok := auo.mutation.IP(); ok {
		if err := audit.IPValidator(v.String()); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`ent: validator failed for field "Audit.ip": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Message(); ok {
		if err := audit.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Audit.message": %w`, err)}
		}
	}
	return nil
}

func (auo *AuditUpdateOne) sqlSave(ctx context.Context) (_node *Audit, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(audit.Table, audit.Columns, sqlgraph.NewFieldSpec(audit.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Audit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, audit.FieldID)
		for _, f := range fields {
			if !audit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != audit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Action(); ok {
		_spec.SetField(audit.FieldAction, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.IP(); ok {
		_spec.SetField(audit.FieldIP, field.TypeString, value)
	}
	if value, ok := auo.mutation.Message(); ok {
		_spec.SetField(audit.FieldMessage, field.TypeString, value)
	}
	_node = &Audit{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
