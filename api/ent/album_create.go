// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"nubuscu/pretense/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AlbumCreate is the builder for creating a Album entity.
type AlbumCreate struct {
	config
	mutation *AlbumMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AlbumCreate) SetCreatedAt(t time.Time) *AlbumCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AlbumCreate) SetNillableCreatedAt(t *time.Time) *AlbumCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AlbumCreate) SetUpdatedAt(t time.Time) *AlbumCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AlbumCreate) SetNillableUpdatedAt(t *time.Time) *AlbumCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AlbumCreate) SetName(s string) *AlbumCreate {
	ac.mutation.SetName(s)
	return ac
}

// AddByIDs adds the "by" edge to the Artist entity by IDs.
func (ac *AlbumCreate) AddByIDs(ids ...int) *AlbumCreate {
	ac.mutation.AddByIDs(ids...)
	return ac
}

// AddBy adds the "by" edges to the Artist entity.
func (ac *AlbumCreate) AddBy(a ...*Artist) *AlbumCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddByIDs(ids...)
}

// AddIncludedInIDs adds the "included_in" edge to the Topic entity by IDs.
func (ac *AlbumCreate) AddIncludedInIDs(ids ...int) *AlbumCreate {
	ac.mutation.AddIncludedInIDs(ids...)
	return ac
}

// AddIncludedIn adds the "included_in" edges to the Topic entity.
func (ac *AlbumCreate) AddIncludedIn(t ...*Topic) *AlbumCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddIncludedInIDs(ids...)
}

// Mutation returns the AlbumMutation object of the builder.
func (ac *AlbumCreate) Mutation() *AlbumMutation {
	return ac.mutation
}

// Save creates the Album in the database.
func (ac *AlbumCreate) Save(ctx context.Context) (*Album, error) {
	var (
		err  error
		node *Album
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Album)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AlbumMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AlbumCreate) SaveX(ctx context.Context) *Album {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AlbumCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AlbumCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AlbumCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := album.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := album.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AlbumCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Album.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Album.updated_at"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Album.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := album.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Album.name": %w`, err)}
		}
	}
	return nil
}

func (ac *AlbumCreate) sqlSave(ctx context.Context) (*Album, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AlbumCreate) createSpec() (*Album, *sqlgraph.CreateSpec) {
	var (
		_node = &Album{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: album.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: album.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(album.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(album.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(album.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := ac.mutation.ByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ByTable,
			Columns: album.ByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.IncludedInIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.IncludedInTable,
			Columns: album.IncludedInPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
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

// AlbumCreateBulk is the builder for creating many Album entities in bulk.
type AlbumCreateBulk struct {
	config
	builders []*AlbumCreate
}

// Save creates the Album entities in the database.
func (acb *AlbumCreateBulk) Save(ctx context.Context) ([]*Album, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Album, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AlbumMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AlbumCreateBulk) SaveX(ctx context.Context) []*Album {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AlbumCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AlbumCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
