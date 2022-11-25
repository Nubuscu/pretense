// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/review"
	"nubuscu/pretense/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicCreate is the builder for creating a Topic entity.
type TopicCreate struct {
	config
	mutation *TopicMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TopicCreate) SetCreatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableCreatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TopicCreate) SetUpdatedAt(t time.Time) *TopicCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TopicCreate) SetNillableUpdatedAt(t *time.Time) *TopicCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TopicCreate) SetName(s string) *TopicCreate {
	tc.mutation.SetName(s)
	return tc
}

// AddReviewedByIDs adds the "reviewed_by" edge to the Review entity by IDs.
func (tc *TopicCreate) AddReviewedByIDs(ids ...int) *TopicCreate {
	tc.mutation.AddReviewedByIDs(ids...)
	return tc
}

// AddReviewedBy adds the "reviewed_by" edges to the Review entity.
func (tc *TopicCreate) AddReviewedBy(r ...*Review) *TopicCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddReviewedByIDs(ids...)
}

// AddIncludeIDs adds the "includes" edge to the Album entity by IDs.
func (tc *TopicCreate) AddIncludeIDs(ids ...int) *TopicCreate {
	tc.mutation.AddIncludeIDs(ids...)
	return tc
}

// AddIncludes adds the "includes" edges to the Album entity.
func (tc *TopicCreate) AddIncludes(a ...*Album) *TopicCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tc.AddIncludeIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tc *TopicCreate) Mutation() *TopicMutation {
	return tc.mutation
}

// Save creates the Topic in the database.
func (tc *TopicCreate) Save(ctx context.Context) (*Topic, error) {
	var (
		err  error
		node *Topic
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Topic)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TopicMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TopicCreate) SaveX(ctx context.Context) *Topic {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TopicCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TopicCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TopicCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := topic.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := topic.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TopicCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Topic.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Topic.updated_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Topic.name"`)}
	}
	return nil
}

func (tc *TopicCreate) sqlSave(ctx context.Context) (*Topic, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TopicCreate) createSpec() (*Topic, *sqlgraph.CreateSpec) {
	var (
		_node = &Topic{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: topic.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(topic.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(topic.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := tc.mutation.ReviewedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   topic.ReviewedByTable,
			Columns: topic.ReviewedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.IncludesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.IncludesTable,
			Columns: topic.IncludesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TopicCreateBulk is the builder for creating many Topic entities in bulk.
type TopicCreateBulk struct {
	config
	builders []*TopicCreate
}

// Save creates the Topic entities in the database.
func (tcb *TopicCreateBulk) Save(ctx context.Context) ([]*Topic, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Topic, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopicMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TopicCreateBulk) SaveX(ctx context.Context) []*Topic {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TopicCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TopicCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
