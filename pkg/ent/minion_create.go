// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/status"
)

// MinionCreate is the builder for creating a Minion entity.
type MinionCreate struct {
	config
	mutation *MinionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (mc *MinionCreate) SetCreateTime(t time.Time) *MinionCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mc *MinionCreate) SetNillableCreateTime(t *time.Time) *MinionCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the "update_time" field.
func (mc *MinionCreate) SetUpdateTime(t time.Time) *MinionCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mc *MinionCreate) SetNillableUpdateTime(t *time.Time) *MinionCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MinionCreate) SetName(s string) *MinionCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetIP sets the "ip" field.
func (mc *MinionCreate) SetIP(s string) *MinionCreate {
	mc.mutation.SetIP(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MinionCreate) SetID(u uuid.UUID) *MinionCreate {
	mc.mutation.SetID(u)
	return mc
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (mc *MinionCreate) AddStatusIDs(ids ...uuid.UUID) *MinionCreate {
	mc.mutation.AddStatusIDs(ids...)
	return mc
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (mc *MinionCreate) AddStatuses(s ...*Status) *MinionCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mc.AddStatusIDs(ids...)
}

// Mutation returns the MinionMutation object of the builder.
func (mc *MinionCreate) Mutation() *MinionMutation {
	return mc.mutation
}

// Save creates the Minion in the database.
func (mc *MinionCreate) Save(ctx context.Context) (*Minion, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MinionCreate) SaveX(ctx context.Context) *Minion {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MinionCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MinionCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MinionCreate) defaults() {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := minion.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := minion.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MinionCreate) check() error {
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Minion.create_time"`)}
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Minion.update_time"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Minion.name"`)}
	}
	if _, ok := mc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "Minion.ip"`)}
	}
	return nil
}

func (mc *MinionCreate) sqlSave(ctx context.Context) (*Minion, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MinionCreate) createSpec() (*Minion, *sqlgraph.CreateSpec) {
	var (
		_node = &Minion{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(minion.Table, sqlgraph.NewFieldSpec(minion.FieldID, field.TypeUUID))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(minion.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.SetField(minion.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(minion.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.IP(); ok {
		_spec.SetField(minion.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if nodes := mc.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   minion.StatusesTable,
			Columns: []string{minion.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MinionCreateBulk is the builder for creating many Minion entities in bulk.
type MinionCreateBulk struct {
	config
	err      error
	builders []*MinionCreate
}

// Save creates the Minion entities in the database.
func (mcb *MinionCreateBulk) Save(ctx context.Context) ([]*Minion, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Minion, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MinionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MinionCreateBulk) SaveX(ctx context.Context) []*Minion {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MinionCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MinionCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
